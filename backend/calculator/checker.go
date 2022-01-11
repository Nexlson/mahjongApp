package main

type winningHands struct {
	winningRules func()
	name         string
	score        int
	exceptions   int
}

// func winningHandGenerator() {
// 	var rules []winningHands

// 	var 大四喜 winningHands
// 	大四喜.name = "大四喜"
// 	大四喜.score = 88
// 	大四喜.winningRules = 大四喜
// }

func checkWin(hand []string) {

}

func checkPong() {

}

func checkChi() {

}

func 大四喜(tiles []string) (bool, string, int) {
	valid := false
	name := "大四喜"
	score := 88
	winningRules := []string{"东", "南", "西", "北", "东", "南", "西", "北", "东", "南", "西", "北", "东", "南", "西", "北"}

	// check
	valid = contained(tiles, winningRules)

	return valid, name, score
}

func 大三元(tiles []string) (bool, string, int) {
	valid := false
	name := "大三元"
	score := 88
	winningRules := []string{"中", "发", "白", "中", "发", "白", "中", "发", "白"}

	// check
	valid = contained(tiles, winningRules)

	return valid, name, score
}

func 绿一色() {
	// only have 2 3 4 6 8 条
}

func 九莲宝灯() {
	// only have 1112345678999
}

func 四杠() {
	// only have 2 3 4 6 8 条
}

func 连七对() {
	// only have 2 3 4 6 8 条
}

func 十三幺() {
	// only have 2 3 4 6 8 条
}

func contained(tiles []string, rule []string) bool {
	counts := len(rule)
	count := 0
	for _, x := range tiles {
		for _, y := range rule {
			if x == y {
				count += 1
				break
			}
		}
	}
	if count == counts {
		return true
	} else {
		return false
	}
}

func main() {
	大四喜([]string{"东", "南", "西", "北", "东", "南", "西", "北", "东", "南", "西", "北", "东", "南", "西", "北", "a", "da", "a"})
}
