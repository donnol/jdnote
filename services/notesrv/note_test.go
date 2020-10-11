package notesrv

import (
	stdctx "context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
)

func TestNoteGetHugoContent(t *testing.T) {
	n := &noteImpl{}
	content := n.getHugoContent("hah", "# hahahah", "2019-12-15", true, []string{}, []string{}, []string{})
	t.Logf("%s\n", content)
	content = n.getHugoContent("hah", "# hahahah", "2019-12-15", false, []string{"Go"}, []string{"Go"}, []string{"Go"})
	t.Logf("%s\n", content)
}

func TestPublish(t *testing.T) {
	n := &noteImpl{}
	sctx := stdctx.Background()
	_, ctx := app.New(sctx)
	if err := n.Publish(ctx, 45); err != nil {
		t.Fatal(err)
	}
}

var (
	noteEntity = notemodel.Entity{
		ID:        1,
		Detail:    "mock",
		Title:     "title",
		UserID:    1,
		CreatedAt: time.Now(),
	}
	mock = notemodel.Mock{
		AddOneHandler: func(ctx context.Context) (id int, err error) {
			return noteEntity.ID, nil
		},
		AddHandler: func(ctx context.Context, entity notemodel.Entity) (id int, err error) {
			return noteEntity.ID, nil
		},
		ModHandler: func(ctx context.Context, id int, entity notemodel.Entity) (err error) {
			return
		},
		DelHandler: func(ctx context.Context, id int) (err error) {
			return
		},
		GetPageHandler: func(ctx context.Context, entity notemodel.Entity, param common.Param) (
			res notemodel.EntityList,
			total int,
			err error,
		) {
			return
		},
		GetHandler: func(ctx context.Context, id int) (entity notemodel.Entity, err error) {
			return noteEntity, nil
		},
		GetListHandler: func(ctx context.Context, ids []int64) (entitys notemodel.EntityList, err error) {
			return notemodel.EntityList{noteEntity}, nil
		},
	}
)

func TestGet(t *testing.T) {
	sctx := stdctx.Background()
	_, ctx := app.New(sctx)

	v := New(mock)
	want := Result{
		NoteID:    noteEntity.ID,
		UserName:  strconv.Itoa(noteEntity.ID),
		Title:     noteEntity.Title,
		Detail:    noteEntity.Detail,
		CreatedAt: noteEntity.CreatedAt.Unix(),
	}
	r, err := v.Get(ctx, noteEntity.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(r, want) {
		t.Fatalf("Bad result: %+v != %+v\n", r, noteEntity)
	}
}
