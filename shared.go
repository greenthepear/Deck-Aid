// Shared properties for the player and enemy - mostly stuff like (de)buffs and health
package main

import "fmt"

type Effect struct {
	impair int //Makes do less damage
	weaken int //Makes take more damage
}

// Should these be global or should I make another struct or something?
var impairMultiplier float64 = 0.75
var weakenMultiplier float64 = 1.30

type Health struct {
	hp    int
	hpMax int
	block int
}

func (effect Effect) totalEffects() int {
	return effect.impair + effect.weaken
}

func (effect Effect) doesHaveAnyEffect() bool {
	return effect.totalEffects() > 0
}

func calcDmgWithEffects(dmg int, attackerEffects Effect, defenderEffects Effect, isPlayer bool) int {
	if attackerEffects.impair != 0 {
		if isPlayer {
			fmt.Printf("As you are impaired, the damage is decreased by %d%%.\n", floatDiffToPercentage(impairMultiplier))
		} else {
			fmt.Printf("As the enemy is impaired, the damage is decreased by %d%%.\n", floatDiffToPercentage(impairMultiplier))
		}
		dmg = int(float64(dmg) * impairMultiplier)
	}
	if defenderEffects.weaken != 0 {
		if isPlayer {
			fmt.Printf("As they are weakened, the damage is increased by %d%%.\n", floatDiffToPercentage(weakenMultiplier))
		} else {
			fmt.Printf("As you are weakened, the damage is increased by %d%%.\n", floatDiffToPercentage(weakenMultiplier))
		}
		dmg = int(float64(dmg) * weakenMultiplier)
	}

	return dmg
}

func (eff *Effect) applyEffectsOn(targetEffects *Effect) {
	if eff.impair != 0 {
		targetEffects.impair += eff.impair
		fmt.Printf("Impaired (does less damage) for %d turns.\n", targetEffects.impair)
	}

	if eff.weaken != 0 {
		targetEffects.weaken += eff.weaken
		fmt.Printf("Weakened (takes more damage) for %d turns.\n", targetEffects.weaken)
	}
}

func (h *Health) takeHit(dmg int) int {
	hpLoss := dmg
	if h.block > 0 {
		blockOriginal := h.block
		fmt.Printf("Blocked ")
		h.block -= dmg
		if h.block <= 0 {
			fmt.Printf("%d damage! Block broken! ", blockOriginal)
			hpLoss = -h.block
			h.block = 0
		} else {
			fmt.Printf("%d damage! ", blockOriginal-h.block)
			hpLoss = 0
		}
	}

	h.hp -= hpLoss

	return hpLoss
}

func (h *Health) addBlock(block int) {
	h.block += block
	fmt.Printf("Block applied: %d. %s\n", block, h.genHPstring())
}

func (h *Health) removeBlock() {
	h.block = 0
}

func (h Health) genHPstring() string {
	r := fmt.Sprintf("(%d/%d)", h.hp, h.hpMax)
	if h.block > 0 {
		r += fmt.Sprintf(" + %d block", h.block)
	}
	return r
}
