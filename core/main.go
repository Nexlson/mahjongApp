package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Nexlson/mahjongApp/backend/calculator"
	// "fmt"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	// managed group
	v1 := r.Group("api/v1")

	// landing page backend
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World from mahjong server backend")
	})

	// calculator
	v1.POST("/calculator", calculate)

	// score logger
	v1.GET("/log", func(c *gin.Context) {
		c.String(200, "Hello World from score logger")
	})

	// docs
	v1.GET("/docs", func(c *gin.Context) {
		c.String(200, "Hello World from mahjong docs")
	})

	r.Run(":3500") // listen and serve on locahost:3500
}

func calculate(c *gin.Context) {
	// define and bind data
	var playerHand calculator.Hands
	c.Bind(&playerHand)

	// calculate
	result := calculator.CalculateScore(playerHand)

	// return result
	c.JSON(200, gin.H{
		"result":  result,
	})
}