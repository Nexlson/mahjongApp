// Rules for winning 

// 88 points
func 大四喜(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := int[]{28, 29, 30, 31}
	
	// check rules
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles

	// compare both list 
	if checkContained(validTiles, handTiles) { // compare if all pong labelled are same as rule
		r.names = append(r.names,"大四喜")
		r.score += 88
		exceptions = append(exceptions,[]int{38, 48, 60, 61, 73}...)
		return r, exceptions
	}

	// return nothing if not matched
	return r, exceptions
}

func 大三元(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int{32,33,34}

	// check rules
	handTiles := extractData(hands.grouped, "pong")

	// compare both list 
	if checkContained(validTiles, handTiles) {
		r.names = append(r.names,"大三元")
		r.score += 88
		exceptions = append(exceptions,[]int{54, 59, 73}...)
		return r, exceptions
	}

	return r, exceptions
}

func 绿一色(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	// contained tiao 2 3 4 6  8 fa 
	validTiles := []int {11,12,13,15,17,33}
	handTiles := removeDuplicateInt(hands.ungrouped)

	// if everything in hand is contained in valid tiles
	if checkContained(validTiles, handTiles) {
		r.names = append(r.names,"绿一色")
		r.score += 88
		if containedInt(handTiles, 33) {
			exceptions = append(exceptions,[]int{49}...)
		}else {
			exceptions = append(exceptions,[]int{49, 22}...)
		}
		return r, exceptions
	}
	return r, exceptions
}

func 九莲宝灯(hands StructuredTiles, r Output, exceptions []int) (Output, []int){
	validPongs := []int{1, 9, 10, 18, 19, 27}
	handPongs := extractData(hands.grouped, "pong")

	samePattern, _ := checkSamePattern(hands.ungrouped)
	contained := checkContained(validPongs, handPongs)

	// remove duplicate, pong tiles, 
	removedDup := removeDuplicateInt(hands.ungrouped)
	removePongs := removeIntFromSlice(removedDup, handPongs)

	if samePattern && contained && len(removePongs) == 7 {
		// only have 1112345678999
		r.names = append(r.names,"九莲宝灯")
		r.score += 88
		exceptions = append(exceptions,[]int{22, 56, 62, 76, 73}...)
		return r, exceptions
	}
	return r, exceptions
}

func 四杠(hands StructuredTiles, r Output, exceptions []int) (Output, []int){
	kongCount := extractData(hands.grouped, "kong")
	if len(kongCount) == 4 {
		r.names = append(r.names,"四杠")
		r.score += 88
		exceptions = append(exceptions,[]int{48, 79}...)
		return r, exceptions
	}
	return r, exceptions
}

func 连七对(hands StructuredTiles, r Output, exceptions []int) (Output, []int){
	samePattern, _ := checkSamePattern(hands.ungrouped)
	if checkPairs(hands.ungrouped) and samePattern {
		r.names = append(r.names,"连七对")
		r.score += 88
		exceptions = append(exceptions,[]int{22, 19, 56, 62, 76, 79}...)
		return r, exceptions
	}
	return r, exceptions
}

func 十三幺(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	if checkContained(validTiles, hands.ungrouped) {
		r.names = append(r.names,"十三幺")
		r.score += 88
		exceptions = append(exceptions,[]int{18, 51, 56, 62, 79}...)
		return r, exceptions
	}
	return r, exceptions
}

// 64 points
func 清幺九(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {1, 9, 10, 18, 19, 27}

	if checkContained(validTiles, removeDuplicateInt(hands.ungrouped)){
		r.names = append(r.names,"清幺九")
		r.score += 64
		exceptions = append(exceptions,[]int{48, 55, 65, 73, 76}...)
		return r, exceptions
	}
	return r, exceptions
}

func 小四喜(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {28, 29, 30, 31}
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles
	handTiles := append(handTiles, extractData(hands.grouped, "pair")...) // extract pair

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"小四喜")
		r.score += 64
		exceptions = append(exceptions,[]int{38, 73}...)
		return r, exceptions
	}
	return r, exceptions
}

func 小三元(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {32, 33, 34}
	handTiles := extractData(hands.grouped, "pong") // extract pong labelled tiles
	handTiles := append(handTiles, extractData(hands.grouped, "pair")...) // extract pair

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"小三元")
		r.score += 64
		exceptions = append(exceptions,[]int{54, 59, 73}...)
		return r, exceptions
	}
	return r, exceptions
}

func 字一色(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	samePattern, pattern := checkSamePattern(hands.ungrouped)
	if samePattern && pattern == 4 {
		r.names = append(r.names,"字一色")
		r.score += 64
		exceptions = append(exceptions,[]int{48, 55, 73}...)
		return r, exceptions
	}
	return r, exceptions
}

func 四暗刻(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	closedTiles := extractData(hands.grouped, "close")

	if len(closedTiles) == 4 {
		r.names = append(r.names,"四暗刻")
		r.score += 64
		exceptions = append(exceptions,[]int{48, 56, 62}...)
		return r, exceptions
	}
	return r, exceptions
}

