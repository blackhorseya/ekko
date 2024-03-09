package timex

import (
	"strconv"
	"strings"
)

// FormatTaiwanDate is a function to format Taiwan date.
func FormatTaiwanDate(date string) string {
	split := strings.Split(date, "/")
	if len(split) != 3 {
		return ""
	}

	year, _ := strconv.Atoi(split[0])

	return strconv.Itoa(year+1911) + "/" + split[1] + "/" + split[2]
}
