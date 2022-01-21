package calculator

import "fmt"

// 88 points
func 大四喜(hands Hands, r Output) Output {
	validTiles := []int{28, 29, 30, 31}
	
	// check rules
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles

	// compare both list 
	if checkContain(validTiles, handTiles, len(validTiles)) && hands.won { // compare if all pong labelled are same as rule
		r.names = append(r.names,"大四喜")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{38, 48, 60, 61, 73}...)
		return r
	}
	return r
}

func 大三元(hands Hands, r Output) Output {
	validTiles := []int{32,33,34}

	// check rules
	handTiles := extractData(hands.grouped, "pong")

	// compare both list 
	if checkContain(validTiles, handTiles, len(validTiles)) && hands.won {
		r.names = append(r.names,"大三元")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{54, 59, 73}...)
		return r
	}

	return r
}

func 绿一色(hands Hands, r Output) Output {
	validTiles := []int {11,12,13,15,17,33}
	handTiles := removeDuplicateInt(hands.ungrouped)

	// if everything in hand is contained in valid tiles
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.won {
		r.names = append(r.names,"绿一色")
		r.score += 88
		if containedInt(handTiles, 33) {
			r.exceptions = append(r.exceptions,[]int{49}...)
		}else {
			r.exceptions = append(r.exceptions,[]int{49, 22}...)
		}
		return r
	}
	return r
}

func 九莲宝灯(hands Hands, r Output) Output {
	validPongs := []int{1, 9, 10, 18, 19, 27}
	handPongs := extractData(hands.grouped, "pong")

	samePattern := checkPattern(extractPattern2List(hands.grouped))
	contained := checkContain(validPongs, handPongs, len(handPongs))

	// remove duplicate, pong tiles, 
	removedDup := removeDuplicateInt(hands.ungrouped)
	removePongs := removeIntFromSlice(removedDup, handPongs)

	if samePattern && contained && len(removePongs) == 7 {
		// only have 1112345678999
		r.names = append(r.names,"九莲宝灯")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{22, 56, 62, 76, 73}...)
		return r
	}
	return r
}

func 四杠(hands Hands, r Output) Output {
	kongCount := extractData(hands.grouped, "kong")
	if len(kongCount) == 4 && hands.won {
		r.names = append(r.names,"四杠")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{48, 79}...)
		return r
	}
	return r
}

func 连七对(hands Hands, r Output) Output {
	samePattern := checkPattern(extractPattern2List(hands.grouped))
	if checkPairs(hands.ungrouped) && samePattern {
		r.names = append(r.names,"连七对")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{22, 19, 56, 62, 76, 79}...)
		return r
	}
	return r
}

func 十三幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	if checkContain(validTiles, removeDuplicateInt(hands.ungrouped), 13) {
		r.names = append(r.names,"十三幺")
		r.score += 88
		r.exceptions = append(r.exceptions,[]int{18, 51, 56, 62, 79}...)
		return r
	}
	return r
}

// 64 points
func 清幺九(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27}

	if checkContained(validTiles, removeDuplicateInt(hands.ungrouped)) && hands.won {
		r.names = append(r.names,"清幺九")
		r.score += 64
		r.exceptions = append(r.exceptions,[]int{48, 55, 65, 73, 76}...)
		return r
	}
	return r
}

func 小四喜(hands Hands, r Output) Output {
	validTiles := []int {28, 29, 30, 31}
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles
	handTiles = append(handTiles, extractData(hands.grouped, "pair")...) // extract pair

	if checkContained(validTiles, handTiles) && hands.won {
		r.names = append(r.names,"小四喜")
		r.score += 64
		r.exceptions = append(r.exceptions,[]int{38, 73}...)
		return r
	}
	return r
}

func 小三元(hands Hands, r Output) Output {
	validTiles := []int {32, 33, 34}
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles
	handTiles = append(handTiles, extractData(hands.grouped, "pair")...) // extract pair

	if checkContained(validTiles, handTiles) && hands.won {
		r.names = append(r.names,"小三元")
		r.score += 64
		r.exceptions = append(r.exceptions,[]int{54, 59, 73}...)
		return r
	}
	return r
}

