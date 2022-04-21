package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	homebrawlapp "github.com/brensch/homebrawl-app"
)

func init() {

	g := &homebrawlapp.Game{
		// Image:      img,
		// states:             states,
		// result:             result,
		ApplianceDimension: 128,

		CurrentView: homebrawlapp.ViewMenu,
		// appliances: appliances,

		Menu: homebrawlapp.InitMenu(),
	}

	mobile.SetGame(g)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
