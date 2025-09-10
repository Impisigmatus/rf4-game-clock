package application

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (gui *GuiApplication) tabClock() *fyne.Container {
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
			fyne.Do(updateTime)
		}
	}()
	updateTime()

	content := container.NewVBox(
		widget.NewLabelWithStyle("🕓 Время в игре", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		realTime,
		gameTime,
		layout.NewSpacer(),
	)

	return content
}