func 一色双龙会(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	samePattern, pattern := checkSamePattern(hands.ungrouped)
	validTiles := PatternLib[pattern]
	validTiles := removeIntFromSlice(validTiles, []int{validTiles[3], validTiles[5]})
	
	if samePattern {
		if checkContained(validTiles, hands.ungrouped){
			r.names = append(r.names,"一色双龙会")
			r.score += 64
			exceptions = append(exceptions,[]int{19, 22, 63, 69, 72, 76}...)
			return r, exceptions
		}
		return r, exceptions
	}
}

func countStatus(grouped []TileGroup, status string) int{
	count := 0 
	for _, group := range grouped {
		if status == "chi" {
			if group.chi == 1{
				count += 1
			}
		}else if status == "pong"{
			if group.pong == 1{
				count += 1
			}
		}
	}
	return count
}

// same pattern for 4 group 
// func 一色四同顺(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
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
// 	return r, exceptions
// }

// func 一色四节高(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
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
// 	return r, exceptions
// }

// func 一色四步高() {
// 	name := "一色四步高"
// 	score := 32
// 	exceptions := []int{48, 55, 73}
// }

func 三杠(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	kongCount := extractData(hands.grouped, "kong")
	if len(kongCount) == 3 {
		r.names = append(r.names,"三杠")
		r.score += 32
		return r, exceptions
	}
	return r, exceptions
}

func 混幺九(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34}

	// 4 pong all related to 
	pongCount := countStatus(hands.grouped, "pong")
	handTiles := removeDuplicateInt(hands.ungrouped)

	if pongCount == 4 && checkContained(validTiles, handTiles) {
		r.names = append(r.names,"混幺九")
		r.score += 32
		exceptions = append(exceptions,[]int{48, 55, 73}...)
		return r, exceptions
	}
	return r, exceptions
}

func 七对(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	if checkPairs(hands.ungrouped) {
		r.names = append(r.names,"七对")
		r.score += 24
		exceptions = append(exceptions,[]int{56, 62, 79}...)
		return r, exceptions
	}
	return r, exceptions
}

func gap2Tiles(group []int) {
	prev := 0 
	for _, t := range group {

	}
}

// func 七星不靠(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
// 	name := "七星不靠"
// 	score := 24
// 	exceptions := []int{34, 51, 56, 62}
// 	validWord = []int{28, 29, 30, 31, 32, 33, 34}
// 	validContinous = [][]int{{}}

// 	// has validWord 
// }

func 全双刻(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles = []int{}
	handTiles = removeDuplicateInt(hands.ungrouped)
	// remove duplicate + compare to validTiles + len is 4 
	if checkContained(validTiles, handTiles) && len(handTiles) == 4 {
		r.names = append(r.names,"全双刻")
		r.score += 24
		exceptions = append(exceptions,[]int{48, 68, 76}...)
		return r, exceptions
	}
	return r, exceptions
}

func 清一色(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	samePattern, pattern := checkSamePattern(hands.ungrouped)
	if samePattern {
		r.names = append(r.names,"清一色")
		r.score += 24
		exceptions = append(exceptions,[]int{76}...)
		return r, exceptions
	}
	return r, exceptions
}

// func 一色三同顺() {
// 	name := "一色三同顺"
// 	score := 24
// 	exceptions := []int{24, 69}
// }

// func 一色三节高() {
// 	name := "一色三节高"
// 	score := 24
// 	exceptions := []int{23}
// }

func 全大(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int{7,8,9,16,17,18,25,26,27}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全大")
		r.score += 24
		exceptions = append(exceptions,[]int{36, 76}...)
		return r, exceptions
	}
	return r, exceptions
}

func 全中(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int{4, 5, 6, 13, 14, 15, 22, 23, 24}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全中")
		r.score += 24
		exceptions = append(exceptions,[]int{68, 76}...)
		return r, exceptions
	}
	return r, exceptions
}

func 全小(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int{1, 2, 3, 10, 11, 12, 19, 20, 21}
	handTiles := removeDuplicateInt(hands.ungrouped)

	if checkContained(validTiles, handTiles){
		r.names = append(r.names,"全小")
		r.score += 24
		exceptions = append(exceptions,[]int{37, 76}...)
		return r, exceptions
	}
	return r, exceptions
}

func 清龙(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := [][]int {{1,2,3,4,5,6,7,8,9}, {10,11,12,13,14,15,16,17,18}, {19,20,21,22,23,24,25,26,27}}
	handTiles := removeDuplicateInt(hands.ungrouped)
	for _, g := range validTiles {
		if checkContained(validTiles, handTiles){
			r.names = append(r.names,"清龙")
			r.score += 16
			exceptions = append(exceptions,[]int{71, 72}...)
			return r, exceptions
		}
	}
	return r, exceptions
}

// func 三色双龙会() {
// 	name := "三色双龙会"
// 	score := 16
// 	exceptions := []int{63, 72, 70, 76}
// }

// func 一色三步高() {
// 	name := "一色三步高"
// 	score := 16
// }

