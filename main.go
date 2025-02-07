package main

import (
	"log"

	goflappy "github.com/chugunov/go-flappy/goflappy"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenHeight = 640
	screenWidth  = 480
)

func main() {
	game, title := goflappy.NewGame(screenHeight, screenWidth)

	ebiten.SetWindowSize(screenHeight, screenWidth)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
