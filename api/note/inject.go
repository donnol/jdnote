// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package note

import (
	"context"

	"github.com/donnol/jdnote/models/note"
	"github.com/google/wire"
)

// InitNote 初始化
func InitNote(ctx context.Context) (note.Noter, error) {
	panic(wire.Build(wire.NewSet(note.New)))
}
