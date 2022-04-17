package server

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestServer_Init_and_ServeHTTP_when_success(t *testing.T) {
	s := NewServer("127.0.0.1:8090")
	Convey("testing server", t, func() {
		Convey("testing Init", func() {
			s.Init()
			So(s.ServeMux, ShouldNotBeNil)
			So(len(s.router), ShouldNotEqual, 0)
			So(len(s.Interceptors), ShouldNotEqual, 0)
		})
		Convey("testing ServeHTTP", func() {
			s.Interceptors = append(s.Interceptors, func(resp http.ResponseWriter, req *http.Request, handler Handler) {
				resp.Header().Set("isAddInterceptorSuccess", "true")
				handler(resp, req)
			})
			respMock := new(httptest.ResponseRecorder)
			reqMock := new(http.Request)
			reqMock.URL = new(url.URL)
			s.ServeHTTP(respMock, reqMock)
			So(respMock.Header().Get("isAddInterceptorSuccess"), ShouldEqual, "true")
		})
	})
}

func TestServer_Run_when_success(t *testing.T) {
	Convey("testing server Run", t, func() {
		Convey("testing Run Success", func() {
			s := NewServer("127.0.0.1:8090")
			s.Init()
			var runErr error
			go func() {
				runErr = s.Run()
			}()
			time.Sleep(100 * time.Millisecond)
			So(runErr, ShouldBeNil)
		})
	})
}

func TestServer_Run_when_fail(t *testing.T) {
	Convey("testing server Run", t, func() {
		Convey("testing Run failed with wrong addr", func() {
			s := NewServer("127.0.0.18090")
			var runErr error
			go func() {
				runErr = s.Run()
			}()
			time.Sleep(100 * time.Millisecond)
			So(runErr, ShouldNotBeNil)
		})
	})
}
