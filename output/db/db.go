package db

import (
	"context"
	"customerProduts/models"
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
		fmt.Println("Erro initializing: ", err)
		panic("Client fail")
	}

	newDataStruct := models.DataStruct{
		ClientID: "1",
		Product:  "1",
		Group:    "1",
		Time:     "2022/02/11 21:00",
		Count:    1,
	}

	dataJSON, err := json.Marshal(newDataStruct)
	js := string(dataJSON)
	_, err = esClient.Index().Index("customer_data").BodyJson(js).Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertCustomer]Insertion Successful")

	return err
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}
