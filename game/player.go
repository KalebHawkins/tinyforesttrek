package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*ebiten.Image
	X, Y      float64
	MoveSpeed float64
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.MoveSpeed
	}

	return nil
}

func (p *Player) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	dst.DrawImage(p.Image, op)
}

func NewPlayer(x, y, moveSpeed float64) *Player {
	p := &Player{
		Image:     ebiten.NewImage(16, 16),
		X:         x,
		Y:         y,
		MoveSpeed: moveSpeed,
	}

	p.Image.Fill(color.White)

	return p
}
