package deck_test

import (
	"fmt"
	"ggpoker/deck"
	"testing"
)

func verify(t *testing.T, received, expected string) {
	t.Helper()
	if received != expected {
		t.Fatalf("expected %s, received %s\n", expected, received)
	}
}

func TestNewDeck(t *testing.T) {
	d := deck.New()
	var (
		cardsIndex    = 0
		numberOfCards = 13
	)

	expected := "ACE of SPADES ♠"
	verify(t, d[cardsIndex].String(), expected)
	cardsIndex++
	for cardIndex := 1; cardIndex < numberOfCards; cardIndex++ {
		expected := fmt.Sprintf("%d of SPADES ♠", cardIndex+1)
		verify(t, d[cardsIndex].String(), expected)
		cardsIndex++
	}

	expected = "ACE of HARTS ♥"
	cardsIndex++
	verify(t, d[13].String(), expected)
	for cardIndex := 1; cardIndex < numberOfCards; cardIndex++ {
		expected = fmt.Sprintf("%d of HARTS ♥", cardIndex+1)
		verify(t, d[cardsIndex].String(), expected)
		cardsIndex++
	}

	expected = "ACE of DIAMONDS ♦"
	verify(t, d[cardsIndex].String(), expected)
	cardsIndex++
	for cardIndex := 1; cardIndex < numberOfCards; cardIndex++ {
		expected = fmt.Sprintf("%d of DIAMONDS ♦", cardIndex+1)
		verify(t, d[cardsIndex].String(), expected)
		cardsIndex++
	}

	expected = "ACE of CLUBS ♣"
	verify(t, d[cardsIndex].String(), expected)
	cardsIndex++
	for cardIndex := 1; cardIndex < numberOfCards; cardIndex++ {
		expected = fmt.Sprintf("%d of CLUBS ♣", cardIndex+1)
		verify(t, d[cardsIndex].String(), expected)
		cardsIndex++
	}
}
