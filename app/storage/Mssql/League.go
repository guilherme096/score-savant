package Mssql

import (
	"database/sql"
	"fmt"
)

func (m *MSqlStorage) GetLeagueList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	pageNumber := page
	pageSize := amount
	orderBy := filters["order"].(string)
	orderDirection := filters["direction"].(string)
	searchLeagueName := filters["leagueName"].(string)
	searchNationName := filters["nationName"].(string)
	minValue := filters["minValue"].(float64)
	maxValue := filters["maxValue"].(float64)

	var maxValUsed *float64 = nil

	if maxValue != -1 {
		maxValUsed = &maxValue
	}

	fmt.Println(maxValUsed)
	rows, err := m.db.Query(`
        SELECT * FROM dbo.GetLeaguesWithPagination(
            @PageNumber,
            @PageSize,
            @OrderBy,
            @OrderDirection,
            @SearchLeagueName,
            @SearchNationName,
            @MinValueTotal,
            @MaxValueTotal
        )`,
		sql.Named("PageNumber", pageNumber),
		sql.Named("PageSize", pageSize),
		sql.Named("OrderBy", orderBy),
		sql.Named("OrderDirection", orderDirection),
		sql.Named("SearchLeagueName", searchLeagueName),
		sql.Named("SearchNationName", searchNationName),
		sql.Named("MinValueTotal", minValue),
		sql.Named("MaxValueTotal", maxValUsed),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Leagues []map[string]interface{} = nil

	for rows.Next() {
		fmt.Println("here")
		var leagueID int
		var leagueName, nation string
		var valueTotal float64

		err := rows.Scan(&leagueID, &leagueName, &nation, &valueTotal)
		if err != nil {
			fmt.Println(err)
		}

		Leagues = append(Leagues, map[string]interface{}{
			"league_id":   leagueID,
			"page_link":   fmt.Sprintf("/league/%d", leagueID),
			"name":        leagueName,
			"nation":      nation,
			"value_total": valueTotal,
		})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(Leagues)

	return Leagues, nil
}

func (m *MSqlStorage) GetLeagueById(id int) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT * FROM GetLeagueById(@league_id)", sql.Named("league_id", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var league map[string]interface{}

	for rows.Next() {
		var leagueID int
		var leagueName, nation string
		var valueTotal float64
		var totalPlayers int
		var totalClubs int
		var totalWage float64
		var avgAtt float64

		err := rows.Scan(&leagueID, &leagueName, &nation, &totalClubs, &totalPlayers, &valueTotal, &totalWage, &avgAtt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		league = map[string]interface{}{
			"league_id":     leagueID,
			"name":          leagueName,
			"nation":        nation,
			"total_value":   valueTotal,
			"total_clubs":   totalClubs,
			"total_players": totalPlayers,
			"total_wage":    totalWage,
			"avg_att":       avgAtt,
		}
	}

	fmt.Println(league)

	return league, nil
}
