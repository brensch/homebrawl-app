package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	buttons []*Button
}

func InitMenu() *Menu {

	startGameButton := InitButton(float64(screenWidth)/4, float64(screenHeight)/4, screenWidth/2, screenHeight/4, "start", color.White, viewGameResult)
	settingsButton := InitButton(float64(screenWidth)/4, float64(screenHeight)/2, screenWidth/2, screenHeight/4, "settings", color.White, viewOptions)

	return &Menu{
		buttons: []*Button{
			startGameButton,
			settingsButton,
		},
	}
	// menu := ebiten.NewImage(screenWidth/2, screenHeight/2)

}

func (g *Game) DrawMenu(screen *ebiten.Image) {

	// fmt.Println(screenHeight, screenWidth)

	for _, button := range g.menu.buttons {
		screen.DrawImage(button.image, button.opts)
	}

}

func (g *Game) UpdateMenu() {

	// screenWidth, screenHeight := g.Layout()

	// fmt.Println(screenHeight, screenWidth)

	cursorX, cursorY := ebiten.CursorPosition()

	for _, button := range g.menu.buttons {
		if button.LocationInside(cursorX, cursorY) {
			button.image = button.highlightedImage
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				g.StartMatch()
				return
			}
		} else {
			button.image = button.normalImage
		}

	}

}
