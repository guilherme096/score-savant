package handlers

import (
	"fmt"
	Player "guilherme096/score-savant/templates/Player"
	"guilherme096/score-savant/templates/Search"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
	. "guilherme096/score-savant/storage"
)

func ListClubsHandler(c echo.Context, s IStorage) error {
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
		maxPlayerCount = 99999999
	}

	if order == "" {
		direction = ""
	}

	if page_err != nil {
		page = 0
	}

	if maxWage == 0 {
		maxWage = 99999999999999.00
	}

	if minValue == 0 {
		minValue = -2
	}

	if maxValue == 0 {
		maxValue = 99999999999999.00
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

	players, err := s.GetClubList(page, 15, filters)

	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}
	return render(c, Search.ClubSearchTable(players))
}

func ListLeaguesHandler(c echo.Context, s IStorage) error {
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

	players, err := s.GetLeagueList(page, 15, filters)

	fmt.Println(players)

	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}
	return render(c, Search.LeagueSearchTable(players))

}

func AddPlayerHandler(c echo.Context, s IStorage) error {
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

	mental_atts_list := s.GetAttributeList("Mental")
	physical_atts_list := s.GetAttributeList("Physical")

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
		gk_atts_list := s.GetAttributeList("Goalkeeping")

		for _, att_name := range gk_atts_list {
			rating, _ := strconv.Atoi(c.FormValue(att_name))
			if rating == 0 {
				continue
			}
			atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
		}

	}

	if playerType == "Outfield" {
		technical_atts_list := s.GetAttributeList("Technical")

		for _, att_name := range technical_atts_list {
			rating, _ := strconv.Atoi(c.FormValue(att_name))
			if rating == 0 {
				continue
			}
			atts = append(atts, fmt.Sprintf("%s:%d", att_name, rating))
		}

	}

	s.AddPlayer(playerName, playerAge, playerWeight, playerHeight, playerNationality, 1, "Premier League", playerClub, playerFoot, int(playerValue), playerType, "Poacher (Attack)", playerWage, contractEnd, int(playerReleaseClause), atts, playerUrl)
	return c.String(200, "OK")

}

func ListNationsHandler(c echo.Context, s IStorage) error {
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

	nations, err := s.GetNationList(page, 15, filters)

	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}

	return render(c, Search.NationSearchTable(nations))

}

func StarPlayerHandler(c echo.Context, s IStorage) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	s.StarPlayer(id)
	return c.String(200, "OK")
}

func ListPlayersHandler(c echo.Context, s IStorage) error {
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

	players, err := s.GetPlayerList(page, 15, filters)

	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}
	return render(c, Search.PlayerSearchTable(players))
}

func GetRandomPlayerHandler(c echo.Context, s IStorage) error {
	name, nation, club, url, playerId, nationId, clubId, err := s.GetRandomPlayer()
	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}
	return render(c, Player.RandomPlayerCard(name, nation, club, url, playerId, nationId, clubId))
}

func ListStaredPlayersHandler(c echo.Context, s IStorage) error {
	page, page_err := strconv.Atoi(c.QueryParam("page"))

	if page_err != nil {
		page = 1
	}

	players, err := s.GetStaredPlayers(page)

	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}

	return render(c, Search.StaredPlayersTable(players))
}

func RemovePlayerHandler(c echo.Context, s IStorage) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	s.RemovePlayer(id)
	return c.String(200, "OK")
}

func RemoveStarHandler(c echo.Context, s IStorage) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	s.RemoveStar(id)
	return c.String(200, "Star Removed")
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
