package main

import (
	"fmt"
	"os"
)

func main() {
	newCards := newDeck()
	newCards.print()
	newCards.shuffle()
	newCards.print()
	hand, remaining := newCards.create_hand(2)
	hand.print()
	remaining.print()
	newCards.saveToFile("cards.save")
	binary_array()
	fmt.Println([]byte("Hello world!"))
	loadedCards := readFromFile("cards.save")
	if loadedCards != nil {
		loadedCards.print()
	} else { // if you move the else to next line it will throw a error!!
		fmt.Println("Couldn't read from file : ", loadedCards)
		os.Exit(-1) //indicate something went wrong!
	}
}

func binary_array() {
	k := []byte{10, 20, 30}
	fmt.Println(k)
}
