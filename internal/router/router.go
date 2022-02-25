package router

import (
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/pkg/logger"
	"webconsole/pkg/respcode"

	v1 "webconsole/internal/router/api/v1"
	v2 "webconsole/internal/router/api/v2"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() (r *gin.Engine, err error) {

	if global.ServerSetting.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()

	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(middleware.Cors())

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1路由组
	apiv1 := r.Group("/api/v1")

	// 注册路由
	apiv1.POST("/signup", v1.SignUpHandler)

	// 登录路由
	apiv1.POST("/login", v1.LoginHandler)

	apiv2 := r.Group("/api/v2")
	apiv2.Use(middleware.JWTAuthMiddleware())

	apiv2.Any("/objects/:name", v2.ObjectHandler)

	//apiv2.Group("/locate", locate.Get)

	//metaGroup := r.Group("meta")
	//{
	//	searchGroup := metaGroup.Group("search")
	//	{
	//		searchGroup.GET("/objects", metadata.GetLastestVersions)
	//		searchGroup.GET("/object/:name", metadata.GetAllVersions)
	//	}
	//	thumbnailGroup := metaGroup.Group("thumbnail")
	//	{
	//		thumbnailGroup.PUT("/:name", metadata.PutThumbnail)
	//		thumbnailGroup.GET("/:name", metadata.GetThumbnail)
	//		thumbnailGroup.DELETE("/:name", metadata.DelThumbnail)
	//	}

	//}

	r.NoRoute(func(c *gin.Context) {
		respcode.ResponseNotFound(c)
	})

	return r, nil
}
