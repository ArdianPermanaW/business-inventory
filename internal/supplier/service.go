package supplier

import (
	"cabbage-inventory/pkg/db"
)

func CreateSupplier(s Supplier) error {
	_, err := db.DB.Exec("INSERT INTO suppliers (name) VALUES ($1)", s.Name)
	return err
}

func GetAllSuppliers() ([]Supplier, error) {
	rows, err := db.DB.Query("SELECT id, name FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []Supplier
	for rows.Next() {
		var s Supplier
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			continue
		}
		suppliers = append(suppliers, s)
	}

	return suppliers, nil
}

func UpdateSupplier(id int, s Supplier) error {
	_, err := db.DB.Exec("UPDATE suppliers SET name = $1 WHERE id = $2", s.Name, id)
	return err
}

func DeleteSupplier(id int) error {
	_, err := db.DB.Exec("DELETE FROM suppliers WHERE id = $1", id)
	return err
}
