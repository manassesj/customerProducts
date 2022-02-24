package models

type DataStruct struct {
	Id       string `json:"id"`
	ClientID string `json:"client"`
	Group    string `json:"group"`
	Product  string `json:"product"`
	Time     string `json:"time"`
}
