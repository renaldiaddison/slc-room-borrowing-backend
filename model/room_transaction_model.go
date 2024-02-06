package model

import "time"

type RoomTransactionBorrowCreateRequest struct {
	BorrowerUsername     string `validate:"required" json:"borrowerUsername"`
	BorrowerDivision     string `validate:"required" json:"borrowerDivision"`
	BorrowerIdentityCode string `validate:"required" json:"borrowerIdentityCode"`
	RoomNumber           string `validate:"required" json:"roomNumber"`
}

type RoomTransactionReturnCreateRequest struct {
	ReturnerUsername     string `validate:"required" json:"returnerUsername"`
	ReturnerDivision     string `validate:"required" json:"returnerDivision"`
	ReturnerIdentityCode string `validate:"required" json:"returnerIdentityCode"`
	RoomNumber           string `validate:"required" json:"roomNumber"`
}

type RoomTransactionResponse struct {
	Id                   string     `json:"id"`
	BorrowerUsername     string     `json:"borrowerUsername"`
	BorrowerDivision     string     `json:"borrowerDivision"`
	BorrowerIdentityCode string     `json:"borrowerIdentityCode"`
	ReturnerUsername     *string    `json:"returnerUsername"`
	ReturnerDivision     *string    `json:"returnerDivision"`
	ReturnerIdentityCode string     `json:"returnerIdentityCode"`
	RoomNumber           string     `json:"roomNumber"`
	RoomIn               time.Time  `json:"roomIn"`
	RoomOut              *time.Time `json:"roomOut"`
}
