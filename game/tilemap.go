package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileMap struct {
	Width, Height int
	Tiles         []int
	TileSize      int
	TileImage     *ebiten.Image
}

func (t *TileMap) Draw(dst *ebiten.Image) {
	tilesPerRow := t.TileImage.Bounds().Dx() / t.TileSize

	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {

			tileId := t.Tiles[y*t.Width+x]
			if tileId < 0 {
				continue
			}

			sx := tileId % tilesPerRow * t.TileSize
			sy := tileId / tilesPerRow * t.TileSize
			src := t.TileImage.SubImage(image.Rect(sx, sy, sx+t.TileSize, sy+t.TileSize)).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*t.TileSize), float64(y*t.TileSize))
			dst.DrawImage(src, op)
		}
	}
}

func NewTileMap(width, height, tileSize int, tiles []int, tileImage *ebiten.Image) *TileMap {
	tm := &TileMap{
		Width:     width,
		Height:    height,
		Tiles:     tiles,
		TileSize:  tileSize,
		TileImage: tileImage,
	}

	return tm
}
