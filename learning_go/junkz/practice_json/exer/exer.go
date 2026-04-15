package execr

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type Order struct {
	TotalPrice   float64 `json:"TotalPrice"`
	IsPaid       bool    `json:"IsPaid"`
	OrderDetails []Item  `json:"OrderDetails"`
	Fragile      bool    `json:"-"`
}

type Customer struct {
	Username      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        Address `json:"ShipTo"`
	PurchaseOrder Order   `json:"PurchaseOrder"`
}
