package deck_test

import (
	"testing"

	"github.com/carloscasalar/go-cards/v2/internal/domain/card"
	"github.com/carloscasalar/go-cards/v2/internal/domain/deck"
	"github.com/carloscasalar/go-cards/v2/pkg/ctypes"
	"github.com/stretchr/testify/assert"
)

var (
	numberOfCardsInDeck int
	expectedCards       []*ctypes.Card
)

func init() {
	numberOfCardsInDeck = 52
	expectedCards = []*ctypes.Card{}

	for _, suit := range ctypes.AllSuits() {
		for _, value := range ctypes.AllValues() {
			newCard := card.NewCard(value, suit)
			expectedCards = append(expectedCards, &newCard)
		}
	}
}

func TestDeck_NewDeck(t *testing.T) {
	newDeck := deck.NewDeck()

	assert.IsType(t, &deck.Deck{}, newDeck)
	assert.Equal(t, numberOfCardsInDeck, len(newDeck.Cards))
	assert.Equal(t, expectedCards, newDeck.Cards)
}

func TestDeck_Shuffle(t *testing.T) {
	deck := deck.NewDeck()
	deck.Shuffle()

	assert.Equal(t, numberOfCardsInDeck, len(deck.Cards))
	assert.NotEqual(t, expectedCards, deck.Cards)

	shuffledCards := make([]*ctypes.Card, len(deck.Cards))
	copy(shuffledCards, deck.Cards)

	// Result must be different from the first shuffle
	deck.Shuffle()
	assert.NotEqual(t, shuffledCards, deck.Cards)
}

func TestDeck_PopCard(t *testing.T) {
	deck := deck.NewDeck()

	c, err := deck.PopCard()

	assert.NoError(t, err)
	assert.Equal(t, numberOfCardsInDeck-1, len(deck.Cards))
	assert.Equal(t, expectedCards[len(expectedCards)-1], c)
}

func TestDeck_PopCard_Error(t *testing.T) {
	newDesck := deck.NewDeck()
	newDesck.Cards = []*ctypes.Card{}

	_, err := newDesck.PopCard()

	assert.Equal(t, deck.ErrNoCards, err)
}
