package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldiaddison/roomborrowingbackend/helper"
	"github.com/renaldiaddison/roomborrowingbackend/model"
	"github.com/renaldiaddison/roomborrowingbackend/service"
	"net/http"
)

type RoomTransactionControllerImpl struct {
	RoomTransactionService service.RoomTransactionService
}

func NewRoomTransaction(roomTransactionService service.RoomTransactionService) RoomTransactionController {
	return &RoomTransactionControllerImpl{
		RoomTransactionService: roomTransactionService,
	}
}

func (controller RoomTransactionControllerImpl) CreateRoomTransactionBorrow(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomTransactionBorrowCreateRequest := model.RoomTransactionBorrowCreateRequest{}
	helper.ReadFromRequestBody(request, &roomTransactionBorrowCreateRequest)

	roomTransactionResponse := controller.RoomTransactionService.CreateRoomTransactionBorrow(request.Context(), roomTransactionBorrowCreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomTransactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomTransactionControllerImpl) CreateRoomTransactionReturn(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomTransactionReturnCreateRequest := model.RoomTransactionReturnCreateRequest{}
	helper.ReadFromRequestBody(request, &roomTransactionReturnCreateRequest)
	roomTransactionResponse := controller.RoomTransactionService.CreateRoomTransactionReturn(request.Context(), roomTransactionReturnCreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomTransactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomTransactionControllerImpl) FindActiveRoomTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	searchQuery := request.URL.Query().Get("roomNumber")
	roomTransactionResponses := controller.RoomTransactionService.FindActiveRoomTransaction(request.Context(), searchQuery)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomTransactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller RoomTransactionControllerImpl) FindAllRoomTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomNumberQuery := request.URL.Query().Get("roomNumber")
	dateQuery := request.URL.Query().Get("date")
	roomTransactionResponses := controller.RoomTransactionService.FindAllRoomTransaction(request.Context(), roomNumberQuery, dateQuery)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomTransactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
