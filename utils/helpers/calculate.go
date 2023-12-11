package helpers

import "github.com/shopspring/decimal"

func TotalTransaction(tax decimal.Decimal, price decimal.Decimal) decimal.Decimal {

	total := tax.Add(price)

	return total
}
