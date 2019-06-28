package model

// FooSvc is struct which provided to service layer and handle layer
type FooSvc struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Bar       string `json:"bar"`
}

// FooCreateForm is struct which provided to service layer and handle layer
type FooCreateForm struct {
	Bar string `form:"bar" binding:"required"` // binding tag for HTTP frame (like: Gin)
}
