package main

import (
	"net/http"

	"src/github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title" : "hogehoge",
			"info" : "",
		})
	})

	router.GET("/test",func(c *gin.Context){
		var keyword string

		keyword = c.Query("keyword")
		ReadDB(keyword)

		c.HTML(http.StatusOK, "index.tpml",gin.H{
			"title" : "test",
			"info" : "",
		})
	})

	router.Run(":8080")
}