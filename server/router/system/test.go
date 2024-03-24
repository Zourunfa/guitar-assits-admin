package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (t *TestRouter) TestUserRouter(Router *gin.RouterGroup) {
	TestUserRouter := Router.Group("test").Use(middleware.OperationRecord())
	TestApi := v1.ApiGroupApp.SystemApiGroup.TestApi
	{
		TestUserRouter.POST("testT", TestApi.TestT)
	}
}
