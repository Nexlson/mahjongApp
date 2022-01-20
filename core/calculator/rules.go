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

func 一色双龙会(hands Hands, r Output) Output {
	samePattern, pattern := checkSamePattern(hands.ungrouped)
	validTiles := PatternLib[pattern]
	validTiles = removeIntFromSlice(validTiles, []int{validTiles[3], validTiles[5]})
	
	if samePattern && hands.won {
		if checkContained(validTiles, hands.ungrouped){
			r.names = append(r.names,"一色双龙会")
			r.score += 64
			r.exceptions = append(r.exceptions,[]int{19, 22, 63, 69, 72, 76}...)
			return r
		}
	}
	return r
}

// same pattern for 4 group 
func 一色四同顺(hands Hands, r Output) Output {
// 	name := "一色四同顺"
// 	score := 48
// 	exceptions := []int{23, 24, 64, 69}

// 	samePattern, pattern := checkSamePattern(hands.ungrouped)

// 	// 4 chi + chi is the same 3 combi
// 	chiCount := countStatus(hands.grouped, "chi")

// 	uniqueTiles := removeDuplicateInt(hands.ungrouped)

// 	if samePattern && chiCount == 4 && len(uniqueTiles) == 4 {
// 		r.names = append(r.names,"一色四同顺")
// 		r.score += 48
// 		exceptions = append(exceptions,[]int{19, 22, 63, 69, 72, 76}...)
// 		return r, exceptions
// 	}
	return r
}

func 一色四节高(hands Hands, r Output) Output {
// 	// 4 pong, same pattern 
// 	samePattern, pattern := checkSamePattern(hands.ungrouped)
// 	pongCount := countStatus(hands.grouped, "pong")
// 	uniqueTiles := removeDuplicateInt(hands.ungrouped)

// 	if samePattern && pongCount == 4 && len(uniqueTiles) == 5 {
// 		r.names = append(r.names,"一色四节高")
// 		r.score += 48
// 		exceptions = append(exceptions,[]int{23, 24, 48}...)
// 		return r, exceptions
// 	}
	return r
}

