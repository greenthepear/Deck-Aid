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
	block       int
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

func (p *Player) moveFromHandToDiscardPileByIndex(index int) {
	p.discardPile = append(p.discardPile, p.hand[index])
	p.hand = removeElement(p.hand, index)
}

func (p *Player) discardHand() {
	p.discardPile = append(p.discardPile, p.hand...)
	p.hand = nil
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

func (p *Player) playCard(card Card, e *Enemy) {
	if card.damage != 0 {
		dmg := card.damage
		if p.effects.impair != 0 {
			fmt.Printf("As you are impaired, the damage is decreased by 25%%.\n")
			dmg = int(float64(dmg) * 0.75) //TODO: Make 0.75 a global thing
		}
		e.hp -= dmg
		fmt.Printf("Enemy attacked for %d! (%d/%d)\n", card.damage, e.hp, e.hpMax)
	}

	if card.block != 0 {
		p.block += card.block
		fmt.Printf("Block applied: %d. (%d/%d) + %d\n", card.block, p.hp, p.hpMax, p.block)
	}

	if card.effects.impair != 0 {
		e.effects.impair += card.effects.impair
		fmt.Printf("Enemy impaired (does less damage) for %d turns.\n", e.effects.impair)
	}

	if card.effects.weaken != 0 {
		e.effects.weaken += card.effects.weaken
		fmt.Printf("Enemy weakened (takes more damage) for %d turns.\n", e.effects.weaken)
	}

	p.energy -= card.cost
}
