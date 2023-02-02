package money

import (
	gomoney "github.com/Rhymond/go-money"
)

type Currency struct {
	internalCurrency     *gomoney.Currency
	smallestDenomination int32
}

var currencies = map[string]*Currency{}

func setCurrency(currencies map[string]*Currency, currency *gomoney.Currency, smallestDenomination int32) {
	currencies[currency.Code] = &Currency{
		internalCurrency:     currency,
		smallestDenomination: smallestDenomination,
	}
}

func getCurrency(code string) *Currency {
	if currency, ok := currencies[code]; ok {
		return currency
	}
	return nil
}

func init() {
	// Need to change Currency TWD Fraction from 2 to 0 in /Rhymond/go-money
	setCurrency(currencies, gomoney.AddCurrency("HKD", "HK$", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("TWD", "NT$", "$1", ".", ",", 0), 1)
	setCurrency(currencies, gomoney.AddCurrency("USD", "US$", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("SGD", "S$", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("EUR", "\u20ac", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("AUD", "A$", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("GBP", "\u00a3", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("PHP", "PHP", "$1", ".", ",", 2), 1) // \u20b1
	setCurrency(currencies, gomoney.AddCurrency("MYR", "RM", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("THB", "\u0e3f", "1 $", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("AED", "DH", "1$", ".", ",", 2), 1) //\u062f.\u0625
	setCurrency(currencies, gomoney.AddCurrency("JPY", "円", "$1", ".", ",", 0), 1)  // \u00a5
	setCurrency(currencies, gomoney.AddCurrency("MMK", "K", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("BND", "B$", "$1", ".", ",", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("KRW", "\u20a9", "$1", ".", ",", 0), 1)
	setCurrency(currencies, gomoney.AddCurrency("IDR", "Rp", "$ 1", ",", ".", 2), 1)
	setCurrency(currencies, gomoney.AddCurrency("VND", "\u20ab", "1 $", ".", ",", 0), 1)
	setCurrency(currencies, gomoney.AddCurrency("CAD", "C$", "$1", ".", ",", 2), 1)
}
