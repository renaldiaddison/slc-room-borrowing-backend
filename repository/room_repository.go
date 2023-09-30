package repository

import (
	"context"
	"database/sql"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, tx *sql.Tx, room entities.Room) entities.Room
	DeleteRoom(ctx context.Context, tx *sql.Tx, room entities.Room)
	FindAllRoom(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.Room
}
