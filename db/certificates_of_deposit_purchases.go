// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertCertificateOfDepositPurchase(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertCertificateOfDepositPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateCertificateOfDepositPurchaseByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateCertificateOfDepositPurchaseByID")
	return m.updateCertificateOfDepositByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllCertificatesOfDepositsPurchases() (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] GetAllCertificatesOfDepositsPurchases")
	return m.getAllCertificatesOfDeposits(purchasesCollection)
}

func (m *mongoSession) GetCertificateOfDepositPurchasesByPortfolioID(id string) (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] GetCertificateOfDepositPurchasesByPortfolioID")
	return m.getCertificateOfDepositByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetCertificateOfDepositPurchaseByID(id string) (*wallet.CertificateOfDeposit, error) {
	log.Debug("[DB] GetCertificateOfDepositPurchaseByID")
	return m.getCertificateOfDepositByID(purchasesCollection, id)
}
