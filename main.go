package main

import (
	"asteroids/internal/asteroids"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

func main() {
	g := asteroids.NewGame(ScreenWidth, ScreenHeight)

	ebiten.SetWindowTitle("Asteriods")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
