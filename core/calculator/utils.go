package calculator

import (
	"sort"
	// "fmt"
)

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
	if count == len(target) {
		return true
	}else{
		return false
	}
}

func checkAnyTileisInSlice(slice []int, slice2 []int) bool {
	for _, s := range slice {
		for _, s2 := range slice2 {
			if s == s2 {
				return true
			}
		}
	}
	return false
}

// compare two slice and return if slice in hand is available in slice in valid 
func checkContain(valid []int, hand []int, rule int) bool {
	count := 0 
	for _, v := range valid {
		for _, h := range hand {
			if v == h {
				count += 1
			}
		}
	}
	if count == rule {
		return true
	}else {
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

func extractPattern2List (group []TileGroup) []int {
	patternList := []int{}
	for _, g := range group {
		patternList = append(patternList,[]int{g.pattern}...)
	}
	return patternList
}

func checkPattern(list []int) bool {
	removeDup := removeDuplicateInt(list)
	if len(removeDup) == 1 {
		return true
	}else {
		return false
	}
}

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

func checkPairs(hands []int) bool {
	return len(removeDuplicateInt(hands)) == 7
}

func removeIntFromSlice(slice []int, hand []int) []int {
	for _, h := range hand {
		slice = removeInt(slice, h)
	}
	return slice
}

func removeInt(slice []int, num int) []int {
	newSlice := []int{}
	for _, s := range slice {
		if s != num {
			newSlice = append(newSlice,[]int{s}...)
		}
	}
	return newSlice
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func extractData(hands []TileGroup, mode string) []int{
	handTiles := []int{}
	for _, g := range hands {
		if mode == "pong"{
			if g.pong {
				handTiles = append(handTiles, []int{g.tiles[0]}...)
			}
		}else if mode == "kong"{
			if g.kong {
				handTiles = append(handTiles, []int{g.tiles[0]}...)
			}
		}else if mode == "pair"{
			if g.pair {
				handTiles = append(handTiles, []int{g.tiles[0]}...)
			}
		}else if mode == "close"{
			if g.open {
				handTiles = append(handTiles, []int{g.tiles[0]}...)
			}
		}
	}
	sort.Ints(handTiles)
	return handTiles
}

func containedInt(slice []int, tile int) bool {
	for _, i := range slice {
		if tile == i {
			return true
		}
	}
	return false
}

func containedInt2(slice []int, tile int) (bool, int){
	for ind, i := range slice {
		if tile == i {
			return true, ind
		}
	}
	return false, 6
}

func GroupOfWinningTile(posWinningTile int, validPositions []int)(bool, int){
	for _, pos := range validPositions {
		if pos == posWinningTile {
			return true, pos -1
		}
	}
	return false, 0
}

func ExtractPatterns(grouped []TileGroup) []int{
	var extracted []int
	for _, group := range grouped {
		extracted = append(extracted, group.pattern)
	}
	return extracted
}

func ExtractChi(grouped []TileGroup) [][]int {
	var extracted [][]int
	for _, group := range grouped {
		if group.chi {
			extracted = append(extracted, group.tiles)
		}
	}
	return extracted
}

func CompareChi(chis [][]int, gap int) bool {
	prev := []int{}
	for _, chi := range chis{
		if len(prev) == 0 {
			prev = chi
		}else{
			if (chi[0] - prev[2] == gap) {
				return true
			}
		}
	}
	return false
}

func checkWon(hand Hands) Hands {
	var count int 
	for _, group := range hand.grouped {
		if group.pong || group.chi {
			count += 1
		}
	}
	if count == 4 {
		hand.won = true
		return hand
	}else {
		hand.won = false
		return hand
	}
}

// func CalculateScore(hands Hands) Output{
// 	finalResult := Output{}

// 	// add initial score
// 	finalResult.score += hands.score

// 	// check won
// 	handsTile = checkWon(hands)

// 	// loop all rules 
// 	for k, v := range RulesMaps{
// 		if !containedInt(finalResult.exceptions, k){
// 			finalResult := v.((func(Hands, Output) Output (hands, finalResult))
// 		}
// 	}
// 	return finalResult
// }

func Hello() string {
	return "Hello World"
}

func countStatus(grouped []TileGroup, status string) int{
	count := 0 
	for _, group := range grouped {
		if status == "chi" {
			if group.chi{
				count += 1
			}
		}else if status == "pong"{
			if group.pong{
				count += 1
			}
		}
	}
	return count
}

func extractTileStatus(hands []TileGroup, mode string) []bool {
	status := []bool {}
	for _, group := range hands {
		if mode == "kong" {
			status = append(status, []bool{group.kong}...)
		}else if mode == "open" {
			status = append(status, []bool{group.open}...)
		}else if mode == "pong" {
			if group.pong == true {
				status = append(status, []bool{group.pong}...)
			}
		}else if mode == "chi" {
			status = append(status, []bool{group.chi}...)
		}
	}
	return status
}

func getCount(slice []int) map[int]int {
	m := make(map[int]int)
	for _, s := range slice {
		m[s]++

	}
	return m
}

func getLargestCountTiles(slice []int) (int, int) {
	larger := 0
	mapSlice := getCount(slice)
	for _,v := range mapSlice {
		if v > larger {
			larger = v
		}
	}
	return getKeyFromValue(mapSlice, larger), larger
}

func getKeyFromValue(maps map[int]int, value int) int {
	for k, v := range maps {
		if v == value {
			return k
		}
	}
	return -1
}

func deValue(tile int) int {
	if tile >= 10 && tile <= 18{
		return tile - 9
	}else if tile >= 19 && tile <= 27{
		return tile - 18
	}else{
		return tile
	}
}

func deValueGroup(slice []int) []int{
	newSlice := []int{}
	for _, s := range slice {
		if s >= 10 && s <= 18{
			newSlice = append(newSlice, s-9)
		}else if s >= 19 && s <= 27{
			newSlice = append(newSlice, s-18)
		}else{
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func sameSlice(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
