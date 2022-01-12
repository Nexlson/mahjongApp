package calculator

import (
	"math/rand"
)

// tiles reference
TileRef := map[int]string{
	1: "一筒",
	2: "二筒",
	3: "三筒",
	4: "四筒",
	5: "五筒",
	6: "六筒",
	7: "七筒",
	8: "八筒",
	9: "九筒",
	10: "一条",
	11: "二条",
	12: "三条",
	13: "四条",
	14: "五条",
	15: "六条",
	16: "七条",
	17: "八条",
	18: "九条",
	19: "一万",
	20: "二万",
	21: "三万",
	22: "四万",
	23: "五万",
	24: "六万",
	25: "七万",
	26: "八万",
	27: "九万",
	28: "东",
	29: "南",
	30: "西",
	31: "北",
	32: "红中",
	33: "禄发",
	34: "白扳",
	35: "春",
	36: "夏",
	37: "秋",
	38: "冬",
	39: "梅",
	40: "兰",
	41: "竹",
	42: "菊",
}

// generate tiles
func tileGenerator() []string {
	var tiles []string
	pattern := [3]string{"条", "万", "筒"}
	number := [9]string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}
	words := [7]string{"东", "南", "西", "北", "中", "发", "白"}
	flower := [8]string{"春", "夏", "秋", "冬", "梅", "兰", "竹", "菊"}

	// 序数牌
	for _, s := range number {
		for _, x := range pattern {
			tiles = append(tiles, s+x)
		}
	}

	// 字牌
	for _, s := range words {
		tiles = append(tiles, s)
	}

	// multiply tiles
	tiles = tileMultiply(tiles, 4)

	// 花牌
	for _, s := range flower {
		tiles = append(tiles, s)
	}

	return tiles
}

// multiply tiles
func tileMultiply(list []string, times int) []string {
	var finalList []string
	for _, item := range list {
		n := 0
		for n < times {
			finalList = append(finalList, item)
			n += 1
		}
	}
	return finalList
}

// shuffleTiles
func shuffleTiles(list []string) []string {
	rand.Seed(100)
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

// filter contained tiles
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

// rules
func 大四喜(tiles []string) (bool, string, int) {
	valid := false
	name := "大四喜"
	score := 88
	winningRules := []int{28, 29, 30,31}
	exceptions := []int{}

	// check
	valid = contained(tiles, winningRules)

	return valid, name, score
}

func 大三元(tiles []string) (bool, string, int) {
	valid := false
	name := "大三元"
	score := 88
	winningRules := []int{32,33,34}

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




