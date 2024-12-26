package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clovers"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	//create two new slices
	//pop out number of cards from deck[0:handsize+1]
	return d[:handsize], d[handsize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bSlice, err := os.ReadFile(filename)
	if err != nil {
		//what do we want to do when an error occurs?
		//option 1 - log error and return a call to newDeck()
		//option 2 - log error and quit program
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
