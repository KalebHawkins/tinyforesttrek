package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Sprite    *ebiten.Image
	X, Y      float64
	MoveSpeed float64
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) Move(dx, dy float64) {
	if dx != 0 || dy != 0 {
		l := math.Hypot(dx, dy)
		dx /= l
		dy /= l
	}

	p.X += dx * p.MoveSpeed
	p.Y += dy * p.MoveSpeed
}

func (p *Player) Draw(dst *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(p.Sprite.Bounds().Dx())/2, -float64(p.Sprite.Bounds().Dy())/2)
	op.GeoM.Scale(camera.ZoomFactor, camera.ZoomFactor)
	op.GeoM.Translate(p.X-camera.X, p.Y-camera.Y)
	dst.DrawImage(p.Sprite, op)
}

func NewPlayer(x, y, moveSpeed float64, sprite *ebiten.Image) *Player {
	p := &Player{
		Sprite:    sprite,
		X:         x,
		Y:         y,
		MoveSpeed: moveSpeed,
	}

	return p
}
