package auth

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/model/user"
	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/utils/apitest"
)

func TestMain(m *testing.M) {
	api.TestMain()

	os.Exit(m.Run())
}

func TestAddLogin(t *testing.T) {
	var at = apitest.NewAT(
		"/v1/auth/login",
		http.MethodPost,
		"登陆",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		route.Error
		Data user.Entity
	}

	t.Run("MakeDoc", func(t *testing.T) {
		t.SkipNow()

		file := "Login.md"
		title := "登陆接口文档"
		f, err := apitest.OpenFile(file, title)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&user.Entity{
				Name:     "jd",
				Password: "13420693396",
			}).
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
	})

	t.Run("Normal", func(t *testing.T) {
		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&user.Entity{
				Name:     "jd",
				Password: "13420693396",
			}).
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
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestAddUser(t *testing.T) {
	var at = apitest.NewAT(
		"/v1/auth/user",
		http.MethodPost,
		"添加",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		route.Error
		Data int
	}

	t.Run("MakeDoc", func(t *testing.T) {
		t.SkipNow()

		file := "Add.md"
		title := "添加用户接口"
		f, err := apitest.OpenFile(file, title)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

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
					return nil
				},
				r.Code, 0,
				r.Msg, "",
				r.Data != 0, true,
			).
			WriteFile(f).
			Err(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Tx", func(t *testing.T) {
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
					return nil
				},
				r.Code, 0,
				r.Msg, "",
				r.Data != 0, true,
			).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestGetUser(t *testing.T) {
	var at = apitest.NewAT(
		"/v1/auth/user",
		http.MethodGet,
		"获取",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		route.Error
		Data user.Entity
	}

	t.Run("MakeDoc", func(t *testing.T) {
		t.SkipNow()

		file := "Get.md"
		title := "获取用户接口"
		f, err := apitest.OpenFile(file, title)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
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
					var u = r.Data
					return at.Equal(
						u.ID != 0, true,
						u.Name, "jd",
					).Err()
				},
				r.Code, 0,
				r.Msg, "",
			).
			WriteFile(f).
			Err(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Normal", func(t *testing.T) {
		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
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
					var u = r.Data
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

	t.Run("PressureRun", func(t *testing.T) {
		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&struct {
				Name string
			}{
				Name: "jd",
			}).
			PressureRun(1000, 100).
			Err(); err != nil {
			t.Fatal(err)
		}
	})
}
