package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/kitara-store/api/configuration"
	"github.com/mariojuzar/kitara-store/api/controller"
	"github.com/mariojuzar/kitara-store/api/entity/path"
	"github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	"github.com/mariojuzar/kitara-store/api/service"
	"net/http"
	"time"
)

func Run() *gin.Engine {
	engine := gin.Default()
	engine.RedirectTrailingSlash = false

	configuration.Initialize()

	var dbSvc = service.NewDatabaseService()
	_ = dbSvc.Initialize()

	var productCtrl = controller.ProductController{ProductService: service.NewProductService()}
	var orderCtrl = controller.OrderController{OrderService: service.NewOrderService()}

	api := engine.Group(path.BaseUrl)
	{
		api.GET(path.Product, productCtrl.GetAllProducts)

		api.POST(path.OrderLock, orderCtrl.Lock)
	}

	engine.NoRoute(func(context *gin.Context) {
		var resp = &response.BaseResponse{
			ServerTime:	time.Now(),
		}

		resp.Code = http.StatusNotFound
		resp.Message = "Route not found"

		context.JSON(http.StatusNotFound, resp)
	})

	return engine
}
