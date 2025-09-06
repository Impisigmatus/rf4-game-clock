package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type Gui struct {
	app *Application

	guiApp fyne.App
	window fyne.Window
}

func NewGui(id string, title string, width float32, height float32) *Gui {
	guiApp := app.NewWithID(id) // Указываем ID, чтобы prefs сохранялись
	window := guiApp.NewWindow(title)
	window.Resize(fyne.NewSize(width, height))

	return &Gui{
		app: NewApplication(),

		guiApp: guiApp,
		window: window,
	}
}

func (gui *Gui) Run() {
	gui.window.SetContent(gui.content())
	gui.window.ShowAndRun()
}

func (gui *Gui) content() *fyne.Container {
	// --- Создаем вкладки ---
	tabs := container.NewAppTabs(
		container.NewTabItem("Время в игре", gui.tab1()),
		container.NewTabItem("Калькулятор", gui.tab2()),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Весь UI: вкладки + тумблер снизу
	return container.NewBorder(nil, gui.themeToggle(), nil, nil, tabs)
}
