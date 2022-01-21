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

func extractPattern2List (group []TileGroup) []int {
	patternList := []int{}
	for _, g := range group {
		patternList = append(patternList,[]int{g.Pattern}...)
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

// check if group contained Chi 
func checkChi(group []int, c int) bool{
	count := 0 
	for _, x := range group {
		for _, y := range group{
			if x != y {
				if y - x == 1 {
					count += 1
				}
			}
		}
	}
	if count == c {
		return true
	}else{
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
		if mode == "Pong"{
			if g.Pong {
				handTiles = append(handTiles, []int{g.Tiles[0]}...)
			}
		}else if mode == "Kong"{
			if g.Kong {
				handTiles = append(handTiles, []int{g.Tiles[0]}...)
			}
		}else if mode == "Pair"{
			if g.Pair {
				handTiles = append(handTiles, []int{g.Tiles[0]}...)
			}
		}else if mode == "close"{
			if g.Open {
				handTiles = append(handTiles, []int{g.Tiles[0]}...)
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

func extractPatterns(grouped []TileGroup) []int{
	var extracted []int
	for _, group := range grouped {
		extracted = append(extracted, group.Pattern)
	}
	return extracted
}

func extractChi(grouped []TileGroup) [][]int {
	var extracted [][]int
	for _, group := range grouped {
		if group.Chi {
			extracted = append(extracted, group.Tiles)
		}
	}
	return extracted
}

func compareChi(chis [][]int, gap int) bool {
	prev := []int{}
	for _, Chi := range chis{
		if len(prev) == 0 {
			prev = Chi
		}else{
			if (Chi[0] - prev[2] == gap) {
				return true
			}
		}
	}
	return false
}

func checkWon(hand Hands) Hands {
	var count int 
	for _, group := range hand.Grouped {
		if group.Pong || group.Chi || group.Pair{
			count += 1
		}
	}
	if count == 5 {
		hand.Won = true
		return hand
	}else {
		hand.Won = false
		return hand
	}
}

func countStatus(grouped []TileGroup, status string) int{
	count := 0 
	for _, group := range grouped {
		if status == "Chi" {
			if group.Chi{
				count += 1
			}
		}else if status == "Pong"{
			if group.Pong{
				count += 1
			}
		}else if status == "Pair"{
			if group.Pair{
				count += 1
			}
		}
	}
	return count
}

func extractTileStatus(hands []TileGroup, mode string) []bool {
	status := []bool {}
	for _, group := range hands {
		if mode == "Kong" {
			status = append(status, []bool{group.Kong}...)
		}else if mode == "Open" {
			status = append(status, []bool{group.Open}...)
		}else if mode == "Pong" {
			if group.Pong == true {
				status = append(status, []bool{group.Pong}...)
			}
		}else if mode == "Chi" {
			status = append(status, []bool{group.Chi}...)
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
	return getKeyFromValue(mapSlice, larger), larger // key value
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

func gapBetweenTile(group []int, gap int) bool{
	for _, x := range group {
		for _, y := range group {
			if !(x == y) {
				if y - x == gap {
					return true
				}
			}
		}
		break
	}
	return false
}

func special147(hands Hands) bool{
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped{
			for _, z := range hands.Grouped{
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Pattern != y.Pattern && x.Pattern != z.Pattern && y.Pattern != z.Pattern {
						if gapBetweenTile(x.Tiles, 3) && gapBetweenTile(y.Tiles, 3) && gapBetweenTile(z.Tiles, 3) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func countOccurance(slice []int, slice2 []int) int{
	count := 0
	for _, s := range slice {
		for _, s2 := range slice2 {
			if s == s2 {
				count += 1
			}
		}
	}
	return count
}

func mergeSlices(slice1 []int, slice2 []int, slice3 []int) []int {
	slice1 = append(slice1, slice2...)
	slice1 = append(slice1, slice3...)
	return slice1
}

func extractTileGroup (hands Hands, mode string) []TileGroup {
	var result []TileGroup
	for _, g := range hands.Grouped {
		if mode == "Pair" && g.Pair {
			result = append(result,g)
		}else if mode == "Chi" && g.Chi {
			result = append(result,g)
		}
	}
	return result
}

func check123789(chiList []TileGroup) []int{
	patternList := []int{}
	for _, g := range chiList {
		for _, y := range chiList{
			if g.Pattern == y.Pattern && !sameSlice(g.Tiles, y.Tiles) {
				merged := append(g.Tiles, y.Tiles...)
				if sameSlice(merged, []int{1,2,3,7,8,9}) || sameSlice(merged, []int{10,11,12,16,17,18}) || sameSlice(merged, []int{19,20,21,25,26,27}){
					patternList = append(patternList,g.Pattern)
				}
			}
		}
	}
	return patternList
}

func checkGap(slice1 []int, slice2 []int, slice3 []int) bool{
	a := deValueGroup(slice1)
	b := deValueGroup(slice2)
	c := deValueGroup(slice3)

	if a[len(a)-1] == b[0] && b[len(b)-1] == c[0]{
		return true
	}else if  a[len(a)-1] == (b[0] + 1) && b[len(b)-1] == (c[0] + 1) {
		return true
	}else{
		return false
	}
}

func checkGap4(slice1 []int, slice2 []int, slice3 []int, slice4 []int) bool{
	a := deValueGroup(slice1)
	b := deValueGroup(slice2)
	c := deValueGroup(slice3)
	d := deValueGroup(slice4)

	if a[len(a)-1] == b[0] && b[len(b)-1] == c[0] && c[len(c)-1] == d[0]{
		return true
	}else if  a[len(a)-1] == (b[0] + 1) && b[len(b)-1] == (c[0] + 1) && c[len(c)-1] == (d[0] + 1){
		return true
	}else{
		return false
	}
}

func CalculateScore(hands Hands) Output{
	finalResult := Output{}

	// add initial score
	finalResult.Score += hands.Score

	// check Won
	handsTile := checkWon(hands)
	allRules := generateFunctions()
	// loop all rules 
	for index, rule := range allRules{
		if !containedInt(finalResult.Exceptions, index){
			finalResult = rule(handsTile, finalResult)
		}
	}
	return finalResult
}