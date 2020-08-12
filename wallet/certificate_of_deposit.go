// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type CertificateOfDeposit struct {
    TradableAsset
}

type CertificateOfDepositList []CertificateOfDeposit

func NewCertificateOfDeposit() *CertificateOfDeposit {
	return &CertificateOfDeposit{TradableAsset:TradableAsset{ItemType: "certificate-of-deposit"}}
}