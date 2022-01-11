package main

import (
	"fmt"

	"github.com/Nexlson/mahjongApp/backend/pkg/calculator"
)

func main() {
	msg := calculator.Hello()

	fmt.Println(msg)
}
