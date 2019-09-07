package auth

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/models/user/userdb"
	"github.com/donnol/jdnote/utils/apitest"
	"github.com/donnol/jdnote/utils/errors"
)

func TestMain(m *testing.M) {
	api.TestMain()

	os.Exit(m.Run())
}

func TestAddLogin(t *testing.T) {
	var at = apitest.NewAT(
		"/auth/login",
		http.MethodPost,
		"登陆",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		errors.Error
		Data userdb.Entity
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
			SetParam(&userdb.Entity{
				Name:     "jd",
				Password: "jd",
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
			SetParam(&userdb.Entity{
				Name:     "jd",
				Password: "jd",
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
		"/auth/user",
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
				Password: "jd",
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
		go func() {
			if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
				SetParam(&struct {
					Name     string
					Password string
				}{
					Name:     "jd",
					Password: "jd",
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
		}()

		// 休眠两秒，然后新建一个请求
		time.Sleep(2 * time.Second)

		if err := at.New().SetPort(fmt.Sprintf(":%d", api.TestPort)).
			SetParam(&struct {
				Name     string
				Password string
			}{
				Name:     "jd",
				Password: "jd",
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
		"/auth/user",
		http.MethodGet,
		"获取",
		http.Header{},
		[]*http.Cookie{},
	)
	var r struct {
		errors.Error
		Data userdb.Entity
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
