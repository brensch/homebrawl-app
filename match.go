package main

import (
	"sync"

	"github.com/brensch/smarthome"
)

type match struct {
	id   string
	turn int8

	yourTeam      []*Sprite
	roundResults  []smarthome.State
	viewedResults bool

	mu sync.Mutex
}

func (g *Game) StartMatch() {

	house := smarthome.HouseState{}
	houseSprite := initSprite(house)

	match := &match{
		// TODO: hook into firestore
		id:   "send",
		turn: 0,
		yourTeam: []*Sprite{
			houseSprite,
		},
	}

	g.mu.Lock()
	g.currentMatch = match
	g.currentView = viewTeamSelection
	g.mu.Unlock()

	// set

}
