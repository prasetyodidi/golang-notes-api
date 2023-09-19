package repository

import (
	"context"

	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/domain"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
}

func NewNoteRepository() NoteRepository {
	return &NoteRepositoryImpl{}
}

func (repository *NoteRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note {
	result := tx.WithContext(ctx).Create(&note)
	helper.PanicIfError(result.Error)

	return note
}

func (repository *NoteRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note {
	tx.WithContext(ctx).Save(&note)

	return note
}

func (repository *NoteRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, note domain.Note) {
	tx.WithContext(ctx).Delete(&note)
}

func (repository *NoteRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, noteId int) (domain.Note, error) {
	note := domain.Note{}
	note.ID = uint(noteId)

	result := tx.WithContext(ctx).First(&note)
	if result.Error != nil {
		return note, result.Error
	}

	return note, nil
}

func (repository *NoteRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Note {
	var notes []domain.Note

	tx.WithContext(ctx).Find(&notes)
	
	return notes
}
