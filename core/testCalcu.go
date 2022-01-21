package main

import (
	"fmt"
	"github.com/Nexlson/mahjongApp/backend/calculator"
	"time"
	"log"
)


func mainTTT() {
	// want := 195
	start := time.Now()
	groupedHand := []calculator.TileGroup{
		{[]int{28, 28, 28}, false, true, false, false, false, 5},
		{[]int{29, 29, 29}, false, true, false, false, false, 5},
		{[]int{30, 30, 30}, false, true, false, false, false, 5},
		{[]int{31, 31, 31}, false, true, false, false, false, 5},
		{[]int{1, 1}, false, false, false, false, true, 1},
	}
	hands := calculator.Hands{
		Grouped: groupedHand,
		Ungrouped: []int{28, 28, 28, 29, 29, 29, 30, 30, 30, 31, 31, 31, 1, 1},
		Won: true,
	}

	// test function 
	got := calculator.CalculateScore(hands) 

	fmt.Println(got.Score)
	fmt.Println(got.Names)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

// func testGin() {
	// type Users struct {
//     Id        int    `json:"id"`
//     Firstname string `json:"firstname"`
//     Lastname  string `json:"lastname"`
// }

	// // test post 
	// r.POST("/user", func(c *gin.Context) {
	// 	var user Users
	// 	c.Bind(&user)
	// 	if user.Firstname != "" && user.Lastname != "" {
	// 		c.JSON(201, gin.H{"success": user})
	// 	} else {
	// 		c.JSON(422, gin.H{"error": "Fields are empty"})
	// 	}
	// })
// }