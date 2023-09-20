package exception

import (
	"net/http"

	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, https *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Status:  false,
			Code:    404,
			Message: "Item not found",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status:  false,
		Code:    500,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
