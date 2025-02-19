package model

type Lead struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Source  string `json:"source"`
	Status  string `json:"status"`
	Notes   string `json:"notes"`
}
