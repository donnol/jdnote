package authapi

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/services/usersrv"
	"github.com/donnol/jdnote/stores/userrolestore"
	"github.com/donnol/jdnote/stores/userstore"
	"github.com/donnol/jdnote/utils/errors"
	"github.com/donnol/tools/apitest"
)

var port = 8820

func TestMain(m *testing.M) {
	appObj := initializers.New()
	appObj.MustRegisterProvider(
		initializers.ProviderOption{
			Provider: userrolestore.New,
		},
		initializers.ProviderOption{
			Provider: userstore.New,
		},
		initializers.ProviderOption{
			Provider: usersrv.New,
		},
	)
	appObj.RegisterRouterWithInject(&Auth{})

	go func() {
		if err := appObj.StartServer(port); err != nil {
			panic(err)
		}
	}()

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
		Data usersrv.Entity
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

		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&usersrv.Entity{
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
		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
			SetParam(&usersrv.Entity{
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

		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
			if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
				log.Fatal(err)
			}
		}()

		// 休眠两秒，然后新建一个请求
		time.Sleep(2 * time.Second)

		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
		Data usersrv.Entity
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

		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
		if err := at.New().SetPort(fmt.Sprintf(":%d", port)).
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
