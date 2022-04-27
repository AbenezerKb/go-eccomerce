package entity

// import "gorm.io/gorm"

type Order struct {
	ID             string `json:"id" gorm:"primary_key; auto_increment" `
	OrderOwner     string `json:"orderowner" binding:"required"`
	Items_quantity int `json:"items_quantity" binding:"required"`
	Total_price    float64 `json:"total_price" binding:"required"`
}

type UpdateOrder struct {
	ID             string `json:"id" `
	OrderOwner     string `json:"orderowner"`
	Items_quantity int `json:"items_quantity" `
	Total_price    float64 `json:"total_price" `
}
