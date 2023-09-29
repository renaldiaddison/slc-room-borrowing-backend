package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldiaddison/roomborrowingbackend/controller"
	"github.com/renaldiaddison/roomborrowingbackend/exception"
	"net/http"
)

func NewRouter(roomTransactionController controller.RoomTransactionController) *httprouter.Router {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		router.ServeHTTP(w, r)
	})

	router.POST("/api/room-transactions/borrow", roomTransactionController.CreateRoomTransactionBorrow)
	router.POST("/api/room-transactions/return", roomTransactionController.CreateRoomTransactionReturn)
	router.GET("/api/room-transactions", roomTransactionController.FindAllRoomTransaction)
	router.GET("/api/room-transactions/active", roomTransactionController.FindActiveRoomTransaction)

	router.PanicHandler = exception.ErrorHandler

	return router
}
