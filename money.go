package money

import (
	"errors"
	"math"

	gomoney "github.com/Rhymond/go-money"
	"github.com/samber/lo"
)

const (
	RoundUp      = "ROUND_UP"
	RoundDown    = "ROUND_DOWN"
	RoundBankers = "ROUND_BANKERS"
)

var (
	// Error
	ErrorDivideByZero = errors.New("invalid operation: division by zero")
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

type MoneyOption func(*Money)

func WithRoundingMode(mode string) MoneyOption {
	return func(m *Money) {
		m.roundingMode = mode
	}
}

func New(cents int64, isoCode string, options ...MoneyOption) *Money {
	nm := gomoney.New(cents, isoCode)
	return newFromGoMoney(nm, options...)
}

func NewFromAmount(dollars float64, isoCode string, options ...MoneyOption) *Money {
	nm := gomoney.NewFromFloat(dollars, isoCode)
	return newFromGoMoney(nm, options...)
}

func newFromGoMoney(nm *gomoney.Money, options ...MoneyOption) *Money {
	money := &Money{
		money:          nm,
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    nm.Currency().Code,
		CurrencySymbol: nm.Currency().Grapheme,
		Label:          nm.Display(),
		roundingMode:   RoundBankers, // Default Round Mode will be RoundBankers
	}
	for _, option := range options {
		option(money)
	}
	return money
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

// Getting the round mode among all Money. If om is not exist, will return the first non-nil round mode. Otherwise will return RoundBankers
func alignRoundingMode(m *Money, ma []*Money) MoneyOption {
	if m.roundingMode != "" {
		return WithRoundingMode(m.roundingMode)
	}
	fm, isFound := lo.Find(ma, func(money *Money) bool {
		return money.roundingMode != ""
	})
	if isFound {
		return WithRoundingMode(fm.roundingMode)
	}
	return WithRoundingMode(RoundBankers)

}

// Rounded function. Default mode will be Banker rounding mode
func (m *Money) RoundByMode(value float64) float64 {
	switch m.roundingMode {
	case RoundUp:
		return math.Ceil(value)
	case RoundDown:
		return math.Floor(value)
	case RoundBankers:
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
		Label:          nm.Display(),
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
		Label:          nm.Display(),
		roundingMode:   m.roundingMode,
	}
}

// Add returns new Money struct with value representing sum of Self and Other Money.
func (m *Money) Add(oms ...*Money) (*Money, error) {
	m.initMoney()
	innerMoney := m.money
	var err error
	for _, om := range oms {
		om.initMoney()
		innerMoney, err = innerMoney.Add(om.money)
		if err != nil {
			return nil, err
		}
	}
	return New(innerMoney.Amount(), m.CurrencyIso, alignRoundingMode(m, oms)), nil

}

// Subtract returns new Money struct with value representing difference of Self and Other Money.
func (m *Money) Subtract(oms ...*Money) (*Money, error) {
	m.initMoney()
	innerMoney := m.money
	var err error
	for _, om := range oms {
		om.initMoney()
		innerMoney, err = innerMoney.Subtract(om.money)
		if err != nil {
			return nil, err
		}
	}
	return New(innerMoney.Amount(), m.CurrencyIso, alignRoundingMode(m, oms)), nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier. And If no rounding mode is setted, banker rounding mode is used
func (m *Money) Multiply(mul float64) *Money {
	m.initMoney()

	cents := m.money.Amount()
	newCents := float64(cents) * mul
	round := m.RoundByMode(newCents)

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
func (m *Money) Divide(div float64) (*Money, error) {
	if div == 0 {
		return nil, ErrorDivideByZero
	}
	m.initMoney()

	cents := m.money.Amount()
	newCents := float64(cents) / div
	round := m.RoundByMode(newCents)

	nm := gomoney.New(int64(round), m.CurrencyIso)

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
		roundingMode:   m.roundingMode,
	}, nil
}
