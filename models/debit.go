package models

type Debit struct {
	Id            int64
	Username      string
	Accountnumber int64
	Expirydate    string
	Balance       float32
}

func (Debit) TableName() string {
	return "debit"
}
