package money

import (
	gomoney "github.com/Rhymond/go-money"
)

func init() {
	// Need to change Currency TWD Fraction from 2 to 0 in /Rhymond/go-money
	gomoney.AddCurrency("HKD", "$", "$1", ".", ",", 2)
	gomoney.AddCurrency("CNY", "\u00a5", "$1", ".", ",", 2)
	gomoney.AddCurrency("TWD", "NT$", "$1", ".", ",", 0)
	gomoney.AddCurrency("USD", "$", "$1", ".", ",", 2)
	gomoney.AddCurrency("SGD", "$", "$1", ".", ",", 2)
	gomoney.AddCurrency("EUR", "\u20ac", "$1", ".", ",", 2)
	gomoney.AddCurrency("AUD", "$", "$1", ".", ",", 2)
	gomoney.AddCurrency("GBP", "\u00a3", "$1", ".", ",", 2)
	gomoney.AddCurrency("PHP", "\u20b1", "$1", ".", ",", 2)
	gomoney.AddCurrency("MYR", "RM", "$1", ".", ",", 2)
	gomoney.AddCurrency("THB", "\u0e3f", "1 $", ".", ",", 2)
	gomoney.AddCurrency("AED", "\u062f.\u0625", "1$", ".", ",", 2)
	gomoney.AddCurrency("JPY", "\u00a5", "$1", ".", ",", 0)
	gomoney.AddCurrency("MMK", "K", "$1", ".", ",", 2)
	gomoney.AddCurrency("BND", "$", "$1", ".", ",", 2)
	gomoney.AddCurrency("KRW", "\u20a9", "$1", ".", ",", 0)
	gomoney.AddCurrency("IDR", "Rp", "$ 1", ",", ".", 2)
	gomoney.AddCurrency("VND", "\u20ab", "1 $", ".", ",", 0)
	gomoney.AddCurrency("CAD", "$", "$1", ".", ",", 2)
}
