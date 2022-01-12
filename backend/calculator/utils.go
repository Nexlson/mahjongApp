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
type ruleFunction func (string, int, int)

rulesMaps := map[int]interface{}{
	1: 大四喜,
	2: 大三元,

}


func checkRules() {

}

func 大四喜(tiles []string) (string, int, int) {
	name := "大四喜"
	score := 88
	winningRules := []int{28, 29, 30,31}
	exceptions := []int{}

	// check
	valid = contained(tiles, winningRules)

	return name, score
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

func 清幺九() {
	// only have 2 3 4 6 8 条
}

func 小四喜() {
	// only have 2 3 4 6 8 条
}


func 小三元() {
	// only have 2 3 4 6 8 条
}

func 字一色() {
	// only have 2 3 4 6 8 条
}

func 四暗刻() {
	// only have 2 3 4 6 8 条
}

func 一色双龙会() {
	// only have 2 3 4 6 8 条
}


func 一色四同顺() {
	// only have 2 3 4 6 8 条
}

func 一色四节高() {
	// only have 2 3 4 6 8 条
}

func 一色四步高() {
	// only have 2 3 4 6 8 条
}

func 三杠() {
	// only have 2 3 4 6 8 条
}

func 混幺九() {
	// only have 2 3 4 6 8 条
}

func 七对() {
	// only have 2 3 4 6 8 条
}

func 七星不靠() {
	// only have 2 3 4 6 8 条
}

func 全双刻() {
	// only have 2 3 4 6 8 条
}

func 清一色() {
	// only have 2 3 4 6 8 条
}

func 一色三同顺() {
	// only have 2 3 4 6 8 条
}

func 一色三节高() {
	// only have 2 3 4 6 8 条
}

func 全大() {
	// only have 2 3 4 6 8 条
}

func 全中() {
	// only have 2 3 4 6 8 条
}

func 全小() {
	// only have 2 3 4 6 8 条
}

func 清龙() {
	// only have 2 3 4 6 8 条
}

func 三色双龙会() {
	// only have 2 3 4 6 8 条
}

func 一色三步高() {
	// only have 2 3 4 6 8 条
}

func 全带五() {
	// only have 2 3 4 6 8 条
}

func 三同刻() {
	// only have 2 3 4 6 8 条
}

func 三暗刻() {
	// only have 2 3 4 6 8 条
}

func 全不靠() {
	// only have 2 3 4 6 8 条
}

func 组合龙() {
	// only have 2 3 4 6 8 条
}

func 大于五() {
	// only have 2 3 4 6 8 条
}


func 小于五() {
	// only have 2 3 4 6 8 条
}

func 三风刻() {
	// only have 2 3 4 6 8 条
}


func 花龙() {
	// only have 2 3 4 6 8 条
}

func 推不倒() {
	// only have 2 3 4 6 8 条
}

func 三色三同顺() {
	// only have 2 3 4 6 8 条
}

func 三色三节高() {
	// only have 2 3 4 6 8 条
}

func 无番和() {
	// only have 2 3 4 6 8 条
}

func 妙手回春() {
	// only have 2 3 4 6 8 条
}

func 海底捞月() {
	// only have 2 3 4 6 8 条
}

func 杠上开花() {
	// only have 2 3 4 6 8 条
}

func 抢杠和() {
	// only have 2 3 4 6 8 条
}

func 碰碰和() {
	// only have 2 3 4 6 8 条
}

func 混一色() {
	// only have 2 3 4 6 8 条
}

func 三色三步高() {
	// only have 2 3 4 6 8 条
}

func 五门齐() {
	// only have 2 3 4 6 8 条
}
func 全求人() {
	// only have 2 3 4 6 8 条
}

func 双暗杠() {
	// only have 2 3 4 6 8 条
}

func 双箭刻() {
	// only have 2 3 4 6 8 条
}

func 全带幺() {
	// only have 2 3 4 6 8 条
}

func 不求人() {
	// only have 2 3 4 6 8 条
}
func 双明杠() {
	// only have 2 3 4 6 8 条
}

func 和绝张() {
	// only have 2 3 4 6 8 条
}

func 箭刻() {
	// only have 2 3 4 6 8 条
}

func 圈风刻() {
	// only have 2 3 4 6 8 条
}

func 门风刻() {
	// only have 2 3 4 6 8 条
}

func 门前清() {
	// only have 2 3 4 6 8 条
}

func 平和() {
	// only have 2 3 4 6 8 条
}

func 四归一() {
	// only have 2 3 4 6 8 条
}

func 双同刻() {
	// only have 2 3 4 6 8 条
}

func 双暗刻() {
	// only have 2 3 4 6 8 条
}

func 暗杠() {
	// only have 2 3 4 6 8 条
}

func 断幺() {
	// only have 2 3 4 6 8 条
}

func 一般高() {
	// only have 2 3 4 6 8 条
}

func 喜相逢() {
	// only have 2 3 4 6 8 条
}

func 连六() {
	// only have 2 3 4 6 8 条
}

func 老少副() {
	// only have 2 3 4 6 8 条
}

func 幺九刻() {
	// only have 2 3 4 6 8 条
}

func 明杠() {
	// only have 2 3 4 6 8 条
}

func 缺一门() {
	// only have 2 3 4 6 8 条
}

func 无字() {
	// only have 2 3 4 6 8 条
}

func 边张() {
	// only have 2 3 4 6 8 条
}

func 坎张() {
	// only have 2 3 4 6 8 条
}

func 单钓() {
	// only have 2 3 4 6 8 条
}

func 自摸() {
	// only have 2 3 4 6 8 条
}

func 花牌() {
	// only have 2 3 4 6 8 条
}









