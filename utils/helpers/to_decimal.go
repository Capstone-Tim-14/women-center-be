package helpers

import "github.com/shopspring/decimal"

func StringToDecimal(input string) decimal.Decimal {
	result, _ := decimal.NewFromString(input)
	return result
}
