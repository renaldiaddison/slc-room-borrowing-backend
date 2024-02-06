package entities

import "time"

type RoomTransaction struct {
	Id                   string
	BorrowerUsername     string
	BorrowerDivision     string
	BorrowerIdentityCode string
	ReturnerUsername     *string
	ReturnerDivision     *string
	ReturnerIdentityCode *string
	RoomNumber           string
	RoomIn               time.Time
	RoomOut              *time.Time
}
