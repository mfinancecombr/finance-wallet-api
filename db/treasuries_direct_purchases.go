// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertTreasuryDirectPurchase(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertTreasuryDirectPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateTreasuryDirectPurchaseByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateTreasuryDirectPurchaseByID")
	return m.updateTreasuryDirectByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllTreasuriesDirectsPurchases() (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetAllTreasuriesDirectsPurchases")
	return m.getAllTreasuriesDirects(purchasesCollection)
}

func (m *mongoSession) GetTreasuryDirectPurchasesByPortfolioID(id string) (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetTreasuryDirectPurchasesByPortfolioID")
	return m.getTreasuryDirectByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetTreasuryDirectPurchaseByID(id string) (*wallet.TreasuryDirect, error) {
	log.Debug("[DB] GetTreasuryDirectPurchaseByID")
	return m.getTreasuryDirectByID(purchasesCollection, id)
}
