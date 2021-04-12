package router

import (
	"FloatingBooks/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit () (r *gin.Engine) {
	r = gin.Default()
	api := r.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("", controller.BorrowBook)
			books.GET("/:bookID", controller.GetBookName)
		}
	}
	return
}