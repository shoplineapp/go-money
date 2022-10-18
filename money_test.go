package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromAmount(t *testing.T) {
	twdM := NewFromAmount(100, "HKD")
	assert.Equal(t, int64(10000), twdM.Cents)

	jpyM := NewFromAmount(100, "JPY")
	assert.Equal(t, int64(100), jpyM.Cents)
}

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

	nm = m.Multiply(1.1)
	assert.Equal(t, int64(1), nm.Cents)

	nm = m.Multiply(1.5)
	assert.Equal(t, int64(2), nm.Cents)

	nm = m.Multiply(1.9)
	assert.Equal(t, int64(2), nm.Cents)

	nm = m.Multiply(2.1)
	assert.Equal(t, int64(2), nm.Cents)

	nm = m.Multiply(2.5)
	assert.Equal(t, int64(2), nm.Cents)

	nm = m.Multiply(2.9)
	assert.Equal(t, int64(3), nm.Cents)
}
