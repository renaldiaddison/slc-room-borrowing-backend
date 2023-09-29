package entities

import "time"

type RoomTransaction struct {
	Id               string
	BorrowerUsername string
	BorrowerDivision string
	ReturnerUsername *string
	ReturnerDivision *string
	RoomNumber       string
	RoomIn           time.Time
	RoomOut          *time.Time
}
