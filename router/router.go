package router

import (
	"FloatingBooks/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit () (r *gin.Engine) {
	r = gin.Default()
	books := r.Group("/books")
	{
		books.POST("", controller.BorrowBook)
		books.GET("/:bookID", controller.GetBookName)
	}
	return
}