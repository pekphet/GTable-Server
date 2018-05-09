package main

import (
	"github.com/gin-gonic/gin"
	App "./app"
)

func main() {
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) { c.String(500, "ERROR!") })
	//router.GET("/yqz/debug/confset", func(c *gin.Context) {
	//	c.Request.ParseForm()
	//	editDebugFile(&c.Request.Form, c)
	//})
	router.GET("/v1/app/update", App.CheckVersion)
	router.GET("/v1/edit/update", App.ChangeUpdateConf)
	router.Run(":14000")
}
