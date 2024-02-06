package models

type Mobile struct {
	Id          string
	Username    string
	Expirydate  string
	Balance     float32
	Phonenumber string
}

func (Mobile) TableName() string {
	return "mobile"
}
