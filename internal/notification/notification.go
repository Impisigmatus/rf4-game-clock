package notification

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type Notification struct {
	Title string
	Icon  []byte

	timer *time.Timer
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

func (notification *Notification) NotifyWithDuration(duration time.Duration, format string, a ...any) {
	if notification.timer != nil {
		notification.timer.Stop()
	}
	notification.timer = time.AfterFunc(duration, func() {
		notification.Notify(format, a...)
	})
}

func (notification *Notification) AlertWithDuration(duration time.Duration, format string, a ...any) {
	if notification.timer != nil {
		notification.timer.Stop()
	}
	notification.timer = time.AfterFunc(duration, func() {
		notification.Alert(format, a...)
	})
}
