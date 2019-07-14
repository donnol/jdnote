package note

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/service/note"
	"github.com/donnol/jdnote/utils/apitest"
	"github.com/donnol/jdnote/utils/errors"
)

func TestMain(m *testing.M) {
	api.TestMain()

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	var at = apitest.NewAT(
		"/note",
		http.MethodPost,
		"添加",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		errors.Error
		Data int
	}

	t.Run("MakeDoc", func(t *testing.T) {
		// t.SkipNow()

		file := "README.md"
		title := "笔记接口"
		f, err := apitest.OpenFile(file, title)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&note.Param{
				Title:  "test title",
				Detail: "test detail",
			}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&r).
			EqualThen(
				func(at *apitest.AT) error {
					t.Log(r.Data)
					return at.Equal(
						r.Data != 0, true,
					).Err()
				},
				r.Code == 0, true,
				r.Msg == "", true,
			).
			WriteFile(f).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}
