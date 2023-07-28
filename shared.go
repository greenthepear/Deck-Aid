// Shared properties for the player and enemy - mostly stuff like (de)buffs and health
package main

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
