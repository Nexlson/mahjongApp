package calculator

import "fmt"

// 88 points
func 大四喜(hands Hands, r Output) Output {
	validTiles := []int{28, 29, 30, 31}
	
	// check rules
	handTiles := extractData(hands.Grouped, "Pong") // extract Pong labelled Tiles

	// compare both list 
	if checkContain(validTiles, handTiles, len(validTiles)) && hands.Won { // compare if all Pong labelled are same as rule
		r.Names = append(r.Names,"大四喜")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{37, 42, 53, 54, 66}...)
		return r
	}
	fmt.Println("testing")
	return r
}

func 大三元(hands Hands, r Output) Output {
	validTiles := []int{32,33,34}

	// check rules
	handTiles := extractData(hands.Grouped, "Pong")

	// compare both list 
	if checkContain(validTiles, handTiles, len(validTiles)) && hands.Won {
		r.Names = append(r.Names,"大三元")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{48, 52, 66}...)
		return r
	}

	return r
}

func 绿一色(hands Hands, r Output) Output {
	validTiles := []int {11,12,13,15,17,33}
	handTiles := removeDuplicateInt(hands.Ungrouped)

	// if everything in hand is contained in valid Tiles
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.Won {
		r.Names = append(r.Names,"绿一色")
		r.Score += 88
		if containedInt(handTiles, 33) {
			r.Exceptions = append(r.Exceptions,[]int{43}...)
		}else {
			r.Exceptions = append(r.Exceptions,[]int{43, 21}...)
		}
		return r
	}
	return r
}

func 九莲宝灯(hands Hands, r Output) Output {
	validPongs := []int{1, 9, 10, 18, 19, 27}
	handPongs := extractData(hands.Grouped, "Pong")

	samePattern := checkPattern(extractPattern2List(hands.Grouped))
	contained := checkContain(validPongs, handPongs, len(handPongs))

	// remove duplicate, Pong Tiles, 
	removedDup := removeDuplicateInt(hands.Ungrouped)
	removePongs := removeIntFromSlice(removedDup, handPongs)

	if samePattern && contained && len(removePongs) == 7 {
		// only have 1112345678999
		r.Names = append(r.Names,"九莲宝灯")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{21, 50, 55, 69, 66}...)
		return r
	}
	return r
}

func 四杠(hands Hands, r Output) Output {
	kongCount := extractData(hands.Grouped, "Kong")
	if len(kongCount) == 4 && hands.Won {
		r.Names = append(r.Names,"四杠")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{42, 72}...)
		return r
	}
	return r
}

func 连七对(hands Hands, r Output) Output {
	samePattern := checkPattern(extractPattern2List(hands.Grouped))
	if checkPairs(hands.Ungrouped) && samePattern {
		r.Names = append(r.Names,"连七对")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{21, 18, 50, 55, 69, 72}...)
		return r
	}
	return r
}

func 十三幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	if checkContain(validTiles, removeDuplicateInt(hands.Ungrouped), 13) {
		r.Names = append(r.Names,"十三幺")
		r.Score += 88
		r.Exceptions = append(r.Exceptions,[]int{17, 45, 50, 55, 72}...)
		return r
	}
	return r
}

// 64 points
func 清幺九(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27}
	handTiles := removeDuplicateInt(hands.Ungrouped)

	if checkContain(validTiles, handTiles, len(handTiles)) && hands.Won {
		r.Names = append(r.Names,"清幺九")
		r.Score += 64
		r.Exceptions = append(r.Exceptions,[]int{42, 49, 58, 66, 69}...)
		return r
	}
	return r
}

func 小四喜(hands Hands, r Output) Output {
	validTiles := []int {28, 29, 30, 31}
	pongTiles := extractData(hands.Grouped, "Pong") // extract Pong labelled Tiles
	pairTiles := extractData(hands.Grouped, "Pair")
	count := 0 
	for _, tile := range pongTiles {
		if containedInt(validTiles, tile){
			count += 1
		}
	}
	pairCheckList := checkContain(validTiles, pairTiles, len(pairTiles)) && len(pairTiles) == 1
	if count == 3 && pairCheckList && hands.Won{
		r.Names = append(r.Names,"小四喜")
		r.Score += 64
		r.Exceptions = append(r.Exceptions,[]int{37, 66}...)
		return r
	}
	return r
}

