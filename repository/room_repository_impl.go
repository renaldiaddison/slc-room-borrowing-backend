package repository

import (
	"context"
	"database/sql"
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
	SQL := "SELECT * FROM rooms WHERE room_number LIKE ? ORDER BY room_number"
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
		rooms = append(rooms, room)
	}

	return rooms
}

func (repository RoomRepositoryImpl) FindActiveRoom(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.Room {
	roomNumber = roomNumber + "%"
	SQL := "SELECT rt.room_number FROM roomtransactions rt WHERE rt.room_out IS NULL AND rt.room_number LIKE ? ORDER BY rt.room_number"
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
		rooms = append(rooms, room)
	}

	return rooms
}

func (repository RoomRepositoryImpl) FindInactiveRoom(ctx context.Context, tx *sql.Tx, roomNumber string) []entities.Room {
	roomNumber = roomNumber + "%"
	SQL := "SELECT r.room_number FROM rooms r WHERE r.room_number NOT IN (SELECT rt.room_number FROM roomtransactions rt WHERE rt.room_out IS NULL) AND r.room_number LIKE ? ORDER BY r.room_number"
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
		rooms = append(rooms, room)
	}

	return rooms
}
