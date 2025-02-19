package deck

import (
	"math/rand"
	"time"

	"github.com/carloscasalar/go-cards/v2/internal/domain/card"
	"github.com/carloscasalar/go-cards/v2/pkg/ctypes"
)

type Deck struct {
	Cards []*ctypes.Card
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
}

func (deck *Deck) PopCard() (*ctypes.Card, error) {
	last := len(deck.Cards) - 1
	if last < 0 {
		return nil, ErrNoCards
	}
	card := deck.Cards[last]
	deck.Cards = deck.Cards[:last]
	return card, nil
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Cards = make([]*ctypes.Card, 0)
	for _, suit := range ctypes.AllSuits() {
		for _, value := range ctypes.AllValues() {
			newCard := card.NewCard(value, suit)
			deck.Cards = append(deck.Cards, &newCard)
		}
	}
	return deck
}
