package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type 'deck'
// Which is slice of strings

type deck []string

// Receiver function
// Any varibaile type of "deck" can access to the function "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// fmt.Println(cardSuits)
	// fmt.Println(cardValues)
	// fmt.Println(cards)
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			// fmt.Println(i, j)
			// fmt.Println(suit)
			// fmt.Println(val)
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffile() {
	// time.Now().UnixNano() for generating int64 number
	source := rand.NewSource(time.Now().UnixNano())
	// fmt.Println(source)
	r := rand.New(source)
	// fmt.Println(r)
	for i := range d {
		newPostion := r.Intn(len(d) - 1)
		d[i], d[newPostion] = d[newPostion], d[i]

	}
}
