package main

import "math/rand"

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
}

func (p *Player) refillDrawPile() {
	p.drawPile = p.discardPile
	p.discardPile = nil
}

func (p *Player) draw(number int) {
	for i := 0; i < number; i++ {
		drawPileLen := len(p.drawPile)
		if drawPileLen == 0 {
			p.refillDrawPile()
			drawPileLen = len(p.drawPile)
		}

		randIndex := rand.Intn(drawPileLen + 1)
		p.hand = append(p.hand, p.drawPile[randIndex])
		p.drawPile = removeElement(p.drawPile, randIndex)
	}
}