func 一色四步高(hands Hands, r Output) Output {
// 	name := "一色四步高"
// 	score := 32
// 	exceptions := []int{48, 55, 73}
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

	if pongCount == 4 && checkContained(validTiles, handTiles) {
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
// 	name := "七星不靠"
// 	score := 24
// 	exceptions := []int{34, 51, 56, 62}
// 	validWord = []int{28, 29, 30, 31, 32, 33, 34}
// 	validContinous = [][]int{{}}

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
// 	name := "一色三同顺"
// 	score := 24
// 	exceptions := []int{24, 69}
	return r
}

func 一色三节高(hands Hands, r Output) Output {
// 	name := "一色三节高"
// 	score := 24
// 	exceptions := []int{23}
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
// 	name := "三色双龙会"
// 	score := 16
// 	exceptions := []int{63, 72, 70, 76}
	return r
}

func 一色三步高(hands Hands, r Output) Output {
// 	name := "一色三步高"
// 	score := 16
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

func gapBetweenTile(group []int, gap int) bool{
	prev := 0 
	for _, tile := range group {
		if prev == 0 {
			tile = prev
		}else {
			if tile - prev != gap {
				return false
			}
		}
	}
	return true
}

func 三同刻(hands Hands, r Output) Output {
// 	name := "三同刻"
// 	score := 16
// 	exceptions := []int{65}

// 3 pong + between each pong is 9 
	return r
}


func 三暗刻(hands Hands, r Output) Output {
// 	name := "三暗刻"
// 	score := 16

// 	// 3 closed pong

	return r
}


func 全不靠(hands Hands, r Output) Output {
// 	name := "全不靠"
// 	score := 12
// 	exceptions := []int{51, 56, 62}
	return r
}

func 组合龙(hands Hands, r Output) Output {
// 	name := "组合龙"
// 	score := 12
	return r
}

func 大于五(hands Hands, r Output) Output {
// 	name := "大于五"
// 	score := 12
// 	exceptions := []int{76}
	return r
}


func 小于五(hands Hands, r Output) Output {
// 	name := "小于五"
// 	score := 12
// 	exceptions := []int{76}
	return r
}

func 三风刻(hands Hands, r Output) Output {
// 	name := "三风刻"
// 	score := 12
// 	exceptions := []int{73}
	return r
}


func 花龙(hands Hands, r Output) Output {
// 	name := "花龙"
// 	score := 8
	return r
}

func 推不倒(hands Hands, r Output) Output {
// 	name := "推不倒"
// 	score := 8
// 	exceptions := []int{75}
	return r
}

func 三色三同顺(hands Hands, r Output) Output {
// 	name := "三色三同顺"
// 	score := 8
// 	exceptions := []int{70}
	return r
}

func 三色三节高(hands Hands, r Output) Output {
// 	name := "三色三节高"
// 	score := 8
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
	if count == 4 {
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
	if len(removeDuplicateInt(patternList)) <= 3 && checkAnyTileisInSlice(patternList, validUniq) && len(remained) == 1 && checkContain(validSeq, remained, len(remained)) {
		r.names = append(r.names,"混一色")
		r.score += 6
	} 
	return r
}

func 三色三步高(hands Hands, r Output) Output {
// 	name := "三色三步高"
// 	score := 6
	return r
}

func 五门齐(hands Hands, r Output) Output {
	patternList := extractPattern2List(hands.grouped)
	if len(patternList) == 5 {
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
	if count == 5 {
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
	if count == 2 {
		r.names = append(r.names,"双暗杠")
		r.score += 6
		return r
	}
	return r
}

func 双箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 2{
		r.names = append(r.names,"双箭刻")
		r.score += 6
		r.exceptions = append(r.exceptions,[]int{59}...)
	}
	return r
}

func 全带幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	for _, g := range hands.grouped {
		if checkAnyTileisInSlice(validTiles, g.tiles) {
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
	if count == 5 {
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
	if count == 2 {
		r.names = append(r.names,"双明杠")
		r.score += 4
		return r
	}
	return r
}

func 箭刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	validTiles := []int{32,33,34}
	if checkAnyTileisInSlice(handPong, validTiles) && len(handPong) == 1{
		r.names = append(r.names,"箭刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 圈风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	if containedInt(handPong, hands.round) {
		r.names = append(r.names,"圈风刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 门风刻(hands Hands, r Output) Output {
	handPong := extractData(hands.grouped, "pong")
	if containedInt(handPong, hands.turn) {
		r.names = append(r.names,"门风刻")
		r.score += 2
		r.exceptions = append(r.exceptions,[]int{73}...)
	}
	return r
}

func 门前清(hands Hands, r Output) Output {
	openStatusList := extractTileStatus(hands.grouped, "open")
	if !openStatusList[0] && !openStatusList[1] && !openStatusList[2] && !openStatusList[3] && openStatusList[4] {
		r.names = append(r.names,"门前清")
		r.score += 2
	}
	return r
}

func 平和(hands Hands, r Output) Output {
	chiStatusList := extractTileStatus(hands.grouped[:4], "chi")
	if chiStatusList[0] && chiStatusList[1] && chiStatusList[2] && chiStatusList[3] {
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
	if !(kongList[0] && kongList[1] && kongList[2] && kongList[3]) && largeTileCount == 4{
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
				if deValue(pongPrev) == deValue(pongCur) {
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
	if count == 2 {
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
	if count == 1 {
		r.names = append(r.names,"暗杠")
		r.score += 2
		return r
	}
	return r
}

func 断幺(hands Hands, r Output) Output {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}
	if !checkAnyTileisInSlice(validTiles, removeDuplicateInt(hands.ungrouped)) {
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
				if count >= 2 {
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
				if sameSlice(devaluePrev, devalueCurr) && (chiPrev.pattern != chiCurrent.pattern) {
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
	handChis := ExtractChi(hands.grouped)
	if CompareChi(handChis, 1) {
		r.names = append(r.names,"连六")
		r.score += 1
		return r
	}
	return r
}

func 老少副(hands Hands, r Output) Output {
	// same pattern of 1 2 3, 7 8 9
	validTiles := [][]int{{1,2,3},{7,8,9}, {10,11,12},{16,17,18}, {19,20,21},{25,26,27}}
	chisGroup := ExtractChi(hands.grouped)
	for _, g := range chisGroup {
		for _, vt := range validTiles{
			if checkContain(g, vt, len(g)){
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
	if checkContain(validTiles, handPongTiles, len(handPongTiles)) {
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
	if count == 1 {
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
	if !checkContained(validPattern, handPattern) {
		r.names = append(r.names,"缺一门")
		r.score += 1
		return r
	}
	return r
}

func 无字(hands Hands, r Output) Output {
	// doesnt contain 字 pattern
	handPattern := ExtractPatterns(hands.grouped)
	if !containedInt(handPattern, 4){
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
	if validity && groupChi{
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
	if validity && groupChi{
		r.names = append(r.names,"坎张")
		r.score += 1
		return r
	}
	return r
}

func 单钓(hands Hands, r Output) Output {
	// winning tile is one of the pair
	validPositions := []int{12, 13}
	if containedInt(validPositions, hands.winnngTile){
		r.names = append(r.names,"单钓")
		r.score += 1
		return r
	}
	return r
}
