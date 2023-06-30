package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	name    string
	cost    int
	damage  int
	block   int
	effects Effect
}

type Effect struct {
	weaken     int
	vulnerable int
}

type Enemy struct {
	name    string
	hp      int
	hpMax   int
	intent  int
	damage  int
	effects Effect
}

func combatStart(player *Player, enemy *Enemy) {
	player.drawPile = player.deck
}

func createCardSliceByReferanceIntSlice(cardTypes []Card, cardIntSlice []int) []Card {
	rSlice := make([]Card, len(cardIntSlice))
	for i, cardNumber := range cardIntSlice {
		rSlice[i] = cardTypes[cardNumber]
	}
	return rSlice
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	cards := []Card{
		{"Strike", 1, 7, 0, Effect{0, 0}},
		{"Duck", 1, 0, 5, Effect{0, 0}},
		{"Shock", 0, 0, 3, Effect{2, 0}},
		{"Crush", 2, 3, 3, Effect{1, 1}},
	}

	startingDeck := []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 3}

	player := Player{
		hp:        0, //start at 0 to test combatStart
		hpMax:     50,
		energy:    0,
		energyMax: 3,
		effects:   Effect{0, 0},
		deck:      createCardSliceByReferanceIntSlice(cards, startingDeck),
	}

	enemy := Enemy{
		name:    "Duende",
		hp:      20,
		hpMax:   20,
		intent:  0,
		damage:  5,
		effects: Effect{0, 0},
	}
	combatStart(&player, &enemy)
	player.draw(5)
	fmt.Print(player.hand)
}
