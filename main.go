package main

import (
	"embed"

	"github.com/Impisigmatus/rf4-game-clock/internal/application"
	"github.com/Impisigmatus/rf4-game-clock/internal/notification"
)

//go:embed assets/**
var assets embed.FS

func main() {
	icon, err := assets.Open("assets/icon.png")
	if err != nil {
		panic(err)
	}
	defer icon.Close()

	if err := notification.Notify("РР4 Игровое время", "Приложение запущено", icon); err != nil {
		panic(err)
	}

	gui := application.NewGui("rf4-game-clock", "РР4 Игровое время", 400, 250)
	gui.Run()
}