func 字一色(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.grouped)
	pattern := removeDuplicateInt(patternList)
	samePattern := checkPattern(extractPattern2List(hands.grouped))

	if samePattern && len(pattern) == 1 && pattern[0] == 4 && hands.won {
		r.names = append(r.names,"字一色")
		r.score += 64
		r.exceptions = append(r.exceptions,[]int{48, 55, 73}...)
		return r
	}
	return r
}

func 四暗刻(hands Hands, r Output) Output {
	threeTileGroup := hands.grouped[:4]
	closedTiles := extractTileStatus(threeTileGroup, "open")
	pongTiles := extractTileStatus(threeTileGroup, "pong")

	if !closedTiles[0] && !closedTiles[1] && !closedTiles[2] && !closedTiles[3] && len(pongTiles) == 4 && hands.won {
		r.names = append(r.names,"四暗刻")
		r.score += 64
		r.exceptions = append(r.exceptions,[]int{48, 56, 62}...)
		return r
	}
	return r
}

// func 一色双龙会(hands Hands, r Output) Output {
// 	samePattern, pattern := checkSamePattern(hands.ungrouped)
// 	validTiles := PatternLib[pattern]
// 	validTiles = removeIntFromSlice(validTiles, []int{validTiles[3], validTiles[5]})
	
// 	if samePattern && hands.won {
// 		if checkContained(validTiles, hands.ungrouped){
// 			r.names = append(r.names,"一色双龙会")
// 			r.score += 64
// 			r.exceptions = append(r.exceptions,[]int{19, 22, 63, 69, 72, 76}...)
// 			return r
// 		}
// 	}
// 	return r
// }

func 一色四同顺(hands Hands, r Output) Output {
	samePattern := [][]int{}
	for _, g := range hands.grouped {
		for _, x := range hands.grouped {
			if sameSlice(g.tiles, x.tiles){
				samePattern = append(samePattern,x.tiles)
			}
		}
		break
	}
	if len(samePattern) == 4{
		r.names = append(r.names,"一色四同顺")
		r.score += 48
		r.exceptions = append(r.exceptions,[]int{23, 24, 64, 69}...)
		return r
	}
	return r
}

func 一色四节高(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped {
			for _, z := range hands.grouped {
				for _, v := range hands.grouped {
					logic1 := !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles) && !sameSlice(x.tiles,v.tiles) && !sameSlice(v.tiles,y.tiles) && !sameSlice(v.tiles,z.tiles)
					if logic1{
						if x.pong && y.pong && z.pong && v.pong{
							newSlice := removeDuplicateInt(append(mergeSlices(deValueGroup(x.tiles), deValueGroup(y.tiles), deValueGroup(z.tiles)), deValueGroup(v.tiles)...))
							if checkChi(newSlice, len(newSlice) - 1){
									if x.pattern == y.pattern && x.pattern == z.pattern && y.pattern == z.pattern && hands.won{
										r.names = append(r.names,"一色四节高")
										r.score += 48
										r.exceptions = append(r.exceptions,[]int{23, 24, 48}...)
										return r
									}
								}
								
						}
					}
				}
			}
		}
	}
	return r
}

func 一色四步高(hands Hands, r Output) Output {
	for _, x := range hands.grouped[:4] {
		for _, y := range hands.grouped[:4] {
			for _, z := range hands.grouped[:4] {
				for _, a := range hands.grouped[:4] {
					logic1 := !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles) && !sameSlice(x.tiles,a.tiles) && !sameSlice(y.tiles,a.tiles) && !sameSlice(a.tiles,z.tiles)
					if logic1 {
						logic2 := x.pattern == y.pattern && y.pattern == z.pattern && x.pattern == z.pattern && x.pattern == a.pattern && y.pattern == a.pattern && a.pattern == z.pattern
						if checkGap4(x.tiles, y.tiles, z.tiles, a.tiles) && logic2 && hands.won{
							r.names = append(r.names,"一色四步高")
							r.score += 32
							r.exceptions = append(r.exceptions,[]int{48, 55, 73}...)
							return r
						}
					}
				}
			}
		}
	}
	return r
}

func 三杠(hands Hands, r Output) Output {
	kongCount := extractData(hands.grouped, "kong")
	if len(kongCount) == 3 {
		r.names = append(r.names,"三杠")
		r.score += 32
		return r
	}
	return r
}

