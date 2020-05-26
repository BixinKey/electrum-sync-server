package main

import (
	"fmt"
)

type DbOpts struct {
	DbPath   string
	Host     string
	User     string
	Dbname   string
	DbType   string
	Password string
}

//SQL Object for wallet
type Wallet struct {
	Id           int       	 `sql:"index" json:"id"`
	XpubId       string    	 `json:"xpubId"`
	WalletId 	 string 	 `json:walletId`
	Xpubs		 string  	 `json:xpubs`
	WalletType   string      `json:walletType`
	WalletName   string		 `json:walletName`
}

type WalletsResponse struct {
	Walltes 	[]Wallet	`json: "wallets"`
}

type WalletRequest struct {
	XpubId		string       `json:xpubId`
	WalletId 	string 		 `json:walletId`
	Xpubs		string  	 `json:xpubs`
	WalletType  string       `json:walletType`
	WalletName  string		 `json:walletName`
}
//SQL Object for tx
type Transaction struct {
	Id           int       	 `sql:"index" json:"id"`
	WalletId 	 string 	 `json:walletId`
	TxHash		 string  	 `json:txHash`
	Tx           string      `json:tx`
}

type TxResponse struct {
	Transactions 	[]Transaction	`json: "transactions"`
}

type TxRequest struct {
	WalletId 	string 		 `json:walletId`
	TxHash		string  	 `json:txHash`
	Tx          string       `json:tx`
}

type TxDelRequest struct {
	WalletId 	string 		 `json:walletId`
	TxHash		string  	 `json:txHash`
}

type TxRbfRequest struct {
	WalletId 	string 		 `json:walletId`
	TxHash		string  	 `json:txHash`
	Tx          string       `json:tx`
	TxHashOld   string       `json:txHashOld`
}
// SQL Object
type Label struct {
	Id             int    `sql:"index" json:"id"`
	ExternalId     string `json:"externalId"`
	EncryptedLabel string `json:"encryptedLabel"`
	Nonce          int    `json:"nonce"`
	WalletId       string `json:"walletId"`
}

// Rest response
type LabelsResponse struct {
	Nonce  int     `json:"nonce"`
	Labels []Label `json:"labels"`
}

// Rest request
type LabelRequest struct {
	EncryptedLabel string `json:"encryptedLabel"`
	ExternalId     string `json:"externalId"`
	WalletId       string `json:"walletId"`
	WalletNonce    int    `json:"walletNonce"`
}

func (self LabelRequest) String() string {
	return fmt.Sprintf(`
Request information:
encryptedLabel: %s
externalId: %s
walletId: %s
walletNonce: %d
	`, self.EncryptedLabel, self.ExternalId, self.WalletId, self.WalletNonce)
}

type LabelsRequest struct {
	WalletNonce int          `json:"walletNonce"`
	WalletId    string       `json:"walletId"`
	Labels      []BatchLabel `json:"labels"`
}

type BatchLabel struct {
	EncryptedLabel string `json:"encryptedLabel"`
	ExternalId     string `json:"externalId"`
}
