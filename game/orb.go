package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Orb struct {
	X, Y         float64
	Sprite       *ebiten.Image
	Collected    bool
	WasCollected bool
	Frame        int
}

func (o *Orb) Update(p *Player) {
	if o.Collected {
		o.Frame++
	}

	if o.CheckCollision(p) {
		o.Collected = true
	}
}

func (o *Orb) Draw(dst *ebiten.Image, cam *Camera) {
	if o.Collected && o.Frame > 30 {
		return
	}

	scale := 1.0 - float64(o.Frame)/30
	alpha := 1.0 - float32(o.Frame)/30

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cam.ZoomFactor, cam.ZoomFactor)
	op.GeoM.Scale(scale, scale)
	op.ColorScale.ScaleAlpha(alpha)
	op.GeoM.Translate(-float64(o.Sprite.Bounds().Dx())/2, -float64(o.Sprite.Bounds().Dy())/2)
	op.GeoM.Translate(o.X-cam.X, o.Y-cam.Y)
	dst.DrawImage(o.Sprite, op)
}

func (o *Orb) CheckCollision(p *Player) bool {
	dist := math.Hypot(p.X-o.X, p.Y-o.Y)
	return dist < 16
}

func (o *Orb) IsDead() bool {
	return o.Collected && o.Frame > 30
}

func NewOrb(x, y float64, sprite *ebiten.Image) *Orb {
	o := &Orb{
		X:         x,
		Y:         y,
		Sprite:    sprite,
		Collected: false,
	}

	return o
}
