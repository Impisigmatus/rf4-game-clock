package application

import (
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type Gui struct {
	app *Application

	guiApp fyne.App
	window fyne.Window

	notification *notification.Notification
}

func NewGui(id string, title string, width float32, height float32, notification *notification.Notification) *Gui {
	guiApp := app.NewWithID(id) // Указываем ID, чтобы prefs сохранялись
	window := guiApp.NewWindow(title)
	window.Resize(fyne.NewSize(width, height))

	return &Gui{
		app: NewApplication(),

		guiApp: guiApp,
		window: window,

		notification: notification,
	}
}

func (gui *Gui) Run() {
	gui.window.SetContent(gui.content())
	gui.window.ShowAndRun()
}

func (gui *Gui) content() *fyne.Container {
	// --- Создаем вкладки ---
	tabs := container.NewAppTabs(
		container.NewTabItem("Время в игре", gui.tabTime()),
		container.NewTabItem("Калькулятор времени", gui.tabCalculator()),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Весь UI: вкладки + тумблер снизу
	return container.NewBorder(nil, gui.themeToggle(), nil, nil, tabs)
}
