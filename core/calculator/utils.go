package calculator

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
	output := []int{}
	for _, h := range hand {
		for , s := range slice {
			if h != s {
				output = append(output,s)
				break
			}
		}
	}
	return output
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
	handTiles := int[]{}
	for _, g := range hands {
		if mode == "pong"{
			if g.pong == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "kong"{
			if g.kong == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "pair"{
			if g.pair == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "close"{
			if g.open == 0 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}
	}
	sort.Ints(handTiles)
	return handTiles
}

func ContainedInt(slice []int, tile int) bool {
	for _, i := range slice {
		if tile == i {
			return true
		}
	}
	return false
}

func GroupOfWinningTile(posWinningTile int, validPositions []int)(bool, int){
	for groupPos, pos := range validPositions {
		if pos == posWinningTile {
			return true, pos
		}
	}
	return false, 0
}

func ExtractPatterns(grouped []TileGroup) []int{
	var extracted []int{}
	for _, group := range grouped {
		extracted = append(extracted, group.pattern)
	}
	return extracted
}

func ExtractChi(grouped []TileGroup) [][]int {
	var extracted [][]int{}
	for _, group := range grouped {
		if group.chi {
			extracted = append(extracted, group)
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
			if chi[0] - prev[2] = gap {
				return true
			}
		}
	}
	return false
}

func checkWon(hand Hands) Hands {
	var count int 
	for _, group := range hands.grouped {
		if group.pong == 1 or group.chi == 1{
			count += 1
		}
	}
	if count == 4 {
		hands.won = true
		return hands
	}else {
		hands.won = false
		return hands
	}
}

func CalculateScore(hands Hands) Output{
	finalResult := Output{}

	// add initial score
	finalResult.score += hands.score

	// check won
	hands := checkWon(hands)

	// loop all rules 
	for k, v := range RulesMaps{
		if !containedInt(finalResult.exceptions, k){
			finalResult := v.((func(Hands, Output) Output) (hands, finalResult))
		}
	}
	return finalResult
}