package controller

import (
	"fresh_api/data/request"
	"fresh_api/data/response"
	"fresh_api/helper"
	"fresh_api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookController struct {
	Bookservice service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{Bookservice: bookService}
}

func (controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)

	controller.Bookservice.Create(requests.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	controller.Bookservice.Update(requests.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.Bookservice.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) FindALl(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.Bookservice.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result := controller.Bookservice.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)

}
