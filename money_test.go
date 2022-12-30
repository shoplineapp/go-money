package money

import (
	"math"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	m := New(100, "TWD", WithRoundingMode(RoundUp))
	assert.Equal(t, int64(100), m.Cents)
	assert.Equal(t, RoundUp, m.GetRoundingMode())
	assert.Equal(t, float64(100), m.Dollars)
	assert.Equal(t, "TWD", m.CurrencyIso)
	assert.Equal(t, "NT$", m.CurrencySymbol)
	assert.Equal(t, "NT$100", m.Label)
}

func TestFromAmount(t *testing.T) {
	// With Default
	hkdM := NewFromAmount(100, "HKD")
	assert.Equal(t, int64(10000), hkdM.Cents)
	assert.Equal(t, RoundBankers, hkdM.GetRoundingMode())
	assert.Equal(t, "HKD", hkdM.CurrencyIso)
	assert.Equal(t, "HK$", hkdM.CurrencySymbol)
	assert.Equal(t, "HK$100.00", hkdM.Label)

	jpyM := NewFromAmount(100, "JPY")
	assert.Equal(t, int64(100), jpyM.Cents)

	// With rounding mode
	usdM := NewFromAmount(28.55, "USD", WithRoundingMode(RoundUp))
	assert.Equal(t, int64(2855), usdM.Cents)
	assert.Equal(t, 28.55, usdM.Dollars)
	assert.Equal(t, RoundUp, usdM.GetRoundingMode())
	assert.Equal(t, "US$28.55", usdM.Label)

	twdM := NewFromAmount(28.55, "TWD", WithRoundingMode(RoundUp))
	assert.Equal(t, int64(29), twdM.Cents)
	assert.Equal(t, float64(29), twdM.Dollars)
	assert.Equal(t, RoundUp, twdM.GetRoundingMode())
	assert.Equal(t, "NT$29", twdM.Label)
}

func TestSetRoundingMode(t *testing.T) {
	m := New(100, "TWD", WithRoundingMode(RoundUp))
	m.SetRoundingMode(RoundDown)
	assert.Equal(t, RoundDown, m.roundingMode)
}

func TestGetRoundingMode(t *testing.T) {
	m := New(100, "TWD", WithRoundingMode(RoundUp))
	m.GetRoundingMode()
	assert.Equal(t, RoundUp, m.roundingMode)
}

func TestInitMoney(t *testing.T) {
	m := New(100, "TWD", WithRoundingMode(RoundUp))
	m.initMoney()
	assert.Equal(t, "TWD", m.money.Currency().Code)
}

func TestAlignRoundingMode(t *testing.T) {
	testTable := []struct {
		mainMoneyRoundingMode  string
		paramMoneyRoundingMode []string
		expected               string
	}{
		{
			mainMoneyRoundingMode:  RoundUp,
			paramMoneyRoundingMode: []string{RoundDown, RoundDown, RoundDown},
			expected:               RoundUp,
		},
		{
			mainMoneyRoundingMode:  "",
			paramMoneyRoundingMode: []string{"", RoundDown, RoundUp},
			expected:               RoundDown,
		},
		{
			mainMoneyRoundingMode:  "",
			paramMoneyRoundingMode: []string{"", "", ""},
			expected:               RoundBankers,
		},
	}
	for _, item := range testTable {
		m := New(1, "TWD", WithRoundingMode(item.mainMoneyRoundingMode))
		ma := lo.Map(item.paramMoneyRoundingMode, func(roundMode string, _ int) *Money {
			return New(1, "TWD", WithRoundingMode(roundMode))
		})
		result := New(1, "TWD", alignRoundingMode(m, ma))
		assert.Equal(t, item.expected, result.GetRoundingMode())
	}

}

func TestRoundByMode(t *testing.T) {
	testTable := []struct {
		roundingMode string
		inputValue   float64
		expected     float64
	}{
		{
			roundingMode: RoundUp,
			inputValue:   1.5,
			expected:     math.Ceil(1.5),
		},
		{
			roundingMode: RoundUp,
			inputValue:   1.4,
			expected:     math.Ceil(1.4),
		},
		{
			roundingMode: RoundDown,
			inputValue:   1.5,
			expected:     math.Floor(1.5),
		},
		{
			roundingMode: RoundDown,
			inputValue:   1.4,
			expected:     math.Floor(1.4),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   1.5,
			expected:     math.RoundToEven(1.5),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   1.4,
			expected:     math.RoundToEven(1.4),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   2.5,
			expected:     math.RoundToEven(2.5),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   2.4,
			expected:     math.RoundToEven(2.4),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   1.5,
			expected:     math.RoundToEven(1.5),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   1.4,
			expected:     math.RoundToEven(1.4),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   2.5,
			expected:     math.RoundToEven(2.5),
		},
		{
			roundingMode: RoundBankers,
			inputValue:   2.4,
			expected:     math.RoundToEven(2.4),
		},
	}
	for _, item := range testTable {
		rd := New(1, "TWD", WithRoundingMode(item.roundingMode))
		assert.Equal(t, item.expected, rd.RoundByMode(item.inputValue))
	}

}

