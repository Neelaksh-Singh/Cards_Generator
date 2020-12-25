package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new Deck type
// which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Takes in deck and returns 2 decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//logs out all the cards present in the deck Slice.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Converts slice of deck type first into string type and then joins
// to form a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

//Saves our deck to file in form of byte type
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Returns us back a deck slice containing our cards from our saved file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//Shuffels our deck randomly everytime we call this function
func (d deck) Shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
