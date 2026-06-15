package main

import (
	"embed"
	"fmt"
	"os"
	"seegolauncher/internal/cache"
	"seegolauncher/internal/localization"
	"seegolauncher/internal/services"
	"seegolauncher/internal/utils"
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
	mu       sync.Mutex
	appState int
	window   *application.WebviewWindow
	view     string
	dialog   *application.MessageDialog
	ready    bool
}

func (a *App) setWindow(w *application.WebviewWindow) {
	a.window = w
}

func (a *App) getWindow() *application.WebviewWindow {
	return a.window
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
	config := services.ConfigService()

	a := &App{appState: Show, dialog: nil}

	app := application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				a.mu.Lock()
				defer a.mu.Unlock()
				log.Info("Tried to open a new instance, but its already running. Showing the instance.")
				w := a.getWindow()
				switch a.appState {
				case Tray:
					w.UnMinimise()
					w.Show()
					w.Focus()
					a.appState = Show
				default:
					w.Focus()
					w.Flash(true)
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
			Handler: application.AssetFileServerFS(assets),
		},
	})

	systray := app.SystemTray.New()
	systray.SetLabel("SeeGO Launcher")
	systray.SetIcon(icon)

	systray.OnClick(func() {
		w := a.getWindow()
		switch a.appState {
		case Tray:
			w.UnMinimise()
			w.Show()
			w.Focus()
			a.appState = Show
		default:
			w.Focus()
			w.Flash(true)
		}
	})

	menu := app.NewMenu()
	menu.Add("SeeGO Launcher")
	menu.AddSeparator()
	menu.Add("Show").OnClick(func(ctx *application.Context) {
		w := a.getWindow()
		switch a.appState {
		case Tray:
			w.UnMinimise()
			w.Show()
			w.Focus()
			a.appState = Show
		default:
			w.Focus()
			w.Flash(true)
		}
	})

	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		app.Quit()
	})
	systray.SetMenu(menu)

	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
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

	a.setWindow(window)
	a.view = "splash"

	window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		switch a.view {
		case "main":
			e.Cancel()
			a.getWindow().Hide()
			a.appState = Tray
			utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, config.GetLanguage()))

		default:
			e.Cancel()
			app.Quit()
			return
		}
	})

	window.OnWindowEvent(events.Common.WindowRuntimeReady, func(e *application.WindowEvent) {

	})

	app.Event.On("minimize", func(e *application.CustomEvent) {
		e.Cancel()
		a.getWindow().Minimise()
		a.appState = Minimized
	})

	app.Event.On("close", func(e *application.CustomEvent) {
		switch a.view {
		case "main":
			e.Cancel()
			a.getWindow().Hide()
			a.appState = Tray
			utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, config.GetLanguage()))

		default:
			e.Cancel()
			app.Quit()
			return
		}
	})

	app.Event.On("terms-declined", func(e *application.CustomEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()

		a.window.SetIgnoreMouseEvents(true)
		if a.dialog == nil {
			dialog := app.Dialog.Warning().
				SetTitle(services.LocalizationService().Get(localization.TermsDeclinedTitle, config.GetLanguage())).
				SetMessage(services.LocalizationService().Get(localization.TermsDeclinedContent, config.GetLanguage()))

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

	app.Event.On("terms-accepted", func(e *application.CustomEvent) {
		config.SetTermsAccepted()

		window.SetSize(1200, int(1200/(16.0/9.0)))
		app.Event.Emit("app:navigate", "main")
		a.view = "main"
		window.Center()
	})

	/*app.Event.On("app-ready", func(e *application.CustomEvent) {
	a.mu.Lock()
	defer a.mu.Unlock()

	}) */

	// https://youtu.be/xXKqODp94VA

	go func() {
		time.Sleep(1 * time.Second)
		if WaitUntil(
			func() bool {
				success := cache.LoadCache()
				return success == true
			},
			5*time.Second,
		) {
			//
			app.Quit()
		}

		//after splash
		window.SetAlwaysOnTop(false)

		if !config.GetTermsAccepted() {
			seeGOInfo1 := services.LocalizationService().Get(localization.SeeGOInfo1, config.GetLanguage())
			seeGOInfo2 := services.LocalizationService().Get(localization.SeeGOInfo2, config.GetLanguage())
			seeGOInfo3 := services.LocalizationService().Get(localization.SeeGOInfo3, config.GetLanguage())
			seeGOInfo4 := services.LocalizationService().Get(localization.SeeGOInfo4, config.GetLanguage())
			app.Dialog.Info().SetTitle("SeeGO Launcher").SetMessage(fmt.Sprintf("%s\n%s\n%s\n%s", seeGOInfo1, seeGOInfo2, seeGOInfo3, seeGOInfo4)).Show()
			log.Info("Terms is not accepted, showing the terms window")

			window.SetSize(620, 480)
			app.Event.Emit("app:navigate", "terms")
			a.view = "terms"
			window.Center()
			return
		}
		window.SetSize(1200, int(1200/(16.0/9.0)))
		app.Event.Emit("app:navigate", "main")
		a.view = "main"
		window.Center()
	}()

	if err := app.Run(); err != nil {
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
