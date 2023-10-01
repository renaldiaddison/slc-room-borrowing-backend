package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RoomController interface {
	CreateRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindActiveRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindInactiveRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
