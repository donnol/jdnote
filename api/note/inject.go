// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package note

import (
	"context"

	notemodel "github.com/donnol/jdnote/models/note"
	"github.com/donnol/jdnote/services/note"
	"github.com/google/wire"
)

// InitNote 初始化
func InitNote(ctx context.Context) (*Note, error) {
	panic(wire.Build(wire.NewSet(notemodel.New, note.New, New)))
}
