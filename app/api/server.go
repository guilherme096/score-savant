package api

import (
	"net/http"

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

func (s *Server) Start() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi there :)")
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
