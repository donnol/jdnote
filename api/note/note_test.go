package note

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/service/note"
	"github.com/donnol/jdnote/utils/apitest"
	"github.com/donnol/jdnote/utils/errors"
)

func TestMain(m *testing.M) {
	api.TestMain()

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	cookie, err := route.MakeCookie(114)
	if err != nil {
		t.Fatal(err)
	}
	h := http.Header{}
	var at = apitest.NewAT(
		"/note",
		http.MethodPost,
		"添加",
		h,
		[]*http.Cookie{&cookie},
	)
	var r struct {
		errors.Error
		Data int `json:"data"`
	}

	var modAT = apitest.NewAT(
		"/note",
		http.MethodPut,
		"修改",
		h,
		[]*http.Cookie{&cookie},
	)

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

		if err := modAT.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&note.ModParam{
				ID: r.Data,
				Param: note.Param{
					Title:  "mod title",
					Detail: "mod detail",
				},
			}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&r).
			EqualThen(
				func(at *apitest.AT) error {
					return nil
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
