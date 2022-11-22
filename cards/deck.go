package main

// how to import multiple libs in go ? here is how :
import (
	"fmt" // no commas or anything here!
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// this function is applied to deck types (d is the reciver of the function)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// tip : "" is different from ” in go
func newDeck() deck {
	res := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Queen", "king"}
	for _, cardSuit := range cardSuits { // _ means we don't care about it
		for _, cardValue := range cardValues {
			res = append(res, cardValue+" of "+cardSuit)
		}
	}

	return res
}

func (d deck) slice() {
	fmt.Println(d[0])    //access the item with index
	fmt.Println(d[0:10]) //use a:b to slice the slice one more time , including a but not b (a, a+1 , ... b-1)
	fmt.Println(d[4:])   //including 4 till the end of the slice
	fmt.Println(d[4 : len(d)-3])
}

func (d deck) create_hand(handsize int) (deck, deck) {
	left := d[:handsize]
	right := d[handsize:]
	return left, right
}

func (d deck) saveToFile(fileName string) error { //what is interface in go?!
	//type conversion in go :
	return ioutil.WriteFile(fileName, []byte(strings.Join(d, ",")), 0666)
}

func readFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName) //byte[] , error
	if err != nil {
		return nil
	}
	stringified := string(bs)
	cards := strings.Split(stringified, ",")
	return deck(cards)
}

// lesssons : random number generation, shuffle list, work with time
func (d deck) shuffle() {
	//set the seed
	random_number_gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		new_position := random_number_gen.Intn(len(d))
		// tmp := d[i]
		// d[i] = d[new_position]
		// d[new_position] = tmp

		//lets use one line swap here :
		d[i], d[new_position] = d[new_position], d[i]
	}
}
