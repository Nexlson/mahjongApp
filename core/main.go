package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/Nexlson/mahjongApp/backend/calculator"
	// "fmt"
)

func main() {
	r := gin.Default()
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
	// get start time
	start := time.Now()

	// define and bind data
	var playerHand calculator.Hands
	c.Bind(&playerHand)

	// calculate
	result := calculator.CalculateScore(playerHand)
	// finished time
	elapsed := time.Since(start)

	// return result
	c.JSON(200, gin.H{
		"result":  result,
		"timeTaken": elapsed,
	})
}