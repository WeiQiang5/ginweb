package routers

import (
	v1 "ginWeb/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(setting.RunMode)
	//新建一个没有任何默认中间件的路由
	r := gin.New()
	//全局中间件
	//Logger中间件将日志写入gin.DefaultWriter,即使你将 GIN_MODE 设置为 release。
	//By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	//Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	v1Api := r.Group("/api/v1")
	{
		//	获取标签列表
		v1Api.GET("/tags", v1.GetTags)
		//	新建标签
		v1Api.POST("/tags", v1.AddTag)
		//	更新制定标签
		v1Api.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		v1Api.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		v1Api.GET("/articles", v1.GetArticles)
		//获取指定文章
		v1Api.GET("/articles/:id", v1.GetArticle)
		//新建文章
		v1Api.POST("/articles", v1.AddArticle)
		//更新指定文章
		v1Api.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		v1Api.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
