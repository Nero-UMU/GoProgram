package routers

import (
	v1 "blog-service/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// 新建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 创建文章和标签的路由
	article := v1.NewArticle()
	tag := v1.NewTag()
	// 新建路由组,向其中添加路由
	apiv1 := r.Group("/api/v1")
	{
		// 添加路由
		apiv1.POST("/tags,", tag.Create) // 定义POST请求的路径,这个请求将会找到对应的方法去实现,后面的同理
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
