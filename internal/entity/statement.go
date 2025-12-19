package entity

import "strings"

type Statement struct {
	UploadID     string `json:"-"`
	ID           string `json:"id"`
	Timestamp    int    `json:"timestamp"`
	Counterparty string `json:"counterparty"`
	Type         string `json:"type"`
	Amount       int    `json:"amount"`
	Status       string `json:"status"`
	Description  string `json:"description"`
}

func (i Statement) IsFailed() bool {
	return strings.EqualFold(i.Status, "failed")
}

func (i Statement) IsCredit() bool {
	return strings.EqualFold(i.Type, "credit")
}

func (i Statement) IsDebit() bool {
	return strings.EqualFold(i.Type, "debit")
}
