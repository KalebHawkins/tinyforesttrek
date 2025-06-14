package main

import (
	"log"

	"github.com/KalebHawkins/tinyforesttrek/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Tiny Forest Trek")

	if err := ebiten.RunGame(game.NewGame(ScreenWidth, ScreenHeight)); err != nil {
		log.Fatal(err)
	}
}
