package Mssql

import (
	"database/sql"
	"fmt"
	Utils "guilherme096/score-savant/utils"
	"strconv"
	"strings"
	"time"
)

func (m *MSqlStorage) LoadPlayerById(id string) (map[string]interface{}, []map[string]interface{}, error) {

	rows, err := m.db.Query("SELECT * FROM GetPlayerById(@player_id)", sql.Named("player_id", id))
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	result := make(map[string]interface{})

	if rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, nil, err
		}

		for i, col := range columns {
			val := values[i]

			if val == nil {
				result[col] = nil
			} else {
				switch v := val.(type) {
				case int64:
					result[col] = int(v)
				case int:
					result[col] = int(v)
				case []uint8:
					// convert []uint8 to string then to float64
					strVal := string(v)
					floatVal, err := strconv.ParseFloat(strVal, 64)
					if err != nil {
						return nil, nil, fmt.Errorf("error converting %s to float64: %v", col, err)
					}
					result[col] = floatVal
				case time.Time:
					result[col] = strings.Split(v.String(), " ")[0]
				default:
					result[col] = val
				}
			}
		}
	} else {
		return nil, nil, fmt.Errorf("player with id %s not found", id)
	}

	attributesRows, err := m.db.Query("SELECT * FROM GetPlayerAttributes(@player_id)", sql.Named("player_id", id))
	if err != nil {
		return nil, nil, err
	}
	defer attributesRows.Close()

	var attributes []map[string]interface{}

	for attributesRows.Next() {
		attributeColumns, err := attributesRows.Columns()
		if err != nil {
			return nil, nil, err
		}

		var attributeValues []interface{}

		attributeValues, err = scanValues(attributesRows, attributeColumns)
		if err != nil {
			return nil, nil, err
		}

		attributeRow := make(map[string]interface{})

		for i, col := range attributeColumns {
			var convertedvalue interface{}
			switch v := attributeValues[i].(type) {
			case int64:
				convertedvalue = int(v)
			default:
				convertedvalue = v
			}
			attributeRow[col] = convertedvalue
		}

		attributes = append(attributes, attributeRow)
	}

	for key, value := range result {
		switch v := value.(type) {
		case int64:
			result[key] = int(v)

		case float64:
			result[key] = v
		}
	}

	return result, attributes, nil
}

func (m *MSqlStorage) GetAttributeList(att_type string) []string {
	query := ""
	switch att_type {
	case "Physical":
		query = "SELECT * FROM Physical_Att"
	case "Mental":
		query = "SELECT * FROM Mental_Att"
	case "Technical":
		query = "SELECT * FROM Technical_Att"
	case "Goalkeeping":
		query = "SELECT * FROM Goalkeeping_Att"
	}

	rows, err := m.db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var atts []string

	for rows.Next() {
		var att string
		err := rows.Scan(&att)
		if err != nil {
			panic(err)
		}
		atts = append(atts, att)
	}

	return atts
}

func (m *MSqlStorage) GetPlayerPosition(id string) (int, string, error) {
	rows, err := m.db.Query("SELECT * FROM GetPositionByPlayerID(@player_id)", sql.Named("player_id", id))
	if err != nil {
		fmt.Println(err)
		return -1, "", err
	}
	defer rows.Close()

	var position_id int
	var position_name string

	if rows.Next() {
		err := rows.Scan(&position_id, &position_name)
		if err != nil {
			return -1, "", err
		}
	} else {
		return -1, "", fmt.Errorf("player with id %s not found", id)
	}

	return position_id, position_name, nil
}

func (m *MSqlStorage) GetRolesByPositionId(PositonId int) []map[string]interface{} {
	rows, err := m.db.Query("SELECT * FROM GetRolesByPositionId(@position_id)", sql.Named("position_id", PositonId))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	var roles []map[string]interface{}

	for rows.Next() {
		var role_id int
		var role_name string
		err := rows.Scan(&role_id, &role_name)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		roles = append(roles, map[string]interface{}{"role_id": role_id, "role_name": role_name})
	}

	return roles
}

func (m *MSqlStorage) GetKeyAttributeList(role_id int) []string {
	rows, err := m.db.Query("SELECT * FROM GetKeyAttributesByRoleId(@role_id)", sql.Named("role_id", role_id))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	var atts []string

	for rows.Next() {
		var att string
		err := rows.Scan(&att)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		atts = append(atts, att)
	}

	return atts
}

func (m *MSqlStorage) GetRoleByPlayerId(player_id int) (string, error) {
	rows, err := m.db.Query("SELECT * FROM GetRoleByPlayerId(@player_id)", sql.Named("player_id", player_id))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer rows.Close()

	var role_id int
	var role_name string

	if rows.Next() {
		err := rows.Scan(&role_id, &role_name)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("player with id %s not found", player_id)
	}

	return role_name, nil
}

