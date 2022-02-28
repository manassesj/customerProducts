package main

import (
	"bufio"
	"customerProduts/customerData"
	"customerProduts/models"
	"customerProduts/utils"
	"encoding/json"
	"fmt"
)

var customerMapData = customerData.GetCustomerDataInstance().Data

func main() {

	file, err := utils.GetFile("./teste.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var dataStruct models.DataStruct
		json.Unmarshal([]byte(scanner.Text()), &dataStruct)
		time := utils.GetData(dataStruct.Time)
		addData(time, dataStruct)
	}
	res2B, _ := json.Marshal(customerMapData)
	fmt.Println(string(res2B))
}

func addData(time string, customerData models.DataStruct) {
	if _, ok := customerMapData[time]; ok {
		addCustomer(time, customerData)
	} else {
		addDate(time, customerData)
	}
}

func addCustomer(time string, customerData models.DataStruct) {
	if customer, ok := customerMapData[time][customerData.ClientID]; ok {
		addCustomerResult(&customer, customerData)
	} else {
		createCustomer(customerMapData[time], customerData)
	}
}

func createCustomer(customer map[string]models.CustomerResult, data models.DataStruct) {
	customer[data.ClientID] = createMapCustomerData(data)
}

func addCustomerResult(customerResult *models.CustomerResult, customerData models.DataStruct) {
	if _, ok := customerResult.UniqueProducts[customerData.Group+"-"+customerData.Product]; !ok {
		customerResult.UniqueProducts[customerData.Group+"-"+customerData.Product] = ""
	}

	customerResult = &models.CustomerResult{
		Count:          customerResult.Count + 1,
		UniqueProducts: customerResult.UniqueProducts,
	}
}

func addDate(time string, customerData models.DataStruct) {
	customerMapData[time] = map[string]models.CustomerResult{customerData.ClientID: createMapCustomerData(customerData)}
}

func createProductData(group string, product string) map[string]string {
	return map[string]string{group + "-" + product: ""}
}

func createMapCustomerData(customerData models.DataStruct) models.CustomerResult {
	customerProductData := createProductData(customerData.Group, customerData.Product)
	return models.CustomerResult{Count: 1, UniqueProducts: customerProductData}
}
