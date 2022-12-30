package money

import (
	gomoney "github.com/Rhymond/go-money"
)

func init() {
	// Need to change Currency TWD Fraction from 2 to 0 in /Rhymond/go-money
	gomoney.AddCurrency("HKD", "HK$", "$1", ".", ",", 2)
	gomoney.AddCurrency("CNY", "CN\u00a5", "$1", ".", ",", 2)
	gomoney.AddCurrency("TWD", "NT$", "$1", ".", ",", 0)
	gomoney.AddCurrency("USD", "US$", "$1", ".", ",", 2)
	gomoney.AddCurrency("SGD", "S$", "$1", ".", ",", 2)
	gomoney.AddCurrency("EUR", "\u20ac", "$1", ".", ",", 2)
	gomoney.AddCurrency("AUD", "A$", "$1", ".", ",", 2)
	gomoney.AddCurrency("GBP", "\u00a3", "$1", ".", ",", 2)
	gomoney.AddCurrency("PHP", "PHP", "$1", ".", ",", 2) // \u20b1
	gomoney.AddCurrency("MYR", "RM", "$1", ".", ",", 2)
	gomoney.AddCurrency("THB", "\u0e3f", "1 $", ".", ",", 2)
	gomoney.AddCurrency("AED", "DH", "1$", ".", ",", 2) //\u062f.\u0625
	gomoney.AddCurrency("JPY", "円", "$1", ".", ",", 0)  // \u00a5
	gomoney.AddCurrency("MMK", "K", "$1", ".", ",", 2)
	gomoney.AddCurrency("BND", "B$", "$1", ".", ",", 2)
	gomoney.AddCurrency("KRW", "\u20a9", "$1", ".", ",", 0)
	gomoney.AddCurrency("IDR", "Rp", "$ 1", ",", ".", 2)
	gomoney.AddCurrency("VND", "\u20ab", "1 $", ".", ",", 0)
	gomoney.AddCurrency("CAD", "C$", "$1", ".", ",", 2)
}
