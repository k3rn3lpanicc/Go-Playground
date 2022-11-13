package main

import (
	"fmt"
)

func slices_main() {
	//how to create a slice
	cards := []string{"salam", "salam2", foos()}
	//how to add a new node to slice?
	cards = append(cards, "matin") //seems that this append function creates a new slice and does not modify the existing one, isn't that supposed to be bad?

	//how to iterate over a slice ?
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}
