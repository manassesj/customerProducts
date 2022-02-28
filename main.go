package main

import (
	"customerProduts/customerData"
	"customerProduts/models"
)

var customerMapData = customerData.GetCustomerDataInstance().Data

func main() {
	aggregator := new(models.Aggregator)
	aggregator.Execute()
}
