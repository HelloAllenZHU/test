package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) { 
		c.String( http.StatusOK, "Welcome!" )
	} )

	r.Run( ":8001" )
}