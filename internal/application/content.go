package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// --- Настройки темы ---
func (gui *GuiApplication) themeToggle() *fyne.Container {
	prefs := gui.guiApp.Preferences()
	themePref := prefs.StringWithFallback("theme", "light")
	if themePref == "dark" {
		gui.guiApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		gui.guiApp.Settings().SetTheme(theme.LightTheme())
	}

	themeToggle := widget.NewCheck("Тёмная тема", func(checked bool) {
		if checked {
			gui.guiApp.Settings().SetTheme(theme.DarkTheme())
			prefs.SetString("theme", "dark")
		} else {
			gui.guiApp.Settings().SetTheme(theme.LightTheme())
			prefs.SetString("theme", "light")
		}
	})
	themeToggle.SetChecked(themePref == "dark")

	return container.NewHBox(themeToggle)
}
