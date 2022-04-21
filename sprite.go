package homebrawlapp

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
		image:     ebiten.NewImage(ScreenWidth, SpriteHeight),
		upgrade:   ebiten.NewImage(ScreenWidth, SpriteHeight),
		opts:      &ebiten.DrawImageOptions{},
	}
}
