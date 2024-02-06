package models

type Credit struct {
	Id           int64
	Username     string
	Creditnumber int64
	Expirydate   string
	Amount       float32
}

func (Credit) TableName() string {
	return "credit"
}
