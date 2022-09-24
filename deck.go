package deck

type value struct {
	n int
	s string
	l string
}

var suits = []string{"Hearts", "Diamonds", "Spades", "Clubs"}

var values = []value{
	// {n: 0, s: "*", l: "Joker"},
	{n: 1, s: "1", l: "One"},
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

// Deck represents a collection of cards that can be leveraged to play games
type Deck struct {
	cards []Card
}

func StandardDeck() Deck {
	cards := generateStandardCards(suits, values, false)
	return Deck{cards: cards}
}

func generateStandardCards(suits []string, values []value, jokers bool) []Card {
	newCards := []Card{}
	for _, s := range suits {
		for _, v := range values {
			newCards = append(newCards, Card{suit: s, value: v})
		}
	}
	if jokers {
		newCards = append(newCards, Card{suit: "*", value: value{n: 0, s: "*", l: "Joker"}})
		newCards = append(newCards, Card{suit: "*", value: value{n: 0, s: "*", l: "Joker"}})
	}
	return newCards
}
