package deck

import "testing"

func verify(t *testing.T, received, expected string) {
	t.Helper()
	if received != expected {
		t.Fatalf("expected '%v', received '%v'\n", expected, received)
	}
}

func TestDeckStringSuccess(t *testing.T) {
	type caseTest struct {
		suit           Suit
		numberOfCard   int
		expectedString string
	}

	tests := []caseTest{
		{Spades, 1, "ACE of SPADES ♠"},
		{Spades, 13, "13 of SPADES ♠"},
		{Spades, 14, "the value of the card cannot be higher then 13"},
		{Spades, 0, "the value of the card cannot be smaller then 1"},

		{Harts, 1, "ACE of HARTS ♥"},
		{Harts, 13, "13 of HARTS ♥"},
		{Harts, 14, "the value of the card cannot be higher then 13"},
		{Harts, 0, "the value of the card cannot be smaller then 1"},

		{Diamonds, 1, "ACE of DIAMONDS ♦"},
		{Diamonds, 13, "13 of DIAMONDS ♦"},
		{Diamonds, 14, "the value of the card cannot be higher then 13"},
		{Diamonds, 0, "the value of the card cannot be smaller then 1"},

		{Clubs, 1, "ACE of CLUBS ♣"},
		{Clubs, 13, "13 of CLUBS ♣"},
		{Clubs, 14, "the value of the card cannot be higher then 13"},
		{Clubs, 0, "the value of the card cannot be smaller then 1"},
	}

	for _, test := range tests {
		received, err := NewCard(test.suit, test.numberOfCard)
		if err != nil {
			verify(t, err.Error(), test.expectedString)
		} else {
			verify(t, received.String(), test.expectedString)
		}
	}
}
