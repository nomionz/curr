package ecb

import (
	"github.com/shopspring/decimal"
)

type Store interface {
	Get(key string) (decimal.Decimal, error)
	GetKeys() []string
	Put(key string, value decimal.Decimal)
}
