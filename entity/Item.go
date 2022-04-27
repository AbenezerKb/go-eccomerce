package entity

// import "gorm.io/gorm"

type Item struct {
	ID       string `json:"id" gorm:"primary_key" `
	Name string `json:"name" binding:"required"`
	Image string `json:"image" `
	StoreID string `json:"storeid" `
}
