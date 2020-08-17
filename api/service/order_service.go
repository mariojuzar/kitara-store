package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mariojuzar/kitara-store/api/entity/model"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/request"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	"github.com/mariojuzar/kitara-store/api/libraries/exception"
)

type OrderService interface {
	LockOrder(request request.LockOrderRequest) (response.OrderResponse, error)
}

type orderService struct {

}

func NewOrderService() OrderService {
	return orderService{}
}

func (o orderService) LockOrder(lockRequest request.LockOrderRequest) (response.OrderResponse, error) {
	var products []model.Product
	storeIds := make(map[uint]uint)

	// begin transaction
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollbackUnlessCommitted()
		}
	}()

	for _, val := range lockRequest.Products {
		product, err := getProduct(val, tx)
		if err == nil {
			products = append(products, product)
			storeIds[product.StoreID] = product.StoreID
		}
	}

	if len(storeIds) > 1 {
		return response.OrderResponse{}, exception.ProductMustInSameStore()
	}

	if len(products) == 0 {
		return response.OrderResponse{}, exception.ProductNotAvailableException()
	} else if  len(products) != len(lockRequest.Products) {
		return response.OrderResponse{}, exception.SomeProductNotAvailableException()
	}

	mapProducts := generateMapProducts(products)
	order, orderDetails, err := generateOrderAndOrderDetail(mapProducts, lockRequest, tx)

	if err != nil {
		return response.OrderResponse{}, err
	}

	// commit transaction
	tx.Commit()

	return constructResponse(order, orderDetails, mapProducts)
}

func generateOrderAndOrderDetail(products map[uint]model.Product, lockRequest request.LockOrderRequest, tx *gorm.DB) (model.Order, []model.OrderDetail, error)  {
	var orderDetails []model.OrderDetail

	order := model.Order{
		BuyerID:  int(lockRequest.UserID),
		Quantity: countAllQuantity(lockRequest.Products),
		StoreID:  int(products[0].StoreID),
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return model.Order{}, nil, err
	}

	for _, val := range lockRequest.Products {
		orderDetail := model.OrderDetail{
			OrderID:   int(order.ID),
			ProductID: int(val.ProductID),
			Quantity:  val.Quantity,
		}

		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			return model.Order{}, nil, err
		}

		product := products[val.ProductID]
		product.CurrentStock =  product.CurrentStock - orderDetail.Quantity
		products[val.ProductID] = product

		orderDetails = append(orderDetails, orderDetail)
	}

	for _, val := range products {
		// last check for available product
		if val.CurrentStock >= 0 {
			if err := tx.Model(&val).Where("id = ?", val.ID).Updates(&val).Error; err != nil {
				tx.Rollback()
				return model.Order{}, nil, err
			}

		} else {
			tx.Rollback()
			return model.Order{}, nil, exception.SomeProductNotAvailableException()
		}
	}

	return order, orderDetails, nil
}


func countAllQuantity(products []request.ProductRequest) int {
	qty := 0
	for _, val := range products {
		qty += val.Quantity
	}
	return qty
}

func generateMapProducts(products []model.Product) map[uint]model.Product  {
	mapProduct := make(map[uint]model.Product)

	for _, val := range products{
		mapProduct[val.ID] = val
	}

	return mapProduct
}


// get product by id and its current stock >= request quantity
func getProduct(request request.ProductRequest, tx *gorm.DB) (model.Product, error) {
	var product model.Product
	tx.Where("id = ? and current_stock >= ?", request.ProductID, request.Quantity).Find(&product)

	if request.ProductID != 0 && request.ProductID == product.ID {
		return product, nil
	} else {
		return model.Product{}, exception.NotExistException()
	}
}

func constructResponse(order model.Order, orderDetails []model.OrderDetail, products map[uint]model.Product) (response.OrderResponse, error) {
	var orderDetailResponses []response.OrderDetailResponse

	for _, val := range orderDetails {
		orderDetailResponse := response.OrderDetailResponse{
			OrderDetailID: int(val.ID),
			ProductID:     val.ProductID,
			ProductName:   products[uint(val.ProductID)].ProductName,
			Quantity:      val.Quantity,
		}
		orderDetailResponses = append(orderDetailResponses, orderDetailResponse)
	}

	orderResponse := response.OrderResponse{
		OrderID:      int(order.ID),
		OrderDetails: orderDetailResponses,
	}

	return orderResponse, nil
}