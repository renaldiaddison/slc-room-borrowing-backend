package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
	"github.com/renaldiaddison/roomborrowingbackend/model"
	"github.com/renaldiaddison/roomborrowingbackend/repository"
)

type RoomServiceImpl struct {
	RoomRepository repository.RoomRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoomService(roomRepository repository.RoomRepository, db *sql.DB, validate *validator.Validate) RoomService {
	return RoomServiceImpl{
		RoomRepository: roomRepository,
		DB:             db,
		Validate:       validate,
	}
}
func (service RoomServiceImpl) CreateRoom(ctx context.Context, request model.RoomCreateRequest) model.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room := entities.Room{
		RoomNumber: request.RoomNumber,
	}

	room = service.RoomRepository.CreateRoom(ctx, tx, room)
	return helper.ToRoomResponse(room)
}

func (service RoomServiceImpl) DeleteRoom(ctx context.Context, request model.RoomDeleteRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room, err := service.RoomRepository.FindRoomById(ctx, tx, request.RoomNumber)
	helper.PanicIfError(err)

	service.RoomRepository.DeleteRoom(ctx, tx, room)
}

func (service RoomServiceImpl) FindAllRoom(ctx context.Context, roomNumber string) []model.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FindAllRoom(ctx, tx, roomNumber)
	return helper.ToRoomResponses(rooms)
}

func (service RoomServiceImpl) FindActiveRoom(ctx context.Context, roomNumber string) []model.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FindActiveRoom(ctx, tx, roomNumber)
	return helper.ToRoomResponses(rooms)
}

func (service RoomServiceImpl) FindInactiveRoom(ctx context.Context, roomNumber string) []model.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FindInactiveRoom(ctx, tx, roomNumber)
	return helper.ToRoomResponses(rooms)
}
