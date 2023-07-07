// Generating CLI "graphics"
package main

import "fmt"

func genEffectString(effect Effect) string {
	r := ""
	if effect.impair > 0 {
		r += "Impair " + fmt.Sprint(effect.impair) + " "
	}

	if effect.weaken > 0 {
		r += "Weaken " + fmt.Sprint(effect.weaken) + " "
	}
	return r
}

func printCards(cards []Card, doPrintNum bool) {
	numberOfCards := len(cards)
	if doPrintNum {
		fmt.Printf("Number of cards: %d", numberOfCards)
	}
	if numberOfCards == 0 {
		return
	}
	fmt.Printf("\nn. Name  \tCost\tDmg\tBlc\tEffects\n")
	for i, card := range cards {
		fmt.Printf("%d. %s\t%d\t%d\t%d\t%s\n", i+1, card.name, card.cost, card.damage, card.block, genEffectString(card.effects))
	}
}

//lint:ignore U1000 Debugging function, ignore being unused
func (p Player) printAllCards(doPrintDeck bool) {
	if doPrintDeck {
		fmt.Printf("\nDECK - ")
		printCards(p.deck, true)
	}

	fmt.Printf("\nHAND - ")
	printCards(p.drawPile, true)

	fmt.Printf("\nDRAW PILE - ")
	printCards(p.drawPile, true)

	fmt.Printf("\nDISCARD PILE - ")
	printCards(p.discardPile, true)
}

func printTurnPrefix(player Player, enemy Enemy) {
	fmt.Printf("Your hp: %s", player.genHPstring())
	if player.block != 0 {
		fmt.Printf(" + %d block", player.block)
	}
	if player.effects.doesHaveAnyEffect() {
		fmt.Printf("\nEffects applied on you: %s", genEffectString(player.effects))
	}

	fmt.Printf("\nThe foe %s (%d/%d)", enemy.name, enemy.hp, enemy.hpMax)
	if enemy.block != 0 {
		fmt.Printf(" + %d block", enemy.block)
	}
	if enemy.effects.doesHaveAnyEffect() {
		fmt.Printf("\nEffects applied the enemy: %s", genEffectString(enemy.effects))
	}
	fmt.Print("\n")

	currentEnemyAction := enemy.actionQueue[enemy.queueIndex]
	if currentEnemyAction.attackDamage != 0 {
		fmt.Printf("The enemy is intending to attack you! Damage: ")
		if currentEnemyAction.attackMultiplier > 1 {
			fmt.Printf("%dx%d\n", currentEnemyAction.attackMultiplier, currentEnemyAction.attackDamage)
		} else {
			fmt.Printf("%d\n", currentEnemyAction.attackDamage)
		}
	}

	if currentEnemyAction.addBlock != 0 {
		fmt.Printf("The enemy is intending to apply block. Blc: %d\n", currentEnemyAction.addBlock)
	}

	if currentEnemyAction.curse.impair != 0 || currentEnemyAction.curse.weaken != 0 {
		fmt.Printf("The enemy is intending to curse you! Effects: %s", genEffectString(currentEnemyAction.curse))
	}
}

func printDrawDiscardNumber(player Player) {
	fmt.Printf("Draw pile: %d cards.\nDiscard pile: %d cards.\n", len(player.drawPile), len(player.discardPile))
}

func printTurnSuffix(player Player, enemy Enemy) {
	fmt.Printf("\nHand: %d cards.", len(player.hand))
	printCards(player.hand, false)
}
