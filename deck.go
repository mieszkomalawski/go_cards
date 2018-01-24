package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// there are no classes and objects and go
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubes"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//_ ignores returned var
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//return two values
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// kinds like method - "receiver"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), "")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func fromFile(filename string) deck {
	bytes, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	sliceOfStrings := strings.Split(string(bytes), ",")
	return deck(sliceOfStrings)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newKey := r.Intn(len(d) - 1)
		d[i], d[newKey] = d[newKey], d[i]
	}
}
