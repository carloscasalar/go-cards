package dealer_test

import (
	"github.com/carloscasalar/go-cards/v2/pkg/ctypes"
	"testing"

	"github.com/carloscasalar/go-cards/v2/pkg/dealer"
	"github.com/stretchr/testify/assert"
)

func TestDealer_NewDealer(t *testing.T) {
	numberOfDecks := uint8(1)
	d1 := dealer.NewDealer(numberOfDecks)

	assert.Equal(t, 52, len(d1.Deck.Cards))

	d2 := dealer.NewDealer(numberOfDecks * 2)
	assert.Equal(t, 104, len(d2.Deck.Cards))
}

func TestDealer_ShuffleCards(t *testing.T) {
	numberOfDecks := uint8(1)
	d1 := dealer.NewDealer(numberOfDecks)

	var copyDeckCards []*ctypes.Card
	copyDeckCards = append(copyDeckCards, d1.Deck.Cards...)

	d1.ShuffleCards()

	assert.NotEqual(t, copyDeckCards, d1.Deck.Cards)
}

func TestDealer_Deal(t *testing.T) {
	numberOfDecks := uint8(1)
	d1 := dealer.NewDealer(numberOfDecks)

	c, err := d1.Deal()

	assert.NoError(t, err)
	assert.Equal(t, 51, len(d1.Deck.Cards))
	assert.NotEmpty(t, c)
}

func TestDealer_Deal_Error(t *testing.T) {
	numberOfDecks := uint8(1)
	d1 := dealer.NewDealer(numberOfDecks)

	for i := 0; i < 52; i++ {
		_, err := d1.Deal()
		assert.NoError(t, err)
	}

	_, err := d1.Deal()
	assert.Error(t, err)
}

func TestDealer_BurnCard(t *testing.T) {
	numberOfDecks := uint8(1)
	d1 := dealer.NewDealer(numberOfDecks)

	d1.BurnCard()

	assert.Equal(t, 51, len(d1.Deck.Cards))
}
