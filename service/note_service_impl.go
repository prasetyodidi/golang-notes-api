package service

import (
	"context"
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/prasetyodidi/golang-notes-api/exception"
	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/domain"
	"github.com/prasetyodidi/golang-notes-api/models/web"
	"github.com/prasetyodidi/golang-notes-api/repository"
	"gorm.io/gorm"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewNoteService(noteRepository repository.NoteRepository, DB *gorm.DB, validator *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB:             DB,
		Validate:       validator,
	}
}

func (service *NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tags, err := json.Marshal(request.Tags)
	helper.PanicIfError(err)

	note := domain.Note{
		Title: request.Title,
		Tags:  string(tags),
		Body:  request.Body,
	}

	note = service.NoteRepository.Save(ctx, service.DB, note)

	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse {
	note, err := service.NoteRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tags, err := json.Marshal(request.Tags)
	helper.PanicIfError(err)

	note.Title = request.Title
	note.Body = request.Body
	note.Tags = string(tags)

	note = service.NoteRepository.Update(ctx, service.DB, note)

	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Delete(ctx context.Context, noteId int) {
	note, err := service.NoteRepository.FindById(ctx, service.DB, noteId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.NoteRepository.Delete(ctx, service.DB, note)
}

func (service *NoteServiceImpl) FindById(ctx context.Context, noteId int) web.NoteResponse {
	note, err := service.NoteRepository.FindById(ctx, service.DB, noteId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) FindAll(ctx context.Context) []web.NoteResponse {
	notes := service.NoteRepository.FindAll(ctx, service.DB)

	return helper.ToNotesResponse(notes)
}
