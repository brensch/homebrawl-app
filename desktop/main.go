package main

import (
	"log"

	homebrawlapp "github.com/brensch/homebrawl-app"
	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

func main() {

	// startingState := smarthome.GetFirstState(houses)
	// states, result := smarthome.PlayGame(startingState)

	// ebiten.SetWindowSize(1920, 1080)
	// img, _, err := ebitenutil.NewImageFromFile("toastersmall.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	homebrawlapp.ScreenWidth, homebrawlapp.ScreenHeight = ebiten.ScreenSizeInFullscreen()

	g := &homebrawlapp.Game{
		// Image:      img,
		// states:             states,
		// result:             result,
		ApplianceDimension: 128,

		CurrentView: homebrawlapp.ViewMenu,
		// appliances: appliances,

		Menu: homebrawlapp.InitMenu(),
	}

	// ebiten.CurrentFPS()

	ebiten.SetWindowTitle("Homebrawl")
	// ebiten.SetWindowResizable(true)
	ebiten.SetFullscreen(true)
	// ebiten.IsFullscreen()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

// func main() {
// 	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
// 	ebiten.SetWindowTitle("Drag & Drop (Ebiten Demo)")
// 	if err := ebiten.RunGame(NewGame()); err != nil {
// 		log.Fatal(err)
// 	}
// }
