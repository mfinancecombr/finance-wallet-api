// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertTreasuryDirectSale(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertTreasuryDirectSale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateTreasuryDirectSaleByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateTreasuryDirectSaleByID")
	return m.updateTreasuryDirectByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllTreasuriesDirectsSales() (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetAllTreasuriesDirectsSales")
	return m.getAllTreasuriesDirects(salesCollection)
}

func (m *mongoSession) GetTreasuryDirectSalesByPortfolioID(id string) (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetTreasuryDirectSalesByPortfolioID")
	return m.getTreasuryDirectByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetTreasuryDirectSaleByID(id string) (*wallet.TreasuryDirect, error) {
	log.Debug("[DB] GetTreasuryDirectSaleByID")
	return m.getTreasuryDirectByID(salesCollection, id)
}
