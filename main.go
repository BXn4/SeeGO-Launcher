package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"seegolauncher/internal/localization"
	"seegolauncher/internal/services"
	"seegolauncher/internal/utils"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/public/images/seego_icon.png
var icon []byte

const (
	Tray int = iota
	Show
	Minimized
	Maximized
)

type App struct {
	mu           sync.Mutex
	app          *application.App
	config       *services.Config
	appState     int
	window       *application.WebviewWindow
	view         string
	dialog       *application.MessageDialog
	ready        bool
	trayNotified bool
}

func (a *App) trayMinimize() {
	a.window.Hide()
	a.appState = Tray
	if !a.trayNotified {
		a.trayNotified = true
		utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, a.config.GetLanguage()))
	}
}

func (a *App) restoreFromTray() {
	a.window.UnMinimise()
	a.window.Show()
	a.window.Focus()
	a.appState = Show
}

func (a *App) minimize() {
	a.window.Minimise()
	a.appState = Minimized
}

func (a *App) setView(w, h int, n, v string) {
	a.window.SetSize(w, h)
	a.app.Event.Emit(n, v)
	a.view = v
}

func init() {
	application.RegisterEvent[string]("time")
}

func main() {
	envFile, err := godotenv.Read(".env")

	hasConfig := err == nil

	if !hasConfig {
		log.Warnf("Cannot find .env file!")
	}

	if services.OA == "" {
		services.OA = envFile["OA"]
	}

	log.Info("SeeGO Launcher by BXn4")

	a := &App{appState: Show, dialog: nil}
	a.config = services.ConfigService()

	frontendFS, err := fs.Sub(assets, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	a.app = application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				a.mu.Lock()
				defer a.mu.Unlock()
				log.Info("Tried to open a new instance, but its already running. Showing the instance.")
				switch a.appState {
				case Tray:
					a.restoreFromTray()
				default:
					a.window.Focus()
					a.window.Flash(true)
				}
			},
		},
		Services: []application.Service{
			application.NewService(services.LocalizationService()),
			application.NewService(services.ConfigService()),
			application.NewService(&services.CacheService{}),
			application.NewService(&services.API{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(frontendFS),
			Middleware: func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if strings.HasSuffix(r.URL.Path, ".css") {
						w.Header().Set("Content-Type", "text/css; charset=utf-8")
					}
					next.ServeHTTP(w, r)
				})
			},
		},
	})

	systray := a.app.SystemTray.New()
	systray.SetLabel("SeeGO Launcher")
	systray.SetIcon(icon)

	systray.OnClick(func() {
		switch a.appState {
		case Tray:
			a.restoreFromTray()
		default:
			a.window.Focus()
			a.window.Flash(true)
		}
	})

	menu := a.app.NewMenu()
	menu.Add("SeeGO Launcher")
	menu.AddSeparator()
	menu.Add("Show").OnClick(func(ctx *application.Context) {
		switch a.appState {
		case Tray:
			a.restoreFromTray()
		default:
			a.window.Focus()
			a.window.Flash(true)
		}
	})

	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		a.app.Quit()
	})
	systray.SetMenu(menu)

	window := a.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:                      "SeeGO Launcher",
		Width:                      476,
		Height:                     300,
		Frameless:                  true,
		AlwaysOnTop:                true,
		DisableResize:              true,
		BackgroundColour:           application.NewRGB(52, 58, 64),
		URL:                        "/",
		Hidden:                     false,
		DevToolsEnabled:            false,
		DefaultContextMenuDisabled: true,
	})
	a.view = "splash"

	window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		switch a.view {
		case "main":
			e.Cancel()
			a.trayMinimize()
		default:
			e.Cancel()
			a.app.Quit()
			return
		}
	})

	window.OnWindowEvent(events.Common.WindowRuntimeReady, func(e *application.WindowEvent) {

	})

	a.window = window

	a.app.Event.On("minimize", func(e *application.CustomEvent) {
		e.Cancel()
		a.minimize()
	})

	a.app.Event.On("close", func(e *application.CustomEvent) {
		switch a.view {
		case "main":
			e.Cancel()
			a.trayMinimize()
		default:
			e.Cancel()
			a.app.Quit()
			return
		}
	})

	a.app.Event.On("terms-declined", func(e *application.CustomEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()

		a.window.SetIgnoreMouseEvents(true)
		if a.dialog == nil {
			dialog := a.app.Dialog.Warning().
				SetTitle(services.LocalizationService().Get(localization.TermsDeclinedTitle, a.config.GetLanguage())).
				SetMessage(services.LocalizationService().Get(localization.TermsDeclinedContent, a.config.GetLanguage()))

			ok := dialog.AddButton("Ok")
			ok.OnClick(
				func() {
					os.Exit(0)
				})

			dialog.SetDefaultButton(ok)

			a.dialog = dialog
		}

		a.dialog.Show()
	})

	a.app.Event.On("terms-accepted", func(e *application.CustomEvent) {
		a.config.SetTermsAccepted()
		a.setView(1200, int(1200/(16.0/9.0)), "app:navigate", "main")
		window.Center()
	})

	// https://youtu.be/xXKqODp94VA

	go func() {
		time.Sleep(2 * time.Second)
		if WaitUntil(
			func() bool {
				success := services.LoadCache()
				return success == true
			},
			5*time.Second,
		) {
			// will notify if the cache was not successed
			a.app.Quit()
		}

		//after splash
		window.SetAlwaysOnTop(false)

		if !a.config.GetTermsAccepted() {
			seeGOInfo1 := services.LocalizationService().Get(localization.SeeGOInfo1, a.config.GetLanguage())
			seeGOInfo2 := services.LocalizationService().Get(localization.SeeGOInfo2, a.config.GetLanguage())
			seeGOInfo3 := services.LocalizationService().Get(localization.SeeGOInfo3, a.config.GetLanguage())
			seeGOInfo4 := services.LocalizationService().Get(localization.SeeGOInfo4, a.config.GetLanguage())
			a.app.Dialog.Info().SetTitle("SeeGO Launcher").SetMessage(fmt.Sprintf("%s\n%s\n%s\n%s", seeGOInfo1, seeGOInfo2, seeGOInfo3, seeGOInfo4)).Show()
			log.Info("Terms is not accepted, showing the terms window")

			a.setView(620, 480, "app:navigate", "terms")
			window.Center()
			return
		}
		a.setView(1200, int(1200/(16.0/9.0)), "app:navigate", "main")
		window.Center()
	}()

	if err := a.app.Run(); err != nil {
		log.Fatal(err)
	}
}

func WaitUntil(condition func() bool, timeout time.Duration) bool {
	startTime := time.Now()
	for !condition() {
		if time.Since(startTime) >= timeout {
			return true
		}
		time.Sleep(10 * time.Millisecond)
	}
	return false
}
