package helper

import (
	"github.com/renaldiaddison/roomborrowingbackend/entities"
	"github.com/renaldiaddison/roomborrowingbackend/model"
)

func ToRoomTransactionResponse(roomTransaction entities.RoomTransaction) model.RoomTransactionResponse {
	return model.RoomTransactionResponse{
		Id:               roomTransaction.Id,
		BorrowerUsername: roomTransaction.BorrowerUsername,
		BorrowerDivision: roomTransaction.BorrowerDivision,
		ReturnerUsername: roomTransaction.ReturnerUsername,
		ReturnerDivision: roomTransaction.ReturnerDivision,
		RoomNumber:       roomTransaction.RoomNumber,
		RoomIn:           roomTransaction.RoomIn,
		RoomOut:          roomTransaction.RoomOut,
	}
}

func ToRoomTransactionResponses(roomTransactions []entities.RoomTransaction) []model.RoomTransactionResponse {
	var roomTransactionResponses []model.RoomTransactionResponse
	for _, roomTransaction := range roomTransactions {
		roomTransactionResponses = append(roomTransactionResponses, ToRoomTransactionResponse(roomTransaction))
	}
	return roomTransactionResponses
}

func ToRoomResponse(room entities.Room) model.RoomResponse {
	return model.RoomResponse{
		RoomNumber: room.RoomNumber,
	}
}

func ToRoomResponses(rooms []entities.Room) []model.RoomResponse {
	var roomResponses []model.RoomResponse
	for _, room := range rooms {
		roomResponses = append(roomResponses, ToRoomResponse(room))
	}
	return roomResponses
}
