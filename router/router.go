package router

import (
	"fmt"
	"fresh_api/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(bookController *controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	apiBook := "/api/book/:bookId"

	router.GET("/api/book", bookController.FindALl)
	router.GET(apiBook, bookController.FindById)
	router.POST("/api/book", bookController.Create)
	router.PATCH(apiBook, bookController.Update)
	router.DELETE(apiBook, bookController.Delete)

	return router
}