func 混幺九(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	// 4 pong all related to 
	pongCount := countStatus(hands.grouped, "pong")
	handTiles := removeDuplicateInt(hands.ungrouped)

	if pongCount == 4 && checkContain(validTiles, handTiles, len(handTiles)) {
		r.names = append(r.names,"混幺九")
		r.score += 32
		r.exceptions = append(r.exceptions,[]int{48, 55, 73}...)
		return r
	}
	return r
}

func 七对(hands Hands, r Output) Output {
	if checkPairs(hands.ungrouped) {
		r.names = append(r.names,"七对")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{56, 62, 79}...)
		return r
	}
	return r
}

func 七星不靠(hands Hands, r Output) Output {
	handTiles := hands.ungrouped
	validTiles := []int{28,29,30,31,32,33,34}
	removedValid := removeIntFromSlice(handTiles,validTiles)
	count := countOccurance(handTiles, validTiles) //count 7 valid tile

	if count == 7 && gapBetweenTile(removedValid, 3){
		r.names = append(r.names,"七星不靠")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{34, 51, 56, 62}...)
	}
	return r
}

func 全双刻(hands Hands, r Output) Output {
	validTiles := []int{2,4,6,8,11,13,15,17,20,22,24,26} // TODO: add valid tiles
	handTiles := removeDuplicateInt(hands.ungrouped)
	// remove duplicate + compare to validTiles + len is 4 
	if checkContain(validTiles, handTiles, len(handTiles)) && len(handTiles) == 5 {
		r.names = append(r.names,"全双刻")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{48, 68, 76}...)
		return r
	}
	return r
}

func 清一色(hands Hands, r Output) Output {
	samePattern := checkPattern(extractPattern2List(hands.grouped))
	if samePattern {
		r.names = append(r.names,"清一色")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{76}...)
		return r
	}
	return r
}

func 一色三同顺(hands Hands, r Output) Output {
	samePattern := [][]int{}
	for _, g := range hands.grouped {
		for _, x := range hands.grouped {
			if sameSlice(g.tiles, x.tiles){
				samePattern = append(samePattern,x.tiles)
			}
		}
		break
	}
	if len(samePattern) == 3{
		r.names = append(r.names,"一色三同顺")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{24, 69}...)
		return r
	}
	return r
}

func 一色三节高(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped {
			for _, z := range hands.grouped {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					if x.pong && y.pong && z.pong {
						newSlice := removeDuplicateInt(mergeSlices(deValueGroup(x.tiles), deValueGroup(y.tiles), deValueGroup(z.tiles)))
						if checkChi(newSlice, len(newSlice) - 1){
								if x.pattern == y.pattern && x.pattern == z.pattern && y.pattern == z.pattern && hands.won{
									r.names = append(r.names,"一色三节高")
									r.score += 24
									r.exceptions = append(r.exceptions,[]int{23}...)
									return r
								}
							}
							
					}
				}
			}
		}
	}
	return r
}

func 全大(hands Hands, r Output) Output {
	validTiles := []int{7,8,9,16,17,18,25,26,27}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全大")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{36, 76}...)
		return r
	}
	return r
}

func 全中(hands Hands, r Output) Output {
	validTiles := []int{4, 5, 6, 13, 14, 15, 22, 23, 24}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全中")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{68, 76}...)
		return r
	}
	return r
}

func 全小(hands Hands, r Output) Output {
	validTiles := []int{1, 2, 3, 10, 11, 12, 19, 20, 21}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全小")
		r.score += 24
		r.exceptions = append(r.exceptions,[]int{37, 76}...)
		return r
	}
	return r
}

func 清龙(hands Hands, r Output) Output {
	validTiles := [][]int {{1,2,3,4,5,6,7,8,9}, {10,11,12,13,14,15,16,17,18}, {19,20,21,22,23,24,25,26,27}}
	handTiles := removeDuplicateInt(hands.ungrouped)
	for _, g := range validTiles {
		if checkContained(g, handTiles){
			r.names = append(r.names,"清龙")
			r.score += 16
			r.exceptions = append(r.exceptions,[]int{71, 72}...)
			return r
		}
	}
	return r
}

