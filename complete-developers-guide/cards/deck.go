package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck which extends a slice of string
type deck []string

// Loops through the deck and prints each card with it's index
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Generates a new deck will every card
func newDeck() deck {
	newDeck := deck{}
	// Loop through every suit ignoring the index
	for _, suit := range suits() {
		// loop through every card in a set ignoring the index
		for _, card := range set() {
			// Append the card to the deck
			newDeck = append(newDeck, card+" of "+suit)
		}
	}
	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Saves the deck to the given file name
func (d deck) saveDeck(fileName string) {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// loads a deck from a file
func loadDeck(fileName string) deck {
	// read the file
	byteDeck, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Convert the byteArray to a string, then split it into a slice of strings
	// and type it as a deck
	return deck(strings.Split(string(byteDeck), ","))
}

func (d deck) shuffle() {
	// Create a new random object with a new source with the current epoch
	// as the seed. Without making a new rand or using rand.Seed the seed
	// will always be 1 meaning it will always give the same numbers
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	deckLen := len(d) - 1
	for i := range d {
		newPos := r.Intn(deckLen)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

// Converts the called deck to a sting seperated by commas
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// suits returns all suits that you would find in a deck
func suits() []string {
	return []string{"Hearts", "Spades", "Clubs", "Diamonds"}
}

// set returns every card you would find in a suits set
func set() []string {
	return []string{
		"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "King", "Queen",
	}
}
