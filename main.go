package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/brensch/smarthome"

	"image/color"
	_ "image/png"
)

type view int

const (
	viewMenu view = iota
	viewOptions
	viewTeamSelection
	viewGameResult
)

// var square *ebiten.Image

type Game struct {
	Text string
	// Image *ebiten.Image
	yourTeam *smarthome.House
	states   []smarthome.State
	result   int8
	turn     int8

	menu *Menu

	applianceDimension int

	currentView view

	currentMatch *match

	mu sync.Mutex

	// appliances []smarthome.Appliance
}

var (
	mplusNormalFont font.Face
	mplusLargeFont  font.Face

	upColour = color.RGBA{
		R: 0x50,
		G: 0x6b,
		B: 0x30,
		A: 0xff,
	}
	downColour = color.RGBA{
		R: 0x6b,
		G: 0x5a,
		B: 0x30,
		A: 0xff,
	}
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	mplusLargeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	// square = ebiten.NewImage(128, 128)
}

func (g *Game) Update() error {

	switch g.currentView {
	case viewMenu:
		g.UpdateMenu()
	case viewGameResult:
		g.UpdateGameResult()
	case viewTeamSelection:
		g.UpdateTeamSelection()
	}
	// g.Text = fmt.Sprint(ebiten.CurrentFPS())
	// fmt.Println(time.Now())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.currentView {
	case viewMenu:
		g.DrawMenu(screen)
	case viewTeamSelection:
		g.DrawTeamSelection(screen)
	case viewGameResult:
		g.DrawGameResult(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

var (
	houses = [2]smarthome.House{
		{
			Appliances: []smarthome.Appliance{
				smarthome.HouseState{
					ObjectState: smarthome.ObjectState{
						Health:   3,
						Strength: 3,
					},
				},
				// GoingUp
				smarthome.Toaster{
					ObjectState: smarthome.ObjectState{
						Location: smarthome.Location{
							X: 0,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
				smarthome.Toaster{
					ObjectState: smarthome.ObjectState{
						Location: smarthome.Location{
							X: 1,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
			},
			// State: ,
		},
		{
			Appliances: []smarthome.Appliance{
				smarthome.HouseState{
					ObjectState: smarthome.ObjectState{
						Health:   3,
						Strength: 3,
					},
				},
				// GoingUp
				smarthome.Toaster{
					ObjectState: smarthome.ObjectState{
						Location: smarthome.Location{
							X: 0,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
				smarthome.Sticky{
					ObjectState: smarthome.ObjectState{
						Location: smarthome.Location{
							X: 1,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
			},
		},
	}
)

var (
	screenWidth  = 640
	screenHeight = 480
)

func main() {

	// startingState := smarthome.GetFirstState(houses)
	// states, result := smarthome.PlayGame(startingState)

	// ebiten.SetWindowSize(1920, 1080)
	// img, _, err := ebitenutil.NewImageFromFile("toastersmall.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()

	g := &Game{
		// Image:      img,
		// states:             states,
		// result:             result,
		applianceDimension: 128,

		currentView: viewMenu,
		// appliances: appliances,

		menu: InitMenu(),
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
// 	ebiten.SetWindowSize(screenWidth, screenHeight)
// 	ebiten.SetWindowTitle("Drag & Drop (Ebiten Demo)")
// 	if err := ebiten.RunGame(NewGame()); err != nil {
// 		log.Fatal(err)
// 	}
// }
