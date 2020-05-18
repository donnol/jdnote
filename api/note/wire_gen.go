// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package note

import (
	"context"
	"github.com/donnol/jdnote/models/note"
	note2 "github.com/donnol/jdnote/services/note"
)

// Injectors from inject.go:

func InitNote(ctx context.Context) (*Note, error) {
	noter := note.New()
	noteNoter := note2.New(noter)
	noteNote := New(noteNoter)
	return noteNote, nil
}
