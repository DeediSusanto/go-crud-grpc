package model

type Product struct {
	ID          int32   `json:"id"` // ✅ Ubah dari uint ke int32
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
