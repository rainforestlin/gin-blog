package routers

import (
	_ "blogWithGin/docs"
	"blogWithGin/middleware/jwt"
	"blogWithGin/pkg/setting"
	v1 "blogWithGin/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", v1.GetAuth)
	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		//get tags
		apiv1.GET("/tags", v1.GetTags)
		//add tag
		apiv1.POST("/tags", v1.AddTag)
		//modify specific tag
		apiv1.PUT("/tags/:id", v1.ModifyTag)
		//delete specific tag
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//	get articles list
		apiv1.GET("/articles", v1.GetArticles)
		//	get specific article
		apiv1.GET("/articles/:id", v1.GetArticle)
		//	add a new article
		apiv1.POST("/articles", v1.AddArticle)
		//	update a specific article
		apiv1.PUT("/articles/:id", v1.ModifyArticle)
		//	delete a specific article
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
