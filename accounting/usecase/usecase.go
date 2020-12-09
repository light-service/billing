package usecase

import "time"

type UseCase interface {
	CreateAccount(accountType string, accountTypeRemark string, userKey string) (accountKey string, err error)
	GetAccountInfo(accountKey string, userKey string) error

	Freeze(transactionID string, accountKey string, amount int) error
	Unfreeze(accountKey string) error
	UpdateCreditLine(accountKey string, amount int) error

	Trade(transactionID string, accountKey string, opponentAccountKey string, Amount int) error
	ReverseTrade(transactionID string) error
	ListAccountFlow(accountFlowID string, transactionID string, createTime time.Time)
}