func 三色双龙会(hands Hands, r Output) Output {
	pairList := extractTileGroup(hands, "pair")
	chiList := extractTileGroup(hands, "chi")
	check1 :=  sameSlice(deValueGroup(pairList[0].tiles), []int{5,5}) 
	pairPattern := pairList[0].pattern
	check123789List := check123789(chiList) 
	check2 := len(check123789List) == 2
	check3 := len(append(check123789List, pairPattern)) == 3

	if check1 && check2 && check3{
		r.names = append(r.names,"三色双龙会")
		r.score += 16
		r.exceptions = append(r.exceptions,[]int{63, 72, 70, 76}...)
		return r
	}

	return r
}

func 一色三步高(hands Hands, r Output) Output {
	for _, x := range hands.grouped[:4] {
		for _, y := range hands.grouped[:4] {
			for _, z := range hands.grouped[:4] {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					logic := x.pattern == y.pattern && y.pattern == z.pattern && x.pattern == z.pattern
					if checkGap(x.tiles, y.tiles, z.tiles) && logic && hands.won{
						r.names = append(r.names,"一色三步高")
						r.score += 16
						return r
					}
				}
			}
		}
	}
	return r
}

func 全带五(hands Hands, r Output) Output {
	validTiles := []int {5, 14, 23}
	fourGroupTiles := hands.grouped[:4]
	pairGroupTiles := hands.grouped[4]
	var FourGcount int
	for _, grouped := range fourGroupTiles {
		if checkContain(validTiles, grouped.tiles, 1) {
			FourGcount += 1
		}
	}
	if FourGcount == 4 && checkContain(validTiles, pairGroupTiles.tiles, 2){
		r.names = append(r.names,"全带五")
		r.score += 16
		r.exceptions = append(r.exceptions,[]int{68, 76}...)
		return r
	}
	return r
}

func 三同刻(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped {
			for _, z := range hands.grouped {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					if x.pong && y.pong && z.pong {
						if sameSlice(deValueGroup(x.tiles), deValueGroup(y.tiles)) && 
							sameSlice(deValueGroup(x.tiles), deValueGroup(z.tiles)) && 
							sameSlice(deValueGroup(y.tiles), deValueGroup(z.tiles)){
								if x.pattern != y.pattern && x.pattern != z.pattern && y.pattern != z.pattern && hands.won{
									r.names = append(r.names,"三同刻")
									r.score += 16
									r.exceptions = append(r.exceptions,[]int{65}...)
									return r
								}
							}
							
					}
				}
			}
		}
	}
	return r
}

func 三暗刻(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if !group.open && group.pong {
			count += 1
		}
	}
	if count == 3 && hands.won{
		r.names = append(r.names,"三暗刻")
		r.score += 16
		return r
	}
	return r
}

func 全不靠(hands Hands, r Output) Output {
	handTiles := removeDuplicateInt(hands.ungrouped)
	validTiles := []int{28,29,30,31,32,33,34}
	count := countOccurance(handTiles, validTiles)
	if special147(hands) && count == 5{
		r.names = append(r.names,"全不靠")
		r.score += 12
		r.exceptions = append(r.exceptions,[]int{51, 56, 62}...)
	}
	return r
}

func 组合龙(hands Hands, r Output) Output {
	cpong := countStatus(hands.grouped, "pong")
	cpair := countStatus(hands.grouped, "pair")
	cchi := countStatus(hands.grouped, "chi")
	logic := cpong != 0 || cchi != 0
	if special147(hands) && logic && cpair != 0{
		r.names = append(r.names,"组合龙")
		r.score += 12
		return r
	}
	return r
}

func 大于五(hands Hands, r Output) Output {
	validTiles := []int {6,7,8,9,15,16,17,18,24,25,26,27}
	handTiles := removeDuplicateInt(hands.ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.won {
		r.names = append(r.names,"大于五")
		r.score += 12
		r.exceptions = append(r.exceptions,[]int{76}...)
	}
	return r
}

func 小于五(hands Hands, r Output) Output {
	validTiles := []int {1,2,3,4,10,11,12,13,19,20,21,22}
	handTiles := removeDuplicateInt(hands.ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.won {
		r.names = append(r.names,"小于五")
		r.score += 12
		r.exceptions = append(r.exceptions,[]int{76}...)
	}
	return r
}

func 三风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	if len(handPong) == 3 && hands.won {
		r.names = append(r.names,"三风刻")
		r.score += 12
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}


func 花龙(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped{
			for _, z := range hands.grouped{
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					if x.chi && y.chi && z.chi {
						if x.pattern != y.pattern && x.pattern != z.pattern && y.pattern != z.pattern && hands.won{
							final := append(deValueGroup(x.tiles), deValueGroup(y.tiles)...)
							final = append(final, deValueGroup(z.tiles)...)
							if sameSlice(final, []int{1,2,3,4,5,6,7,8,9}) && hands.won{
								r.names = append(r.names,"花龙")
								r.score += 8
								r.exceptions = append(r.exceptions,[]int{70}...)
								return r
							}
						}
					}
				}
			}
		}
	}
	return r
}

