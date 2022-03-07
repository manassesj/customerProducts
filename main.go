package main

import (
	"customerProduts/customerData"
	"customerProduts/output/db"
	"customerProduts/output/file"
	"customerProduts/service"
	"os"
)

var customerMapData = customerData.GetCustomerDataInstance().Data

func main() {
	dataPath := os.Args[1:][0]
	output := os.Args[1:][1]
	timeFormat := os.Args[1:][2]
	aggregator := new(service.Aggregator)
	switch output {
	case "db":
		aggregator.OutputProvider = new(db.DbOutput)
	default:
		aggregator.OutputProvider = new(file.FileOutput)
	}
	aggregator.Execute(dataPath, output, timeFormat)

}
