package api

import (
	"fmt"
	Player "guilherme096/score-savant/templates/Player"

	storage "guilherme096/score-savant/storage"

	Insertions "guilherme096/score-savant/templates/InsertionPages/PlayerInsertion"

	Search "guilherme096/score-savant/templates/Search"

	Club "guilherme096/score-savant/templates/Club"

	Utils "guilherme096/score-savant/utils"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
	"sort"
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
			fmt.Println(err)
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

		// separate the attributes into the respective categories (technical, mental, physical)
		for _, att := range atts {
			ok := false
			for _, att_name := range gk_atts_list {
				if att["att_id"].(string) == att_name {
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

		PositionId, PlayerPosition, err := s.storage.GetPlayerPosition(id)

		Roles := s.storage.GetRolesByPositionId(PositionId)

		PreferedRole, err := s.storage.GetRoleByPlayerId(PositionId)

		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}

		RoleKeyAtts := make(map[string][]string)

		for _, role := range Roles {
			RoleKeyAtts[role["role_name"].(string)] = s.storage.GetKeyAttributeList(role["role_id"].(int))
		}

		RolesRating := make([]map[string]interface{}, len(Roles))

		for i, role := range Roles {
			RolesRating[i] = map[string]interface{}{
				"role_name":   role["role_name"],
				"role_rating": Utils.CalculateRoleRating(atts, RoleKeyAtts[role["role_name"].(string)]),
			}
		}

		sort.Slice(RolesRating, func(i, j int) bool {
			return RolesRating[i]["role_rating"].(int) > RolesRating[j]["role_rating"].(int)
		})

		if err != nil {
			return c.String(500, "Internal Server Error")
		}
		return render(c, Player.Player(player, technical_atts, mental_atts, physical_atts, PlayerPosition, PreferedRole, RolesRating))
	})

	e.GET("/player-insertion", func(c echo.Context) error {
		return render(c, Insertions.PlayerInsertion())
	})

	e.GET("/search-player", func(c echo.Context) error {
		return render(c, Search.PlayerSearchPage())
	})

	e.GET("/club", func(c echo.Context) error {
		return render(c, Club.ClubPage())
	})

	e.GET("/search-club", func(c echo.Context) error {
		return render(c, Search.ClubSearchPage())
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