func 推不倒(hands Hands, r Output) Output {
	validTiles := []int {1,2,3,4,5,8,9,11,13,14,15,17,18,34}
	handTiles := removeDuplicateInt(hands.ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.won{
		r.names = append(r.names,"推不倒")
		r.score += 8
		r.exceptions = append(r.exceptions,[]int{75}...)
	}
	return r
}

func 三色三同顺(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped {
			for _, z := range hands.grouped {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					if x.chi && y.chi && z.chi {
						if sameSlice(deValueGroup(x.tiles), deValueGroup(y.tiles)) && 
							sameSlice(deValueGroup(x.tiles), deValueGroup(z.tiles)) && 
							sameSlice(deValueGroup(y.tiles), deValueGroup(z.tiles)){
								if x.pattern != y.pattern && x.pattern != z.pattern && y.pattern != z.pattern && hands.won{
									r.names = append(r.names,"三色三同顺")
									r.score += 8
									r.exceptions = append(r.exceptions,[]int{70}...)
									return r
								}
							}
							
					}
				}
			}
		}
	}
	return r
}

func 三色三节高(hands Hands, r Output) Output {
	for _, x := range hands.grouped {
		for _, y := range hands.grouped {
			for _, z := range hands.grouped {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					if x.pong && y.pong && z.pong {
						newSlice := removeDuplicateInt(mergeSlices(deValueGroup(x.tiles), deValueGroup(y.tiles), deValueGroup(z.tiles)))
						if checkChi(newSlice, len(newSlice) - 1){
								if x.pattern != y.pattern && x.pattern != z.pattern && y.pattern != z.pattern && hands.won{
									r.names = append(r.names,"三色三节高")
									r.score += 8
									return r
								}
							}
							
					}
				}
			}
		}
	}
	return r
}

func 无番和(hands Hands, r Output) Output {
	if hands.won && len(r.names) == 0{
		r.names = append(r.names,"无番和")
		r.score += 8
		return r
	}
	return r
}

func 碰碰和(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if group.pong {
			count += 1
		}
	}
	if count == 4 && hands.won{
		r.names = append(r.names,"碰碰和")
		r.score += 6
		r.exceptions = append(r.exceptions,[]int{79}...)
		return r
	}
	return r
}

func 混一色(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.grouped)
	validSeq := []int{1,2,3}
	validUniq := []int{4,5}	
	remained := removeDuplicateInt(removeIntFromSlice(patternList, validUniq))
	if len(removeDuplicateInt(patternList)) <= 3 && checkAnyTileisInSlice(patternList, validUniq) && len(remained) == 1 && 
		checkContain(validSeq, remained, len(remained)) && hands.won{
			r.names = append(r.names,"混一色")
			r.score += 6
	} 
	return r
}

func 三色三步高(hands Hands, r Output) Output {
	for _, x := range hands.grouped[:4] {
		for _, y := range hands.grouped[:4] {
			for _, z := range hands.grouped[:4] {
				if !sameSlice(x.tiles,y.tiles) && !sameSlice(y.tiles,z.tiles) && !sameSlice(x.tiles,z.tiles){
					logic := x.pattern != y.pattern && y.pattern != z.pattern && x.pattern != z.pattern
					if checkGap(x.tiles, y.tiles, z.tiles) && logic && hands.won{
						r.names = append(r.names,"三色三步高")
						r.score += 6
						return r
					}
				}
			}
		}
	}
	return r
}

func 五门齐(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.grouped)
	if len(patternList) == 5 && hands.won{
		r.names = append(r.names,"五门齐")
		r.score += 6
	}
	return r
}

