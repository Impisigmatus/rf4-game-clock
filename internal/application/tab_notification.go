package application

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func (gui *GuiApplication) tabNotification() *fyne.Container {
	// Создаем выборку часов
	hoursOptions := make([]string, 24)
	for i := 0; i <= 23; i++ {
		hoursOptions[i] = fmt.Sprintf("%02d", i)
	}
	hoursSelect := widget.NewSelect(hoursOptions, nil)
	hoursSelect.PlaceHolder = "HH"
	hoursSelect.Selected = "00"

	// Создаем выборку минут
	minutesOptions := make([]string, 2)
	for i := 0; i < 60; i += 30 {
		minutesOptions[i/30] = fmt.Sprintf("%02d", i)
	}
	minutesSelect := widget.NewSelect(minutesOptions, nil)
	minutesSelect.PlaceHolder = "MM"
	minutesSelect.Selected = "00"

	// Создаем текстовое поле
	result := widget.NewLabel("")

	// Создаем кнопку
	var notifyButton *widget.Button
	notifyButton = widget.NewButton("Уведомить", func() {
		if notifyButton.Importance == widget.LowImportance {
			return
		}

		notifyButton.Importance = widget.LowImportance
		notifyButton.Refresh()

		hours, minutes, err := gui.app.ConvertTime(hoursSelect.Selected, minutesSelect.Selected)
		if err != nil {
			return
		}
		gui.notification.Alert("Не забудьте про рыбалку в %02d:%02d", hours, minutes)

		if gui.timer != nil {
			gui.timer.Stop()
		}
		gui.timer = time.AfterFunc(gui.app.GetDurationToTarget(hours, minutes), func() {
			gui.notification.Alert("Пора на рыбалку!")
		})
	})
	notifyButton.Importance = widget.HighImportance

	// Создаем переключатель
	notifyHourlyCheck := widget.NewCheck("Уведомлять каждый час", func(checked bool) {
		if checked {
			notifyButton.Importance = widget.LowImportance // Серый цвет
		} else {
			notifyButton.Importance = widget.HighImportance // Обычный/акцентный цвет
		}
		notifyButton.Refresh()

		logrus.Infof("TODO: Уведомлять каждый час %t", checked)
	})

	setTime := func() {
		if notifyButton.Importance == widget.LowImportance {
			notifyButton.Importance = widget.HighImportance
			notifyButton.Refresh()

			notifyHourlyCheck.SetChecked(false)
			notifyHourlyCheck.Refresh()
		}

		hours, minutes, err := gui.app.ConvertTime(hoursSelect.Selected, minutesSelect.Selected)
		if err != nil {
			result.SetText(err.Error())
			return
		}

		result.SetText(fmt.Sprintf("Это игровое время наступит в %02d:%02d", hours, minutes))
	}

	hoursSelect.OnChanged = func(_ string) { setTime() }
	minutesSelect.OnChanged = func(_ string) { setTime() }

	content := container.NewVBox(
		widget.NewLabelWithStyle("Оповещения о рыбалке", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(hoursSelect, minutesSelect),
		result,
		notifyHourlyCheck,
		notifyButton,
		layout.NewSpacer(),
	)
	return content
}
