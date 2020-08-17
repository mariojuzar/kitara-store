package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/request"
	baseResponse "github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	"github.com/mariojuzar/kitara-store/api/service"
	"net/http"
	"time"
)

type OrderController struct {
	service.OrderService
}

func (oc *OrderController) Lock(c *gin.Context)  {
	var response = &baseResponse.BaseResponse{}

	var lockRequest request.LockOrderRequest
	err := c.Bind(&lockRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		response.ServerTime = time.Now()

		c.JSON(http.StatusBadRequest, response)
	} else {
		lockRequest.Products = checkProductRequest(lockRequest.Products)
		order, err := oc.OrderService.LockOrder(lockRequest)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			response.ServerTime = time.Now()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.ServerTime = time.Now()
			response.Data = order

			c.JSON(http.StatusOK, response)
		}
	}
}

// validate all request, if has duplicate the join the duplicate product
func checkProductRequest(productRequest []request.ProductRequest) []request.ProductRequest {
	setProducts := make(map[uint]int)
	var newRequest []request.ProductRequest

	for _, val := range productRequest {
		if setProducts[val.ProductID] == 0 {
			setProducts[val.ProductID] = val.Quantity
		} else {
			setProducts[val.ProductID] = setProducts[val.ProductID] + val.Quantity
		}
	}

	for key, val := range setProducts {
		newRequest = append(newRequest, request.ProductRequest{
			ProductID: key,
			Quantity:  val,
		})
	}

	return newRequest
}