func 全求人(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if group.open {
			count += 1
		}
	}
	if count == 5 && hands.won{
		r.names = append(r.names,"全求人")
		r.score += 6
		r.exceptions = append(r.exceptions,[]int{79}...)
		return r
	}
	return r
}

func 双暗杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if !group.open && group.kong {
			count += 1
		}
	}
	if count == 2 && hands.won{
		r.names = append(r.names,"双暗杠")
		r.score += 6
		return r
	}
	return r
}

func 双箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 2 && hands.won{
		r.names = append(r.names,"双箭刻")
		r.score += 6
		r.exceptions = append(r.exceptions,[]int{59}...)
	}
	return r
}

func 全带幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	for _, g := range hands.grouped {
		if checkAnyTileisInSlice(validTiles, g.tiles) && hands.won{
			r.names = append(r.names,"全带幺")
			r.score += 4
			return r
		}
	}
	return r
}

func 不求人(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if !group.open {
			count += 1
		}
	}
	if count == 5 && hands.won{
		r.names = append(r.names,"不求人")
		r.score += 4
		r.exceptions = append(r.exceptions,[]int{80}...)
		return r
	}
	return r
}

func 双明杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if group.open && group.kong {
			count += 1
		}
	}
	if count == 2 && hands.won{
		r.names = append(r.names,"双明杠")
		r.score += 4
		return r
	}
	return r
}

func 箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 1 && hands.won{
		r.names = append(r.names,"箭刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 圈风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	if containedInt(handPong, hands.round) && hands.won{
		r.names = append(r.names,"圈风刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 门风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	if containedInt(handPong, hands.turn) && hands.won{
		r.names = append(r.names,"门风刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 门前清(hands Hands, r Output) Output {
	openStatusList := extractTileStatus(hands.grouped, "open")
	if !openStatusList[0] && !openStatusList[1] && !openStatusList[2] && !openStatusList[3] && openStatusList[4] && hands.won{
		r.names = append(r.names,"门前清")
		r.score += 2
	}
	return r
}

func 平和(hands Hands, r Output) Output {
	chiStatusList := extractTileStatus(hands.grouped[:4], "chi")
	if chiStatusList[0] && chiStatusList[1] && chiStatusList[2] && chiStatusList[3] && hands.won{
		r.names = append(r.names,"平和")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{76}...)
		return r
	}
	return r
}

func 四归一(hands Hands, r Output) Output {
	kongList := extractTileStatus(hands.grouped[:4], "kong")
	_, largeTileCount := getLargestCountTiles(hands.ungrouped)
	if !(kongList[0] && kongList[1] && kongList[2] && kongList[3]) && largeTileCount == 4 && hands.won{
		r.names = append(r.names,"四归一")
		r.score += 2
		return r
	}
	return r
}

func 双同刻(hands Hands, r Output) Output {
	pongList := extractData(hands.grouped, "pong")
	for _, pongPrev := range pongList {
		for _, pongCur := range pongList {
			if pongPrev != pongCur {
				if deValue(pongPrev) == deValue(pongCur) && hands.won{
					r.names = append(r.names,"双同刻")
					r.score += 2
					return r
				}
			}
		}
		break
	}
	return r
}

func 双暗刻(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if !group.open && group.pong {
			count += 1
		}
	}
	if count == 2 && hands.won{
		r.names = append(r.names,"双暗刻")
		r.score += 2
		return r
	}
	return r
}

func 暗杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.grouped {
		if !group.open && group.kong {
			count += 1
		}
	}
	if count == 1 && hands.won{
		r.names = append(r.names,"暗杠")
		r.score += 2
		return r
	}
	return r
}

func 断幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	if !checkAnyTileisInSlice(validTiles, removeDuplicateInt(hands.ungrouped)) && hands.won{
		r.names = append(r.names,"断幺")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{76}...)
		return r
	}
	return r
}

func 一般高(hands Hands, r Output) Output {
// 	// same pattern same value
	count := 0
	for _, chiPrev := range hands.grouped {
		for _, chiCurrent := range hands.grouped{
			if sameSlice(chiPrev.tiles, chiCurrent.tiles) && (chiPrev.pattern == chiCurrent.pattern) {
				count += 1
				if count >= 2 && hands.won{
					r.names = append(r.names,"一般高")
					r.score += 1
					return r
				}
			}
		}
		break
	}
	return r
}

