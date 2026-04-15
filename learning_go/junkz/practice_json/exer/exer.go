package execr

import "encoding/json"

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type Item struct {
	Name        string  `json:"itemname"`
	Description string  `json:"desc,omitempty"`
	Quantity    int     `json:"qty"`
	Price       float64 `json:"price"`
}

type Order struct {
	TotalPrice  float64 `json:"totalprice,omitempty"`
	IsPaid      bool    `json:"paid"`
	OrderDetail []Item  `json:"orderdetail"`
	Fragile     bool    `json:"-"`
}

type Customer struct {
	Username      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        Address `json:"shipto"`
	PurchaseOrder Order   `json:"order"`
}

func ParseJson(data []byte, cust *Customer) error {
	err := json.Unmarshal(data, cust)
	if err != nil {
		return err
	}
	return nil
}
