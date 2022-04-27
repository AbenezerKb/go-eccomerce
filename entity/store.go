package entity

type Store struct {
	ID       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	Image    string `json:"image" binding:"required"`
	Status   string `json:"status" binding:"required"`
	// Items    []Item
}

type UpdateStore struct {
	ID       string `json:"id"  `
	Name     string `json:"name" `
	Location string `json:"location" `
	Image    string `json:"image" `
	Status   string `json:"status" `
	// Items    []Item
}

type StoreRole struct{
		ClerkID string `json:"clerkid" binding:"required"`
		StoreID string `json:"storeid" binding:"required"`
		Role string `json:"role" binding:"required"`
		Resource string `json:"resource" binding:"required"`
		Operation string `json:"operation" binding:"required"`

}
