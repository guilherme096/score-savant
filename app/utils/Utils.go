package utils

import (
	"fmt"
	"strconv"
)

func FormatNumber(num float64) string {
	switch {
	case num >= 1000000000:
		return fmt.Sprintf("%.1fB", num/1000000000)
	case num >= 1000000:
		return fmt.Sprintf("%.1fM", num/1000000)
	case num >= 1000:
		return fmt.Sprintf("%.1fk", num/1000)
	default:
		return fmt.Sprintf("%.1f", num)
	}
}

func AttributeColor(str_val string) string {
	value, err := strconv.Atoi(str_val)
	prefix := "text-"
	color := "black"
	if err != nil {
		color = "danger"
	}
	switch {
	case value >= 16:
		color = "success"
	case value >= 12:
		color = "info"
	case value >= 8:
		color = "warning"
	default:
		color = "danger"
	}
	return prefix + color
}

func CalculateRoleRating(atts_ratings []map[string]interface{}, role_key_atts []string) int {
	role_rating := 0
	lenght := len(role_key_atts)
	for _, att := range atts_ratings {
		for _, key_att := range role_key_atts {
			if att["att_id"].(string) == key_att {
				role_rating += att["rating"].(int)
				break
			}
		}
	}
	return role_rating / lenght
}
