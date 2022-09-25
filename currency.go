package money

import (
	"errors"
	"strings"
)

type Currency struct {
	Code string

	config currencyConfig
}

type currencyConfig struct {
	priority             int
	isoCode              string
	name                 string
	symbol               string
	disambiguateSymbol   string
	alternateSymbols     []string
	subunit              string
	subunitToUnit        int
	symbolFirst          bool
	htmlEntity           string
	decimalMark          string
	thousandsSeparator   string
	isoNumeric           string
	smallestDenomination int
}

var currencyConfigMap = map[string]currencyConfig{
	"hkd": {
		priority:           100,
		isoCode:            "HKD",
		name:               "Hong Kong Dollar",
		symbol:             "$",
		disambiguateSymbol: "HK$",
		alternateSymbols: []string{
			"HK$",
		},
		subunit:              "Cent",
		subunitToUnit:        100,
		symbolFirst:          true,
		htmlEntity:           "$",
		decimalMark:          ".",
		thousandsSeparator:   ",",
		isoNumeric:           "344",
		smallestDenomination: 10,
	},
}

func newCurrency(isoCode string) *Currency {
	config, ok := currencyConfigMap[strings.ToLower(isoCode)]
	if !ok {
		panic(errors.New("unsupported country code"))
	}

	return &Currency{Code: config.isoCode, config: config}
}
