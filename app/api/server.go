package api

import (
	"fmt"
	Player "guilherme096/score-savant/templates/Player"

	storage "guilherme096/score-savant/storage"

	Insertions "guilherme096/score-savant/templates/InsertionPages/PlayerInsertion"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
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

	s.storage.Start()
	defer s.storage.Stop()

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/player/:id", func(c echo.Context) error {
		id := c.Param("id")
		player, atts, err := s.storage.LoadPlayerById(id)
		if err != nil {
			return c.String(500, "Internal Server Error")
		}
		if player == nil {
			return c.String(404, "Not Found")
		}

		technical_atts_list := s.storage.GetAttributeList("Technical")
		mental_atts_list := s.storage.GetAttributeList("Mental")
		physical_atts_list := s.storage.GetAttributeList("Physical")
		gk_atts_list := s.storage.GetAttributeList("Goalkeeping")

		var technical_atts []map[string]interface{}
		var mental_atts []map[string]interface{}
		var physical_atts []map[string]interface{}

		fmt.Println(atts)
		// separate the attributes into the respective categories (technical, mental, physical)
		for _, att := range atts {
			ok := false
			for _, att_name := range gk_atts_list {
				if att["att_id"].(string) == att_name {
					fmt.Println("A gk att")
					technical_atts = append(technical_atts, map[string]interface{}{"att_id": att["att_id"].(string), "rating": att["rating"].(int)})
					ok = true
					break
				}
			}
			for _, att_name := range technical_atts_list {
				if att["att_id"].(string) == att_name {
					technical_atts = append(technical_atts, map[string]interface{}{"att_id": att["att_id"].(string), "rating": att["rating"].(int)})
					ok = true
					break
				}
			}
			if !ok {
				for _, att_name := range mental_atts_list {
					if att["att_id"].(string) == att_name {
						mental_atts = append(mental_atts, map[string]interface{}{"att_id": att["att_id"].(string), "rating": att["rating"].(int)})
						ok = true
						break
					}
				}
			}

			if !ok {
				for _, att_name := range physical_atts_list {
					if att["att_id"].(string) == att_name {
						physical_atts = append(physical_atts, map[string]interface{}{"att_id": att["att_id"].(string), "rating": att["rating"].(int)})
						ok = true
						break
					}
				}
			}

		}

		_, PlayerPosition, err := s.storage.GetPlayerPosition(id)
		fmt.Println(PlayerPosition)
		if err != nil {
			return c.String(500, "Internal Server Error")
		}
		return render(c, Player.Player(player, technical_atts, mental_atts, physical_atts, PlayerPosition, nil))
	})

	e.GET("/player-insertion", func(c echo.Context) error {
		return render(c, Insertions.PlayerInsertion())
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
