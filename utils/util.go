package utils

import (
	"os"
	"strconv"
	"time"
)

func GetData(unixTime string) string {
	int64UnixTime, _ := strconv.ParseInt(unixTime, 10, 64)
	tm := time.Unix(int64UnixTime, 0)
	keyMap := tm.Format("2006-01-02")
	return keyMap
}

func GetFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	return file, err
}
