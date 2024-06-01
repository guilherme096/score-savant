package Mssql

import (
	"database/sql"
	"fmt"
	"strconv"
)

func (m *MSqlStorage) GetClubList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error) {

	if page < 1 {
		page = 1
	}
	pageNumber := page
	pageSize := amount
	orderBy := filters["order"].(string)
	orderDirection := filters["direction"].(string)
	searchClubName := filters["clubName"].(string)
	searchNationName := filters["nationName"].(string)
	searchLeagueName := filters["leagueName"].(string)
	minWage := filters["minWage"].(float64)
	maxWage := filters["maxWage"].(float64)
	minValue := filters["minValue"].(float64)
	maxValue := filters["maxValue"].(float64)
	minPlayerCount := filters["minPlayerCount"].(int)
	maxPlayerCount := filters["maxPlayerCount"].(int)

	rows, err := m.db.Query(`
        SELECT * FROM dbo.GetClubsWithPagination(
            @PageNumber,
            @PageSize,
            @OrderBy,
            @OrderDirection,
            @SearchClubName,
            @SearchLeagueName,
            @SearchNationName,
            @MinPlayerCount,
            @MaxPlayerCount,
            @MinWageTotal,
            @MaxWageTotal,
            @MinValueTotal,
            @MaxValueTotal
        )`,
		sql.Named("PageNumber", pageNumber),
		sql.Named("PageSize", pageSize),
		sql.Named("OrderBy", orderBy),
		sql.Named("OrderDirection", orderDirection),
		sql.Named("SearchClubName", searchClubName),
		sql.Named("SearchLeagueName", searchLeagueName),
		sql.Named("SearchNationName", searchNationName),
		sql.Named("MinPlayerCount", minPlayerCount),
		sql.Named("MaxPlayerCount", maxPlayerCount),
		sql.Named("MinWageTotal", minWage),
		sql.Named("MaxWageTotal", maxWage),
		sql.Named("MinValueTotal", minValue),
		sql.Named("MaxValueTotal", maxValue),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Players []map[string]interface{} = nil

	for rows.Next() {
		var clubID int
		var clubName, nation, league string
		var playerCount int
		var wageTotal, valueTotal float64

		err := rows.Scan(&clubID, &clubName, &nation, &league, &playerCount, &wageTotal, &valueTotal)
		if err != nil {
			fmt.Println(err)
		}

		Players = append(Players, map[string]interface{}{
			"club_id":      clubID,
			"page_link":    fmt.Sprintf("/club/%d", clubID),
			"name":         clubName,
			"nation":       nation,
			"league":       league,
			"value_total":  valueTotal,
			"wage_total":   wageTotal,
			"player_count": playerCount,
		})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return Players, nil

}

func (m *MSqlStorage) GetClubById(id int) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT * FROM GetClubById(@club_id)", sql.Named("club_id", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var club map[string]interface{}

	if rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		values, err := scanValues(rows, columns)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		club = make(map[string]interface{})

		for i, col := range columns {
			var convertedvalue interface{}
			switch v := values[i].(type) {
			case int64:
				convertedvalue = int(v)
			case []uint8:
				strVal := string(v)
				floatVal, err := strconv.ParseFloat(strVal, 64)
				if err != nil {
					return nil, fmt.Errorf("error converting %s to float64: %v", col, err)
				}
				convertedvalue = floatVal
			default:
				convertedvalue = v
			}
			club[col] = convertedvalue
		}
	} else {
		return nil, fmt.Errorf("club with id %d not found", id)
	}

	return club, nil
}
