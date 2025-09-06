package notification

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

type Notification struct {
	Title    string
	IconPath string
}

func NewNotification(title string, iconPath string) *Notification {
	return &Notification{
		Title:    title,
		IconPath: iconPath,
	}
}

func (notification *Notification) Notify(format string, a ...any) {
	_ = beeep.Notify(notification.Title, fmt.Sprintf(format, a...), notification.IconPath)
}

func (notification *Notification) Alert(format string, a ...any) {
	_ = beeep.Alert(notification.Title, fmt.Sprintf(format, a...), notification.IconPath)
}
