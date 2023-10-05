package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldiaddison/roomborrowingbackend/controller"
	"github.com/renaldiaddison/roomborrowingbackend/exception"
)

func NewRouter(roomController controller.RoomController, roomTransactionController controller.RoomTransactionController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/rooms", roomController.CreateRoom)
	router.GET("/api/rooms", roomController.FindAllRoom)
	router.DELETE("/api/rooms", roomController.DeleteRoom)
	router.GET("/api/rooms/active", roomController.FindActiveRoom)
	router.GET("/api/rooms/inactive", roomController.FindInactiveRoom)

	router.POST("/api/room-transactions/borrow", roomTransactionController.CreateRoomTransactionBorrow)
	router.POST("/api/room-transactions/return", roomTransactionController.CreateRoomTransactionReturn)
	router.GET("/api/room-transactions", roomTransactionController.FindAllRoomTransaction)
	router.GET("/api/room-transactions/active", roomTransactionController.FindActiveRoomTransaction)

	router.PanicHandler = exception.ErrorHandler

	return router
}
