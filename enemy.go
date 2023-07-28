// Enemy definitions and behavior
package main

import (
	"math/rand"
)

type Enemy struct {
	name           string
	health         Health
	effects        Effect
	actionQueue    []EnemyAction
	randomizeQueue bool
	queueIndex     int
}

type EnemyAction struct {
	attackDamage     int
	attackMultiplier int
	addBlock         int
	curse            Effect
}

func (e *Enemy) doHit(baseDamage int, p *Player) {
	dmg := baseDamage
	if p.effects.weaken != 0 {
		//fmt.Printf("As you are weakened, the damage is increased by 30%%.\n")
		dmg = int(float64(baseDamage) * weakenMultiplier)
	}
	p.takeHit(dmg)

}

func (e *Enemy) increaseQueueIndex() {
	actionsNumber := len(e.actionQueue)
	if e.randomizeQueue {
		e.queueIndex = rand.Intn(actionsNumber)
	} else {
		e.queueIndex = (e.queueIndex + 1) % actionsNumber
	}
}

func (e *Enemy) doNextAction(p *Player) {
	action := e.actionQueue[e.queueIndex]

	if action.attackDamage != 0 { //Attacking
		//attackMultiplier = 0 means to not multiply the attack, so 1 is the same thing
		multiplier := 1
		if action.attackMultiplier != 0 {
			multiplier = action.attackMultiplier
		}

		updatedDamage := calcDmgWithEffects(action.attackDamage, e.effects, p.effects, false)

		for i := 0; i < multiplier; i++ {
			e.doHit(updatedDamage, p)
		}
	}

	if action.addBlock != 0 { //Adding block
		e.health.addBlock(action.addBlock)
	}

	if action.curse.doesHaveAnyEffect() { //Applying effects
		action.curse.applyEffectsOn(&p.effects)
	}
}