func TestDisplay(t *testing.T) {
	testTable := []struct {
		cents    int64
		currency string
		expected string
	}{
		{
			cents:    100000,
			currency: "HKD",
			expected: "HK$1,000.00",
		},
		{
			cents:    100000,
			currency: "CNY",
			expected: "CN¥1,000.00",
		},
		{
			cents:    100000,
			currency: "TWD",
			expected: "NT$100,000",
		},
		{
			cents:    100000,
			currency: "USD",
			expected: "US$1,000.00",
		},
		{
			cents:    100000,
			currency: "SGD",
			expected: "S$1,000.00",
		},
		{
			cents:    100000,
			currency: "EUR",
			expected: "€1,000.00",
		},
		{
			cents:    100000,
			currency: "AUD",
			expected: "A$1,000.00",
		},
		{
			cents:    100000,
			currency: "GBP",
			expected: "£1,000.00",
		},
		{
			cents:    100000,
			currency: "PHP",
			expected: "PHP1,000.00",
		},
		{
			cents:    100000,
			currency: "MYR",
			expected: "RM1,000.00",
		},
		{
			cents:    100000,
			currency: "THB",
			expected: "1,000.00 ฿",
		},
		{
			cents:    100000,
			currency: "AED",
			expected: "1,000.00DH",
		},
		{
			cents:    100000,
			currency: "JPY",
			expected: "円100,000",
		},
		{
			cents:    100000,
			currency: "MMK",
			expected: "K1,000.00",
		},
		{
			cents:    100000,
			currency: "BND",
			expected: "B$1,000.00",
		},
		{
			cents:    100000,
			currency: "KRW",
			expected: "₩100,000",
		},
		{
			cents:    100000,
			currency: "IDR",
			expected: "Rp 1.000,00",
		},
		{
			cents:    100000,
			currency: "VND",
			expected: "100,000 ₫",
		},
		{
			cents:    100000,
			currency: "CAD",
			expected: "C$1,000.00",
		},
	}
	for _, item := range testTable {
		m := New(item.cents, item.currency, WithRoundingMode(RoundDown))
		nm := m.Display()
		assert.Equal(t, item.expected, nm, item.currency)
	}
}

func TestEqual_SameCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    0,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    -1,
			currency: "TWD",
			expected: false,
		},
	}
	for _, item := range testTable {
		m2 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm, err := m1.Equals(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm)

	}
}

func TestEqual_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.Equals(m2)
	assert.Error(t, err)
	assert.Equal(t, false, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestGreaterThan_SameCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: false,
		},
	}
	for _, item := range testTable {
		m2 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm, err := m1.GreaterThan(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm)

	}
}

func TestGreaterThan_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.GreaterThan(m2)
	assert.Error(t, err, err)
	assert.Equal(t, false, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestGreaterThanOrEqual_SameCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: false,
		},
	}
	for _, item := range testTable {
		m2 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm, err := m1.GreaterThanOrEqual(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm)

	}
}

func TestGreaterThanOrEqual_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.GreaterThanOrEqual(m2)
	assert.Error(t, err, err)
	assert.Equal(t, false, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestLessThan_SameCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: true,
		},
	}
	for _, item := range testTable {
		m2 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm, err := m1.LessThan(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm)

	}
}

func TestLessThan_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.LessThan(m2)
	assert.Error(t, err, err)
	assert.Equal(t, false, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestLessThanOrEqual_SameCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: true,
		},
	}
	for _, item := range testTable {
		m2 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm, err := m1.LessThanOrEqual(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm)

	}
}

func TestLessThanOrEqual_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.LessThanOrEqual(m2)
	assert.Error(t, err, err)
	assert.Equal(t, false, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestIsZero(t *testing.T) {
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: false,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm := m1.IsZero()
		assert.Equal(t, item.expected, nm)
	}
}

func TestIsNegative(t *testing.T) {
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: true,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: false,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm := m1.IsNegative()
		assert.Equal(t, item.expected, nm)
	}
}

func TestIsPositive(t *testing.T) {
	testTable := []struct {
		cents    float64
		currency string
		expected bool
	}{
		{
			cents:    -1,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    0,
			currency: "TWD",
			expected: false,
		},
		{
			cents:    1,
			currency: "TWD",
			expected: true,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents), item.currency, WithRoundingMode(RoundDown))
		nm := m1.IsPositive()
		assert.Equal(t, item.expected, nm)
	}
}

