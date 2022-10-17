package money

import (
	"math"

	gomoney "github.com/Rhymond/go-money"
	"github.com/samber/lo"
)

func init() {
	// Need to change Currency TWD Fraction from 2 to 0 in /Rhymond/go-money
	gomoney.AddCurrency("TWD", "NT$", "$1", ".", ",", 0)
}

const (
	RoundUp     = "ROUND_UP"
	RoundDown   = "ROUND_DOWN"
	BankerRound = "ROUND_HALF_EVEN"
)

type Money struct {
	Cents          int64   `json:"cents" bson:"cents"`
	CurrencySymbol string  `json:"currency_symbol" bson:"currency_symbol"`
	CurrencyIso    string  `json:"currency_iso" bson:"currency_iso"`
	Label          string  `json:"label" bson:"label"`
	Dollars        float64 `json:"dollars" bson:"dollars"`

	roundingMode string
	money        *gomoney.Money
}

func New(cents int64, isoCode string, roundingMode string) *Money {
	nm := gomoney.New(cents, isoCode)
	return &Money{
		money:          nm,
		Cents:          cents,
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    isoCode,
		CurrencySymbol: nm.Currency().Grapheme,
		roundingMode:   roundingMode,
	}
}

// Setting the roundingMode of the money object
func (m *Money) SetRoundingMode(mode string) {
	m.roundingMode = mode
}

// Getting the roundingMode of the money object
func (m *Money) GetRoundingMode() string {
	return m.roundingMode
}

func (m *Money) initMoney() {
	if m.money == nil {
		m.money = gomoney.New(m.Cents, m.CurrencyIso)
	}
}

// Getting the roound mode among all Money. If om is not exist, will return the first non-nil round mode. Otherwise will return ""
func (m *Money) getRoundingModeAmongMoneys(ma []*Money) string {
	if m.roundingMode != "" {
		return m.roundingMode
	}
	fm, _, isFound := lo.FindIndexOf(ma, func(money *Money) bool {
		return money.roundingMode != ""
	})
	if isFound {
		return fm.roundingMode
	}
	return ""

}

// Rounded function. Default mode will be Banker rounding mode
func (m *Money) roundingByModeSet(value float64) float64 {
	switch m.roundingMode {
	case RoundUp:
		return math.Ceil(value)
	case RoundDown:
		return math.Floor(value)
	case BankerRound:
		return math.RoundToEven(value)
	default:
		return math.RoundToEven(value)

	}
}

func (m *Money) Display() string {
	return m.money.Display()
}

// Equals checks equality between two Money types.
func (m *Money) Equals(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.Equals(om.money)
}

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.GreaterThan(om.money)
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.GreaterThanOrEqual(om.money)
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.LessThan(om.money)
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.LessThanOrEqual(om.money)
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool {
	m.initMoney()
	return m.money.IsZero()
}

// IsPositive returns boolean of whether the value of Money is positive.
func (m *Money) IsPositive() bool {
	m.initMoney()
	return m.money.IsPositive()
}

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool {
	m.initMoney()
	return m.money.IsNegative()
}

// Absolute returns new Money struct from given Money using absolute monetary value.
func (m *Money) Absolute() *Money {
	m.initMoney()
	nm := m.money.Absolute()
	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
		roundingMode:   m.roundingMode,
	}
}

// Negative returns new Money struct from given Money using negative monetary value.
func (m *Money) Negative() *Money {
	m.initMoney()
	nm := m.money.Negative()
	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
		roundingMode:   m.roundingMode,
	}
}

// Add returns new Money struct with value representing sum of Self and Other Money.
func (m *Money) Add(oms ...*Money) (*Money, error) {
	result := m
	for _, om := range oms {
		result.initMoney()
		om.initMoney()
		nm, err := result.money.Add(om.money)
		if err != nil {
			return nil, err
		}
		result = &Money{
			Cents:          nm.Amount(),
			Dollars:        nm.AsMajorUnits(),
			CurrencyIso:    m.CurrencyIso,
			CurrencySymbol: m.CurrencySymbol,
			roundingMode:   m.getRoundingModeAmongMoneys(oms),
		}
	}
	return result, nil
}

// Subtract returns new Money struct with value representing difference of Self and Other Money.
func (m *Money) Subtract(oms ...*Money) (*Money, error) {
	result := m
	for _, om := range oms {
		result.initMoney()
		om.initMoney()
		nm, err := result.money.Subtract(om.money)
		if err != nil {
			return nil, err
		}

		result = &Money{
			Cents:          nm.Amount(),
			Dollars:        nm.AsMajorUnits(),
			CurrencyIso:    m.CurrencyIso,
			CurrencySymbol: m.CurrencySymbol,
			roundingMode:   m.getRoundingModeAmongMoneys(oms),
		}
	}
	return result, nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier. And If no rounding mode is setted, banker rounding mode is used
func (m *Money) Multiply(mul float64) *Money {
	m.initMoney()

	cents := m.money.Amount()
	newCents := float64(cents) * mul
	round := m.roundingByModeSet(newCents)

	nm := gomoney.New(int64(round), m.CurrencyIso)

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
		roundingMode:   m.roundingMode,
	}
}

// Divide returns new Money struct with value representing Self divided value by dividsor. And If no rounding mode is setted, banker rounding mode is used
func (m *Money) Divide(div float64) *Money {
	m.initMoney()

	cents := m.money.Amount()
	newCents := float64(cents) / div
	round := m.roundingByModeSet(newCents)

	nm := gomoney.New(int64(round), m.CurrencyIso)

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
		roundingMode:   m.roundingMode,
	}
}