func 小三元(hands Hands, r Output) Output {
	validTiles := []int {32, 33, 34}
	handTiles := extractData(hands.Grouped, "Pong") // extract Pong labelled Tiles
	handTiles = append(handTiles, extractData(hands.Grouped, "Pair")...) // extract Pair

	if checkContained(validTiles, handTiles) && hands.Won {
		r.Names = append(r.Names,"小三元")
		r.Score += 64
		r.Exceptions = append(r.Exceptions,[]int{48, 52, 66}...)
		return r
	}
	return r
}

func 字一色(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.Grouped)
	Pattern := removeDuplicateInt(patternList)
	samePattern := checkPattern(extractPattern2List(hands.Grouped))

	if samePattern && len(Pattern) == 1 && Pattern[0] == 4 && hands.Won {
		r.Names = append(r.Names,"字一色")
		r.Score += 64
		r.Exceptions = append(r.Exceptions,[]int{42, 49, 66}...)
		return r
	}
	return r
}

func 四暗刻(hands Hands, r Output) Output {
	threeTileGroup := hands.Grouped[:4]
	closedTiles := extractTileStatus(threeTileGroup, "Open")
	pongTiles := extractTileStatus(threeTileGroup, "Pong")
	if len(closedTiles) == 4 {
		if !closedTiles[0] && !closedTiles[1] && !closedTiles[2] && !closedTiles[3] && len(pongTiles) == 4 && hands.Won {
			r.Names = append(r.Names,"四暗刻")
			r.Score += 64
			r.Exceptions = append(r.Exceptions,[]int{42, 50, 55}...)
			return r
		}
	}
	return r
}

func 一色双龙会(hands Hands, r Output) Output {
	pairList := extractTileGroup(hands, "Pair")
	chiList := extractTileGroup(hands, "Chi")
	check1 := false
	if len(pairList) != 0 {
		check1 =  sameSlice(deValueGroup(pairList[0].Tiles), []int{5,5}) 
	}
	count123 := 0 
	count789:= 0 
	for _, g := range chiList {
		if sameSlice(g.Tiles, []int{1,2,3}){
			count123 += 1
		}else if sameSlice(g.Tiles, []int{7,8,9}){
			count789 += 1
		}
	}

	if check1 && count123 == 2 && count789 == 2 && hands.Won{
		r.Names = append(r.Names,"一色双龙会")
		r.Score += 64
		r.Exceptions = append(r.Exceptions,[]int{18, 21, 56, 62, 65, 69}...)
		return r
	}

	return r
}

func 一色四同顺(hands Hands, r Output) Output {
	samePattern := [][]int{}
	for _, g := range hands.Grouped {
		for _, x := range hands.Grouped {
			if sameSlice(g.Tiles, x.Tiles){
				samePattern = append(samePattern,x.Tiles)
			}
		}
		break
	}
	if len(samePattern) == 4 && hands.Won{
		r.Names = append(r.Names,"一色四同顺")
		r.Score += 48
		r.Exceptions = append(r.Exceptions,[]int{22, 23, 57, 62}...)
		return r
	}
	return r
}

