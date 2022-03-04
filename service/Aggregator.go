package service

import (
	"bufio"
	"customerProduts/output"
	"customerProduts/utils"
	"encoding/json"
	"fmt"
)

type Aggregator struct {
	OutputProvider output.OutputProvider
}

func (aggregator *Aggregator) Execute(dataPath string, output string, timeFormat string) {
	file, err := utils.GetFile(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	customerMap := make(map[string]int64)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var payload map[string]interface{}
		err = json.Unmarshal([]byte(scanner.Text()), &payload)
		buildMap(customerMap, payload, timeFormat)
	}

	err = aggregator.OutputProvider.FlushData(customerMap, output)

}

func buildMap(customerMap map[string]int64, payload map[string]interface{}, timeFormat string) {
	key := utils.BuildKey(payload, timeFormat)
	if _, ok := customerMap[key]; !ok {
		customerMap[key] = 1
	} else {
		customerMap[key] += 1
	}
}
