package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/Impisigmatus/rf4-game-clock/internal/application"
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"
	"github.com/gen2brain/beeep"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})
}

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
		logrus.Panicf("Invalid loading notification icon: %s", err)
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
