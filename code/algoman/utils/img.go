package utils

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"bytes"
	"log"
)

func ImgFromByte(imgByte []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return ebitenImg
}
