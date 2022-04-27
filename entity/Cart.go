package entity

// import "gorm.io/gorm"

type Cart struct {
	ID             string `json:"id" gorm:"primary_key; auto_increment" `
	OrderOwner     string `json:"orderowner" binding:"required"`
	Items_quantity int `json:"items_quantity" `
	Total_price    float64 `json:"total_price" `
}

