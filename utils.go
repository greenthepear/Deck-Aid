// General purpase utility functions
package main

func removeElement(slice []Card, index int) []Card {
	slice = append(slice[:index], slice[index+1:]...)

	return slice
}
