// General purpase utility functions
package main

func removeElement(slice []Card, index int) []Card {
	slice = append(slice[:index], slice[index+1:]...)

	return slice
}

// Float differences, for example 1.30 becomes 30, 0.75 becomes 25
func floatDiffToPercentage(f float64) int {
	if f < 1 {
		return int((1 - f) * 100)
	} else {
		return int((f - 1) * 100)
	}
}
