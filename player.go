package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	health      Health
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

func (p *Player) decreaseDebuffEffects() {
	if p.effects.impair > 0 {
		p.effects.impair--
	}

	if p.effects.weaken > 0 {
		p.effects.weaken--
	}
}

func (p *Player) playCard(card Card, e *Enemy) {
	if card.damage != 0 {
		dmg := card.damage
		dmg = calcDmgWithEffects(dmg, p.effects, e.effects, true)
		fmt.Printf("Attacking enemy for %d...\n", dmg)
		actualDamage := e.health.takeHit(dmg)
		fmt.Printf("Enemy hit for %d! %s\n", actualDamage, e.health.genHPstring())
	}

	if card.block != 0 {
		p.health.addBlock(card.block)
	}

	if card.effects.doesHaveAnyEffect() {
		card.effects.applyEffectsOn(&e.effects)
	}

	p.energy -= card.cost
}

func (p *Player) takeHit(dmg int) {
	hpLoss := p.health.takeHit(dmg)
	fmt.Printf("You got hit for %d! %s\n", hpLoss, p.health.genHPstring())
}
