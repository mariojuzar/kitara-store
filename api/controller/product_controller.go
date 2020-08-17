package controller

import (
	"github.com/gin-gonic/gin"
	baseResponse "github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	"github.com/mariojuzar/kitara-store/api/service"
	"net/http"
	"time"
)

type ProductController struct {
	service.ProductService
}

func (pc *ProductController) GetAllProducts(c *gin.Context)  {
	soldiers, err := pc.ProductService.GetAllProduct()

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = soldiers

		c.JSON(http.StatusOK, response)
	}
}