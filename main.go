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

func cardInput(player *Player, enemy *Enemy) bool {
	invalidChoice := true
	var chosenCard int
	for invalidChoice {
		fmt.Printf("You have -%d- energy left. Choose a card to play, 0 to end turn: ", player.energy)
		_, err := fmt.Scanf("%d", &chosenCard)
		if err != nil {
			fmt.Printf("Invalid input!\n")
			continue
		}

		if chosenCard == 0 {
			fmt.Print("Ending turn...\n")
			return true
		}

		if chosenCard < 1 || chosenCard > len(player.hand) {
			fmt.Printf("Not a card!\n")
			continue
		}

		cardToPlay := player.hand[chosenCard-1]
		if cardToPlay.cost > player.energy {
			fmt.Printf("Not enough energy to play %s (%d)...\n", cardToPlay.name, cardToPlay.cost)
			continue
		}
		fmt.Printf("Playing %s...\n", cardToPlay.name)
		player.playCard(cardToPlay, enemy)
		if enemy.hp <= 0 {
			return true
		}
		player.moveFromHandToDiscardPileByIndex(chosenCard - 1)
		invalidChoice = false
		printTurnSuffix(*player, *enemy)
	}
	return false
}

func doTurn(player *Player, enemy *Enemy) {
	printTurnPrefix(*player, *enemy)
	fmt.Printf("Your turn starts, drawing %d cards...\n\n", player.drawNumber)
	player.draw(player.drawNumber)
	printDrawDiscardNumber(*player)
	printTurnSuffix(*player, *enemy)
	player.energy = player.energyMax
	for {
		if cardInput(player, enemy) {
			player.discardHand()
			break
		}
	}
}

func gameLoop(player *Player, enemy *Enemy) {
	combatStart(player, enemy)

	for {
		if enemy.hp <= 0 {
			fmt.Printf("%s has been defeated!\n", enemy.name)
			return
		}
		doTurn(player, enemy)
	}
}

func createCardSliceByReferanceIntSlice(cardTypes []Card, cardIntSlice []int) []Card {
	rSlice := make([]Card, len(cardIntSlice))
	for i, cardNumber := range cardIntSlice {
		rSlice[i] = cardTypes[cardNumber]
	}
	return rSlice
}

func main() {
	//Commented out for reproductible results
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.New(rand.NewSource(0))

	cards := []Card{
		{"Strike", 1, 7, 0, Effect{0, 0}},
		{"Duck  ", 1, 0, 5, Effect{0, 0}}, //Spaces for better print formatting, temporary fix
		{"Shock", 0, 0, 3, Effect{2, 0}},
		{"Crush", 2, 3, 3, Effect{1, 1}},
	}

	startingDeck := []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 3}

	player := Player{
		hp:         50,
		hpMax:      50,
		energy:     0,
		energyMax:  3,
		effects:    Effect{0, 0},
		deck:       createCardSliceByReferanceIntSlice(cards, startingDeck),
		drawNumber: 5,
	}

	enemy := Enemy{
		name:    "Duende",
		hp:      20,
		hpMax:   20,
		intent:  0,
		damage:  5,
		effects: Effect{0, 0},
	}

	gameLoop(&player, &enemy)
}
