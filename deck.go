package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, number int) (first deck, second deck) {
	return d[:number], d[number:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		// fmt.Println("Error:- ", err)
		panic(err)
	}
	s := strings.Split(string(bytes), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for index := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[newPosition], d[index] = d[index], d[newPosition]
	}
}
