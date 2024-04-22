package utils

import "fmt"

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
