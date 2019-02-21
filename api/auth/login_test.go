package auth

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/donnol/jdnote/model/user"
	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
	"github.com/donnol/jdnote/utils/apitest"
)

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
		if err := at.New().SetPort(":8810").
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
		if err := at.New().SetPort(":8810").
			SetParam(&struct {
				Name string
			}{
				Name: "jd",
			}).
			Debug().
			Run().
			EqualCode(http.StatusOK).
			Result(&r).
			EqualThen(
				func(at *apitest.AT) error {
					b, _ := json.Marshal(r.Data)
					var u userao.User
					json.Unmarshal(b, &u)
					return at.Equal(
						u.ID, 11,
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
