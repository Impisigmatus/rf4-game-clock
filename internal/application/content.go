package application

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (gui *Gui) tab1() *fyne.Container {
	// --- Вкладка 1: Текущее время ---
	realTime := widget.NewLabel("")
	gameTime := widget.NewLabel("")

	updateTime := func() {
		now := time.Now()
		hours := now.Hour()
		minutes := now.Minute()
		gameHours, gameMinutes := gui.app.CurrentGameTime()

		realTime.SetText(fmt.Sprintf("Реальное время: %02d:%02d", hours, minutes))
		gameTime.SetText(fmt.Sprintf("Игровое время:  %02d:%02d", gameHours, gameMinutes))
	}
	// --- Таймер обновления ---
	go func() {
		for range time.Tick(time.Second) {
			updateTime()
		}
	}()
	updateTime()

	tab1Content := container.NewVBox(
		widget.NewLabelWithStyle("🕓 Время в игре", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		realTime,
		gameTime,
		layout.NewSpacer(),
	)

	return tab1Content
}

func (gui *Gui) tab2() *fyne.Container {
	// --- Вкладка 2: Калькулятор времени ---
	hoursOptions := make([]string, 24)
	for i := 0; i <= 23; i++ {
		hoursOptions[i] = fmt.Sprintf("%02d", i)
	}
	minutesOptions := make([]string, 2)
	for i := 0; i < 60; i += 30 {
		minutesOptions[i/30] = fmt.Sprintf("%02d", i)
	}

	hoursSelect := widget.NewSelect(hoursOptions, nil)
	hoursSelect.PlaceHolder = "HH"
	hoursSelect.Selected = "00"

	minutesSelect := widget.NewSelect(minutesOptions, nil)
	minutesSelect.PlaceHolder = "MM"
	minutesSelect.Selected = "00"

	result := widget.NewLabel("")

	setTime := func() {
		hour, minute, err := gui.app.ConvertTime(hoursSelect.Selected, minutesSelect.Selected)
		if err != nil {
			result.SetText(err.Error())
			return
		}
		result.SetText(fmt.Sprintf("Это игровое время наступит в %02d:%02d", hour, minute))
	}

	hoursSelect.OnChanged = func(_ string) { setTime() }
	minutesSelect.OnChanged = func(_ string) { setTime() }

	tab2Content := container.NewVBox(
		widget.NewLabelWithStyle("🧮 Калькулятор игрового времени", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(hoursSelect, minutesSelect),
		result,
		layout.NewSpacer(),
	)
	return tab2Content
}

func (gui *Gui) themeToggle() *fyne.Container {
	// --- Настройки темы ---
	prefs := gui.guiApp.Preferences()
	themePref := prefs.StringWithFallback("theme", "light")

	if themePref == "dark" {
		gui.guiApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		gui.guiApp.Settings().SetTheme(theme.LightTheme())
	}

	// --- Тумблер темы ---
	themeToggle := widget.NewCheck("Тёмная тема", func(checked bool) {
		if checked {
			gui.guiApp.Settings().SetTheme(theme.DarkTheme())
			prefs.SetString("theme", "dark") // сохраняем выбор
		} else {
			gui.guiApp.Settings().SetTheme(theme.LightTheme())
			prefs.SetString("theme", "light") // сохраняем выбор
		}
	})
	themeToggle.SetChecked(themePref == "dark")

	return container.NewHBox(themeToggle)
}
