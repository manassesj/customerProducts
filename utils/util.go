package utils

import (
	"os"
	"strconv"
	"time"
)

func GetData(unixTime string) string {
	int64UnixTime, _ := strconv.ParseInt(unixTime, 10, 64)
	tm := time.Unix(int64UnixTime, 0)
	time := tm.Format("2006/01/02")
	return time
}

func GetFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	return file, err
}

func BuildKey(payload map[string]interface{}) string {
	return GetData(payload["time"].(string)) + "_" + payload["client"].(string) +
		"_" + payload["group"].(string) + "_" + payload["product"].(string)
}
