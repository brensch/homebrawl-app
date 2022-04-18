package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Button struct {
	x      float64
	y      float64
	height int
	width  int

	view view

	image *ebiten.Image

	normalImage      *ebiten.Image
	highlightedImage *ebiten.Image
	opts             *ebiten.DrawImageOptions
}

func InitButton(x, y float64, width, height int, buttonText string, bgColor color.Color, view view) *Button {
	button := &Button{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		view:   view,
	}

	button.normalImage = ebiten.NewImage(width, height)
	button.normalImage.Fill(bgColor)
	text.Draw(button.normalImage, buttonText, mplusLargeFont, 100, mplusLargeFont.Metrics().Height.Ceil()+button.height/2, color.Black)

	button.highlightedImage = ebiten.NewImage(width, height)
	button.highlightedImage.Fill(bgColor)
	text.Draw(button.highlightedImage, buttonText, mplusLargeFont, 100, mplusLargeFont.Metrics().Height.Ceil()+button.height/2, color.RGBA{G: 255, A: 255})

	button.opts = &ebiten.DrawImageOptions{}
	button.opts.GeoM.Translate(x, y)

	// start with normal image
	button.image = button.normalImage

	return button

}

func (b *Button) LocationInside(x, y int) bool {

	return x >= int(b.x) && x <= int(b.x)+b.width &&
		y >= int(b.y) && y <= int(b.y)+b.height
}
