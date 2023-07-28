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
		if p.effects.impair != 0 {
			fmt.Printf("As you are impaired, the damage is decreased by %d%%.\n", int((1-impairMultiplier)*100))
			dmg = int(float64(dmg) * impairMultiplier) //TODO: Make 0.75 a global thing
		}
		if e.effects.weaken != 0 {
			fmt.Printf("As the enemy is weakened, the damage is increased by %d%%.\n", int((weakenMultiplier-1)*100))
			dmg = int(float64(dmg) * weakenMultiplier)
		}
		e.health.hp -= dmg
		fmt.Printf("Enemy attacked for %d! (%d/%d)\n", dmg, e.health.hp, e.health.hpMax)
	}

	if card.block != 0 {
		p.health.block += card.block
		fmt.Printf("Block applied: %d. %s\n", card.block, p.genHPstring())
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

func (p Player) genHPstring() string {
	r := fmt.Sprintf("(%d/%d)", p.health.hp, p.health.hpMax)
	if p.health.block > 0 {
		r += fmt.Sprintf(" + %d block", p.health.block)
	}
	return r
}

func (p *Player) takeHit(dmg int) {
	hpLoss := dmg
	if p.health.block > 0 {
		blockOriginal := p.health.block
		fmt.Printf("Blocked ")
		p.health.block -= dmg
		if p.health.block <= 0 {
			fmt.Printf("%d damage! Block broken! ", blockOriginal)
			hpLoss = -p.health.block
			p.health.block = 0
		} else {
			fmt.Printf("%d damage! ", blockOriginal-p.health.block)
			hpLoss = 0
		}
	}

	p.health.hp -= hpLoss

	fmt.Printf("You got hit for %d! %s\n", hpLoss, p.genHPstring())
}
