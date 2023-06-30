package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	hp          int
	hpMax       int
	energy      int
	energyMax   int
	effects     Effect
	deck        []Card
	drawPile    []Card
	discardPile []Card
	hand        []Card
	drawNumber  int
}

func (p *Player) refillDrawPile() {
	p.drawPile = p.discardPile
	p.discardPile = nil
}

func (p *Player) draw(number int) {
	if number > len(p.deck) {
		number = len(p.deck)
	}

	for i := 0; i < number; i++ {
		drawPileLen := len(p.drawPile)
		if drawPileLen == 0 {
			fmt.Printf("Draw pile empty, shuffling from discard pile...\n")
			p.refillDrawPile()
			drawPileLen = len(p.drawPile)
		}

		randIndex := rand.Intn(drawPileLen)
		p.hand = append(p.hand, p.drawPile[randIndex])
		p.drawPile = removeElement(p.drawPile, randIndex)
	}
}
