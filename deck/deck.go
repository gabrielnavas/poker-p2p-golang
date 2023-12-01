package deck

import "math/rand"

type Deck [52]Card

func New() Deck {
	var (
		err           error
		numberOfSuits = 4
		numberOfCards = 13
		cardsIndex    = 0
		cards         = [52]Card{}
	)
	for suitIndex := 0; suitIndex < numberOfSuits; suitIndex++ {
		for cardIndex := 0; cardIndex < numberOfCards; cardIndex++ {
			cards[cardsIndex], err = NewCard(Suit(suitIndex), cardIndex+1)
			if err != nil {
				panic(err)
			}
			cardsIndex++
		}
	}
	return cards
}

func Shurffle(d Deck) Deck {
	for i := 0; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}
	return d
}
