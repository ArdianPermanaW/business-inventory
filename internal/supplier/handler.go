package supplier

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetSuppliersHandler(w http.ResponseWriter, r *http.Request) {
	suppliers, err := GetAllSuppliers()
	if err != nil {
		http.Error(w, "Failed to fetch suppliers", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(suppliers)
}

func CreateSupplierHandler(w http.ResponseWriter, r *http.Request) {
	var s Supplier
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := CreateSupplier(s); err != nil {
		http.Error(w, "Failed to create supplier", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Supplier created"))
}

func UpdateSupplierHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s Supplier
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := UpdateSupplier(id, s); err != nil {
		http.Error(w, "Failed to update supplier", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Supplier updated"))
}

func DeleteSupplierHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := DeleteSupplier(id); err != nil {
		http.Error(w, "Failed to delete supplier", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Supplier deleted"))
}
