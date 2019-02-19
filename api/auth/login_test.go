package auth

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/service/user"
	"github.com/donnol/jdnote/utils/apitest"
)

func TestLogin(t *testing.T) {
	var at = apitest.NewAT(
		"/login",
		http.MethodGet,
		"登陆",
		http.Header{},
		[]*http.Cookie{},
	)
	var r route.Result

	t.Run("Get", func(t *testing.T) {
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
					var u user.User
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
					var u user.User
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
