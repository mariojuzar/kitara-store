package response

type ProductResponse struct {
	ID				uint			`json:"id"`
	ProductName		string			`json:"product_name"`
	CurrentStock 	int				`json:"current_stock"`
	Store 			StoreResponse	`json:"store"`
}

type StoreResponse struct {
	StoreName 	string	`json:"store_name"`
}

type OrderResponse struct {
	OrderID 		int 					`json:"order_id"`
	OrderDetails 	[]OrderDetailResponse	`json:"order_details"`
}

type OrderDetailResponse struct {
	OrderDetailID	int		`json:"order_detail_id"`
	ProductID		int		`json:"product_id"`
	ProductName		string	`json:"product_name"`
	Quantity 		int		`json:"quantity"`
}