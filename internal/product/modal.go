package product

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	SupplierID int    `json:"supplier_id"`
}

//huh no getter setter boilerplate thats cool ig
