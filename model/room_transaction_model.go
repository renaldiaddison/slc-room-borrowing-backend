package model

import "time"

type RoomTransactionBorrowCreateRequest struct {
	BorrowerUsername string `validate:"required" json:"borrowerUsername"`
	BorrowerDivision string `validate:"required" json:"borrowerDivision"`
	RoomNumber       string `validate:"required" json:"roomNumber"`
}

type RoomTransactionReturnCreateRequest struct {
	Id               string `validate:"required" json:"id"`
	ReturnerUsername string `validate:"required" json:"returnerUsername"`
	ReturnerDivision string `validate:"required" json:"returnerDivision"`
	RoomNumber       string `validate:"required" json:"roomNumber"`
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
