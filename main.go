package main

import (
	"embed"
	"io"
	"os"

	"github.com/Impisigmatus/rf4-game-clock/internal/application"
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"
)

//go:embed assets/**
var assets embed.FS

func main() {
	const (
		title  = "РР4 Игровое время"
		width  = 400
		height = 250

		iconAssetPath = "assets/icon.png"
	)

	iconPath, err := loadIcon(iconAssetPath)
	if err != nil {
		panic(err)
	}

	notification := notification.NewNotification(title, iconPath)
	gui := application.NewGui("rf4-game-clock", title, width, height, notification)
	gui.Run()
}

func loadIcon(iconAssetPath string) (string, error) {
	file, err := assets.Open(iconAssetPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	tmp, err := os.CreateTemp("", "*.png")
	if err != nil {
		return "", err
	}
	defer tmp.Close()

	if _, err = io.Copy(tmp, file); err != nil {
		return "", err
	}

	return tmp.Name(), nil
}
