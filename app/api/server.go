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

	e.GET("/api/get-random-player", func(c echo.Context) error {
		name, nation, club, url, playerId, nationId, clubId, err := s.storage.GetRandomPlayer()
		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, Player.RandomPlayerCard(name, nation, club, url, playerId, nationId, clubId))
	})

	e.GET("/api/list-players", func(c echo.Context) error {
		page, page_err := strconv.Atoi(c.QueryParam("page"))
		order := c.QueryParam("sort")
		direction := c.QueryParam("direction")
		playerName := c.QueryParam("playerName")
		clubName := c.QueryParam("clubName")
		positionName := c.QueryParam("positionName")
		nationName := c.QueryParam("nationName")
		leagueName := c.QueryParam("leagueName")
		minWage, err := strconv.ParseFloat(c.QueryParam("minWage"), 64)
		maxWage, err := strconv.ParseFloat(c.QueryParam("maxWage"), 64)
		minValue, err := strconv.ParseFloat(c.QueryParam("minValue"), 64)
		maxValue, err := strconv.ParseFloat(c.QueryParam("maxValue"), 64)
		minAge, err := strconv.Atoi(c.QueryParam("minAge"))
		maxAge, err := strconv.Atoi(c.QueryParam("maxAge"))
		minReleaseClause, err := strconv.ParseFloat(c.QueryParam("minReleaseClause"), 64)
		maxReleaseClause, err := strconv.ParseFloat(c.QueryParam("maxReleaseClause"), 64)

		if order == "" {
			direction = ""
		}

		if page_err != nil {
			page = 0
		}

		if maxAge == 0 {
			maxAge = 99
		}

		if maxWage == 0 {
			maxWage = 99999999999999
		}

		if minValue == 0 {
			minValue = -2
		}

		if maxValue == 0 {
			maxValue = 99999999999999
		}

		if minReleaseClause == 0 {
			minReleaseClause = -2
		}

		if maxReleaseClause == 0 {
			maxReleaseClause = 99999999999999
		}

		filters := make(map[string]interface{})

		filters["playerName"] = playerName
		filters["clubName"] = clubName
		filters["positionName"] = positionName
		filters["nationName"] = nationName
		filters["leagueName"] = leagueName
		filters["minWage"] = minWage
		filters["maxWage"] = maxWage
		filters["minValue"] = minValue
		filters["maxValue"] = maxValue
		filters["minAge"] = minAge
		filters["maxAge"] = maxAge
		filters["minReleaseClause"] = minReleaseClause
		filters["maxReleaseClause"] = maxReleaseClause
		filters["order"] = order
		filters["direction"] = direction

		players, err := s.storage.GetPlayerList(page, 15, filters)

		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, Search.PlayerSearchTable(players))
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

	e.GET("/api/list-clubs", func(c echo.Context) error {
		page, page_err := strconv.Atoi(c.QueryParam("page"))
		order := c.QueryParam("sort")
		direction := c.QueryParam("direction")
		clubName := c.QueryParam("clubName")
		nationName := c.QueryParam("nationName")
		leagueName := c.QueryParam("leagueName")
		minWage, err := strconv.ParseFloat(c.QueryParam("minWage"), 64)
		maxWage, err := strconv.ParseFloat(c.QueryParam("maxWage"), 64)
		minValue, err := strconv.ParseFloat(c.QueryParam("minValue"), 64)
		maxValue, err := strconv.ParseFloat(c.QueryParam("maxValue"), 64)
		minPlayerCount, err := strconv.Atoi(c.QueryParam("minPlayerCount"))
		maxPlayerCount, err := strconv.Atoi(c.QueryParam("maxPlayerCount"))

		if maxPlayerCount == 0 {
			maxPlayerCount = 99999999999999
		}

		if order == "" {
			direction = ""
		}

		if page_err != nil {
			page = 0
		}

		if maxWage == 0 {
			maxWage = 99999999
		}

		if minValue == 0 {
			minValue = -2
		}

		if maxValue == 0 {
			maxValue = 99999999
		}

		filters := make(map[string]interface{})

		filters["clubName"] = clubName
		filters["nationName"] = nationName
		filters["leagueName"] = leagueName
		filters["minWage"] = minWage
		filters["maxWage"] = maxWage
		filters["minValue"] = minValue
		filters["maxValue"] = maxValue
		filters["minPlayerCount"] = minPlayerCount
		filters["maxPlayerCount"] = maxPlayerCount
		filters["order"] = order
		filters["direction"] = direction

		players, err := s.storage.GetClubList(page, 15, filters)

		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, Search.ClubSearchTable(players))
	})

	e.GET("/api/list-leagues", func(c echo.Context) error {
		page, page_err := strconv.Atoi(c.QueryParam("page"))
		order := c.QueryParam("sort")
		direction := c.QueryParam("direction")
		clubName := c.QueryParam("clubName")
		nationName := c.QueryParam("nationName")
		leagueName := c.QueryParam("leagueName")
		minWage, err := strconv.ParseFloat(c.QueryParam("minWage"), 64)
		maxWage, err := strconv.ParseFloat(c.QueryParam("maxWage"), 64)
		minValue, err := strconv.ParseFloat(c.QueryParam("minValue"), 64)
		maxValue, err := strconv.ParseFloat(c.QueryParam("maxValue"), 64)
		minPlayerCount, err := strconv.Atoi(c.QueryParam("minPlayerCount"))
		maxPlayerCount, err := strconv.Atoi(c.QueryParam("maxPlayerCount"))

		if maxPlayerCount == 0 {
			maxPlayerCount = 99999999999999
		}

		if order == "" {
			direction = ""
		}

		if page_err != nil {
			page = 0
		}

		if maxWage == 0 {
			maxWage = 99999999
		}

		if minValue == 0 {
			minValue = -2
		}

		if maxValue == 0 {
			maxValue = -1
		}

		filters := make(map[string]interface{})

		filters["clubName"] = clubName
		filters["nationName"] = nationName
		filters["leagueName"] = leagueName
		filters["minWage"] = minWage
		filters["maxWage"] = maxWage
		filters["minValue"] = minValue
		filters["maxValue"] = maxValue
		filters["minPlayerCount"] = minPlayerCount
		filters["maxPlayerCount"] = maxPlayerCount
		filters["order"] = order
		filters["direction"] = direction

		players, err := s.storage.GetLeagueList(page, 15, filters)

		fmt.Println(players)

		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}
		return render(c, Search.LeagueSearchTable(players))
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

	e.POST("/api/add-player", func(c echo.Context) error {
		playerType := c.FormValue("playerType")

		var playerName, playerUrl, playerFoot, playerNationality, contractEnd, playerClub string
		var playerAge, playerHeight, playerWeight int
		var playerWage, playerValue, playerReleaseClause float64

		playerName = c.FormValue("playerName")
		playerUrl = c.FormValue("playerUrl")
		playerFoot = c.FormValue("playerFoot")
		playerNationality = c.FormValue("playerNationality")
		contractEnd = c.FormValue("contractEnd")
		playerClub = c.FormValue("playerClub")
		playerAge, _ = strconv.Atoi(c.FormValue("playerAge"))
		playerHeight, _ = strconv.Atoi(c.FormValue("playerHeight"))
		playerWeight, _ = strconv.Atoi(c.FormValue("playerWeight"))
		playerWage, _ = strconv.ParseFloat(c.FormValue("playerWage"), 64)
		playerValue, _ = strconv.ParseFloat(c.FormValue("playerValue"), 64)
		playerReleaseClause, _ = strconv.ParseFloat(c.FormValue("playerReleaseClause"), 64)

		fmt.Println(playerName, playerUrl, playerFoot, playerNationality, contractEnd, playerClub, playerAge, playerHeight, playerWeight, playerWage, playerValue, playerReleaseClause)

		mental_atts_list := s.storage.GetAttributeList("Mental")
		physical_atts_list := s.storage.GetAttributeList("Physical")

		var atts []string

		for _, att_name := range mental_atts_list {
			rating, _ := strconv.Atoi(c.FormValue(att_name))
			if rating == 0 {
				continue
			}
			atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
		}

		for _, att_name := range physical_atts_list {
			rating, _ := strconv.Atoi(c.FormValue(att_name))
			if rating == 0 {
				continue
			}
			atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
		}

		if playerType == "Goalkeeper" {
			gk_atts_list := s.storage.GetAttributeList("Goalkeeping")

			for _, att_name := range gk_atts_list {
				rating, _ := strconv.Atoi(c.FormValue(att_name))
				if rating == 0 {
					continue
				}
				atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
			}

		}

		if playerType == "Outfield" {
			technical_atts_list := s.storage.GetAttributeList("Technical")

			for _, att_name := range technical_atts_list {
				rating, _ := strconv.Atoi(c.FormValue(att_name))
				if rating == 0 {
					continue
				}
				atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
			}

		}

		s.storage.AddPlayer(playerName, playerAge, playerWeight, playerHeight, playerNationality, 1, "Premier League", playerClub, playerFoot, int(playerValue), playerType, "Poacher (Attack)", playerWage, contractEnd, int(playerReleaseClause), atts, playerUrl)
		return c.String(200, "OK")
	})

	e.GET("/api/delete-player", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		s.storage.DeletePlayer(id)
		return c.String(200, "OK")
	})

	e.GET("/api/list-nations", func(c echo.Context) error {
		page, page_err := strconv.Atoi(c.QueryParam("page"))
		order := c.QueryParam("sort")
		direction := c.QueryParam("direction")
		nationName := c.QueryParam("nationName")
		minValue, err := strconv.ParseFloat(c.QueryParam("minValue"), 64)
		maxValue, err := strconv.ParseFloat(c.QueryParam("maxValue"), 64)

		if order == "" {
			direction = ""
		}

		if page_err != nil {
			page = 1
		}

		if minValue == 0 {
			minValue = -2
		}

		if maxValue == 0 {
			maxValue = 99999999
		}

		filters := make(map[string]interface{})

		filters["nationName"] = nationName
		filters["minValue"] = minValue
		filters["maxValue"] = maxValue
		filters["order"] = order
		filters["direction"] = direction

		nations, err := s.storage.GetNationList(page, 15, filters)

		if err != nil {
			fmt.Println(err)
			return c.String(500, "Internal Server Error")
		}

		return render(c, Search.NationSearchTable(nations))

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
		return render(c, Nation.NationPage(nation))
	})
	e.Logger.Fatal(e.Start(s.listen_add))
}
