package helper

import (
	"encoding/json"

	"github.com/prasetyodidi/golang-notes-api/models/domain"
	"github.com/prasetyodidi/golang-notes-api/models/web"
)

func ToNoteResponse(note domain.Note) web.NoteResponse {
	var tags []string
	err := json.Unmarshal([]byte(note.Tags), &tags)
	PanicIfError(err)

	return web.NoteResponse{
		Id:    int(note.ID),
		Title: note.Title,
		Tags:  tags,
		Body:  note.Body,
	}
}

func ToNotesResponse(notes []domain.Note) []web.NoteResponse {
	var notesResponse []web.NoteResponse
	for _, note := range notes {
		notesResponse = append(notesResponse, ToNoteResponse(note))
	}

	return notesResponse
}
