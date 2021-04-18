package controller

import (
	"FloatingBooks/db"
	"FloatingBooks/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
)

func BorrowBook (c *gin.Context) {
	data := c.PostForm("data")
	var borrowInfo model.BorrowBook
	if err := json.Unmarshal([]byte(data), &borrowInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "数据有误！",
		})
		return
	}

	// 验证手机号
	ok, err := regexp.Match("^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$", []byte(borrowInfo.Phone))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "手机号格式有误！",
		})
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 验证学号
	ok, err = regexp.Match("[UMDumd](20)((1[6-9])|(20))\\d{5}", []byte(borrowInfo.StudentID))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "手机号格式有误！",
		})
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if ok, msg := db.BorrowABook(&borrowInfo); ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": "出错！",
		})
		return
	}
}

func GetBookName (c *gin.Context) {
	bookID, err := strconv.ParseInt(c.Param("bookID"), 10, 16)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "数据有误！",
		})
		return
	}

	bookName, err := db.GetBookName(int16(bookID))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"msg": "没有这本书哦~",
		})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": gin.H{
			"book_name": bookName,
		},
	})
}

func ReturnBook (c *gin.Context) {
	data := c.PostForm("data")
	var returnInfo model.ReturnBook
	if err := json.Unmarshal([]byte(data), &returnInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "数据有误！",
		})
		return
	}

	if ok := db.ReturnBook(&returnInfo);  ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "还书成功！",
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": "出错！",
		})
		return
	}
}