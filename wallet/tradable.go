// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type Tradable interface {
	GetPrice() float64
	GetShares() float64
	GetComission() float64
	GetType() string
	GetBrokerSlug() string
}
