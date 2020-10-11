package noteapi

import (
	stdctx "context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/errors"
	"github.com/donnol/jdnote/utils/route"
	"github.com/donnol/tools/apitest"
)

var appObj *app.App
var port = 8820

func TestMain(m *testing.M) {
	ctx := stdctx.Background()
	var cctx context.Context
	appObj, cctx = app.New(ctx)
	appObj.Register(cctx, &Note{})

	go func() {
		if err := appObj.StartServer(port); err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	cookie, err := appObj.MakeCookie(1)
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
		Data route.AddResult `json:"data"`
	}

	var modAT = apitest.NewAT(
		"/note",
		http.MethodPut,
		"修改",
		h,
		[]*http.Cookie{&cookie},
	)
	var modR errors.Error

	var pageAT = apitest.NewAT(
		"/note/page",
		http.MethodGet,
		"获取分页",
		h,
		[]*http.Cookie{&cookie},
	)
	var pageR struct {
		errors.Error
		Data notesrv.PageResult `json:"data"`
	}
	var detailAT = apitest.NewAT(
		"/note",
		http.MethodGet,
		"获取详情",
		h,
		[]*http.Cookie{&cookie},
	)
	var detailR struct {
		errors.Error
		Data notesrv.Result `json:"data"`
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

		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&struct{}{}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&r).
			EqualThen(
				func(at *apitest.AT) error {
					return at.Equal(
						r.Data.ID != 0, true,
					).Err()
				},
				r.Code == 0, true,
				r.Msg == "", true,
			).
			WriteFile(f).
			Err(); err != nil {
			t.Fatal(err)
		}

		if err := modAT.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&notesrv.ModParam{
				NoteID: r.Data.ID,
				Param: notesrv.Param{
					Title:  "mod title",
					Detail: "mod detail",
				},
			}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&modR).
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

		pageParam := notesrv.PageParam{}
		pageParam.PageSize = 10
		if err := pageAT.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&pageParam).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&pageR).
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

		detailParam := struct {
			NoteID int `json:"noteID"`
		}{
			NoteID: r.Data.ID,
		}
		if err := detailAT.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&detailParam).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&detailR).
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
