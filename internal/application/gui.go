package application

import (
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"
	"github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type GuiApplication struct {
	app *Application

	guiApp fyne.App
	window fyne.Window

	notification *notification.Notification
}

func NewGuiApplication(id string, title string, width float32, height float32, notification *notification.Notification) *GuiApplication {
	guiApp := app.NewWithID(id) // Указываем ID, чтобы prefs сохранялись
	window := guiApp.NewWindow(title)
	window.Resize(fyne.NewSize(width, height))

	return &GuiApplication{
		app: NewApplication(),

		guiApp: guiApp,
		window: window,

		notification: notification,
	}
}

func (gui *GuiApplication) Run() {
	gui.window.SetContent(gui.content())
	logrus.Info("Starting application")
	gui.window.ShowAndRun()
}

func (gui *GuiApplication) content() *fyne.Container {
	// --- Создаем вкладки ---
	tabs := container.NewAppTabs(
		container.NewTabItem("Игровое время", gui.tabClock()),
		container.NewTabItem("Оповещения", gui.tabNotification()),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Весь UI: вкладки + тумблер снизу
	return container.NewBorder(nil, gui.themeToggle(), nil, nil, tabs)
}
