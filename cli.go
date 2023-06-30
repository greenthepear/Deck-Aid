// Generating CLI "graphics"
package main

import "fmt"

func genEffectString(card Card) string {
	r := ""
	if card.effects.vulnerable > 0 {
		r += "Vulnerable " + fmt.Sprint(card.effects.vulnerable) + " "
	}

	if card.effects.weaken > 0 {
		r += "Weaken " + fmt.Sprint(card.effects.weaken) + " "
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
		fmt.Printf("%d. %s\t%d\t%d\t%d\t%s\n", i+1, card.name, card.cost, card.damage, card.block, genEffectString(card))
	}
}

func (p Player) printAllCards(doPrintDeck bool) { //for debugging
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
	fmt.Printf("Your hp: (%d/%d)\n", player.hp, player.hpMax)
	fmt.Printf("The foe %s (%d/%d)", enemy.name, enemy.hp, enemy.hpMax)
	switch enemy.intent {
	case 0:
		fmt.Printf(" is going to attack you!\n")
	default:
		fmt.Printf(" is doing something you have no clue about.\n")
	}
}

func printDrawDiscardNumber(player Player) {
	fmt.Printf("Draw pile: %d cards.\nDiscard pile: %d cards.\n", len(player.drawPile), len(player.discardPile))
}

func printTurnSuffix(player Player, enemy Enemy) {
	fmt.Printf("\nHand: %d cards.", len(player.hand))
	printCards(player.hand, false)
}
