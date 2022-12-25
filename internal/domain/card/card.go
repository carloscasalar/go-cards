package card

import (
	"fmt"

	"github.com/carloscasalar/go-cards/pkg/ctypes"
)
type card struct {
	Value ctypes.Value
	Suit  ctypes.Suit
}

func (c *card) String() string {
	return fmt.Sprintf("%s[%s]", c.Value, c.Suit)
}

func NewCard(value ctypes.Value, suit ctypes.Suit) ctypes.Card {
	card :=&card{Value: value, Suit: suit}
	return card
}
