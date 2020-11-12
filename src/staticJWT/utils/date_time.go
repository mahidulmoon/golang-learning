package utils

import (
	"strings"
	"time"
)

func FormatTimeStamp(inputTime string) string {

	splitted := strings.Split(inputTime, ".")
	t, _ := time.Parse("2006-01-02 15:04:05", splitted[0])

	return t.Format("2006-01-02 15:04:05")
}
