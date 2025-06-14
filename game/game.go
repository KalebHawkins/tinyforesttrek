package game

import (
	"github.com/KalebHawkins/tinyforesttrek/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenWidth, ScreenHeight int
	Player                    *Player
	TileMap                   *TileMap
	Camera                    *Camera
}

func (g *Game) Update() error {
	g.Camera.Follow(g.Player.X, g.Player.Y)

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
	g.TileMap.Draw(screen, g.Camera)
	g.Player.Draw(screen, g.Camera)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	g.Player = NewPlayer(0, 0, 5.0)
	g.TileMap = NewTileMap(40, 30, 64, tiles, assets.Load("tiles.png"))
	g.Camera = NewCamera(
		screenWidth,
		screenHeight,
		g.TileMap.Width*g.TileMap.TileSize,
		g.TileMap.Height*g.TileMap.TileSize,
	)

	return g
}
