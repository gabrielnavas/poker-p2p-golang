package deck

import "testing"

func TestSuitToUnicode(t *testing.T) {
	type caseTest struct {
		suit           Suit
		expectedString string
	}

	tests := []caseTest{
		{Spades, "♠"},
		{Harts, "♥"},
		{Diamonds, "♦"},
		{Clubs, "♣"},
	}

	for _, test := range tests {
		received := suitToUnicode(test.suit)
		verify(t, received, test.expectedString)
	}
}
