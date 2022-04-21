package homebrawlapp

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) DrawTeamSelection(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "team select")

}

func (g *Game) UpdateTeamSelection() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

		if int(g.turn) == len(g.states)-1 {
			return
		}
		g.turn++

	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {

		if int(g.turn) == 0 {
			return
		}
		g.turn--

	}
}
