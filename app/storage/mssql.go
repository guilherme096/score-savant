package storage

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	Utils "guilherme096/score-savant/utils"

	_ "github.com/microsoft/go-mssqldb"
)

type MSqlStorage struct {
	connectionString string
	db               *sql.DB
}

func NewMSqlStorage(username string, password string, host string, port int, databaseName string) *MSqlStorage {

	conString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, username, password, port, databaseName)

	return &MSqlStorage{connectionString: conString}
}

func (m *MSqlStorage) Start() {
	fmt.Printf("Connecting to SQL Server: %s\n", m.connectionString)
	db, err := sql.Open("sqlserver", m.connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to SQL Server")
	m.db = db
}

func (m *MSqlStorage) Stop() {
	m.db.Close()
	fmt.Println("Disconnected from SQL Server")
}

func (m *MSqlStorage) LoadPlayerById(id string) (map[string]interface{}, []map[string]interface{}, error) {

	// execute stored procedure
	rows, err := m.db.Query("SELECT * FROM GetPlayerById(@player_id)", sql.Named("player_id", id))
	// close when the function ends
	defer rows.Close()

	// get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// create a slice of interfaces to store the values from the database
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	result := make(map[string]interface{})

	// get the values from each row
	if rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, nil, err
		}

		for i, col := range columns {
			val := values[i]

			// If the value is nil, set it to a zero value
			if val == nil {
				result[col] = nil
			} else {
				switch v := val.(type) {
				case int64:
					result[col] = int(v)
				case int:
					result[col] = int(v)
				case []uint8:
					// Convert []uint8 to string then to float64
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

	// Execute query to fetch player attributes
	attributesRows, err := m.db.Query("SELECT * FROM GetPlayerAttributes(@player_id)", sql.Named("player_id", id))
	if err != nil {
		return nil, nil, err
	}
	defer attributesRows.Close()

	// Initialize a slice of maps to hold the attribute data
	var attributes []map[string]interface{}

	// Iterate over each row of attributes
	for attributesRows.Next() {
		// Get column names for attributes
		attributeColumns, err := attributesRows.Columns()
		if err != nil {
			return nil, nil, err
		}

		// Create a slice to hold the values of each attribute row
		var attributeValues []interface{}

		// Get the values from the current row
		attributeValues, err = scanValues(attributesRows, attributeColumns)
		if err != nil {
			return nil, nil, err
		}

		// Create a map for the current attribute row
		attributeRow := make(map[string]interface{})

		// Populate attributeRow map with column names and values
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

		// Add attributeRow to the slice of attributes
		attributes = append(attributes, attributeRow)
	}

	// convert int64 to int if needed
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

	// Execute the function
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

	// Process the results
	for rows.Next() {
		var playerID int
		var playerName, position, club, nation, league, url string
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
	//minWage := filters["minWage"].(float64)
	//maxWage := filters["maxWage"].(float64)
	minValue := filters["minValue"].(float64)
	maxValue := filters["maxValue"].(float64)
	//minPlayerCount := filters["minPlayerCount"].(int)
	//maxPlayerCount := filters["maxPlayerCount"].(int)

	// Execute the function
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
		sql.Named("MinPlayerCount", -1),
		sql.Named("MaxPlayerCount", 20),
		sql.Named("MinWageTotal", -1.00),
		sql.Named("MaxWageTotal", 99999999.00),
		sql.Named("MinValueTotal", minValue),
		sql.Named("MaxValueTotal", maxValue),
	)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var Players []map[string]interface{} = nil

	// Process the results
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
	// Execute the function
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

	// Process the results
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

func (m *MSqlStorage) AddPlayer(name string, age int, weight int, height int, nation string, nation_league_id int, league string, club string, foot string, value int, position string, role string, wage float64, contract_end string, release_clause int, atts []string, url string) {
	atts_flat := "Corners:8,Crossing:8,Dribbling:8,Finishing:8,First_Touch:8,Free_Kick_Taking:8,Heading:8,Long_Shots:8,Long_Throws:8,Marking:8,Passing:8,Penalty_Taking:8,Tackling:8,Technique:8,Aggression:8,Anticipation:8,Bravery:8,Composure:8,Concentration:8,Decisons:8,Determination:8,Flair:8,Leadership:8,Off_The_Ball:8,Positioning:8,Teamwork:8,Vision:8,Work_Rate:8,Acceleration:8,Agility:8,Balance:8,Jumping_Reach:8,Natural_Fitness:8,Pace:8,Stamina:8,Strength:8"

	fmt.Println(atts_flat)

	_, err := m.db.Exec("AddPlayer", name, age, weight, height, nation, nation_league_id, league, club, foot, value, "STC", role, wage, contract_end, release_clause, atts_flat, url)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func (m *MSqlStorage) DeletePlayer(id int) {
	_, err := m.db.Exec("DELETE FROM Player WHERE PlayerID = @player_id", sql.Named("player_id", id))
	if err != nil {
		fmt.Println(err)
	}

	return
}

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

// Function to scan values from a row into a slice of interfaces
func scanValues(rows *sql.Rows, columns []string) ([]interface{}, error) {
	// Create a slice to hold the values of each row
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Scan the values from the row into the slice of interfaces
	err := rows.Scan(valuePtrs...)
	if err != nil {
		return nil, err
	}

	return values, nil
}
