package file

import (
	"customerProduts/models"
	"encoding/json"
	"os"
	"strings"
)

type FileOutput struct {
}

func (*FileOutput) FlushData(customerMap map[string]int64, filePath string) error {
	f, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer f.Close()

	for key, value := range customerMap {
		dataStruct := createDataStruct(key, value)
		data, _ := json.Marshal(dataStruct)
		_, err2 := f.Write(data)
		f.WriteString("\n")

		if err2 != nil {
			return err2
		}
	}
	return nil
}

func createDataStruct(key string, value int64) models.DataStruct {
	properties := strings.Split(key, "_")

	return models.DataStruct{
		Time:     properties[0],
		ClientID: properties[1],
		Group:    properties[2],
		Product:  properties[3],
		Count:    value,
	}

}
