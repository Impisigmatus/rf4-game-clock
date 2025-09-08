package notification

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

type Notification struct {
	Title string
	Icon  []byte
}

func NewNotification(title string, icon []byte) *Notification {
	return &Notification{
		Title: title,
		Icon:  icon,
	}
}

func (notification *Notification) Notify(format string, a ...any) {
	_ = beeep.Notify(notification.Title, fmt.Sprintf(format, a...), notification.Icon)
}

func (notification *Notification) Alert(format string, a ...any) {
	_ = beeep.Alert(notification.Title, fmt.Sprintf(format, a...), notification.Icon)
}
