package db

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

type DbOutput struct {
}

func (*DbOutput) FlushData(customerMap map[string]int64, filePath string) error {

	client := GetESClient()
	return nil
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:5601"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}
