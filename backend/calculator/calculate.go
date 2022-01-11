package calculator

type result struct {
	name string
	score int
	valid bool
}

func groupTiles(tiles []int) [][]int{

}

func checkGroup(groups []int, count int) bool{
	var prev int
	var counts int 
	for _, id := range groups {
		if prev == 0 { // set previous tile
			prev = id
		}else if prev == id { // check pong
			counts += 1
		}else if  obj - prev == 1 { // check chi
			counts += 1
		}
	}
	if counts == count {
		return true
	}else{
		return false
	}
}	

func Calculate(hand [][]int) result{
	var output result
	var totalScore int
	
	if len(calculate) != 4 {
		output.valid = false
		return output
	}
	// check win hands
	for _, group := range hand {
		win := checkGroup(group, len(group))
	}
	if !win {
		output.valid = false
		return output
	}

	// check winning score and name
}

// input => [{1,1,1}, {2,2,2}, {3,3,3}, {4,4,4}, {5,5}]}