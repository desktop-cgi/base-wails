package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	// "github.com/wailsapp/wails/v2/pkg/options/linux"
    // "github.com/wailsapp/wails/v2/pkg/options/mac"
    // "github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	// https://wails.io/docs/reference/options
	err := wails.Run(&options.App{
		Title:  "Desktop-CGI",
		Width:  1280,
		Height: 1024,
		DisableResize:      false,
        Fullscreen:         true,
        Frameless:          false,
        MinWidth:           400,
        MinHeight:          400,
        MaxWidth:           1280,
        MaxHeight:          1024,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		
		// Logger:             nil,
        // LogLevel:           logger.DEBUG,
        // LogLevelProduction: logger.ERROR,
        // OnStartup:          app.startup,
		OnStartup:        app.startup,
        // OnDomReady:         app.domready,
        // OnShutdown:         app.shutdown,
        // OnBeforeClose:      app.beforeClose,
        // WindowStartState:   options.Maximised,
        // CSSDragProperty:   "--wails-draggable",
        // CSSDragValue:      "drag",
        // ZoomFactor:           1.0,
        // IsZoomControlEnabled: false,
		Bind: []interface{}{
			app,
		},
		// Windows: &windows.Options{
        //     WebviewIsTransparent:              false,
        //     WindowIsTranslucent:               false,
        //     BackdropType:                      windows.Mica,
        //     DisableWindowIcon:                 false,
        //     DisableFramelessWindowDecorations: false,
        //     WebviewUserDataPath:               "",
        //     WebviewBrowserPath:                "",
        //     Theme:                             windows.SystemDefault,
        //     CustomTheme: &windows.ThemeSettings{
        //         DarkModeTitleBar:   windows.RGB(20, 20, 20),
        //         DarkModeTitleText:  windows.RGB(200, 200, 200),
        //         DarkModeBorder:     windows.RGB(20, 0, 20),
        //         LightModeTitleBar:  windows.RGB(200, 200, 200),
        //         LightModeTitleText: windows.RGB(20, 20, 20),
        //         LightModeBorder:    windows.RGB(200, 200, 200),
        //     },
        //     // User messages that can be customised
        //     Messages *windows.Messages
        //     // OnSuspend is called when Windows enters low power mode
        //     OnSuspend func()
        //     // OnResume is called when Windows resumes from low power mode
        //     OnResume func()
        // },
        // Mac: &mac.Options{
        //     TitleBar: &mac.TitleBar{
        //         TitlebarAppearsTransparent: true,
        //         HideTitle:                  false,
        //         HideTitleBar:               false,
        //         FullSizeContent:            false,
        //         UseToolbar:                 false,
        //         HideToolbarSeparator:       true,
        //     },
        //     Appearance:           mac.NSAppearanceNameDarkAqua,
        //     WebviewIsTransparent: true,
        //     WindowIsTranslucent:  false,
        //     About: &mac.AboutInfo{
        //         Title:   "My Application",
        //         Message: "© 2021 Me",
        //         Icon:    icon,
        //     },
        // },
        // Linux: &linux.Options{
        //     Icon: icon,
        //     WindowIsTranslucent: false,
        // },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
