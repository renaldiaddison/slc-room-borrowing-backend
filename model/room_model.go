package model

type RoomCreateRequest struct {
	RoomNumber string `validate:"required" json:"roomNumber"`
}

type RoomDeleteRequest struct {
	RoomNumber string `validate:"required" json:"roomNumber"`
}

type RoomResponse struct {
	RoomNumber string `validate:"required" json:"roomNumber"`
}
