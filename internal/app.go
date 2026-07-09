package internal

import (
	"embed"
	"fmt"
	"seegolauncher/internal/localization"
	"seegolauncher/internal/services"
	"seegolauncher/internal/utils"
	"time"

	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

const (
	Tray int = iota
	Show
	Minimized
	Maximized
)

type App struct {
	app          *application.App
	tray         *application.SystemTray
	config       *services.Config
	appState     int
	window       *application.WebviewWindow
	view         string
	dialog       *application.MessageDialog
	trayNotified bool
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) Quit() {
	log.Info("App quit called, quiting")
	a.app.Quit()
}

func (a *App) OnReady() {
	a.window.Show()

	a.EmitEvent("splash:setCurrentProgress", "splash-loading")

	go func() {
		if utils.WaitUntil(
			func() bool {
				success := services.LoadCache()
				return success == nil
			},
			5*time.Second,
		) {
			// will notify, if the cache was not successed
			a.Quit()
		}

		// after splash
		a.window.SetAlwaysOnTop(false)

		if !a.config.GetTermsAccepted() {
			seeGOInfo1 := services.LocalizationService().Get(localization.SeeGOInfo1, a.config.GetLanguage())
			seeGOInfo2 := services.LocalizationService().Get(localization.SeeGOInfo2, a.config.GetLanguage())
			seeGOInfo3 := services.LocalizationService().Get(localization.SeeGOInfo3, a.config.GetLanguage())
			seeGOInfo4 := services.LocalizationService().Get(localization.SeeGOInfo4, a.config.GetLanguage())
			a.app.Dialog.Info().SetTitle("SeeGO Launcher").SetMessage(fmt.Sprintf("%s\n%s\n%s\n%s", seeGOInfo1, seeGOInfo2, seeGOInfo3, seeGOInfo4)).Show()
			log.Info("Terms is not accepted, showing the terms window")

			a.SetView(620, 480, AppNavigate, "terms")
			a.window.Center()
			return
		}

		a.SetView(1200, int(1200/(16.0/9.0)), AppNavigate, "main")
		a.window.Center()
	}()
}

func (a *App) CreateNewWindow(options application.WebviewWindowOptions) *application.WebviewWindow {
	window := a.app.Window.NewWithOptions(options)

	window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		if a.view != "main" {
			log.Info("Application closed on not main view")
			e.Cancel()
			a.Quit()
			return
		}

		e.Cancel()
		a.ToTray()
	})

	window.OnWindowEvent(events.Common.WindowRestore, func(e *application.WindowEvent) {
		a.window.Show()
		a.EmitEvent("app:active", nil)
	})

	return window
}

func (a *App) ToTray() {
	a.window.Hide()
	a.appState = Tray
	if !a.trayNotified {
		a.trayNotified = true
		utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, a.config.GetLanguage()))
	}
}

func (a *App) RestoreFromTray() {
	a.window.UnMinimise()
	a.window.Show()
	a.window.Focus()
	a.appState = Show

	a.EmitEvent("app:active", nil)
}

func (a *App) Minimize() {
	a.window.Minimise()
	a.appState = Minimized

	a.EmitEvent("app:notActive", nil)
}

func (a *App) SetView(w, h int, n, v string) {
	a.window.SetSize(w, h)
	a.app.Event.Emit(n, v)
	a.view = v
}

func CreateApp(assets embed.FS, icon []byte) App {
	app := App{appState: Show, dialog: nil}
	app.config = services.ConfigService()

	app.app = application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				log.Info("Tried to open a new instance, but its already running. Showing the instance.")
				switch app.appState {
				case Tray:
					app.RestoreFromTray()
				default:
					app.window.Focus()
					app.window.Flash(true)
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

	app.RegisterListeners()

	log.Info("Created the new app!")

	app.window = app.CreateNewWindow(application.WebviewWindowOptions{
		Title:                      "SeeGO Launcher",
		Width:                      476,
		Height:                     300,
		Frameless:                  true,
		AlwaysOnTop:                true,
		DisableResize:              true,
		BackgroundColour:           application.NewRGB(52, 58, 64),
		URL:                        "/",
		Hidden:                     true,
		DevToolsEnabled:            false,
		DefaultContextMenuDisabled: true,
	})

	app.CreateSysTray(icon)

	return app
}

func (a *App) CreateSysTray(icon []byte) {
	systray := a.app.SystemTray.New()
	systray.SetLabel("SeeGO Launcher")
	systray.SetIcon(icon)
	systray.OnDoubleClick(func() {
		switch a.appState {
		case Tray:
			a.RestoreFromTray()
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
			a.RestoreFromTray()
		default:
			a.window.Focus()
			a.window.Flash(true)
		}
	})

	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		a.app.Quit()
	})
	systray.SetMenu(menu)

	log.Info("Created the Systray!")
}

func (a *App) RegisterListeners() {
	a.app.Event.On(AppReady, func(e *application.CustomEvent) {
		log.Info("App is ready to show!")
		a.OnReady()
	})

	a.app.Event.On(TermsAccept, func(e *application.CustomEvent) {
		a.config.SetTermsAccepted()
		a.SetView(1200, int(1200/(16.0/9.0)), AppNavigate, "main")
		a.window.Center()
	})

	a.app.Event.On(TermsDecline, func(e *application.CustomEvent) {
		a.window.SetIgnoreMouseEvents(true)
		if a.dialog == nil {
			dialog := a.app.Dialog.Warning().
				SetTitle(services.LocalizationService().Get(localization.TermsDeclinedTitle, a.config.GetLanguage())).
				SetMessage(services.LocalizationService().Get(localization.TermsDeclinedContent, a.config.GetLanguage()))

			ok := dialog.AddButton("Ok")
			ok.OnClick(
				func() {
					a.Quit()
				})

			dialog.SetDefaultButton(ok)

			a.dialog = dialog
		}

		a.dialog.Show()
	})

	a.app.Event.On(FeedbackEvent, func(e *application.CustomEvent) {
		log.Info(e.Data)
	})

	a.app.Event.On(AppMinimize, func(e *application.CustomEvent) {
		a.Minimize()
	})

	a.app.Event.On(AppClose, func(e *application.CustomEvent) {
		a.ToTray()
	})

}

func (a *App) EmitEvent(name string, data any) {
	a.window.EmitEvent(name, data)
}
