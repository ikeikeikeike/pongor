package main

import (
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	serv := echo.New()
	r := pongor.GetRenderer()
	// r := pongor.GetRenderer(pongor.PongorOption{
	// 	Reload: true, // if you want to reload template every request, set Reload to true.
	// })
	serv.SetRenderer(r)
	serv.Static("/static", "./static")
	serv.Get("/", func() echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Render(200, "index.html", map[string]interface{}{
				"title": "你好，世界",
			})
			return nil
		}
	}())

	serv.Run(standard.New("127.0.0.1:9000"))
}
