package card

import (
	"fmt"

	"github.com/carloscasalar/go-cards/pkg/ctypes"
)

type Card struct {
	Value ctypes.Value
	Suit  ctypes.Suit
}

func (card *Card) String() string {
	return fmt.Sprintf("%s[%s]", card.Value, card.Suit)
}

func NewCard(value ctypes.Value, suit ctypes.Suit) *Card {
	card := &Card{Value: value, Suit: suit}
	return card
}
