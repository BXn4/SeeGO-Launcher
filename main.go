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
	Show int = iota
	Minimized
)

type App struct {
	mu         sync.Mutex
	appState   int
	window     *application.WebviewWindow
	splashDone bool
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

	a := &App{appState: Show}

	app := application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				log.Info("Tried to open a new instance, but its already running. Showing the instance.")
				w := a.getWindow()
				a.mu.Lock()
				state := a.appState
				a.mu.Unlock()
				switch state {
				case Minimized:
					w.UnMinimise()
					w.Show()
					w.Focus()
					a.mu.Lock()
					a.appState = Show
					a.mu.Unlock()
				case Show:
					w.Focus()
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

	splash := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "SeeGO Launcher",
		Width:            476,
		Height:           300,
		Frameless:        true,
		AlwaysOnTop:      true,
		DisableResize:    true,
		BackgroundColour: application.NewRGB(52, 58, 64),
		URL:              "/",
		Hidden:           false,
	})

	main := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "SeeGO Launcher",
		Width:            1200,
		Height:           int(1200 / (16.0 / 9.0)),
		Frameless:        true,
		DisableResize:    false,
		BackgroundColour: application.NewRGB(52, 58, 64),
		URL:              "/",
		Hidden:           true,
	})

	terms := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "SeeGO Launcher",
		Width:            800,
		Height:           600,
		Frameless:        true,
		AlwaysOnTop:      true,
		DisableResize:    true,
		BackgroundColour: application.NewRGB(52, 58, 64),
		URL:              "/",
		Hidden:           true,
	})

	a.setWindow(splash)

	splash.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		a.mu.Lock()
		done := a.splashDone
		a.mu.Unlock()
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
		e.Cancel()
		main.Hide()
		a.mu.Lock()
		a.appState = Minimized
		a.mu.Unlock()
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
		e.Cancel()
		main.Hide()
		a.mu.Lock()
		a.appState = Minimized
		a.mu.Unlock()
		utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, config.GetLanguage()))
	})

	app.Event.On("minimize", func(e *application.CustomEvent) {
		e.Cancel()
		a.getWindow().Minimise()
		a.mu.Lock()
		a.appState = Minimized
		a.mu.Unlock()
	})

	app.Event.On("terms-declined", func(e *application.CustomEvent) {
		dialog := app.Dialog.Warning().
			SetTitle(services.LocalizationService().Get(localization.TermsDeclinedTitle, config.GetLanguage())).
			SetMessage(services.LocalizationService().Get(localization.TermsDeclinedContent, config.GetLanguage()))

		ok := dialog.AddButton("Ok")
		ok.OnClick(
			func() {
				os.Exit(0)
			})

		dialog.SetDefaultButton(ok)

		dialog.Show()
	})

	app.Event.On("terms-accepted", func(e *application.CustomEvent) {
		config.SetTermsAccepted()
		main.Show()
		main.Focus()
		a.setWindow(main)

		time.Sleep(100 * time.Millisecond)
		app.Event.Emit("navigate", "main")

		terms.Close()
	})

	go func() {
		time.Sleep(2 * time.Second)

		if !config.GetTermsAccepted() {
			log.Info("Terms is not accepted!")

			terms.Show()
			terms.Focus()
			a.setWindow(terms)

			time.Sleep(100 * time.Millisecond)
			app.Event.Emit("navigate", "terms")

			a.mu.Lock()
			a.splashDone = true
			a.mu.Unlock()
			splash.Close()
			return
		}

		main.Show()
		main.Focus()
		a.setWindow(main)

		time.Sleep(100 * time.Millisecond)
		app.Event.Emit("navigate", "main")

		a.mu.Lock()
		a.splashDone = true
		a.mu.Unlock()
		splash.Close()
	}()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
