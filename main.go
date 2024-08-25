package main

import (
	"fmt"
	"fresh_api/config"
	"fresh_api/controller"
	"fresh_api/helper"
	"fresh_api/repository"
	"fresh_api/router"
	"fresh_api/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	serverPort := "localhost:8088"
	fmt.Printf("Start server")
	// database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: serverPort, Handler: routes}

	log.Printf("Server berjalan di port %s\n", serverPort)
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
