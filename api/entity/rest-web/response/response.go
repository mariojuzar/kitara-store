package response

type ProductResponse struct {
	ID				uint				`json:"id"`
	ProductName		string			`json:"product_name"`
	CurrentStock 	int				`json:"current_stock"`
	Store 			StoreResponse	`json:"store"`
}

type StoreResponse struct {
	StoreName 	string	`json:"store_name"`
}