package models

type Order struct {
    Uuid   		string `json:"uuid"`
    TaxTotal  int 	 `json:"taxTotal"`
    Total 		string `json:"total"`
}