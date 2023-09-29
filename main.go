package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renaldiaddison/roomborrowingbackend/app"
	"github.com/renaldiaddison/roomborrowingbackend/controller"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
	"net/http"
	//"github.com/renaldiaddison/roomborrowingbackend/middleware"
	"github.com/renaldiaddison/roomborrowingbackend/repository"
	"github.com/renaldiaddison/roomborrowingbackend/service"
)

func main() {

	db := app.NewDatabase()
	validate := validator.New()
	roomTransactionRepository := repository.NewRoomTransactionRepository()
	roomTransactionService := service.NewRoomTransaction(roomTransactionRepository, db, validate)
	roomTransactionController := controller.NewRoomTransaction(roomTransactionService)
	router := app.NewRouter(roomTransactionController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
