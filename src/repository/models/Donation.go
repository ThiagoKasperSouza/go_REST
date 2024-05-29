package models

type Donation struct {
	CrudModel
	Id   string `json:"id"`
	Nome string `json:"nome"`
	Link string `json:"link"`
}
