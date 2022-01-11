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
