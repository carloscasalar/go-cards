package dealer

import (
	"github.com/carloscasalar/go-cards/internal/domain/deck"
    "github.com/carloscasalar/go-cards/pkg/ctypes"
)

type Dealer struct {
	Deck deck.Deck
}

func (dealer *Dealer) ShuffleCards() {
	dealer.Deck.Shuffle()
}

func (dealer *Dealer) Deal() (*ctypes.Card, error) {
	card, err := dealer.Deck.PopCard()
	if err != nil {
		return nil, deck.ErrNoCards
	}
	return card, nil
}

func (dealer *Dealer) BurnCard() {
	_, _ = dealer.Deal()
}

func NewDealer(numberOfDecks uint8) *Dealer {
	dealer := &Dealer{}
	for i := uint8(0); i < numberOfDecks; i++ {
		dealer.Deck.Cards = append(dealer.Deck.Cards, deck.NewDeck().Cards...)
	}
	return dealer
}
