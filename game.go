package homebrawlapp

import (
	"image/color"
	"sync"

	"github.com/brensch/smarthome"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type view int

const (
	ViewMenu view = iota
	ViewOptions
	ViewTeamSelection
	ViewGameResult
)

// var square *ebiten.Image

type Game struct {
	Text string
	// Image *ebiten.Image
	yourTeam *smarthome.House
	states   []smarthome.State
	result   int8
	turn     int8

	Menu *Menu

	ApplianceDimension int

	CurrentView view

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

	switch g.CurrentView {
	case ViewMenu:
		g.UpdateMenu()
	case ViewGameResult:
		g.UpdateGameResult()
	case ViewTeamSelection:
		g.UpdateTeamSelection()
	}
	// g.Text = fmt.Sprint(ebiten.CurrentFPS())
	// fmt.Println(time.Now())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.CurrentView {
	case ViewMenu:
		g.DrawMenu(screen)
	case ViewTeamSelection:
		g.DrawTeamSelection(screen)
	case ViewGameResult:
		g.DrawGameResult(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
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
	ScreenWidth  = 640
	ScreenHeight = 480
)
