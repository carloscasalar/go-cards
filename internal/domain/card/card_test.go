package card_test

import (
    "testing"

    "github.com/carloscasalar/go-cards/internal/domain/card"
    "github.com/carloscasalar/go-cards/pkg/ctypes"
    "github.com/stretchr/testify/assert"
)

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card ctypes.Card
		want string
	}{
		{
			name: "Ace of Spade",
			card: card.NewCard(ctypes.Ace, ctypes.Spade),
			want: "A[♠]",
		},
		{
			name: "Jack of Diamond",
			card: card.NewCard(ctypes.Jack, ctypes.Diamond),
			want: "J[♦]",
		},
		{
			name: "Ten of Club",
			card: card.NewCard(ctypes.Ten, ctypes.Club),
			want: "10[♣]",
		},
		{
			name: "Two of Heart",
			card: card.NewCard(ctypes.Two, ctypes.Heart),
			want: "2[♥]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.card.String())
		})
	}
}