func (m *MSqlStorage) GetPlayerList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error) {

	if page < 1 {
		page = 1
	}
	pageNumber := page
	pageSize := amount
	orderBy := filters["order"].(string)
	orderDirection := filters["direction"].(string)
	searchPlayerName := filters["playerName"].(string)
	searchClubName := filters["clubName"].(string)
	searchPositionName := filters["positionName"].(string)
	searchNationName := filters["nationName"].(string)
	searchLeagueName := filters["leagueName"].(string)
	minWage := filters["minWage"].(float64)
	maxWage := filters["maxWage"].(float64)
	minValue := filters["minValue"].(float64)
	maxValue := filters["maxValue"].(float64)
	minAge := filters["minAge"].(int)
	maxAge := filters["maxAge"].(int)
	minReleaseClause := filters["minReleaseClause"].(float64)
	maxReleaseClause := filters["maxReleaseClause"].(float64)

	rows, err := m.db.Query(`
        SELECT * FROM dbo.GetPlayersWithPagination(
            @PageNumber,
            @PageSize,
            @OrderBy,
            @OrderDirection,
            @SearchPlayerName,
            @SearchClubName,
            @SearchPositionName,
            @SearchNationName,
            @SearchLeagueName,
            @MinWage,
            @MaxWage,
            @MinValue,
            @MaxValue,
            @MinAge,
            @MaxAge,
            @MinReleaseClause,
            @MaxReleaseClause
        )`,
		sql.Named("PageNumber", pageNumber),
		sql.Named("PageSize", pageSize),
		sql.Named("OrderBy", orderBy),
		sql.Named("OrderDirection", orderDirection),
		sql.Named("SearchPlayerName", searchPlayerName),
		sql.Named("SearchClubName", searchClubName),
		sql.Named("SearchPositionName", searchPositionName),
		sql.Named("SearchNationName", searchNationName),
		sql.Named("SearchLeagueName", searchLeagueName),
		sql.Named("MinWage", minWage),
		sql.Named("MaxWage", maxWage),
		sql.Named("MinValue", minValue),
		sql.Named("MaxValue", maxValue),
		sql.Named("MinAge", minAge),
		sql.Named("MaxAge", maxAge),
		sql.Named("MinReleaseClause", minReleaseClause),
		sql.Named("MaxReleaseClause", maxReleaseClause),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Players []map[string]interface{} = nil

	for rows.Next() {
		var playerName, position, club, nation, league, url string
		var playerID int
		var wage, value, releaseClause float64
		var age int

		err := rows.Scan(&playerID, &url, &playerName, &position, &club, &wage, &value, &nation, &league, &age, &releaseClause)
		if err != nil {
			fmt.Println(err)
		}

		PlayerValue := ""

		if value < 0 {
			PlayerValue = "Not For Sale"
		} else {
			PlayerValue = Utils.FormatNumber(value)
		}

		Players = append(Players, map[string]interface{}{
			"player_id":      playerID,
			"page_link":      fmt.Sprintf("/player/%d", playerID),
			"name":           playerName,
			"url":            url,
			"position":       position,
			"club":           club,
			"nation":         nation,
			"league":         league,
			"wage":           wage,
			"value":          PlayerValue,
			"age":            age,
			"release_clause": releaseClause,
		})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return Players, nil

}

func (m *MSqlStorage) GetRandomPlayer() (string, string, string, string, int, int, int, error) {
	rows, err := m.db.Query("SELECT * FROM RandomPlayer")
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", -1, -1, -1, err
	}
	defer rows.Close()

	var name, nation, club, url string
	var playerId, nationId, clubId int

	if rows.Next() {
		err := rows.Scan(&playerId, &name, &nationId, &nation, &clubId, &club, &url)
		if err != nil {
			return "", "", "", "", -1, -1, -1, err
		}
	} else {
		return "", "", "", "", -1, -1, -1, fmt.Errorf("no random player found")
	}

	return name, nation, club, url, playerId, nationId, clubId, nil
}

func (m *MSqlStorage) AddPlayer(name string, age int, weight int, height int, nation string, nation_league_id int, league string, club string, foot string, value int, position string, role string, wage float64, contract_end string, release_clause int, atts []string, url string) {

	atts_flat := strings.Join(atts, ",")

	fmt.Println(atts_flat)

	_, err := m.db.Exec("AddPlayer", name, age, weight, height, nation, nation_league_id, league, club, foot, value, "STC", role, wage, contract_end, release_clause, atts_flat, url)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func (m *MSqlStorage) GetStaredPlayers(pageNumber int) ([]map[string]interface{}, error) {
	if pageNumber < 1 {
		pageNumber = 1
	}
	pageSize := 15

	rows, err := m.db.Query(`
        SELECT * FROM dbo.GetStaredPlayersWithPagination(
            @PageNumber,
            @PageSize
        )`,
		sql.Named("PageNumber", pageNumber),
		sql.Named("PageSize", pageSize),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Players []map[string]interface{} = nil

	for rows.Next() {
		var playerName, position, club, nation, league, url string
		var playerID int
		var wage, value, releaseClause float64
		var age int

		err := rows.Scan(&playerID, &playerName, &url, &position, &club, &wage, &value, &nation, &league, &age, &releaseClause)
		if err != nil {
			fmt.Println(err)
		}

		PlayerValue := ""

		if value < 0 {
			PlayerValue = "Not For Sale"
		} else {
			PlayerValue = Utils.FormatNumber(value)
		}

		Players = append(Players, map[string]interface{}{
			"player_id":      playerID,
			"page_link":      fmt.Sprintf("/player/%d", playerID),
			"name":           playerName,
			"url":            url,
			"position":       position,
			"club":           club,
			"nation":         nation,
			"league":         league,
			"wage":           wage,
			"value":          PlayerValue,
			"age":            age,
			"release_clause": releaseClause,
		})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return Players, nil

}

func (m *MSqlStorage) RemovePlayer(id int) {
	_, err := m.db.Exec("DELETE FROM Player WHERE player_id = @player_id", sql.Named("player_id", id))
	if err != nil {
		fmt.Println(err)
	}

	return
}

func (m *MSqlStorage) RemoveStar(id int) {
	_, err := m.db.Exec("RemoveStaredPlayer", id)
	if err != nil {
		fmt.Println(err)
	}

	return
}
