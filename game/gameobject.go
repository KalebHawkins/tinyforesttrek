package game

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Update(p *Player)
	Draw(dst *ebiten.Image, cam *Camera)
	CheckCollision(p *Player) bool
	IsDead() bool
}
