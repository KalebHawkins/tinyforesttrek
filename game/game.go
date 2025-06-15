package game

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/KalebHawkins/tinyforesttrek/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type GameState int

const (
	GameStatePlaying = iota
	GameStateWin
	GameStateLose
)

var AudioCtx = audio.NewContext(44100)

type Game struct {
	ScreenWidth, ScreenHeight int
	Player                    *Player
	TileMap                   *TileMap
	Camera                    *Camera
	GameObjects               []GameObject
	State                     GameState
	FontFaceSource            *text.GoTextFaceSource
	FontFace                  *text.GoTextFace
	FlashTimer                int
	TotalOrbs                 int
	PickupAudio               []byte
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

	var filtered []GameObject
	var orbsRemaining int
	for _, o := range g.GameObjects {
		o.Update(g.Player)

		// Check orb collision
		if orb, ok := o.(*Orb); ok {
			if orb.Collected && !orb.WasCollected {
				p := AudioCtx.NewPlayerFromBytes(g.PickupAudio)
				p.Play()
			}
			if ok && !orb.Collected {
				orbsRemaining++
			}
			orb.WasCollected = orb.Collected
		}

		// Check Thorn Collision
		thorn, ok := o.(*Thorn)
		if ok && thorn.IsDead() {
			g.State = GameStateLose
			g.FlashTimer = 30
		}

		if !o.IsDead() {
			filtered = append(filtered, o)
		}
	}

	if orbsRemaining == 0 {
		g.State = GameStateWin
	}

	if g.FlashTimer > 0 {
		g.FlashTimer--
	}

	g.GameObjects = filtered

	if g.State != GameStatePlaying && ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case GameStatePlaying:
		g.TileMap.Draw(screen, g.Camera)

		var orbsRemaining int
		for _, o := range g.GameObjects {
			o.Draw(screen, g.Camera)

			if _, ok := o.(*Orb); ok {
				orbsRemaining++
			}
		}

		collected := g.TotalOrbs - orbsRemaining
		top := &text.DrawOptions{}
		text.Draw(screen, fmt.Sprintf("Orbs: %d / %d", collected, g.TotalOrbs), g.FontFace, top)
		g.Player.Draw(screen, g.Camera)

	case GameStateWin:
		top := &text.DrawOptions{}
		top.LineSpacing = 16
		top.PrimaryAlign = text.AlignCenter
		top.GeoM.Translate(float64(g.ScreenWidth)/2, float64(g.ScreenHeight)/2)
		text.Draw(screen, "You Win\nPress R to Restart", g.FontFace, top)
	case GameStateLose:

		if g.FlashTimer > 0 && g.FlashTimer%5 == 0 {
			alpha := float64(g.FlashTimer) / 30.0
			overlay := ebiten.NewImage(g.ScreenWidth, g.ScreenHeight)
			overlay.Fill(color.RGBA{255, 0, 0, uint8(100 * alpha)})
			screen.DrawImage(overlay, nil)
			return
		}

		top := &text.DrawOptions{}
		top.LineSpacing = 16
		top.PrimaryAlign = text.AlignCenter
		top.GeoM.Translate(float64(g.ScreenWidth)/2, float64(g.ScreenHeight)/2)
		text.Draw(screen, "You Lose\nPress R to Restart", g.FontFace, top)
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Player:       NewPlayer(0, 0, 5.0, assets.LoadImage("player.png")),
		TileMap:      NewTileMap(40, 30, 64, tiles, assets.LoadImage("tiles.png")),
		PickupAudio:  assets.LoadAudio("orb_pickup.wav"),
	}

	g.Camera = NewCamera(
		screenWidth,
		screenHeight,
		g.TileMap.Width*g.TileMap.TileSize,
		g.TileMap.Height*g.TileMap.TileSize,
	)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	g.FontFaceSource = s
	g.FontFace = &text.GoTextFace{
		Source:    s,
		Direction: 0,
		Size:      12,
	}

	orb := NewOrb(100, 100, assets.LoadImage("orb.png"))
	orb2 := NewOrb(400, 350, assets.LoadImage("orb.png"))
	thorn := NewThorn(200, 100, assets.LoadImage("thorn.png"))
	g.GameObjects = append(g.GameObjects, orb, orb2, thorn)

	for _, o := range g.GameObjects {
		if _, ok := o.(*Orb); ok {
			g.TotalOrbs++
		}
	}

	return g
}

func (g *Game) Reset() {
	*g = *NewGame(g.ScreenWidth, g.ScreenHeight)
}
