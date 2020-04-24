// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertCertificateOfDepositSale(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertCertificateOfDepositSale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateCertificateOfDepositSaleByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateCertificateOfDepositSaleByID")
	return m.updateCertificateOfDepositByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllCertificatesOfDepositsSales() (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] GetAllCertificatesOfDepositsSales")
	return m.getAllCertificatesOfDeposits(salesCollection)
}

func (m *mongoSession) GetCertificateOfDepositSalesByPortfolioID(id string) (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] GetCertificateOfDepositSalesByPortfolioID")
	return m.getCertificateOfDepositByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetCertificateOfDepositSaleByID(id string) (*wallet.CertificateOfDeposit, error) {
	log.Debug("[DB] GetCertificateOfDepositSaleByID")
	return m.getCertificateOfDepositByID(salesCollection, id)
}