func 一色四节高(hands Hands, r Output) Output {
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped {
			for _, z := range hands.Grouped {
				for _, v := range hands.Grouped {
					logic1 := !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles) && !sameSlice(x.Tiles,v.Tiles) && !sameSlice(v.Tiles,y.Tiles) && !sameSlice(v.Tiles,z.Tiles)
					if logic1{
						if x.Pong && y.Pong && z.Pong && v.Pong{
							newSlice := removeDuplicateInt(append(mergeSlices(deValueGroup(x.Tiles), deValueGroup(y.Tiles), deValueGroup(z.Tiles)), deValueGroup(v.Tiles)...))
							key,_ := getLargestCountTiles(newSlice)
							if checkChi(newSlice, len(newSlice) - 1) && key < 28{
									if x.Pattern == y.Pattern && x.Pattern == z.Pattern && y.Pattern == z.Pattern && hands.Won{
										r.Names = append(r.Names,"一色四节高")
										r.Score += 48
										r.Exceptions = append(r.Exceptions,[]int{22, 23, 42}...)
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
	for _, x := range hands.Grouped[:4] {
		for _, y := range hands.Grouped[:4] {
			for _, z := range hands.Grouped[:4] {
				for _, a := range hands.Grouped[:4] {
					logic1 := !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles) && !sameSlice(x.Tiles,a.Tiles) && !sameSlice(y.Tiles,a.Tiles) && !sameSlice(a.Tiles,z.Tiles)
					if logic1 {
						logic2 := x.Pattern == y.Pattern && y.Pattern == z.Pattern && x.Pattern == z.Pattern && x.Pattern == a.Pattern && y.Pattern == a.Pattern && a.Pattern == z.Pattern
						newSlice := removeDuplicateInt(append(mergeSlices(deValueGroup(x.Tiles), deValueGroup(y.Tiles), deValueGroup(z.Tiles)), deValueGroup(a.Tiles)...))
						key,_ := getLargestCountTiles(newSlice)
						if checkGap4(x.Tiles, y.Tiles, z.Tiles, a.Tiles) && logic2 && key < 28 && hands.Won{
							r.Names = append(r.Names,"一色四步高")
							r.Score += 32
							r.Exceptions = append(r.Exceptions,[]int{29, 64, 65}...)
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
	kongCount := extractData(hands.Grouped, "Kong")
	if len(kongCount) == 3 && hands.Won{
		r.Names = append(r.Names,"三杠")
		r.Score += 32
		return r
	}
	return r
}

func 混幺九(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	// 4 Pong all related to 
	pongCount := countStatus(hands.Grouped, "Pong")
	handTiles := removeDuplicateInt(hands.Ungrouped)

	if pongCount == 4 && checkContain(validTiles, handTiles, len(handTiles)) && hands.Won{
		r.Names = append(r.Names,"混幺九")
		r.Score += 32
		r.Exceptions = append(r.Exceptions,[]int{42, 49, 66}...)
		return r
	}
	return r
}

func 七对(hands Hands, r Output) Output {
	if checkPairs(hands.Ungrouped) {
		r.Names = append(r.Names,"七对")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{50, 55, 72}...)
		return r
	}
	return r
}

func 七星不靠(hands Hands, r Output) Output {
	handTiles := hands.Ungrouped
	validTiles := []int{28,29,30,31,32,33,34}
	removedValid := removeIntFromSlice(handTiles,validTiles)
	count := countOccurance(handTiles, validTiles) //count 7 valid tile

	if count == 7 && gapBetweenTile(removedValid, 3){
		r.Names = append(r.Names,"七星不靠")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{33, 45, 50, 55}...)
	}
	return r
}

func 全双刻(hands Hands, r Output) Output {
	validTiles := []int{2,4,6,8,11,13,15,17,20,22,24,26} // TODO: add valid Tiles
	handTiles := removeDuplicateInt(hands.Ungrouped)
	// remove duplicate + compare to validTiles + len is 4 
	if checkContain(validTiles, handTiles, len(handTiles)) && len(handTiles) == 5 && hands.Won{
		r.Names = append(r.Names,"全双刻")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{42, 61, 69}...)
		return r
	}
	return r
}

func 清一色(hands Hands, r Output) Output {
	samePattern := checkPattern(extractPattern2List(hands.Grouped))
	if samePattern && hands.Won{
		r.Names = append(r.Names,"清一色")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{69}...)
		return r
	}
	return r
}

func 一色三同顺(hands Hands, r Output) Output {
	samePattern := [][]int{}
	for _, g := range hands.Grouped {
		for _, x := range hands.Grouped {
			if sameSlice(g.Tiles, x.Tiles){
				samePattern = append(samePattern,x.Tiles)
			}
		}
		break
	}
	if len(samePattern) == 3 && hands.Won{
		r.Names = append(r.Names,"一色三同顺")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{23, 62}...)
		return r
	}
	return r
}

func 一色三节高(hands Hands, r Output) Output {
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped {
			for _, z := range hands.Grouped {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Pong && y.Pong && z.Pong {
						newSlice := removeDuplicateInt(mergeSlices(deValueGroup(x.Tiles), deValueGroup(y.Tiles), deValueGroup(z.Tiles)))
						key,_ := getLargestCountTiles(newSlice)
						if checkChi(newSlice, len(newSlice) - 1) && key < 28{
								if x.Pattern == y.Pattern && x.Pattern == z.Pattern && y.Pattern == z.Pattern && hands.Won{
									r.Names = append(r.Names,"一色三节高")
									r.Score += 24
									r.Exceptions = append(r.Exceptions,[]int{22}...)
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
	handTiles := removeDuplicateInt(hands.Ungrouped)

	if checkContained(validTiles, handTiles) && hands.Won{
		r.Names = append(r.Names,"全大")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{35, 69}...)
		return r
	}
	return r
}

func 全中(hands Hands, r Output) Output {
	validTiles := []int{4, 5, 6, 13, 14, 15, 22, 23, 24}
	handTiles := removeDuplicateInt(hands.Ungrouped)

	if checkContained(validTiles, handTiles) && hands.Won{
		r.Names = append(r.Names,"全中")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{61, 69}...)
		return r
	}
	return r
}

func 全小(hands Hands, r Output) Output {
	validTiles := []int{1, 2, 3, 10, 11, 12, 19, 20, 21}
	handTiles := removeDuplicateInt(hands.Ungrouped)

	if checkContained(validTiles, handTiles) && hands.Won{
		r.Names = append(r.Names,"全小")
		r.Score += 24
		r.Exceptions = append(r.Exceptions,[]int{36, 69}...)
		return r
	}
	return r
}

func 清龙(hands Hands, r Output) Output {
	validTiles := [][]int {{1,2,3,4,5,6,7,8,9}, {10,11,12,13,14,15,16,17,18}, {19,20,21,22,23,24,25,26,27}}
	handTiles := removeDuplicateInt(hands.Ungrouped)
	for _, g := range validTiles {
		if checkContained(g, handTiles) && hands.Won{
			r.Names = append(r.Names,"清龙")
			r.Score += 16
			r.Exceptions = append(r.Exceptions,[]int{64, 65}...)
			return r
		}
	}
	return r
}

func 三色双龙会(hands Hands, r Output) Output {
	pairList := extractTileGroup(hands, "Pair")
	chiList := extractTileGroup(hands, "Chi")
	check1 := false
	check3 := false
	check123789List := check123789(chiList) 
	if len(pairList) != 0 {
		check1 =  sameSlice(deValueGroup(pairList[0].Tiles), []int{5,5}) 
		pairPattern := pairList[0].Pattern
		check3 = len(append(check123789List, pairPattern)) == 3
	}
	check2 := len(check123789List) == 2
	if check1 && check2 && check3 && hands.Won{
		r.Names = append(r.Names,"三色双龙会")
		r.Score += 16
		r.Exceptions = append(r.Exceptions,[]int{56, 65, 63, 69}...)
		return r
	}

	return r
}

func 一色三步高(hands Hands, r Output) Output {
	for _, x := range hands.Grouped[:4] {
		for _, y := range hands.Grouped[:4] {
			for _, z := range hands.Grouped[:4] {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					logic := x.Pattern == y.Pattern && y.Pattern == z.Pattern && x.Pattern == z.Pattern
					newSlice := removeDuplicateInt(mergeSlices(deValueGroup(x.Tiles), deValueGroup(y.Tiles), deValueGroup(z.Tiles)))
					key,_ := getLargestCountTiles(newSlice)
					if checkGap(x.Tiles, y.Tiles, z.Tiles) && logic && key < 28 && hands.Won{
						r.Names = append(r.Names,"一色三步高")
						r.Score += 16
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
	fourGroupTiles := hands.Grouped[:4]
	pairGroupTiles := hands.Grouped[4]
	var FourGcount int
	for _, Grouped := range fourGroupTiles {
		if checkContain(validTiles, Grouped.Tiles, 1) {
			FourGcount += 1
		}
	}
	if FourGcount == 4 && checkContain(validTiles, pairGroupTiles.Tiles, 2) && hands.Won{
		r.Names = append(r.Names,"全带五")
		r.Score += 16
		r.Exceptions = append(r.Exceptions,[]int{61, 69}...)
		return r
	}
	return r
}

func 三同刻(hands Hands, r Output) Output {
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped {
			for _, z := range hands.Grouped {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Pong && y.Pong && z.Pong {
						if sameSlice(deValueGroup(x.Tiles), deValueGroup(y.Tiles)) && 
							sameSlice(deValueGroup(x.Tiles), deValueGroup(z.Tiles)) && 
							sameSlice(deValueGroup(y.Tiles), deValueGroup(z.Tiles)){
								if x.Pattern != y.Pattern && x.Pattern != z.Pattern && y.Pattern != z.Pattern && hands.Won{
									r.Names = append(r.Names,"三同刻")
									r.Score += 16
									r.Exceptions = append(r.Exceptions,[]int{58}...)
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
	for _, group := range hands.Grouped {
		if !group.Open && group.Pong {
			count += 1
		}
	}
	if count == 3 && hands.Won{
		r.Names = append(r.Names,"三暗刻")
		r.Score += 16
		return r
	}
	return r
}

func 全不靠(hands Hands, r Output) Output {
	handTiles := removeDuplicateInt(hands.Ungrouped)
	validTiles := []int{28,29,30,31,32,33,34}
	count := countOccurance(handTiles, validTiles)
	if special147(hands) && count == 5{
		r.Names = append(r.Names,"全不靠")
		r.Score += 12
		r.Exceptions = append(r.Exceptions,[]int{45, 50, 55}...)
	}
	return r
}

func 组合龙(hands Hands, r Output) Output {
	cpong := countStatus(hands.Grouped, "Pong")
	cpair := countStatus(hands.Grouped, "Pair")
	cchi := countStatus(hands.Grouped, "Chi")
	logic := cpong != 0 || cchi != 0
	if special147(hands) && logic && cpair != 0 {
		r.Names = append(r.Names,"组合龙")
		r.Score += 12
		return r
	}
	return r
}

func 大于五(hands Hands, r Output) Output {
	validTiles := []int {6,7,8,9,15,16,17,18,24,25,26,27}
	handTiles := removeDuplicateInt(hands.Ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.Won {
		r.Names = append(r.Names,"大于五")
		r.Score += 12
		r.Exceptions = append(r.Exceptions,[]int{69}...)
	}
	return r
}

func 小于五(hands Hands, r Output) Output {
	validTiles := []int {1,2,3,4,10,11,12,13,19,20,21,22}
	handTiles := removeDuplicateInt(hands.Ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.Won {
		r.Names = append(r.Names,"小于五")
		r.Score += 12
		r.Exceptions = append(r.Exceptions,[]int{69}...)
	}
	return r
}

func 三风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.Grouped, "Pong")
	if len(handPong) == 3 && hands.Won {
		r.Names = append(r.Names,"三风刻")
		r.Score += 12
		r.Exceptions = append(r.Exceptions,[]int{66}...)
	}
	return r
}


func 花龙(hands Hands, r Output) Output {
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped{
			for _, z := range hands.Grouped{
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Chi && y.Chi && z.Chi {
						if x.Pattern != y.Pattern && x.Pattern != z.Pattern && y.Pattern != z.Pattern{
							final := append(deValueGroup(x.Tiles), deValueGroup(y.Tiles)...)
							final = append(final, deValueGroup(z.Tiles)...)
							if sameSlice(final, []int{1,2,3,4,5,6,7,8,9}) && hands.Won{
								r.Names = append(r.Names,"花龙")
								r.Score += 8
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
	handTiles := removeDuplicateInt(hands.Ungrouped)
	if checkContain(validTiles, handTiles, len(handTiles)) && hands.Won{
		r.Names = append(r.Names,"推不倒")
		r.Score += 8
		r.Exceptions = append(r.Exceptions,[]int{68}...)
	}
	return r
}

func 三色三同顺(hands Hands, r Output) Output {
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped {
			for _, z := range hands.Grouped {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Chi && y.Chi && z.Chi {
						if sameSlice(deValueGroup(x.Tiles), deValueGroup(y.Tiles)) && 
							sameSlice(deValueGroup(x.Tiles), deValueGroup(z.Tiles)) && 
							sameSlice(deValueGroup(y.Tiles), deValueGroup(z.Tiles)){
								if x.Pattern != y.Pattern && x.Pattern != z.Pattern && y.Pattern != z.Pattern && hands.Won{
									r.Names = append(r.Names,"三色三同顺")
									r.Score += 8
									r.Exceptions = append(r.Exceptions,[]int{63}...)
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
	for _, x := range hands.Grouped {
		for _, y := range hands.Grouped {
			for _, z := range hands.Grouped {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					if x.Pong && y.Pong && z.Pong {
						newSlice := removeDuplicateInt(mergeSlices(deValueGroup(x.Tiles), deValueGroup(y.Tiles), deValueGroup(z.Tiles)))
						if checkChi(newSlice, len(newSlice) - 1){
								if x.Pattern != y.Pattern && x.Pattern != z.Pattern && y.Pattern != z.Pattern && hands.Won{
									r.Names = append(r.Names,"三色三节高")
									r.Score += 8
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
	if hands.Won && len(r.Names) == 0{
		r.Names = append(r.Names,"无番和")
		r.Score += 8
		return r
	}
	return r
}

func 碰碰和(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if group.Pong {
			count += 1
		}
	}
	if count == 4 && hands.Won{
		r.Names = append(r.Names,"碰碰和")
		r.Score += 6
		return r
	}
	return r
}

func 混一色(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.Grouped)
	validSeq := []int{1,2,3}
	validUniq := []int{4,5}	
	remained := removeDuplicateInt(removeIntFromSlice(patternList, validUniq))
	if len(removeDuplicateInt(patternList)) <= 3 && checkAnyTileisInSlice(patternList, validUniq) && len(remained) == 1 && 
		checkContain(validSeq, remained, len(remained)) && hands.Won{
			r.Names = append(r.Names,"混一色")
			r.Score += 6
	} 
	return r
}

func 三色三步高(hands Hands, r Output) Output {
	for _, x := range hands.Grouped[:4] {
		for _, y := range hands.Grouped[:4] {
			for _, z := range hands.Grouped[:4] {
				if !sameSlice(x.Tiles,y.Tiles) && !sameSlice(y.Tiles,z.Tiles) && !sameSlice(x.Tiles,z.Tiles){
					logic := x.Pattern != y.Pattern && y.Pattern != z.Pattern && x.Pattern != z.Pattern
					if checkGap(x.Tiles, y.Tiles, z.Tiles) && logic && hands.Won{
						r.Names = append(r.Names,"三色三步高")
						r.Score += 6
						return r
					}
				}
			}
		}
	}
	return r
}

func 五门齐(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.Grouped)
	if len(removeDuplicateInt(patternList)) == 5 && hands.Won{
		r.Names = append(r.Names,"五门齐")
		r.Score += 6
	}
	return r
}

func 全求人(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if group.Open {
			count += 1
		}
	}
	if count == 5 && hands.Won{
		r.Names = append(r.Names,"全求人")
		r.Score += 6
		r.Exceptions = append(r.Exceptions,[]int{72}...)
		return r
	}
	return r
}

func 双暗杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if !group.Open && group.Kong {
			count += 1
		}
	}
	if count == 2 && hands.Won{
		r.Names = append(r.Names,"双暗杠")
		r.Score += 6
		return r
	}
	return r
}

func 双箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.Grouped, "Pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 2 && hands.Won{
		r.Names = append(r.Names,"双箭刻")
		r.Score += 6
		r.Exceptions = append(r.Exceptions,[]int{52, 66}...)
	}
	return r
}

func 全带幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	for _, g := range hands.Grouped {
		if checkAnyTileisInSlice(validTiles, g.Tiles) && hands.Won{
			r.Names = append(r.Names,"全带幺")
			r.Score += 4
			return r
		}
	}
	return r
}

func 不求人(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if !group.Open {
			count += 1
		}
	}
	if count == 5 && hands.Won{
		r.Names = append(r.Names,"不求人")
		r.Score += 4
		return r
	}
	return r
}

func 双明杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if group.Open && group.Kong {
			count += 1
		}
	}
	if count == 2 && hands.Won{
		r.Names = append(r.Names,"双明杠")
		r.Score += 4
		return r
	}
	return r
}

func 箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.Grouped, "Pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 1 && hands.Won{
		r.Names = append(r.Names,"箭刻")
		r.Score += 2
		r.Exceptions = append(r.Exceptions,[]int{66}...)
	}
	return r
}

func 圈风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.Grouped, "Pong")
	if containedInt(handPong, hands.Round) && hands.Won{
		r.Names = append(r.Names,"圈风刻")
		r.Score += 2
		r.Exceptions = append(r.Exceptions,[]int{66}...)
	}
	return r
}

func 门风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.Grouped, "Pong")
	if containedInt(handPong, hands.Turn) && hands.Won{
		r.Names = append(r.Names,"门风刻")
		r.Score += 2
		r.Exceptions = append(r.Exceptions,[]int{66}...)
	}
	return r
}

func 门前清(hands Hands, r Output) Output {
	openStatusList := extractTileStatus(hands.Grouped, "Open")
	if len(openStatusList) == 5{
		if !openStatusList[0] && !openStatusList[1] && !openStatusList[2] && !openStatusList[3] && openStatusList[4] && hands.Won{
			r.Names = append(r.Names,"门前清")
			r.Score += 2
		}
	}
	return r
}

func 平和(hands Hands, r Output) Output {
	chiStatusList := extractTileStatus(hands.Grouped[:4], "Chi")
	if chiStatusList[0] && chiStatusList[1] && chiStatusList[2] && chiStatusList[3] && hands.Won{
		r.Names = append(r.Names,"平和")
		r.Score += 2
		r.Exceptions = append(r.Exceptions,[]int{69}...)
		return r
	}
	return r
}

func 四归一(hands Hands, r Output) Output {
	kongList := extractTileStatus(hands.Grouped[:4], "Kong")
	_, largeTileCount := getLargestCountTiles(hands.Ungrouped)
	if !(kongList[0] && kongList[1] && kongList[2] && kongList[3]) && largeTileCount == 4 && hands.Won{
		r.Names = append(r.Names,"四归一")
		r.Score += 2
		return r
	}
	return r
}

func 双同刻(hands Hands, r Output) Output {
	pongList := extractData(hands.Grouped, "Pong")
	for _, pongPrev := range pongList {
		for _, pongCur := range pongList {
			if pongPrev != pongCur {
				if deValue(pongPrev) == deValue(pongCur) && hands.Won{
					r.Names = append(r.Names,"双同刻")
					r.Score += 2
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
	for _, group := range hands.Grouped {
		if !group.Open && group.Pong {
			count += 1
		}
	}
	if count == 2 && hands.Won{
		r.Names = append(r.Names,"双暗刻")
		r.Score += 2
		return r
	}
	return r
}

func 暗杠(hands Hands, r Output) Output {
	count := 0 
	for _, group := range hands.Grouped {
		if !group.Open && group.Kong {
			count += 1
		}
	}
	if count == 1 && hands.Won{
		r.Names = append(r.Names,"暗杠")
		r.Score += 2
		return r
	}
	return r
}

func 断幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	if !checkAnyTileisInSlice(validTiles, removeDuplicateInt(hands.Ungrouped)) && hands.Won{
		r.Names = append(r.Names,"断幺")
		r.Score += 2
		r.Exceptions = append(r.Exceptions,[]int{69}...)
		return r
	}
	return r
}

func 一般高(hands Hands, r Output) Output {
// 	// same Pattern same value
	count := 0
	for _, chiPrev := range hands.Grouped {
		for _, chiCurrent := range hands.Grouped{
			if sameSlice(chiPrev.Tiles, chiCurrent.Tiles) && (chiPrev.Pattern == chiCurrent.Pattern) {
				count += 1
				if count >= 2 && hands.Won{
					r.Names = append(r.Names,"一般高")
					r.Score += 1
					return r
				}
			}
		}
		break
	}
	return r
}

func 喜相逢(hands Hands, r Output) Output {
// 	// same consecutive but different Pattern
	for _, chiPrev := range hands.Grouped {
		for _, chiCurrent := range hands.Grouped{
			if !sameSlice(chiPrev.Tiles, chiCurrent.Tiles) {
				devaluePrev := deValueGroup(chiPrev.Tiles)
				devalueCurr := deValueGroup(chiCurrent.Tiles)
				if sameSlice(devaluePrev, devalueCurr) && (chiPrev.Pattern != chiCurrent.Pattern) && hands.Won{
					r.Names = append(r.Names,"喜相逢")
					r.Score += 1
					return r
				}
			}
		}
		break
	}
	return r
}

func 连六(hands Hands, r Output) Output {
	handChis := extractChi(hands.Grouped)
	if compareChi(handChis, 1) && hands.Won{
		r.Names = append(r.Names,"连六")
		r.Score += 1
		return r
	}
	return r
}

func 老少副(hands Hands, r Output) Output {
	// same Pattern of 1 2 3, 7 8 9
	validTiles := [][]int{{1,2,3},{7,8,9}, {10,11,12},{16,17,18}, {19,20,21},{25,26,27}}
	chisGroup := extractChi(hands.Grouped)
	for _, g := range chisGroup {
		for _, vt := range validTiles{
			if checkContain(g, vt, len(g)) && hands.Won{
					r.Names = append(r.Names,"老少副")
					r.Score += 1
					return r
			}
		}
	}
	return r
}

func 幺九刻(hands Hands, r Output) Output {
	// Pong of 19 and zhi
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	handPongTiles := extractData(hands.Grouped, "Pong")
	if checkContain(validTiles, handPongTiles, len(handPongTiles)) && hands.Won{
		r.Names = append(r.Names,"幺九刻")
		r.Score += 1
		return r
	}
	return r
}

func 明杠(hands Hands, r Output) Output {
	// Kong and Open is 1
	count := 0 
	for _, group := range hands.Grouped {
		if group.Open && group.Kong {
			count += 1
		}
	}
	if count == 1 && hands.Won{
		r.Names = append(r.Names,"明杠")
		r.Score += 1
		return r
	}
	return r
}

func 缺一门(hands Hands, r Output) Output {
	// doesnt have one of the Pattern
	handPattern := removeDuplicateInt(extractPatterns(hands.Grouped))
	if len(handPattern) == 4 && hands.Won {
		r.Names = append(r.Names,"缺一门")
		r.Score += 1
		return r
	}
	return r
}

func 无字(hands Hands, r Output) Output {
	// doesnt contain 字 Pattern
	handPattern := extractPatterns(hands.Grouped)
	if !containedInt(handPattern, 4) && !containedInt(handPattern, 5) && hands.Won{
		r.Names = append(r.Names,"无字")
		r.Score += 1
		return r
	}
	return r
}

func 边张(hands Hands, r Output) Output {
	// winning tile is third of Grouped 3 tile 
	var groupChi bool
	validPositions := []int{2, 5, 8, 11}
	validity, groupPosition := containedInt2(validPositions, hands.WiningTile)
	if groupPosition <=5 {
		groupChi = hands.Grouped[groupPosition].Chi
	}else{
		groupChi = false
	}
	if validity && groupChi && hands.Won{
		r.Names = append(r.Names,"边张")
		r.Score += 1
		return r
	}
	return r
}

func 坎张(hands Hands, r Output) Output {
	// winning tile is middle of Grouped 3 tile 
	var groupChi bool
	validPositions := []int{1, 4, 7, 10}
	validity, groupPosition := containedInt2(validPositions, hands.WiningTile)
	if groupPosition <=5 {
		groupChi = hands.Grouped[groupPosition].Chi
	}else{
		groupChi = false
	}
	// group must be Chi
	if validity && groupChi && hands.Won{
		r.Names = append(r.Names,"坎张")
		r.Score += 1
		return r
	}
	return r
}

func 单钓(hands Hands, r Output) Output {
	// winning tile is one of the Pair
	validPositions := []int{12, 13}
	if containedInt(validPositions, hands.WiningTile) && hands.Won{
		r.Names = append(r.Names,"单钓")
		r.Score += 1
		return r
	}
	return r
}
