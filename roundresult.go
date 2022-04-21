package homebrawlapp

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/brensch/smarthome"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) DrawGameResult(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	// op.GeoM.Scale(.5, .5)
	// offset := time.Now().UnixMilli() % int64(screen.Bounds().Dx())

	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	op.GeoM.Translate(float64(x), float64(y))
	// }

	// text.Draw(g.Image, "lumb", mplusNormalFont, imageWidth/2, imageHeight/2, color.Black)
	// screen.DrawImage(g.Image, op)
	var houses [2]smarthome.HouseState

	for _, appliance := range g.states[g.turn].Appliances {
		house, ok := appliance.(smarthome.HouseState)
		if ok {
			if house.Team == -1 {
				houses[0] = house
			} else {
				houses[1] = house
			}
			continue
		}

		applianceSquare := ebiten.NewImage(g.ApplianceDimension, g.ApplianceDimension)
		op := &ebiten.DrawImageOptions{}

		// add 1 to y for house
		imageX, imageY := float64(appliance.State().Location.X)*float64(g.ApplianceDimension), float64(appliance.State().Location.Y+1)*float64(g.ApplianceDimension)
		// fmt.Println(appliance.Type(), imageX, imageY)
		op.GeoM.Translate(imageX, imageY)
		// op.ColorM.Apply(color.White)
		if appliance.State().Team == -1 {
			applianceSquare.Fill(downColour)
		} else {
			applianceSquare.Fill(upColour)
		}

		// op.C
		label := fmt.Sprintf("(%d,%d) %s\nhealth: %d\nstrength: %d", appliance.State().Location.X, appliance.State().Location.Y, appliance.Type(), appliance.State().Health, appliance.State().Strength)
		text.Draw(applianceSquare, label, mplusNormalFont, 0, mplusNormalFont.Metrics().Height.Ceil(), color.Black)

		screen.DrawImage(applianceSquare, op)
	}

	// draw events sidepanel
	eventsSquare := ebiten.NewImage(g.ApplianceDimension*2, g.ApplianceDimension*2)
	eventsString := bytes.NewBuffer(nil)
	for _, event := range g.states[g.turn].Events {
		eventsString.WriteString(fmt.Sprintf("%s - %d,%d\n", event.Type(), event.Base().Target.X, event.Base().Target.Y))
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(3*g.ApplianceDimension), 0)
	text.Draw(eventsSquare, eventsString.String(), mplusNormalFont, 0, mplusNormalFont.Metrics().Height.Ceil(), color.White)
	screen.DrawImage(eventsSquare, op)

	if int(g.turn) == len(g.states)-1 {
		resultSquare := ebiten.NewImage(g.ApplianceDimension*2, g.ApplianceDimension*2)
		op.GeoM.Translate(float64(3*g.ApplianceDimension), 0)
		text.Draw(resultSquare, "lumb", mplusNormalFont, 0, mplusNormalFont.Metrics().Height.Ceil(), color.White)

		screen.DrawImage(resultSquare, op)
	}

	// draw houses
	upHouse := ebiten.NewImage(g.ApplianceDimension*3, g.ApplianceDimension)
	upHouse.Fill(upColour)
	upHouseOp := &ebiten.DrawImageOptions{}
	upHouseState := fmt.Sprintf("Team UP\nhealth: %d", houses[1].Health)
	text.Draw(upHouse, upHouseState, mplusNormalFont, 0, mplusNormalFont.Metrics().Height.Ceil(), color.Black)
	screen.DrawImage(upHouse, upHouseOp)

	downHouse := ebiten.NewImage(g.ApplianceDimension*3, g.ApplianceDimension)
	downHouse.Fill(downColour)
	downHouseOp := &ebiten.DrawImageOptions{}
	downHouseOp.GeoM.Translate(0, float64(7*g.ApplianceDimension))
	downHouseState := fmt.Sprintf("Team DOWN\nhealth: %d", houses[0].Health)
	text.Draw(downHouse, downHouseState, mplusNormalFont, 0, mplusNormalFont.Metrics().Height.Ceil(), color.Black)

	screen.DrawImage(downHouse, downHouseOp)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f, %d:%d %t\nturn:%d", ebiten.CurrentFPS(), x, y, ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft), g.turn))

}

func (g *Game) UpdateGameResult() {
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
