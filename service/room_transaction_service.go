package service

import (
	"context"
	"github.com/renaldiaddison/roomborrowingbackend/model"
)

type RoomTransactionService interface {
	CreateRoomTransactionBorrow(ctx context.Context, request model.RoomTransactionBorrowCreateRequest) model.RoomTransactionResponse
	CreateRoomTransactionReturn(ctx context.Context, request model.RoomTransactionReturnCreateRequest) model.RoomTransactionResponse
	FindActiveRoomTransaction(ctx context.Context, roomNumber string) []model.RoomTransactionResponse
	FindAllRoomTransaction(ctx context.Context, roomNumber string, date string) []model.RoomTransactionResponse
}
