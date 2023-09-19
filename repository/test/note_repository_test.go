package repository_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/prasetyodidi/golang-notes-api/helper"
	"github.com/prasetyodidi/golang-notes-api/models/domain"
	"github.com/prasetyodidi/golang-notes-api/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/golang_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func truncateNote(db *gorm.DB) {
	db.Exec("TRUNCATE notes")
}

func TestSaveNoteSuccess(t *testing.T) {
	db := setupTestDB()
	db.AutoMigrate(&domain.Note{})
	truncateNote(db)

	ctx := context.Background()

	tags := []string{"tag1", "tag2", "tag3"}
	tagsString, err := json.Marshal(tags)
	helper.PanicIfError(err)

	note := domain.Note{
		Title: "Golang",
		Tags:  string(tagsString),
		Body:  "catatan golang",
	}

	noteRepository := repository.NewNoteRepository()
	resultNote := noteRepository.Save(ctx, db, note)

	assert.Equal(t, note.Title, resultNote.Title)
	assert.Equal(t, note.Tags, resultNote.Tags)
	assert.Equal(t, note.Body, resultNote.Body)
}

func TestUpdateNoteSuccess(t *testing.T) {
	db := setupTestDB()
	db.AutoMigrate(&domain.Note{})

	truncateNote(db)

	noteRepository := repository.NewNoteRepository()

	ctx := context.Background()

	tags := []string{"tag1", "tag2", "tag3"}
	tagsString, err := json.Marshal(tags)
	helper.PanicIfError(err)

	note1 := domain.Note{
		Title: "Goroutine",
		Tags:  string(tagsString),
		Body:  "Ini adalah contoh catatan",
	}

	resultNote1 := noteRepository.Save(ctx, db, note1)

	updateNote := domain.Note{
		Title: "Judul Satu Update",
		Body:  "Ini adalah contoh catatan updated",
	}
	updateNote.ID = resultNote1.ID
	resultUpdateNote := noteRepository.Update(ctx, db, updateNote)

	assert.Equal(t, updateNote.Title, resultUpdateNote.Title)
	assert.Equal(t, updateNote.Tags, resultUpdateNote.Tags)
	assert.Equal(t, updateNote.Body, resultUpdateNote.Body)
}

func TestDeleteNoteSuccess(t *testing.T) {
	db := setupTestDB()
	db.AutoMigrate(&domain.Note{})

	truncateNote(db)

	noteRepository := repository.NewNoteRepository()

	ctx := context.Background()

	tags := []string{"tag1", "tag2", "tag3"}
	tagsString, err := json.Marshal(tags)
	helper.PanicIfError(err)

	note1 := domain.Note{
		Title: "Delete Goroutine",
		Tags:  string(tagsString),
		Body:  "catatan delete",
	}

	resultNote1 := noteRepository.Save(ctx, db, note1)

	noteRepository.Delete(ctx, db, resultNote1)
}

func TestFindByIdNoteSuccess(t *testing.T) {
	db := setupTestDB()
	db.AutoMigrate(&domain.Note{})

	truncateNote(db)

	noteRepository := repository.NewNoteRepository()

	ctx := context.Background()

	tags := []string{"tag1", "tag2", "tag3"}
	tagsString, err := json.Marshal(tags)
	helper.PanicIfError(err)

	note1 := domain.Note{
		Title: "find by id Golang",
		Tags:  string(tagsString),
		Body:  "catatan find by id",
	}

	resultNote1 := noteRepository.Save(ctx, db, note1)

	note, err := noteRepository.FindById(ctx, db, int(resultNote1.ID))
	helper.PanicIfError(err)

	assert.Equal(t, note1.Title, note.Title)
	assert.Equal(t, note1.Tags, note.Tags)
	assert.Equal(t, note1.Body, note.Body)
}

func TestFindAllNoteSuccess(t *testing.T) {
	db := setupTestDB()
	db.AutoMigrate(&domain.Note{})

	truncateNote(db)

	noteRepository := repository.NewNoteRepository()

	ctx := context.Background()

	tags := []string{"tag1", "tag2", "tag3"}
	tagsString, err := json.Marshal(tags)
	helper.PanicIfError(err)

	note1 := domain.Note{
		Title: "find all Golang 1",
		Tags:  string(tagsString),
		Body:  "catatan find all",
	}

	resultNote1 := noteRepository.Save(ctx, db, note1)

	note2 := domain.Note{
		Title: "find all Golang 2",
		Tags:  string(tagsString),
		Body:  "catatan find all",
	}

	resultNote2 := noteRepository.Save(ctx, db, note2)

	notes := noteRepository.FindAll(ctx, db)
	result1 := notes[0]
	result2 := notes[1]

	assert.Equal(t, resultNote1.Title, result1.Title)
	assert.Equal(t, resultNote2.Title, result2.Title)
}
