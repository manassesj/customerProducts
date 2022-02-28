package models

import (
	"bufio"
	"customerProduts/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Aggregator struct {
}

func (*Aggregator) Execute() {
	file, err := utils.GetFile("./test.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	customerMap := make(map[string]int64)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var payload map[string]interface{}
		err = json.Unmarshal([]byte(scanner.Text()), &payload)
		buildMap(customerMap, payload)
	}
	fmt.Print(customerMap)

	err = buildFile(customerMap, "./teste.txt")
	fmt.Println(err)
}

func buildMap(customerMap map[string]int64, payload map[string]interface{}) {
	key := utils.BuildKey(payload)
	if _, ok := customerMap[key]; !ok {
		customerMap[key] = 1
	} else {
		customerMap[key] += 1
	}
}

func buildFile(customerMap map[string]int64, filePath string) error {
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

func createDataStruct(key string, value int64) DataStruct {
	properties := strings.Split(key, "_")

	return DataStruct{
		Time:     properties[0],
		ClientID: properties[1],
		Group:    properties[2],
		Product:  properties[3],
		Count:    value,
	}

}
