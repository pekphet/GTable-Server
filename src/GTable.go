package main

import (
	"github.com/gin-gonic/gin"
	App "./app"
	Result "./app/result"
	Award "./app/award"
	Account "./app/account"
)

func main() {
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) { c.String(500, "ERROR!") })
	//router.GET("/yqz/debug/confset", func(c *gin.Context) {
	//	c.Request.ParseForm()
	//	editDebugFile(&c.Request.Form, c)
	//})
	router.GET("/v1/app/update", App.CheckVersionServer)
	router.GET("/v1/edit/update", App.ChangeUpdateConfServer)

	router.GET("/v1/account/check", Account.CheckNameServer)
	router.POST("/v1/account/new", Account.AddNameServer)
	router.POST("/v1/account/store", Account.StoreDataServer)

	router.GET("/v1/result/list", Result.RetResultListServer)
	router.POST("/v1/result/upload", Result.ReceiveResultListServer)

	router.GET("/v1/award/code", Award.ExchangeCodeServer)
	router.Run(":14000")
}
