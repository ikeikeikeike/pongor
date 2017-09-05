package pongor

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/test"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func request(method, path string, e *echo.Echo) (int, string) {
	req := test.NewRequest(method, path, nil)
	rec := test.NewResponseRecorder()
	e.ServeHTTP(req, rec)
	return rec.Status(), rec.Body.String()
}

func TestRenderHtml(t *testing.T) {
	Convey("Render HTML", t, func() {
		e := echo.New()
		r := GetRenderer(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/echo", func() echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return ctx.Render(http.StatusOK, "echo.html", nil)
			}
		}())
		status, body := request("GET", "/echo", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello world</h1>\n")
	})

	Convey("Render HTML with Reload", t, func() {
		e := echo.New()
		r := GetRenderer(PongorOption{
			Directory: "test",
			Reload:    true,
		})
		e.SetRenderer(r)
		e.Get("/echo", func() echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return ctx.Render(http.StatusOK, "echo.html", nil)
			}
		}())
		status, body := request("GET", "/echo", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello world</h1>\n")
	})

	Convey("Render HTML with Context", t, func() {
		e := echo.New()
		r := GetRenderer(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/echo", func() echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return ctx.Render(http.StatusOK, "echo_markup.html", map[string]interface{}{
					"name": "echo",
				})
			}
		}())
		status, body := request("GET", "/echo", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, echo</h1>\n")
	})

	Convey("Render HTML with Context and Reload", t, func() {
		e := echo.New()
		r := GetRenderer(PongorOption{
			Directory: "test",
			Reload:    true,
		})
		e.SetRenderer(r)
		e.Get("/echo", func() echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return ctx.Render(http.StatusOK, "echo_markup.html", map[string]interface{}{
					"name": "echo",
				})
			}
		}())
		status, body := request("GET", "/echo", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, echo</h1>\n")
	})
}

func ExampleRender() {
	e := echo.New()
	r := GetRenderer()
	e.SetRenderer(r)
	e.Get("/", func() echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// render ./templates/index.html file.
			ctx.Render(200, "index.html", map[string]interface{}{
				"title": "你好，世界",
			})
			return nil
		}
	}())
}
