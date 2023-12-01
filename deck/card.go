package deck

import (
	"errors"
	"fmt"
	"strconv"
)

type Card struct {
	suit  Suit
	value int
}

func NewCard(s Suit, v int) (Card, error) {
	if v > 13 {
		return Card{}, errors.New("the value of the card cannot be higher then 13")
	}
	if v < 1 {
		return Card{}, errors.New("the value of the card cannot be smaller then 1")
	}
	return Card{
		suit:  s,
		value: v,
	}, nil
}

func (c Card) String() string {
	value := strconv.Itoa(c.value)
	if c.value == 1 {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, c.suit, suitToUnicode(c.suit))
}