func TestAbsolute(t *testing.T) {
	testTable := []struct {
		cents    float64
		expected int64
	}{
		{
			cents:    -1,
			expected: 1,
		},
		{
			cents:    0,
			expected: 0,
		},
		{
			cents:    1,
			expected: 1,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents), "TWD", WithRoundingMode(RoundDown))
		nm := m1.Absolute()
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestNegative(t *testing.T) {
	testTable := []struct {
		cents    float64
		expected int64
	}{
		{
			cents:    -1,
			expected: -1,
		},
		{
			cents:    0,
			expected: -0,
		},
		{
			cents:    1,
			expected: -1,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents), "TWD", WithRoundingMode(RoundDown))
		nm := m1.Negative()
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestAdd_SingleValue(t *testing.T) {
	testTable := []struct {
		cents1   int64
		cents2   int64
		currency string
		expected int64
	}{
		{
			cents1:   5,
			cents2:   5,
			expected: 10,
		},
		{
			cents1:   10,
			cents2:   5,
			expected: 15,
		},
		{
			cents1:   1,
			cents2:   -1,
			expected: 0,
		},
		{
			cents1:   -1,
			cents2:   -1,
			expected: -2,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents1), "TWD", WithRoundingMode(RoundDown))
		m2 := New(int64(item.cents2), "TWD", WithRoundingMode(RoundDown))
		nm, err := m1.Add(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestAdd_MultiValue(t *testing.T) {
	testTable := []struct {
		centsArray []int64
		currency   string
		expected   int64
	}{
		{
			centsArray: []int64{1, 2, 3, 4},
			expected:   10,
		},
		{
			centsArray: []int64{1, 0, -1, 1},
			expected:   1,
		},
		{
			centsArray: []int64{-1, -2, -3, -4},
			expected:   -10,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(0), "TWD", WithRoundingMode(RoundDown))
		// nm, err := m1.Add(m2)
		moneys := lo.Map(item.centsArray, func(cents int64, _ int) *Money {
			return New(cents, "TWD", WithRoundingMode(RoundDown))
		})
		nm, err := m1.Add(moneys...)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestAdd_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.Add(m2)
	assert.Error(t, err)
	assert.Nil(t, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestSubtract_SingleValue(t *testing.T) {
	testTable := []struct {
		cents1   int64
		cents2   int64
		currency string
		expected int64
	}{
		{
			cents1:   5,
			cents2:   5,
			expected: 0,
		},
		{
			cents1:   5,
			cents2:   -5,
			expected: 10,
		},
		{
			cents1:   -5,
			cents2:   5,
			expected: -10,
		},
		{
			cents1:   -5,
			cents2:   -5,
			expected: -0,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(item.cents1), "TWD", WithRoundingMode(RoundDown))
		m2 := New(int64(item.cents2), "TWD", WithRoundingMode(RoundDown))
		nm, err := m1.Subtract(m2)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestSubtract_MultiValue(t *testing.T) {
	testTable := []struct {
		centsArray []int64
		currency   string
		expected   int64
	}{
		{
			centsArray: []int64{1, 2, 3, 4},
			expected:   0,
		},
		{
			centsArray: []int64{-1, -2, -3, -4},
			expected:   20,
		},
		{
			centsArray: []int64{-1, -2, 1, 2},
			expected:   10,
		},
	}
	for _, item := range testTable {
		m1 := New(int64(10), "TWD", WithRoundingMode(RoundDown))
		moneys := lo.Map(item.centsArray, func(cents int64, _ int) *Money {
			return New(cents, "TWD", WithRoundingMode(RoundDown))
		})
		nm, err := m1.Subtract(moneys...)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestSubtract_DifferentCurrencies(t *testing.T) {
	m1 := New(0, "TWD", WithRoundingMode(RoundDown))
	m2 := New(0, "USD", WithRoundingMode(RoundDown))
	nm, err := m1.Subtract(m2)
	assert.Error(t, err)
	assert.Nil(t, nm)
	assert.Equal(t, "currencies don't match", err.Error())
}

func TestMultiply(t *testing.T) {
	testTable := []struct {
		multiplier float64
		roundMode  string
		expected   int64
	}{
		{
			multiplier: 0.4,
			roundMode:  RoundUp,
			expected:   1,
		},
		{
			multiplier: 0.4,
			roundMode:  RoundDown,
			expected:   0,
		},
		{
			multiplier: 0.4,
			roundMode:  RoundBankers,
			expected:   0,
		},
		{
			multiplier: 0.4,
			roundMode:  "unknown",
			expected:   0,
		},
		{
			multiplier: 0.5,
			roundMode:  RoundUp,
			expected:   1,
		},
		{
			multiplier: 0.5,
			roundMode:  RoundDown,
			expected:   0,
		},
		{
			multiplier: 0.5,
			roundMode:  RoundBankers,
			expected:   0,
		},
		{
			multiplier: 0.5,
			roundMode:  "unknown",
			expected:   0,
		},
		{
			multiplier: 0.6,
			roundMode:  RoundUp,
			expected:   1,
		},
		{
			multiplier: 0.6,
			roundMode:  RoundDown,
			expected:   0,
		},
		{
			multiplier: 0.6,
			roundMode:  RoundBankers,
			expected:   1,
		},
		{
			multiplier: 0.6,
			roundMode:  "unknown",
			expected:   1,
		},
		{
			multiplier: 1.4,
			roundMode:  RoundUp,
			expected:   2,
		},
		{
			multiplier: 1.4,
			roundMode:  RoundDown,
			expected:   1,
		},
		{
			multiplier: 1.4,
			roundMode:  RoundBankers,
			expected:   1,
		},
		{
			multiplier: 1.4,
			roundMode:  "unknown",
			expected:   1,
		},
		{
			multiplier: 1.5,
			roundMode:  RoundUp,
			expected:   2,
		},
		{
			multiplier: 1.5,
			roundMode:  RoundDown,
			expected:   1,
		},
		{
			multiplier: 1.5,
			roundMode:  RoundBankers,
			expected:   2,
		},
		{
			multiplier: 1.5,
			roundMode:  "unknown",
			expected:   2,
		},
		{
			multiplier: 1.6,
			roundMode:  RoundUp,
			expected:   2,
		},
		{
			multiplier: 1.6,
			roundMode:  RoundDown,
			expected:   1,
		},
		{
			multiplier: 1.6,
			roundMode:  RoundBankers,
			expected:   2,
		},
		{
			multiplier: 1.6,
			roundMode:  "unknown",
			expected:   2,
		},
		{
			multiplier: 2.4,
			roundMode:  RoundUp,
			expected:   3,
		},
		{
			multiplier: 2.4,
			roundMode:  RoundDown,
			expected:   2,
		},
		{
			multiplier: 2.4,
			roundMode:  RoundBankers,
			expected:   2,
		},
		{
			multiplier: 2.4,
			roundMode:  "unknown",
			expected:   2,
		},
		{
			multiplier: 2.5,
			roundMode:  RoundUp,
			expected:   3,
		},
		{
			multiplier: 2.5,
			roundMode:  RoundDown,
			expected:   2,
		},
		{
			multiplier: 2.5,
			roundMode:  RoundBankers,
			expected:   2,
		},
		{
			multiplier: 2.5,
			roundMode:  "unknown",
			expected:   2,
		},
		{
			multiplier: 2.6,
			roundMode:  RoundUp,
			expected:   3,
		},
		{
			multiplier: 2.6,
			roundMode:  RoundDown,
			expected:   2,
		},
		{
			multiplier: 2.6,
			roundMode:  RoundBankers,
			expected:   3,
		},
		{
			multiplier: 2.6,
			roundMode:  "unknown",
			expected:   3,
		},
	}
	for _, item := range testTable {
		m := New(int64(1), "HKD", WithRoundingMode(item.roundMode))
		nm := m.Multiply(item.multiplier)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestDivide_NoError(t *testing.T) {
	testTable := []struct {
		dividend  float64
		roundMode string
		expected  int64
	}{
		{
			dividend:  4,
			roundMode: RoundUp,
			expected:  1,
		},
		{
			dividend:  4,
			roundMode: RoundDown,
			expected:  0,
		},
		{
			dividend:  4,
			roundMode: RoundBankers,
			expected:  0,
		},
		{
			dividend:  4,
			roundMode: "unknown",
			expected:  0,
		},
		{
			dividend:  5,
			roundMode: RoundUp,
			expected:  1,
		},
		{
			dividend:  5,
			roundMode: RoundDown,
			expected:  0,
		},
		{
			dividend:  5,
			roundMode: RoundBankers,
			expected:  0,
		},
		{
			dividend:  5,
			roundMode: "unknown",
			expected:  0,
		},
		{
			dividend:  6,
			roundMode: RoundUp,
			expected:  1,
		},
		{
			dividend:  6,
			roundMode: RoundDown,
			expected:  0,
		},
		{
			dividend:  6,
			roundMode: RoundBankers,
			expected:  1,
		},
		{
			dividend:  6,
			roundMode: "unknown",
			expected:  1,
		},
	}
	for _, item := range testTable {
		m := New(int64(item.dividend), "HKD", WithRoundingMode(item.roundMode))
		nm, err := m.Divide(10)
		assert.NoError(t, err)
		assert.Equal(t, item.expected, nm.Cents)
	}
}

func TestDivide_WithError(t *testing.T) {
	m := New(int64(1), "HKD", WithRoundingMode(RoundUp))
	_, err := m.Divide(0)
	assert.Error(t, err)
	assert.ErrorIs(t, ErrorDivideByZero, err)
}
