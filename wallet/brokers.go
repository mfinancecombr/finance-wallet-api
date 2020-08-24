// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type Broker struct {
	CNPJ string `json:"CNPJ" bson:"CNPJ"`
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name" validate:"required"`
	Slug string `json:"slug" bson:"slug" validate:"required"`
}

type BrokersList struct {
	Brokers []Broker `json:"brokers" bson:"brokers"`
}

func (s Broker) GetCollectionName() string {
	return "brokers"
}

func (s Broker) GetItemType() string {
	return ""
}
