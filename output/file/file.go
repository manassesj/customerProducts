package file

import (
	"customerProduts/utils"
	"encoding/json"
	"os"
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
		dataStruct := utils.CreateDataStruct(key, value)
		data, _ := json.Marshal(dataStruct)
		_, err2 := f.Write(data)
		f.WriteString("\n")

		if err2 != nil {
			return err2
		}
	}
	return nil
}
