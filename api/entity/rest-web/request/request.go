package request

type LockOrderRequest struct {
	UserID 		uint				`json:"user_id" binding:"required"`
	Products	[]ProductRequest	`json:"products" binding:"required"`
}

type ProductRequest struct {
	ProductID 	uint	`json:"product_id" binding:"required"`
	Quantity 	int		`json:"quantity" binding:"required"`
}