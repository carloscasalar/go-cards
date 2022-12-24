package card_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/carloscasalar/go-cards/internal/domain/card"
)

func Test_Errors(t *testing.T) {
	t.Run("ErrInvalidCard", func(t *testing.T) {
		assert.Equal(t, "invalid card", card.ErrInvalidCard.Error())
	})
}
