package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	ScreenWidth, ScreenHeight int
	Player                    *Player
}

func (g *Game) Update() error {
	var dx, dy float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dy -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dx -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		dy += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		dx += 1
	}

	g.Player.Move(dx, dy)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	g.Player = NewPlayer(float64(g.ScreenWidth)/2, float64(screenHeight)/2, 5.0)

	return g
}
