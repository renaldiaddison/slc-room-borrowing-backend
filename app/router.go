package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldiaddison/roomborrowingbackend/controller"
	"github.com/renaldiaddison/roomborrowingbackend/exception"
)

func NewRouter(roomTransactionController controller.RoomTransactionController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/room-transactions/borrow", roomTransactionController.CreateRoomTransactionBorrow)
	router.POST("/api/room-transactions/return", roomTransactionController.CreateRoomTransactionReturn)
	router.GET("/api/room-transactions", roomTransactionController.FindAllRoomTransaction)
	router.GET("/api/room-transactions/active", roomTransactionController.FindActiveRoomTransaction)

	router.PanicHandler = exception.ErrorHandler

	return router
}
