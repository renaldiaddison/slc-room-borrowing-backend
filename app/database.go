package app

import (
	"database/sql"
	"fmt"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
)

var db *sql.DB

func NewDatabase() *sql.DB {
	if db == nil {
		database, err := sql.Open("mysql", "root@tcp(localhost:3306)/slc_room_borrowing?parseTime=true")
		fmt.Println(err)
		helper.PanicIfError(err)
		db = database
	}
	return db
}
