package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prasetyodidi/golang-notes-api/controller"
	"github.com/prasetyodidi/golang-notes-api/exception"
	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/web"
)

func NewRouter(noteController controller.NoteController) chi.Router {
	router := chi.NewRouter()
	router.Use(panicHandler)
	router.NotFound(methodNotAllowedHandler)
	router.MethodNotAllowed(routeNotFoundHandler)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {

		webResponse := web.WebResponse{
			Status:  true,
			Code:    200,
			Message: "Request Success",
			Data:    "Hello World",
		}

		helper.WriteToResponseBody(writer, webResponse)
	})

	router.Route("/notes", func(r chi.Router) {
		r.Get("/", noteController.FindAll)
		r.Post("/", noteController.Create)
		r.Route("/{noteId}", func(r chi.Router) {
			r.Get("/", noteController.FindById)
			r.Put("/", noteController.Update)
			r.Delete("/", noteController.Delete)
		})
	})

	return router
}

func routeNotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	webResponse := web.WebResponse{
		Status:  false,
		Code:    405,
		Message: "Method not allowed",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	webResponse := web.WebResponse{
		Status:  false,
		Code:    404,
		Message: "Route not found",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func panicHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				exception.ErrorHandler(w, r, rvr)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
