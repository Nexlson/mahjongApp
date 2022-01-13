// package calculator

// type tileGroup struct {
// 	tiles  []int
// 	status int // group status 0: Open, 1: Close, 2: Open Kong, 3: Close Kong
// }

// type result struct {
// 	name string
// 	score int
// }

// // group tiles
// func groupTiles(tiles []int, status []int) []tileGroup{
// 	// sort inputs
// 	sort.Sort(sort.IntSlice(tiles))

// 	groups := [][]int{tiles[0:3], tiles[3:6], tiles[6:9], tiles[9:12], tiles[12:14]}

// 	hands := []tileGroup
// 	for index, g := range groups {
// 		hand := tileGroup
// 		hand.tiles = g
// 		hand.status = status[index]
// 		hands = hands.append(hand)
// 	}
// 	return hands
// }

// // check pong
// func checkPong([]int) []int{

// }

// // check chi
// func checkChi([]int) []int{

// }

// // compare hand tiles with rules
// func checkScore(groups [][]int) result {

// }

// // calculate tiles 
// func Calculate([]tileGroup) result{
// 	var output result
// 	var totalScore int
	
// 	if len(calculate) != 4 {
// 		output.valid = false
// 		return output
// 	}
// 	// check win hands
// 	for _, group := range hand {
// 		win := checkGroup(group, len(group))
// 	}
// 	if !win {
// 		output.valid = false
// 		return output
// 	}

// 	// check winning score and name
// }
