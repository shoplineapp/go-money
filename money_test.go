package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMul(t *testing.T) {

	m := Money{
		Cents:       1,
		CurrencyIso: "USD",
	}

	nm := m.Multiply(0.8)
	assert.Equal(t, int64(1), nm.Cents)

	nm = m.Multiply(0.4)
	assert.Equal(t, int64(0), nm.Cents)

	nm = m.Multiply(2.6)
	assert.Equal(t, int64(3), nm.Cents)
}
