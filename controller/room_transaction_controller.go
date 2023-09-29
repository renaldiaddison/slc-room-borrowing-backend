package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RoomTransactionController interface {
	CreateRoomTransactionBorrow(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateRoomTransactionReturn(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindActiveRoomTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllRoomTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
