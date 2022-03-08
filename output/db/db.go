package db

import (
	"context"
	"customerProduts/utils"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type DbOutput struct {
}

func (*DbOutput) FlushData(customerMap map[string]int64, filePath string) error {

	ctx := context.Background()
	esClient, err := GetESClient()

	if err != nil {
		fmt.Println("Error initializing: ", err)
		panic("Client fail")
	}
	for key, value := range customerMap {
		dataJSON, err := json.Marshal(utils.CreateDataStruct(key, value))
		js := string(dataJSON)
		_, err = esClient.Index().Index("customer_data").BodyJson(js).Do(ctx)

		if err != nil {
			panic(err)
		}
		fmt.Println("[Elastic][InsertCustomer]Insertion Successful: ", js)
	}

	return err
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}
