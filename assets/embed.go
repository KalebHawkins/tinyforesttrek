package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var Assets embed.FS

func LoadImage(file string) *ebiten.Image {
	f, err := Assets.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func LoadAudio(file string) []byte {
	f, err := Assets.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
