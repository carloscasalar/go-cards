package card_test

import (
    "testing"

    "github.com/carloscasalar/go-cards/internal/domain/card"
    "github.com/stretchr/testify/assert"
)

func Test_Errors(t *testing.T) {
	t.Run("ErrInvalidCard", func(t *testing.T) {
		assert.Equal(t, "invalid card", card.ErrInvalidCard.Error())
	})
}
