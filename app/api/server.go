package api

import (
	Player "guilherme096/score-savant/templates/Player"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
	storage "guilherme096/score-savant/storage"
)

type Server struct {
	listen_add string
	storage    storage.IStorage
}

func New_server(listen_add string, storage storage.IStorage) *Server {
	return &Server{
		listen_add: listen_add,
		storage:    storage,
	}
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func (s *Server) Start() {

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/player/:id", func(c echo.Context) error {
		id := c.Param("id")
		player, err := s.storage.LoadPlayerById(id)
		if err != nil {
			return c.String(500, "Internal Server Error")
		}
		if player == nil {
			return c.String(404, "Not Found")
		}
		return render(c, Player.Player(*player))
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
