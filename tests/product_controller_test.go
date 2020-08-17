package tests

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/mariojuzar/kitara-store/api/controller"
	"github.com/mariojuzar/kitara-store/api/entity/path"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	mock_service "github.com/mariojuzar/kitara-store/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProducts(t *testing.T)  {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productSvc := mock_service.NewMockProductService(ctrl)
	productCtrl := controller.ProductController{ProductService: productSvc}

	t.Run("Get All Success", func(t *testing.T) {
		productSvc.EXPECT().GetAllProduct().Return([]response.ProductResponse{}, nil)
		req := httptest.NewRequest("GET", path.BaseUrl + path.Product, nil)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		productCtrl.GetAllProducts(ctx)

		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("Get All Failed", func(t *testing.T) {
		productSvc.EXPECT().GetAllProduct().Return([]response.ProductResponse{}, errors.New("error happened"))
		req := httptest.NewRequest("GET", path.BaseUrl + path.Product, nil)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		productCtrl.GetAllProducts(ctx)

		assert.Equal(t, rec.Code, http.StatusInternalServerError)
	})
}