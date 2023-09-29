package repository

import (
	"context"
	"database/sql"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
)

type RoomTransactionRepository interface {
	CreateRoomTransactionBorrow(ctx context.Context, tx *sql.Tx, roomTransaction entities.RoomTransaction) entities.RoomTransaction
	CreateRoomTransactionReturn(ctx context.Context, tx *sql.Tx, roomTransaction entities.RoomTransaction) entities.RoomTransaction
	FindActiveRoomTransaction(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.RoomTransaction
	FindAllRoomTransaction(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.RoomTransaction
	FindRoomTransactionById(ctx context.Context, tx *sql.Tx, roomTransactionId string) (entities.RoomTransaction, error)
}
