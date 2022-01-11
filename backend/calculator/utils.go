package calculator

import (
	"math/rand"
)

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

func shuffleTiles(list []string) []string {
	rand.Seed(100)
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
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
