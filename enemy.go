// Enemy definitions and behavior
package main

import (
	"fmt"
	"math/rand"
)

type Enemy struct {
	name           string
	hp             int
	hpMax          int
	block          int
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

	if action.attackDamage != 0 {
		multiplier := 1
		if action.attackMultiplier != 0 {
			multiplier = action.attackMultiplier
		}

		if p.effects.weaken != 0 {
			fmt.Printf("As you are weakened, the damage is increased by 30%%.\n")
		}

		for i := 0; i < multiplier; i++ {
			e.doHit(action.attackDamage, p)
		}
	}
}
