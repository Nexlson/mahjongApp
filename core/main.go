package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// landing page backend
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World from mahjong server backend")
	})
	
	// calculator
	// r.POST("calculate" calculator.calculate)

	// docs
	// get data from backend
	r.GET("/docs", getDocs)

	// score logger

	r.Run(":3500") // listen and serve on locahost:8080
}

func getDocs(c *gin.Context) {

}

func postCalculate(c *gin.Context) {

}