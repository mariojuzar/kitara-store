package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/mariojuzar/kitara-store/api/controller"
	"github.com/mariojuzar/kitara-store/api/entity/path"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	"github.com/mariojuzar/kitara-store/api/libraries/exception"
	mock_service "github.com/mariojuzar/kitara-store/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLockOrder(t *testing.T)  {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock_service.NewMockOrderService(ctrl)
	orderCtrl := controller.OrderController{OrderService: orderSvc}

	t.Run("Order Lock Request Bind Failed", func(t *testing.T) {
		req := httptest.NewRequest("POST", path.BaseUrl + path.OrderLock, bytes.NewBufferString("{}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.Lock(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Order Lock Some Product Not Available", func(t *testing.T) {
		orderSvc.EXPECT().LockOrder(gomock.Any()).Return(response.OrderResponse{}, exception.SomeProductNotAvailableException())
		req := httptest.NewRequest("POST", path.BaseUrl + path.OrderLock, bytes.NewBufferString("{\"user_id\":2,\"products\":[{\"product_id\":1,\"quantity\":2},{\"product_id\":2,\"quantity\":1}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.Lock(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Order Lock All Product Not Available", func(t *testing.T) {
		orderSvc.EXPECT().LockOrder(gomock.Any()).Return(response.OrderResponse{}, exception.ProductNotAvailableException())
		req := httptest.NewRequest("POST", path.BaseUrl + path.OrderLock, bytes.NewBufferString("{\"user_id\":2,\"products\":[{\"product_id\":1,\"quantity\":2},{\"product_id\":2,\"quantity\":1}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.Lock(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Order Lock Product Must Be In The Same Store", func(t *testing.T) {
		orderSvc.EXPECT().LockOrder(gomock.Any()).Return(response.OrderResponse{}, exception.ProductMustInSameStore())
		req := httptest.NewRequest("POST", path.BaseUrl + path.OrderLock, bytes.NewBufferString("{\"user_id\":2,\"products\":[{\"product_id\":1,\"quantity\":2},{\"product_id\":2,\"quantity\":1}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.Lock(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Order Lock Success", func(t *testing.T) {
		orderSvc.EXPECT().LockOrder(gomock.Any()).Return(response.OrderResponse{}, nil)
		req := httptest.NewRequest("POST", path.BaseUrl + path.OrderLock, bytes.NewBufferString("{\"user_id\":2,\"products\":[{\"product_id\":1,\"quantity\":2},{\"product_id\":2,\"quantity\":1}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.Lock(ctx)

		assert.Equal(t, rec.Code, http.StatusOK)
	})


}