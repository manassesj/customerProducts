package customerData

import "customerProduts/models"

type CustomerMapData struct {
	Data map[string]map[string]models.CustomerResult
}

var singleInstanceCustomerData = new()

func GetCustomerDataInstance() CustomerMapData {
	return singleInstanceCustomerData
}

func new() CustomerMapData {
	return CustomerMapData{
		Data: make(map[string]map[string]models.CustomerResult),
	}
}
