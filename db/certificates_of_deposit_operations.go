// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertCertificateOfDepositOperation(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertCertificateOfDepositOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateCertificateOfDepositOperationByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateCertificateOfDepositOperationByID")
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllCertificatesOfDepositsOperations() (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] GetAllCertificatesOfDepositsOperations")
	return m.getAllCertificatesOfDeposits(operationsCollection)
}

func (m *mongoSession) GetCertificateOfDepositOperationByID(id string) (*wallet.CertificateOfDeposit, error) {
	log.Debug("[DB] GetCertificateOfDepositOperationByID")
	certificateOfDeposit := &wallet.CertificateOfDeposit{}
	if err := m.getOperationByID(operationsCollection, id, certificateOfDeposit); err != nil {
		return nil, err
	}
	if certificateOfDeposit.Symbol == "" {
		return nil, nil
	}
	return certificateOfDeposit, nil
}
