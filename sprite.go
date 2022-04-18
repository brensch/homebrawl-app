package main

import (
	"github.com/brensch/smarthome"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	appliance smarthome.Appliance

	image   *ebiten.Image
	upgrade *ebiten.Image
	opts    *ebiten.DrawImageOptions
}

const (
	SpriteWidth  = 256
	SpriteHeight = 256
)

func initSprite(appliance smarthome.Appliance) *Sprite {
	return &Sprite{
		appliance: appliance,
		image:     ebiten.NewImage(screenWidth, SpriteHeight),
		upgrade:   ebiten.NewImage(screenWidth, SpriteHeight),
		opts:      &ebiten.DrawImageOptions{},
	}
}
