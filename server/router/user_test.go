package router

import (
	"errors"
	. "github.com/agiledragon/gomonkey"
	"github.com/liaojuntao/controller"
	"github.com/liaojuntao/infrastruct"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// 测试创建用户
func Test_CreateUserRouter_Handler(t *testing.T) {
	patch := NewPatches()

	Convey("testing CreateUserRouter handler", t, func() {
		Convey("testing CreateUserRouter when success", func() {
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "CreateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return nil
				})

			userRou := GetUserRouter().CreateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "succeed")
			So(respMock.Code, ShouldEqual, 200)
		})
		Convey("testing CreateUserRouter handler when fail with wrongParam", func() {
			userRou := GetUserRouter().CreateUserRouter()
			postData := "{\"userName\":1000,\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4001,\"errMessage\":\"参数校验错误\"}")
			So(respMock.Code, ShouldEqual, 200)
		})
		Convey("testing CreateUserRouter handler when fail with optErr", func() {
			patch.Reset()
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "CreateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return errors.New("该用户已存在")
				})
			userRou := GetUserRouter().CreateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4002,\"errMessage\":\"该用户已存在\"}")
			So(respMock.Code, ShouldEqual, 200)
		})
	})
}
