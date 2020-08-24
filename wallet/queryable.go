// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type Queryable interface {
	GetCollectionName() string
	GetItemType() string
}
