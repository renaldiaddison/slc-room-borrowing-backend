package model

import "time"

type RoomTransactionBorrowCreateRequest struct {
	BorrowerUsername string `json:"borrowerUsername"`
	BorrowerDivision string `json:"borrowerDivision"`
	RoomNumber       string `json:"roomNumber"`
}

type RoomTransactionReturnCreateRequest struct {
	Id               string  `json:"id"`
	ReturnerUsername *string `json:"returnerUsername"`
	ReturnerDivision *string `json:"returnerDivision"`
	RoomNumber       string  `json:"roomNumber"`
}

type RoomTransactionResponse struct {
	Id               string     `json:"id"`
	BorrowerUsername string     `json:"borrowerUsername"`
	BorrowerDivision string     `json:"borrowerDivision"`
	ReturnerUsername *string    `json:"returnerUsername"`
	ReturnerDivision *string    `json:"returnerDivision"`
	RoomNumber       string     `json:"roomNumber"`
	RoomIn           time.Time  `json:"roomIn"`
	RoomOut          *time.Time `json:"roomOut"`
}
