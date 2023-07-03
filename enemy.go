// Enemy definitions and behavior
package main

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

func (e *Enemy) doNextAction(p *Player) {
	action := e.actionQueue[e.queueIndex]

	if action.attackDamage != 0 {

	}
}
