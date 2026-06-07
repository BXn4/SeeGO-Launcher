package main

import (
	"embed"
	"seegolauncher/internal/localization"
	"seegolauncher/internal/services"
	"seegolauncher/internal/utils"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

const (
	Show int = iota
	Minimized
)

type App struct {
	appState int
	view     string
	window   *application.WebviewWindow
	app      *application.App
	config   *utils.Config
}

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	log.Info("SeeGO Launcher by BXn4")
	config, err := utils.LoadConfig()
	if err != nil {
		log.Error("Failed to load config, using default.", "err", err)
	}

	a := &App{config: config}
	initStyles()

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "seego-launcher",
		Description: "Opensource alternative launcher for SeeRPG server",
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "1000",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				switch a.appState {
				case Minimized:
					{
						a.window.Show()
						a.window.UnMinimise()
						a.window.Focus()
						a.UpdateAppState()
						log.Info("Tried to open a new instance, but its already running. Showing the instance.")
					}
				}
			},
		},
		Services: []application.Service{
			// application.NewService(&GreetService{}),
			application.NewService(services.LocalizationService())},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	a.app = app

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
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

	a.window = splash

	main := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "SeeGO Launcher",
		Width:            1200,
		Height:           1200 / (16.0 / 9.0),
		MinWidth:         800,
		MinHeight:        400,
		Frameless:        true,
		AlwaysOnTop:      false,
		DisableResize:    false,
		BackgroundColour: application.NewRGB(52, 58, 64),
		URL:              "/",
		Hidden:           true,
	})

	main.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		if a.appState == Show {
			e.Cancel()
			main.Hide()
			a.appState = Minimized
			log.Info("The app is minimized")
			utils.Notify(services.LocalizationService().Get(localization.LauncherMinimized, a.config.Language))
		}
	})

	go func() {
		time.Sleep(2 * time.Second)
		main.Show()
		a.window = main
		splash.Close()
		app.Event.Emit("navigate", "main")
	}()

	// Run the application. This blocks until the application has been exited.
	err = app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) UpdateAppState() {
	switch a.appState {
	case Minimized:
		a.appState = Show
	default:
		a.appState = Minimized
	}
}

func initStyles() {
	styles := log.DefaultStyles()

	// Default
	s := lipgloss.NewStyle().Bold(true)
	styles.Levels[log.ErrorLevel] = s.Copy().SetString("ERROR").Foreground(lipgloss.Color("#ff0000"))
	styles.Levels[log.WarnLevel] = s.Copy().SetString("WARN").Foreground(lipgloss.Color("#ffff00"))
	styles.Levels[log.InfoLevel] = s.Copy().SetString("INFO").Foreground(lipgloss.Color("#33ffcc"))
	styles.Levels[log.DebugLevel] = s.Copy().SetString("DEBUG").Foreground(lipgloss.Color("#7e9edf"))

	// Custom
	styles.Levels[log.Level(-3)] = s.Copy().SetString("SENT").Foreground(lipgloss.Color("#ffffed"))
	styles.Levels[log.Level(-2)] = s.Copy().SetString("ANNOUNCE").Foreground(lipgloss.Color("#fff3db"))
	styles.Levels[log.Level(-1)] = s.Copy().SetString("BROADCAST").Foreground(lipgloss.Color("#fff3db"))
	styles.Levels[log.Level(-5)] = s.Copy().SetString("RECEIVED").Foreground(lipgloss.Color("#c8eec8"))

	log.SetStyles(styles)
}
