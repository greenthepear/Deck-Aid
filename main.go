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

func combatStart(player *Player, enemy *Enemy) {
	player.drawPile = player.deck
	enemy.queueIndex = 0
}

func cardInput(player *Player, enemy *Enemy) bool {
	var chosenCard int
	for {
		fmt.Printf("You have -%d- energy left. Choose a card to play, 0 to end turn: ", player.energy)
		_, err := fmt.Scanf("%d", &chosenCard)
		if err != nil {
			fmt.Printf("Invalid input!\n")
			continue
		}

		if chosenCard == 0 {
			fmt.Print("Ending turn...\n\n")
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
		if enemy.health.hp <= 0 {
			return true
		}
		player.moveFromHandToDiscardPileByIndex(chosenCard - 1)
		printTurnSuffix(*player, *enemy)
		return false
	}
}

func doEnemyTurn(player *Player, enemy *Enemy) {
	fmt.Printf("Enemy turn starts!\n")
	enemy.health.removeBlock()
	enemy.doNextAction(player)
	enemy.increaseQueueIndex()
}

func doTurn(player *Player, enemy *Enemy) {
	player.decreaseDebuffEffects()
	player.health.removeBlock()
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
		if player.health.hp <= 0 {
			fmt.Printf("You have been defeated...\n")
			return
		}
		doTurn(player, enemy)
		if enemy.health.hp <= 0 {
			fmt.Printf("\n%s has been defeated!\n", enemy.name)
			return
		}
		doEnemyTurn(player, enemy)
	}
}

func createCardSliceByReferanceIntSlice(cardTypes []Card, cardIntSlice []int) []Card { //is it descriptive enough?
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
		health:     Health{50, 50, 0},
		energy:     0,
		energyMax:  3,
		effects:    Effect{0, 0},
		deck:       createCardSliceByReferanceIntSlice(cards, startingDeck),
		drawNumber: 5,
	}

	enemy := Enemy{
		name:    "Duende",
		health:  Health{20, 20, 0},
		effects: Effect{0, 0},
		//Attack 2x4, Block 12, Attack 9, Curse with both for 2
		actionQueue: []EnemyAction{
			{attackDamage: 4, attackMultiplier: 2},
			{addBlock: 12},
			{attackDamage: 9},
			{curse: Effect{2, 2}},
		},
		randomizeQueue: false,
		queueIndex:     0,
	}
	gameLoop(&player, &enemy)
}
