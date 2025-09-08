package main

import (
	"embed"
	"io"

	"github.com/Impisigmatus/rf4-game-clock/internal/application"
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"
	"github.com/gen2brain/beeep"
)

//go:embed assets/**
var assets embed.FS

func main() {
	const (
		appName = "rf4-game-clock"
		title   = "РР4 Игровое время"
		width   = 400
		height  = 250

		iconPath = "assets/icon.png"
	)
	beeep.AppName = appName

	icon, err := loadIcon(iconPath)
	if err != nil {
		panic(err)
	}

	notification := notification.NewNotification(title, icon)
	gui := application.NewGuiApplication(appName, title, width, height, notification)
	gui.Run()
}

func loadIcon(path string) ([]byte, error) {
	file, err := assets.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
