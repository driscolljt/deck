package deck

import (
	"math/rand"
	"time"
)

type value struct {
	n int
	s string
	l string
}

var suits = []string{"Spades", "Diamonds", "Clubs", "Hearts"}

var values = []value{
	// {n: 0, s: "*", l: "Joker"},
	{n: 1, s: "1", l: "Ace"},
	{n: 2, s: "2", l: "Two"},
	{n: 3, s: "3", l: "Three"},
	{n: 4, s: "4", l: "Four"},
	{n: 5, s: "5", l: "Five"},
	{n: 6, s: "6", l: "Six"},
	{n: 7, s: "7", l: "Sever"},
	{n: 8, s: "8", l: "Eight"},
	{n: 9, s: "9", l: "Nine"},
	{n: 10, s: "10", l: "Ten"},
	{n: 11, s: "J", l: "Jack"},
	{n: 12, s: "Q", l: "Queen"},
	{n: 13, s: "K", l: "King"},
}

// Card represents a single card in the deck
type Card struct {
	suit  string
	value value
}

// String returns a string representation of the Card
func (c Card) String() string {
	if c.value.n != 0 {
		return ("" + c.value.l + " of " + c.suit)
	} else {
		return ("" + c.value.l + " | " + c.suit)
	}
}

// Deck represents a collection of cards that can be leveraged to play games
type Deck struct {
	Cards []Card
}

// String returns a string representation of the entire Deck
func (d Deck) String() string {
	ret := ""
	for _, c := range d.Cards {
		ret += c.String() + "\n"
	}
	return ret
}

func (d Deck) Shuffle() Deck {
	// shuffle the cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
	return d
}

// StandardDeck returns a standard 52 card deck for games like solitaire and poker
func StandardDeck() Deck {
	cards := generateStandardCards(suits, values, false)
	return Deck{Cards: cards}
}

// StandardDeckWithJokers returns a standard 52 card deck, plus two Jokers
func StandardDeckWithJokers() Deck {
	cards := generateStandardCards(suits, values, true)
	return Deck{Cards: cards}
}

func generateStandardCards(suits []string, values []value, jokers bool) []Card {
	newCards := []Card{}
	if jokers {
		newCards = append(newCards, Card{suit: "*", value: value{n: 0, s: "*", l: "Joker"}})
		newCards = append(newCards, Card{suit: "*", value: value{n: 0, s: "*", l: "Joker"}})
	}
	for _, s := range suits {
		for _, v := range values {
			newCards = append(newCards, Card{suit: s, value: v})
		}
	}
	return newCards
}
