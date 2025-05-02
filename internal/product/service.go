package product

import (
	"log"

	"cabbage-inventory/pkg/db"
)

func CreateProduct(p Product) error {
	query := `INSERT INTO products (name, quantity, supplier_id) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(query, p.Name, p.Quantity, p.SupplierID)
	return err
}

func GetAllProducts() ([]Product, error) {
	query := `
		SELECT 
			p.id, p.name, p.quantity, p.supplier_id,
			s.id, s.name
		FROM products p
		LEFT JOIN suppliers s ON p.supplier_id = s.id
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		var s SupplierDetails
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.SupplierID, &s.ID, &s.Name); err != nil {
			log.Println("Error scanning product:", err)
			continue
		}
		p.Supplier = &s
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(id int, p Product) error {
	query := `UPDATE products SET name = $1, quantity = $2, supplier_id = $3 WHERE id = $4`
	_, err := db.DB.Exec(query, p.Name, p.Quantity, p.SupplierID, id)
	return err
}

func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}
