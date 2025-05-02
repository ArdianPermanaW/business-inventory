package product

type Product struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Quantity   int              `json:"quantity"`
	SupplierID int              `json:"supplier_id"`
	Supplier   *SupplierDetails `json:"supplier,omitempty"`
}

type SupplierDetails struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
