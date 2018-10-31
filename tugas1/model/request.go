package model

type RegisterRequest struct {
	UserID			  string     `json:"user_id"`
	Nama              string     `json:"nama"`
}

type TransferRequest struct {
	UserID            string     `json:"user_id"`
	Nilai             uint       `json:"nilai"`
}

type SaldoRequest struct {
	UserID            string     `json:"user_id"`
}