package application

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func (gui *GuiApplication) tabNotification() *fyne.Container {
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
		gui.notification.Alert("Не забудьте про рыбалку в %02d:%02d", hour, minute)
	}

	hoursSelect.OnChanged = func(_ string) { setTime() }
	minutesSelect.OnChanged = func(_ string) { setTime() }

	var buttonDisabled bool
	// Создаем кнопку
	notifyButton := widget.NewButton("Уведомить", func() {
		if buttonDisabled {
			return
		}

		logrus.Info("TODO: Уведомить")
	})
	notifyButton.Importance = widget.HighImportance

	// Создаем переключатель
	notifyHourlyCheck := widget.NewCheck("Уведомлять каждый час", func(checked bool) {
		buttonDisabled = checked
		if checked {
			notifyButton.Importance = widget.LowImportance // Серый цвет
		} else {
			notifyButton.Importance = widget.HighImportance // Обычный/акцентный цвет
		}

		logrus.Infof("TODO: Уведомлять каждый час %t", checked)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle("Оповещения о рыбалке", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(hoursSelect, minutesSelect),
		notifyHourlyCheck,
		notifyButton,
		result,
		layout.NewSpacer(),
	)
	return content
}
