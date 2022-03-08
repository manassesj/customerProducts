package utils

import (
	"customerProduts/models"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetData(unixTime string, format string) string {
	int64UnixTime, _ := strconv.ParseInt(unixTime, 10, 64)
	tm := time.Unix(int64UnixTime, 0)

	switch format {
	case "d":
		t := int64UnixTime / 216000
		result := t * 216000
		tm = time.Unix(result, 0)
		return tm.Format("2006/01/02 15:04")

	case "h":
		t := int64UnixTime / 3600
		result := t * 3600
		tm = time.Unix(result, 0)
		return tm.Format("2006/01/02 15:04")
	case "m":
		t := int64UnixTime / 60
		result := t * 60
		tm = time.Unix(result, 0)
		return tm.Format("2006/01/02 15:04")
	}
	return tm.Format("2006/01/02 15:04")
}

func GetFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	return file, err
}

func BuildKey(payload map[string]interface{}, timeFormat string) string {
	return GetData(payload["time"].(string), timeFormat) + "_" + payload["client"].(string) +
		"_" + payload["group"].(string) + "_" + payload["product"].(string)
}

func CreateDataStruct(key string, value int64) models.DataStruct {
	properties := strings.Split(key, "_")

	return models.DataStruct{
		Time:     properties[0],
		ClientID: properties[1],
		Group:    properties[2],
		Product:  properties[3],
		Count:    value,
	}

}