func 喜相逢(hands Hands, r Output) Output {
// 	// same consecutive but different pattern
	for _, chiPrev := range hands.grouped {
		for _, chiCurrent := range hands.grouped{
			if !sameSlice(chiPrev.tiles, chiCurrent.tiles) {
				devaluePrev := deValueGroup(chiPrev.tiles)
				devalueCurr := deValueGroup(chiCurrent.tiles)
				if sameSlice(devaluePrev, devalueCurr) && (chiPrev.pattern != chiCurrent.pattern) && hands.won{
					r.names = append(r.names,"喜相逢")
					r.score += 1
					return r
				}
			}
		}
		break
	}
	fmt.Println("Hello")
	return r
}

func 连六(hands Hands, r Output) Output {
	handChis := extractChi(hands.grouped)
	if CompareChi(handChis, 1) && hands.won{
		r.names = append(r.names,"连六")
		r.score += 1
		return r
	}
	return r
}

func 老少副(hands Hands, r Output) Output {
	// same pattern of 1 2 3, 7 8 9
	validTiles := [][]int{{1,2,3},{7,8,9}, {10,11,12},{16,17,18}, {19,20,21},{25,26,27}}
	chisGroup := extractChi(hands.grouped)
	for _, g := range chisGroup {
		for _, vt := range validTiles{
			if checkContain(g, vt, len(g)) && hands.won{
					r.names = append(r.names,"老少副")
					r.score += 1
					return r
			}
		}
	}
	return r
}

func 幺九刻(hands Hands, r Output) Output {
	// pong of 19 and zhi
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	handPongTiles := extractData(hands.grouped, "pong")
	if checkContain(validTiles, handPongTiles, len(handPongTiles)) && hands.won{
		r.names = append(r.names,"幺九刻")
		r.score += 1
		return r
	}
	return r
}

func 明杠(hands Hands, r Output) Output {
	// kong and open is 1
	count := 0 
	for _, group := range hands.grouped {
		if group.open && group.kong {
			count += 1
		}
	}
	if count == 1 && hands.won{
		r.names = append(r.names,"明杠")
		r.score += 1
		return r
	}
	return r
}

func 缺一门(hands Hands, r Output) Output {
	// doesnt have one of the pattern
	validPattern := []int{1,2,3,4}
	handPattern := ExtractPatterns(hands.grouped)
	if !checkContained(validPattern, handPattern) && hands.won{
		r.names = append(r.names,"缺一门")
		r.score += 1
		return r
	}
	return r
}

func 无字(hands Hands, r Output) Output {
	// doesnt contain 字 pattern
	handPattern := ExtractPatterns(hands.grouped)
	if !containedInt(handPattern, 4) && hands.won{
		r.names = append(r.names,"无字")
		r.score += 1
		return r
	}
	return r
}

func 边张(hands Hands, r Output) Output {
	// winning tile is third of grouped 3 tile 
	var groupChi bool
	validPositions := []int{2, 5, 8, 11}
	validity, groupPosition := containedInt2(validPositions, hands.winnngTile)
	if groupPosition <=5 {
		groupChi = hands.grouped[groupPosition].chi
	}else{
		groupChi = false
	}
	if validity && groupChi && hands.won{
		r.names = append(r.names,"边张")
		r.score += 1
		return r
	}
	return r
}

func 坎张(hands Hands, r Output) Output {
	// winning tile is middle of grouped 3 tile 
	var groupChi bool
	validPositions := []int{1, 4, 7, 10}
	validity, groupPosition := containedInt2(validPositions, hands.winnngTile)
	if groupPosition <=5 {
		groupChi = hands.grouped[groupPosition].chi
	}else{
		groupChi = false
	}
	// group must be chi
	if validity && groupChi && hands.won{
		r.names = append(r.names,"坎张")
		r.score += 1
		return r
	}
	return r
}

func 单钓(hands Hands, r Output) Output {
	// winning tile is one of the pair
	validPositions := []int{12, 13}
	if containedInt(validPositions, hands.winnngTile) && hands.won{
		r.names = append(r.names,"单钓")
		r.score += 1
		return r
	}
	return r
}
