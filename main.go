package main

import (
	"customerProduts/customerData"
	"customerProduts/models"
	"os"
)

var customerMapData = customerData.GetCustomerDataInstance().Data

func main() {
	datapath := os.Args[1:][0]
	output := os.Args[1:][1]
	aggregator := new(models.Aggregator)
	aggregator.Execute(datapath, output)

}
