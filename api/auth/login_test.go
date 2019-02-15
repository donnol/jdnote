package auth

import (
	"net/http"
	"testing"

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
	var r user.User

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
			Equal(
				r.ID, 5,
				r.Name, "jd",
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
	var r user.User

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
			Equal(
				r.ID, 5,
				r.Name, "jd",
			).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}