func 全带五(hands StructuredTiles, r Output, exceptions []int) (Output, []int) {
	validTiles := []int {5, 14, 23}

	for _, grouped := range hands.grouped {
		if ! checkContained(validTiles, grouped) {
			return r, exceptions
		}
	}
	r.names = append(r.names,"全带五")
	r.score += 16
	exceptions = append(exceptions,[]int{68, 76}...)
	return r, exceptions
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

// func 三同刻() {
// 	name := "三同刻"
// 	score := 16
// 	exceptions := []int{65}

// 3 pong + between each pong is 9 
// }

// func 三暗刻() {
// 	name := "三暗刻"
// 	score := 16

// closed count = 3 
// }

// func 全不靠() {
// 	name := "全不靠"
// 	score := 12
// 	exceptions := []int{51, 56, 62}
// }

// func 组合龙() {
// 	name := "组合龙"
// 	score := 12
// }

// func 大于五() {
// 	name := "大于五"
// 	score := 12
// 	exceptions := []int{76}
// }


// func 小于五() {
// 	name := "小于五"
// 	score := 12
// 	exceptions := []int{76}
// }

// func 三风刻() {
// 	name := "三风刻"
// 	score := 12
// 	exceptions := []int{73}
// }


// func 花龙() {
// 	name := "花龙"
// 	score := 8
// }

// func 推不倒() {
// 	name := "推不倒"
// 	score := 8
// 	exceptions := []int{75}
// }

// func 三色三同顺() {
// 	name := "三色三同顺"
// 	score := 8
// 	exceptions := []int{70}
// }

// func 三色三节高() {
// 	name := "三色三节高"
// 	score := 8
// }

// func 无番和() {
// 	name := "无番和"
// 	score := 8
// }

// func 妙手回春() {
// 	name := "妙手回春"
// 	score := 8
// 	exceptions := []int{80}
// }

// func 海底捞月() {
// 	name := "海底捞月"
// 	score := 8
// }

// func 杠上开花() {
// 	name := "杠上开花"
// 	score := 8
// 	exceptions := []int{80}
// }

// func 抢杠和() {
// 	name := "杠上开花"
// 	score := 8
// 	exceptions := []int{58}
// }

// func 碰碰和() {
// 	name := "杠上开花"
// 	score := 6
// }

// func 混一色() {
// 	name := "混一色"
// 	score := 6
// }

// func 三色三步高() {
// 	name := "三色三步高"
// 	score := 6
// }

// func 五门齐() {
// 	name := "五门齐"
// 	score := 6
// }
// func 全求人() {
// 	name := "全求人"
// 	score := 6
// 	exceptions := []int{79}
// }

// func 双暗杠() {
// 	name := "双暗杠"
// 	score := 6
// }

// func 双箭刻() {
// 	name := "双箭刻"
// 	score := 6
// 	exceptions := []int{59}
// }

// func 全带幺() {
// 	name := "全带幺"
// 	score := 4
// }

// func 不求人() {
// 	name := "不求人"
// 	score := 4
// 	exceptions := []int{80}
// }
// func 双明杠() {
// 	name := "双明杠"
// 	score := 4
// }

// func 和绝张() {
// 	name := "和绝张"
// 	score := 4
// }

// func 箭刻() {
// 	name := "箭刻"
// 	score := 2
// 	exceptions := []int{73}
// }

// func 圈风刻() {
// 	name := "圈风刻"
// 	score := 2
// 	exceptions := []int{73}
// }

// func 门风刻() {
// 	name := "门风刻"
// 	score := 2
// 	exceptions := []int{73}
// }

// func 门前清() {
// 	name := "门前清"
// 	score := 2
// }

// func 平和() {
// 	name := "平和"
// 	score := 2
// 	exceptions := []int{76}
// }

// func 四归一() {
// 	name := "四归一"
// 	score := 2
// }

// func 双同刻() {
// 	name := "双同刻"
// 	score := 2
// }

// func 双暗刻() {
// 	name := "双暗刻"
// 	score := 2
// }

// func 暗杠() {
// 	name := "暗杠"
// 	score := 2
// }

// func 断幺() {
// 	name := "断幺"
// 	score := 2
// 	exceptions := []int{76}
// }

// func 一般高() {
// 	name := "一般高"
// 	score := 1
// }

// func 喜相逢() {
// 	name := "喜相逢"
// 	score := 1
// }

// func 连六() {
// 	name := "连六"
// 	score := 1
// }

// func 老少副() {
// 	name := "老少副"
// 	score := 1
// }

// func 幺九刻() {
// 	name := "幺九刻"
// 	score := 1
// }

// func 明杠() {
// 	name := "明杠"
// 	score := 1
// }

// func 缺一门() {
// 	name := "缺一门"
// 	score := 1
// }

// func 无字() {
// 	name := "无字"
// 	score := 1
// }

// func 边张() {
// 	name := "边张"
// 	score := 1
// }

// func 坎张() {
// 	name := "坎张"
// 	score := 1
// }

// func 单钓() {
// 	name := "单钓"
// 	score := 1
// }

// func 自摸() {
// 	name := "自摸"
// 	score := 1
// }

// func 花牌() {
// 	name := "花牌"
// 	score := 1
// }
