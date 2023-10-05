package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renaldiaddison/roomborrowingbackend/app"
	"github.com/renaldiaddison/roomborrowingbackend/controller"
	"log"
	"net/http"
	"os"

	"github.com/renaldiaddison/roomborrowingbackend/repository"
	"github.com/renaldiaddison/roomborrowingbackend/service"
	//"github.com/renaldiaddison/roomborrowingbackend/middleware"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := app.NewDatabase()
	validate := validator.New()

	roomRepository := repository.NewRoomRepository()
	roomService := service.NewRoomService(roomRepository, db, validate)
	roomController := controller.NewRoomController(roomService)

	roomTransactionRepository := repository.NewRoomTransactionRepository()
	roomTransactionService := service.NewRoomTransaction(roomTransactionRepository, db, validate)
	roomTransactionController := controller.NewRoomTransaction(roomTransactionService)

	router := app.NewRouter(roomController, roomTransactionController)

	handler := cors.Default().Handler(router)

	log.Printf("Server is at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
