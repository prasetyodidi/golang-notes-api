package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/prasetyodidi/golang-notes-api/app"
	"github.com/prasetyodidi/golang-notes-api/controller"
	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/domain"
	"github.com/prasetyodidi/golang-notes-api/repository"
	"github.com/prasetyodidi/golang-notes-api/service"
)

func main() {
	validator := validator.New()
	db := app.NewDatabase()
	db.AutoMigrate(&domain.Note{})
	noteRepository := repository.NewNoteRepository()
	noteService := service.NewNoteService(noteRepository, db, validator)
	noteController := controller.NewNotControllerImpl(noteService)
	router := app.NewRouter(noteController)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
