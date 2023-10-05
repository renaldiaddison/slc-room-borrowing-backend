package service

import (
	"context"
	"github.com/renaldiaddison/roomborrowingbackend/model"
)

type RoomService interface {
	CreateRoom(ctx context.Context, request model.RoomCreateRequest) model.RoomResponse
	DeleteRoom(ctx context.Context, request model.RoomDeleteRequest)
	FindAllRoom(ctx context.Context, roomNumber string) []model.RoomResponse
	FindActiveRoom(ctx context.Context, roomNumber string) []model.RoomResponse
	FindInactiveRoom(ctx context.Context, roomNumber string) []model.RoomResponse
}
