package Mssql

import (
	"database/sql"
	"fmt"
	"strings"
)

func (m *MSqlStorage) GetNationList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	pageNumber := page
	pageSize := amount
	orderBy := filters["order"].(string)
	orderDirection := filters["direction"].(string)
	searchNationName := filters["nationName"].(string)
	minValue := filters["minValue"].(float64)
	//maxValue := filters["maxValue"].(float64)

	// Execute the function
	rows, err := m.db.Query(`
        SELECT * FROM dbo.GetNationsWithPagination(
            @PageNumber,
            @PageSize,
            @OrderBy,
            @OrderDirection,
            @SearchNationName,
            @MinValueTotal,
            @MaxValueTotal
        )`,
		sql.Named("PageNumber", pageNumber),
		sql.Named("PageSize", pageSize),
		sql.Named("OrderBy", orderBy),
		sql.Named("OrderDirection", orderDirection),
		sql.Named("SearchNationName", searchNationName),
		sql.Named("MinValueTotal", minValue),
		sql.Named("MaxValueTotal", nil),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Nations []map[string]interface{} = nil

	// Process the results
	for rows.Next() {
		var nationID int
		var nationName string
		var valueTotal float64

		err := rows.Scan(&nationID, &nationName, &valueTotal)
		if err != nil {
			fmt.Println(err)
		}

		Nations = append(Nations, map[string]interface{}{
			"nation_id":   nationID,
			"page_link":   fmt.Sprintf("/nation/%d", nationID),
			"name":        nationName,
			"value_total": valueTotal,
		})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return Nations, nil
}

func (m *MSqlStorage) GetNationById(id int) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT * FROM GetNationById(@nation_id)", sql.Named("nation_id", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var nation map[string]interface{}

	for rows.Next() {
		var nationID int
		var nationName string
		var valueTotal float64
		var totalPlayers int
		var totalLeagues int
		var leagueNames string

		err := rows.Scan(&nationID, &nationName, &totalLeagues, &leagueNames, &valueTotal)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		nation = map[string]interface{}{
			"nation_id":     nationID,
			"name":          nationName,
			"total_leagues": totalLeagues,
			"total_players": totalPlayers,
			"total_value":   valueTotal,
			"league_names":  strings.Split(leagueNames, ","),
		}
	}

	fmt.Println(nation)

	return nation, nil
}

func (m *MSqlStorage) StarPlayer(id int) {
	_, err := m.db.Exec("AddStaredPlayer", id)
	if err != nil {
		fmt.Println(err)
	}

	return
}
