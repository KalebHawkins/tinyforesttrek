package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Thorn struct {
	X, Y   float64
	Sprite *ebiten.Image
	Dead   bool
}

func (t *Thorn) Update(p *Player) {
	if t.CheckCollision(p) {
		t.Dead = true
	}
}

func (t *Thorn) Draw(dst *ebiten.Image, cam *Camera) {
	if t.Dead {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cam.ZoomFactor, cam.ZoomFactor)
	op.GeoM.Translate(-float64(t.Sprite.Bounds().Dx())/2, -float64(t.Sprite.Bounds().Dy())/2)
	op.GeoM.Translate(t.X-cam.X, t.Y-cam.Y)
	dst.DrawImage(t.Sprite, op)
}

func (t *Thorn) CheckCollision(p *Player) bool {
	dist := math.Hypot(p.X-t.X, p.Y-t.Y)
	return dist < 16
}

func (t *Thorn) IsDead() bool {
	return t.Dead
}

func NewThorn(x, y float64, sprite *ebiten.Image) *Thorn {
	t := &Thorn{
		X:      x,
		Y:      y,
		Sprite: sprite,
		Dead:   false,
	}

	return t
}
