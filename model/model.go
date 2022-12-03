package model

type Family struct {
	Id      string `json:"id"`
	Mother  string `json:"mother"`
	Brother string `json:"brother"`
	Prise   int    `json:"prise"`
}
