package main

import (
	"fmt"
)

// check if group contained chi 
func checkChi(group []int) bool{
	prev := 0
	for _, tile := range group {
		if prev == 0 {
			prev = tile
		}else{
			if tile - prev != 1 {
				return false
			}else{
				prev = tile
			}
		}
	}
	return true
}

// check if group contained pong
func checkPong(group []int) bool{
	sum := 0
	for _, tile := range group {
		sum += tile
	}
	if sum / len(group) == group[0] {
		return true
	}else {
		return false
	}
}

// get min max of slices
func minMax(array []int) []int {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return []int {min, max}
}

// check if hand contained in target
func checkContained(target []int, hand []int) bool {
	count := 0 
	for _, i := range hand {
		for _, j := range target {
			if i == j {
				count++
				break
			}
		}
	}
	if count == len(hand) {
		return true
	}else{
		return false
	}
}

// check all tiles is same pattern
func checkSamePattern (tiles []int) (bool, int) {
	minMax := minMax(tiles)
	for k, v := range PatternLib {
		if checkContained(v, minMax){
			return true, k
		}
	}
	return false, 0
}

func checkWon(grouped []TileGroup) bool {
	for _, group := range grouped {
		if ! (checkChi(group.tiles) and checkPong(group.tiles)) {
			return false
		}
	}
	return true
}

func calculateScore(hand StructuredTiles) Output{
	result := Output {
		names: []int{}, 
		score: 0,
	}
	// check if won 
	if checkNormalWon(hand.grouped) {

	}else{
		// check special pattern

	}
}

func main() {
	fmt.Println(checkSamePattern([]int{10, 11}))
}