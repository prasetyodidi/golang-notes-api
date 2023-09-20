package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/web"
	"github.com/prasetyodidi/golang-notes-api/service"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNotControllerImpl(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (controller *NoteControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	noteCreateRequest := web.NoteCreateRequest{}
	helper.ReadFromRequestBody(request, &noteCreateRequest)

	noteResponse := controller.NoteService.Create(request.Context(), noteCreateRequest)

	webResponse := web.WebResponse{
		Status:  true,
		Code:    201,
		Message: "Success create new note",
		Data: map[string]interface{}{
			"note": noteResponse,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NoteControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(request, "noteId"))
	helper.PanicIfError(err)

	noteUpdateRequest := web.NoteUpdateRequest{}
	noteUpdateRequest.Id = noteId
	helper.ReadFromRequestBody(request, &noteUpdateRequest)

	noteResponse := controller.NoteService.Update(request.Context(), noteUpdateRequest)

	webResponse := web.WebResponse{
		Status:  true,
		Code:    200,
		Message: "Success update note",
		Data: map[string]interface{}{
			"note": noteResponse,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NoteControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(request, "noteId"))
	helper.PanicIfError(err)

	controller.NoteService.Delete(request.Context(), noteId)

	webResponse := web.WebResponse{
		Status:  true,
		Code:    200,
		Message: "Success delete note",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NoteControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(request, "noteId"))
	helper.PanicIfError(err)

	note := controller.NoteService.FindById(request.Context(), noteId)

	webResponse := web.WebResponse{
		Status:  true,
		Code:    200,
		Message: "Success get note by id",
		Data: map[string]interface{}{
			"note": note,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NoteControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	notes := controller.NoteService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Status:  true,
		Code:    200,
		Message: "Success get note by id",
		Data: map[string]interface{}{
			"notes": notes,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
