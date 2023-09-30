package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/renaldiaddison/roomborrowingbackend/entities"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
)

type RoomRepositoryImpl struct {
}

func NewRoomRepository() RoomRepository {
	return &RoomRepositoryImpl{}
}

func (repository RoomRepositoryImpl) CreateRoom(ctx context.Context, tx *sql.Tx, room entities.Room) entities.Room {
	SQL := "INSERT INTO `rooms`(`room_number`) VALUES (?)"
	_, err := tx.ExecContext(ctx, SQL, room.RoomNumber)
	helper.PanicIfError(err)
	return room
}

func (repository RoomRepositoryImpl) DeleteRoom(ctx context.Context, tx *sql.Tx, room entities.Room) {
	SQL := "DELETE FROM `rooms` WHERE room_number = ?"
	_, err := tx.ExecContext(ctx, SQL, room.RoomNumber)
	helper.PanicIfError(err)
}

func (repository RoomRepositoryImpl) FindAllRoom(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.Room {
	roomNumber = roomNumber + "%"
	SQL := "SELECT * FROM rooms WHERE room_number LIKE ?"
	rows, err := tx.QueryContext(ctx, SQL, roomNumber)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var rooms []entities.Room
	for rows.Next() {
		room := entities.Room{}
		err := rows.Scan(&room.RoomNumber)
		helper.PanicIfError(err)
		fmt.Println(room.RoomNumber)
		rooms = append(rooms, room)
	}

	return rooms
}
