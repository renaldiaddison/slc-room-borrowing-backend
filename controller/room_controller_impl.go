package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
	"github.com/renaldiaddison/roomborrowingbackend/model"
	"github.com/renaldiaddison/roomborrowingbackend/service"
	"net/http"
)

type RoomControllerImpl struct {
	RoomService service.RoomService
}

func NewRoomController(roomService service.RoomService) RoomController {
	return RoomControllerImpl{
		RoomService: roomService,
	}
}

func (controller RoomControllerImpl) CreateRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomCreateRequest := model.RoomCreateRequest{}
	helper.ReadFromRequestBody(request, &roomCreateRequest)

	roomResponse := controller.RoomService.CreateRoom(request.Context(), roomCreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomControllerImpl) DeleteRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomNumber := params.ByName("roomNumber")

	controller.RoomService.DeleteRoom(request.Context(), roomNumber)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomControllerImpl) FindAllRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	searchQuery := request.URL.Query().Get("roomNumber")
	roomResponses := controller.RoomService.FindAllRoom(request.Context(), searchQuery)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomControllerImpl) FindActiveRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	searchQuery := request.URL.Query().Get("roomNumber")
	roomResponses := controller.RoomService.FindActiveRoom(request.Context(), searchQuery)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomControllerImpl) FindInactiveRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	searchQuery := request.URL.Query().Get("roomNumber")
	roomResponses := controller.RoomService.FindInactiveRoom(request.Context(), searchQuery)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
