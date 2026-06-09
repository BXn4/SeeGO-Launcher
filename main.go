package main

import (
	"embed"
	"os"
	"seegolauncher/internal/localization"
	"seegolauncher/internal/services"
	"seegolauncher/internal/utils"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	Tray int = iota
	Show
	Minimized
	Maximized
)

type App struct {
	mu         sync.Mutex
	appState   int
	window     *application.WebviewWindow
	splashDone bool
	dialog     *application.MessageDialog
}

func (a *App) setWindow(w *application.WebviewWindow) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.window = w
}

func (a *App) getWindow() *application.WebviewWindow {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.window
}

func init() {
	application.RegisterEvent[string]("time")
}

func main() {
	log.Info("SeeGO Launcher by BXn4")
	config := services.ConfigService()

	a := &App{appState: Show, dialog: nil}

	// cache.LoadCache()

	app := application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				log.Info("Tried to open a new instance, but its already running. Showing the instance.")
				w := a.getWindow()
				a.mu.Lock()
				defer a.mu.Unlock()
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
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	systray := app.SystemTray.New()
	systray.SetLabel("SeeGO Launcher")

	systray.OnClick(func() {
		a.appState = Show
		a.window.Show()
		a.window.Focus()
	})

	menu := app.NewMenu()
	menu.Add("Show").OnClick(func(ctx *application.Context) {
		a.appState = Show
		a.window.Show()
		a.window.Focus()
	})
	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		app.Quit()
	})
	systray.SetMenu(menu)

	splash := app.Window.NewWithOptions(application.WebviewWindowOptions{
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

	main := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:                      "SeeGO Launcher",
		Width:                      1200,
		Height:                     int(1200 / (16.0 / 9.0)),
		MinWidth:                   600,
		MinHeight:                  400,
		Frameless:                  true,
		DisableResize:              false,
		BackgroundColour:           application.NewRGB(52, 58, 64),
		URL:                        "/",
		Hidden:                     true,
		DevToolsEnabled:            false,
		DefaultContextMenuDisabled: true,
	})

	terms := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:                      "SeeGO Launcher",
		Width:                      800,
		Height:                     600,
		Frameless:                  true,
		AlwaysOnTop:                false,
		DisableResize:              true,
		BackgroundColour:           application.NewRGB(52, 58, 64),
		URL:                        "/",
		Hidden:                     true,
		DevToolsEnabled:            false,
		DefaultContextMenuDisabled: true,
	})

	a.setWindow(splash)

	splash.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()
		done := a.splashDone
		if !done {
			log.Info("Splash was closed, exiting the app")
			os.Exit(0)
		}
	})

	terms.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		if !config.GetTermsAccepted() {
			a.mu.Lock()
			done := a.splashDone
			a.mu.Unlock()
			if done {
				log.Info("Terms was closed, exiting the app")
				os.Exit(0)
			}
		}
	})

	main.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()
		e.Cancel()
		main.Hide()
		a.appState = Tray
		utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, config.GetLanguage()))
	})

	splash.OnWindowEvent(events.Common.WindowRuntimeReady, func(e *application.WindowEvent) {
		app.Event.Emit("update-text", map[string]string{
			"id":    "splash-alt",
			"value": localization.SplashLoading,
		})
	})

	app.Event.On("shutdown", func(e *application.CustomEvent) {
		os.Exit(0)
	})

	// im on hyprland, wails v3 not supports tray on linux, and also minimize not works for me, so i need to test it on win
	app.Event.On("close", func(e *application.CustomEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()
		e.Cancel()
		main.Hide()
		a.appState = Tray
		utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, config.GetLanguage()))
	})

	app.Event.On("minimize", func(e *application.CustomEvent) {
		e.Cancel()
		a.getWindow().Minimise()
		a.mu.Lock()
		a.appState = Minimized
		a.mu.Unlock()
	})

	app.Event.On("toggle-maximize", func(e *application.CustomEvent) {
		e.Cancel()
		switch a.appState {
		case Show:
			a.getWindow().Maximise()
			a.mu.Lock()
			a.appState = Maximized
			a.mu.Unlock()
		case Maximized:
			a.getWindow().UnMaximise()
			a.mu.Lock()
			a.appState = Show
			a.mu.Unlock()
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
		time.Sleep(100 * time.Millisecond)
		app.Event.Emit("navigate", "main")

		terms.Close()

		main.Show()
		main.Focus()
		a.setWindow(main)
	})

	go func() {
		time.Sleep(2 * time.Second)

		if !config.GetTermsAccepted() {
			log.Info("Terms is not accepted!")

			a.mu.Lock()
			a.splashDone = true
			a.mu.Unlock()
			splash.Close()

			time.Sleep(100 * time.Millisecond)
			app.Event.Emit("navigate", "terms")

			terms.Show()
			terms.Focus()
			a.setWindow(terms)

			return
		}

		a.mu.Lock()
		a.splashDone = true
		a.mu.Unlock()
		splash.Close()

		time.Sleep(100 * time.Millisecond)
		app.Event.Emit("navigate", "main")

		main.Show()
		main.Focus()
		a.setWindow(main)
	}()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
