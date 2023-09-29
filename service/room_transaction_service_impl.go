package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
	"github.com/renaldiaddison/roomborrowingbackend/exception"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
	"github.com/renaldiaddison/roomborrowingbackend/model"
	"github.com/renaldiaddison/roomborrowingbackend/repository"
	"time"
)

type RoomTransactionServiceImpl struct {
	RoomTransactionRepository repository.RoomTransactionRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewRoomTransaction(roomTransactionRepository repository.RoomTransactionRepository, db *sql.DB, validator *validator.Validate) RoomTransactionService {
	return &RoomTransactionServiceImpl{
		RoomTransactionRepository: roomTransactionRepository,
		DB:                        db,
		Validate:                  validator,
	}
}

func (service RoomTransactionServiceImpl) CreateRoomTransactionBorrow(ctx context.Context, request model.RoomTransactionBorrowCreateRequest) model.RoomTransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roomTransaction := entities.RoomTransaction{
		Id:               uuid.New().String(),
		BorrowerUsername: request.BorrowerUsername,
		BorrowerDivision: request.BorrowerDivision,
		ReturnerUsername: nil,
		ReturnerDivision: nil,
		RoomNumber:       request.RoomNumber,
		RoomIn:           time.Now(),
		RoomOut:          nil,
	}

	roomTransaction = service.RoomTransactionRepository.CreateRoomTransactionBorrow(ctx, tx, roomTransaction)
	return helper.ToRoomTransactionResponse(roomTransaction)
}

func (service RoomTransactionServiceImpl) CreateRoomTransactionReturn(ctx context.Context, request model.RoomTransactionReturnCreateRequest) model.RoomTransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	now := time.Now()
	timePtr := &now

	roomTransaction, err := service.RoomTransactionRepository.FindRoomTransactionById(ctx, tx, request.Id)
	if err != nil {

		panic(exception.NewNotFoundError(err.Error()))
	}

	returnerUsername := request.ReturnerUsername
	returnerDivision := request.ReturnerDivision

	roomTransaction.ReturnerUsername = &returnerUsername
	roomTransaction.ReturnerDivision = &returnerDivision
	roomTransaction.RoomOut = timePtr

	roomTransaction = service.RoomTransactionRepository.CreateRoomTransactionReturn(ctx, tx, roomTransaction)

	return helper.ToRoomTransactionResponse(roomTransaction)
}

func (service RoomTransactionServiceImpl) FindActiveRoomTransaction(ctx context.Context, roomNumber string) []model.RoomTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activeRoomTransactions := service.RoomTransactionRepository.FindActiveRoomTransaction(ctx, tx, roomNumber)
	return helper.ToRoomTransactionResponses(activeRoomTransactions)
}

func (service RoomTransactionServiceImpl) FindAllRoomTransaction(ctx context.Context, roomNumber string) []model.RoomTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notActiveRoomTransactions := service.RoomTransactionRepository.FindAllRoomTransaction(ctx, tx, roomNumber)
	return helper.ToRoomTransactionResponses(notActiveRoomTransactions)
}
