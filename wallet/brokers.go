// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type Broker struct {
	ID   string `json:"id" bson:"_id" validate:"required"`
	Name string `json:"name" bson:"name" validate:"required"`
	CNPJ string `json:"CNPJ" bson:"CNPJ"`
}

type BrokersList struct {
	Brokers []Broker `json:"brokers" bson:"brokers"`
}
