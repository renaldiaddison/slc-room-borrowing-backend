package repository

import (
	"context"
	"database/sql"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
	"github.com/renaldiaddison/roomborrowingbackend/exception"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
)

type RoomTransactionRepositoryImpl struct {
}

func NewRoomTransactionRepository() RoomTransactionRepository {
	return &RoomTransactionRepositoryImpl{}
}

func (repository *RoomTransactionRepositoryImpl) CreateRoomTransactionBorrow(ctx context.Context, tx *sql.Tx, roomTransaction entities.RoomTransaction) entities.RoomTransaction {
	SQL := "INSERT INTO `roomtransactions`(`id`, `borrower_username`, `borrower_division`, `borrower_identity_code`, `room_number`, `room_in`) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, roomTransaction.Id, roomTransaction.BorrowerUsername, roomTransaction.BorrowerDivision, roomTransaction.BorrowerIdentityCode, roomTransaction.RoomNumber, roomTransaction.RoomIn.Format("2006-01-02 15:04:05"))
	helper.PanicIfError(err)
	return roomTransaction
}

func (repository *RoomTransactionRepositoryImpl) CreateRoomTransactionReturn(ctx context.Context, tx *sql.Tx, roomTransaction entities.RoomTransaction) entities.RoomTransaction {
	SQL := "UPDATE `roomtransactions` SET `returner_username`= ?,`returner_division`= ?, `returner_identity_code` = ?, `room_out`= ? WHERE `id` = ?"
	_, err := tx.ExecContext(ctx, SQL, roomTransaction.ReturnerUsername, roomTransaction.ReturnerDivision, roomTransaction.ReturnerIdentityCode, roomTransaction.RoomOut.Format("2006-01-02 15:04:05"), roomTransaction.Id)
	helper.PanicIfError(err)
	return roomTransaction
}

func (repository *RoomTransactionRepositoryImpl) FindRoomTransactionById(ctx context.Context, tx *sql.Tx, roomTransactionId string) (entities.RoomTransaction, error) {
	SQL := "SELECT * FROM roomtransactions WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roomTransactionId)

	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	roomTransaction := entities.RoomTransaction{}
	if rows.Next() {
		err := rows.Scan(&roomTransaction.Id, &roomTransaction.BorrowerUsername, &roomTransaction.BorrowerDivision, &roomTransaction.BorrowerIdentityCode, &roomTransaction.ReturnerUsername, &roomTransaction.ReturnerDivision, &roomTransaction.ReturnerIdentityCode, &roomTransaction.RoomNumber, &roomTransaction.RoomIn, &roomTransaction.RoomOut)
		helper.PanicIfError(err)
		return roomTransaction, nil
	} else {
		return roomTransaction, exception.NewNotFoundError("room transaction not found")
	}
}

func (repository *RoomTransactionRepositoryImpl) FindOneActiveRoomTransaction(ctx context.Context, tx *sql.Tx, roomNumber string) (entities.RoomTransaction, error) {
	SQL := "SELECT * FROM roomtransactions WHERE room_out IS NULL AND room_number = ?"
	rows, err := tx.QueryContext(ctx, SQL, roomNumber)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var activeRoomTransaction entities.RoomTransaction
	if rows.Next() {
		err := rows.Scan(&activeRoomTransaction.Id, &activeRoomTransaction.BorrowerUsername, &activeRoomTransaction.BorrowerDivision, &activeRoomTransaction.BorrowerIdentityCode, &activeRoomTransaction.ReturnerUsername, &activeRoomTransaction.ReturnerDivision, &activeRoomTransaction.ReturnerIdentityCode, &activeRoomTransaction.RoomNumber, &activeRoomTransaction.RoomIn, &activeRoomTransaction.RoomOut)
		helper.PanicIfError(err)
		return activeRoomTransaction, nil
	} else {
		return activeRoomTransaction, exception.NewNotFoundError("room transaction not found")
	}
}

func (repository *RoomTransactionRepositoryImpl) FindActiveRoomTransaction(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.RoomTransaction {
	roomNumber = roomNumber + "%"
	SQL := "SELECT * FROM roomtransactions WHERE room_out IS NULL AND room_number LIKE ?"
	rows, err := tx.QueryContext(ctx, SQL, roomNumber)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var activeRoomTransactions []entities.RoomTransaction
	for rows.Next() {
		activeRoomTransaction := entities.RoomTransaction{}
		err := rows.Scan(&activeRoomTransaction.Id, &activeRoomTransaction.BorrowerUsername, &activeRoomTransaction.BorrowerDivision, &activeRoomTransaction.BorrowerIdentityCode, &activeRoomTransaction.ReturnerUsername, &activeRoomTransaction.ReturnerDivision, &activeRoomTransaction.ReturnerIdentityCode, &activeRoomTransaction.RoomNumber, &activeRoomTransaction.RoomIn, &activeRoomTransaction.RoomOut)
		helper.PanicIfError(err)
		activeRoomTransactions = append(activeRoomTransactions, activeRoomTransaction)
	}

	return activeRoomTransactions
}
func (repository *RoomTransactionRepositoryImpl) FindAllRoomTransaction(ctx context.Context, tx *sql.Tx, roomNumber string, date string) []entities.RoomTransaction {
	roomNumber = roomNumber + "%"
	date = date + "%"
	SQL := "SELECT * FROM roomtransactions WHERE room_number LIKE ? AND DATE(room_in) LIKE ?"
	rows, err := tx.QueryContext(ctx, SQL, roomNumber, date)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var activeRoomTransactions []entities.RoomTransaction
	for rows.Next() {
		activeRoomTransaction := entities.RoomTransaction{}
		err := rows.Scan(&activeRoomTransaction.Id, &activeRoomTransaction.BorrowerUsername, &activeRoomTransaction.BorrowerDivision, &activeRoomTransaction.BorrowerIdentityCode, &activeRoomTransaction.ReturnerUsername, &activeRoomTransaction.ReturnerDivision, &activeRoomTransaction.ReturnerIdentityCode, &activeRoomTransaction.RoomNumber, &activeRoomTransaction.RoomIn, &activeRoomTransaction.RoomOut)
		helper.PanicIfError(err)
		activeRoomTransactions = append(activeRoomTransactions, activeRoomTransaction)
	}

	return activeRoomTransactions
}
