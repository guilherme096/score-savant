package api

import (
	Hello "guilherme096/score-savant/templates/Hello"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

type Server struct {
	listen_add string
}

func New_server(listen_add string) *Server {
	return &Server{
		listen_add: listen_add,
	}
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func (s *Server) Start() {

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return render(c, Hello.Hello())
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
