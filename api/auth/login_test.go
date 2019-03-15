package auth

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/model/user"
	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
	"github.com/donnol/jdnote/utils/apitest"
)

func TestMain(m *testing.M) {
	api.TestMain()

	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	var at = apitest.NewAT(
		"/login",
		http.MethodPost,
		"登陆",
		http.Header{},
		[]*http.Cookie{},
	)
	var r route.Result

	t.Run("Login", func(t *testing.T) {
		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&userao.User{
				User: user.User{
					Name:     "jd",
					Password: "13420693396",
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
				r.Code, 0,
				r.Msg, "",
			).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestAdd(t *testing.T) {
	var at = apitest.NewAT(
		"/add",
		http.MethodPost,
		"添加",
		http.Header{},
		[]*http.Cookie{},
	)
	var r route.Result

	t.Run("Post", func(t *testing.T) {
		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&struct {
				Name     string
				Password string
			}{
				Name:     "jd",
				Password: "13420693396",
			}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&r).
			EqualThen(
				func(at *apitest.AT) error {
					var u userao.User
					if err := r.PresentData(&u); err != nil {
						return err
					}
					return at.Equal(
						u.ID != 0, true,
						u.Name, "jd",
					).Err()
				},
				r.Code, 0,
				r.Msg, "",
			).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}
