package homebrawlapp

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	buttons []*Button
}

func InitMenu() *Menu {

	startGameButton := InitButton(float64(ScreenWidth)/4, float64(ScreenHeight)/4, ScreenWidth/2, ScreenHeight/4, "start", color.White, ViewGameResult)
	settingsButton := InitButton(float64(ScreenWidth)/4, float64(ScreenHeight)/2, ScreenWidth/2, ScreenHeight/4, "settings", color.White, ViewOptions)

	return &Menu{
		buttons: []*Button{
			startGameButton,
			settingsButton,
		},
	}
	// Menu := ebiten.NewImage(ScreenWidth/2, ScreenHeight/2)

}

func (g *Game) DrawMenu(screen *ebiten.Image) {

	// fmt.Println(ScreenHeight, ScreenWidth)

	for _, button := range g.Menu.buttons {
		screen.DrawImage(button.image, button.opts)
	}

}

func (g *Game) UpdateMenu() {

	// ScreenWidth, ScreenHeight := g.Layout()

	// fmt.Println(ScreenHeight, ScreenWidth)

	cursorX, cursorY := ebiten.CursorPosition()

	for _, button := range g.Menu.buttons {
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
