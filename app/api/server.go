package api

import (
	"fmt"

	Player "guilherme096/score-savant/templates/Player"

	"strconv"

	storage "guilherme096/score-savant/storage"

	Insertions "guilherme096/score-savant/templates/InsertionPages/PlayerInsertion"

	Search "guilherme096/score-savant/templates/Search"

	Club "guilherme096/score-savant/templates/Club"

	League "guilherme096/score-savant/templates/League"

	Nation "guilherme096/score-savant/templates/Nation"

	Utils "guilherme096/score-savant/utils"

	Home "guilherme096/score-savant/templates/Home"

	TopPlayers "guilherme096/score-savant/templates/TopPlayers"

	Handlers "guilherme096/score-savant/handlers"

	"sort"

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

	api := e.Group("/api")

	api.GET("/list-clubs", func(c echo.Context) error {
		return Handlers.ListClubsHandler(c, s.storage)
	})

	api.GET("/list-leagues", func(c echo.Context) error {
		return Handlers.ListLeaguesHandler(c, s.storage)
	})

	api.GET("/list-nations", func(c echo.Context) error {
		return Handlers.ListNationsHandler(c, s.storage)
	})

	api.GET("/list-players", func(c echo.Context) error {
		return Handlers.ListPlayersHandler(c, s.storage)
	})

	api.POST("/add-player", func(c echo.Context) error {
		return Handlers.AddPlayerHandler(c, s.storage)
	})

	api.GET("/get-random-player", func(c echo.Context) error {
		return Handlers.GetRandomPlayerHandler(c, s.storage)
	})

	api.GET("/list-nations", func(c echo.Context) error {
		return Handlers.ListNationsHandler(c, s.storage)
	})
	api.GET("/star-player", func(c echo.Context) error {
		return Handlers.StarPlayerHandler(c, s.storage)
	})
	api.GET("/list-stared-players", func(c echo.Context) error {
		return Handlers.ListStaredPlayersHandler(c, s.storage)
	})
	api.POST("/player/remove/:id", func(c echo.Context) error {
		return Handlers.RemovePlayerHandler(c, s.storage)
	})
	api.POST("/star/remove/:id", func(c echo.Context) error {
		return Handlers.RemoveStarHandler(c, s.storage)
	})

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

		fmt.Println(player)

		if player["value"] == -1 {
			player["value"] = "N/A"
		} else {
			player["value"] = Utils.FormatNumber(float64(player["value"].(int)))
		}

		if player["release_clause"] == -1 {
			player["release_clause"] = "N/A"
		} else {
			player["release_clause"] = Utils.FormatNumber(float64(player["release_clause"].(int)))
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

		PreferedRole, err := s.storage.GetRoleByPlayerId(player["player_id"].(int))

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

	e.GET("/club/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}

		club, err := s.storage.GetClubById(id)
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, Club.ClubPage(club))
	})

	e.GET("/search-league", func(c echo.Context) error {
		return render(c, Search.LeagueSearchPage())
	})

	e.GET("/search-club", func(c echo.Context) error {
		return render(c, Search.ClubSearchPage())
	})

	e.GET("/search-nation", func(c echo.Context) error {
		return render(c, Search.NationSearchPage())
	})

	e.GET("/league/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		league, err := s.storage.GetLeagueById(id)
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, League.LeaguePage(league))
	})

	e.GET("/", func(c echo.Context) error {
		return render(c, Home.HomePage())
	})

	e.GET("/nation/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		nation, err := s.storage.GetNationById(id)
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}

		if nation["league_names"].([]string)[0] == "" {
			nation["league_names"] = []string{"N/A"}
		}
		return render(c, Nation.NationPage(nation))
	})

	e.GET("/stared-players", func(c echo.Context) error {
		return render(c, Search.GetStaredPlayers())
	})

	e.GET("/top-players", func(c echo.Context) error {
		return render(c, TopPlayers.TopPlayersPage())
	})

	e.Logger.Fatal(e.Start(s.listen_add))
}
