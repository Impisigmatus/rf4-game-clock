package application

import (
	"fmt"
	"strconv"
	"time"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (*Application) CurrentGameTime() (int, int) {
	now := time.Now()
	minutes := now.Minute()
	seconds := now.Second()

	gameTotalMinutes := float64(minutes)*24 + float64(seconds)*24/60
	return int(gameTotalMinutes) / 60, int(gameTotalMinutes) % 60
}

func (*Application) ConvertTime(hours string, minutes string) (int, int, error) {
	gameHours, err := strconv.Atoi(hours)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid hours value: %s", hours)
	}

	gameMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid minutes value: %s", minutes)
	}

	totalGameMinutes := gameHours*60 + gameMinutes
	totalRealMinutes := float64(totalGameMinutes) * (2.5 / 60)

	now := time.Now()
	if now.Minute() > int(totalRealMinutes) {
		now = now.Add(time.Hour)
	}

	return now.Hour(), int(totalRealMinutes), nil
}
