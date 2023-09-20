package service

import (
	"context"

	"github.com/prasetyodidi/golang-notes-api/models/web"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse
	Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) web.NoteResponse
	FindAll(ctx context.Context) []web.NoteResponse
}
