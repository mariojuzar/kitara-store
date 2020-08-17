package service

import (
	"github.com/mariojuzar/kitara-store/api/entity/model"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
)

type ProductService interface {
	GetAllProduct() ([]response.ProductResponse, error)
}

type productService struct {

}

func NewProductService() ProductService {
	return productService{}
}

func (p productService) GetAllProduct() ([]response.ProductResponse, error) {
	var products []model.Product
	var stores []model.Store
	var storeIds []uint
	var productResponses []response.ProductResponse
	mapStores := make(map[uint]model.Store)

	database.DB.Find(&products)

	for _, val := range products{
		storeIds = append(storeIds, val.StoreID)
	}

	database.DB.Where("id IN (?)", storeIds).Find(&stores)

	for _, val := range stores {
		mapStores[val.ID] = val
	}

	for _, product := range products {
		storeResponse := response.StoreResponse{StoreName: mapStores[product.StoreID].StoreName}

		productResponse := response.ProductResponse{
			ID:           product.ID,
			ProductName:  product.ProductName,
			CurrentStock: product.CurrentStock,
			Store:        storeResponse,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}