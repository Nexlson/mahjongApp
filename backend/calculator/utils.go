package calculator

import (
	"math/rand"
)

type TileGroup struct {
	tiles []int
	status int
}

type Output struct {
	names []string
	score int
}

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

// maps ruleFunc to id
rulesMaps := map[int]func{
	1: 大四喜,
	2: 大三元,
	3: 绿一色,
	4: 九莲宝灯,
	5: 四杠,
	6: 连七对,
	7: 十三幺,
	8: 清幺九,
	9: 小四喜,
	10: 小三元,
	11: 字一色,
	12: 四暗刻,
	13: 一色双龙会,
	14: 一色四同顺,
	15: 一色四节高,
	16: 一色四步高,
	17: 三杠,
	18: 混幺九,
	19: 七对,
	20: 七星不靠,
	21: 全双刻,
	22: 清一色,
	23: 一色三同顺,
	24: 一色三节高,
	25: 全大,
	26: 全中,
	27: 全小,
	28: 清龙,
	29: 三色双龙会,
	30: 一色三步高,
	31: 全带五,
	32: 三同刻,
	33: 三暗刻,
	34: 全不靠,
	35: 组合龙,
	36: 大于五,
	37: 小于五,
	38: 三风刻,
	39: 花龙,
	40: 推不倒,
	41: 三色三同顺,
	42: 三色三节高,
	43: 无番和,
	44: 妙手回春,
	45: 海底捞月,
	46: 杠上开花,
	47: 抢杠和,
	48: 碰碰和,
	49: 混一色,
	50: 三色三步高,
	51: 五门齐,
	52: 全求人,
	53: 双暗杠,
	54: 双箭刻,
	55: 全带幺,
	56: 不求人,
	57: 双明杠,
	58: 和绝张,
	59: 箭刻,
	60: 圈风刻,
	61: 门风刻,
	62: 门前清,
	63: 平和,
	64: 四归一,
	65: 双同刻,
	66: 双暗刻,
	67: 暗杠,
	68: 断幺,
	69: 一般高,
	70: 喜相逢,
	71: 连六,
	72: 老少副,
	73: 幺九刻,
	74: 明杠,
	75: 缺一门,
	76: 无字,
	77: 边张,
	78: 坎张,
	79: 单钓,
	80: 自摸,
	81: 花牌,
}

func checkRules() {

}

func 大四喜(tiles []string) (string, int, []int) {
	// winning hands has pong of 28,29,30,31
	name := "大四喜"
	score := 88
	exceptions := []int{38, 48, 60, 61, 73}

	// check rules

	return name, score, exceptions
}

func checkPong(group tileGroup, tile int) (bool){
	// return pong boolean and status
	for _, t = range group.tiles {
		if t != tile {
			return false
		}
	}
	return true
}

func 大三元(hands []tileGroup) (bool, string, int) {
	name := "大三元"
	score := 88
	exceptions := []int{54, 59, 73}

	return valid, name, score
}

func 绿一色() {
	name := "绿一色"
	score := 88
	exceptions := []int{49}

	// if no fa -> can count 22
}

func 九莲宝灯() {
	// only have 1112345678999
	name := "九莲宝灯"
	score := 88
	exceptions := []int{22, 56, 62, 76, 73}
}

func 四杠() {
	name := "四杠"
	score := 88
	exceptions := []int{48, 79}
}

func 连七对() {
	name := "连七对"
	score := 88
	exceptions := []int{22, 19, 56, 62, 76, 79}
}

func 十三幺() {
	name := "十三幺"
	score := 88
	exceptions := []int{18, 51, 56, 62, 79}
}

func 清幺九() {
	name := "清幺九"
	score := 64
	exceptions := []int{48, 55, 65, 73, 76}
}

func 小四喜() {
	name := "小四喜"
	score := 64
	exceptions := []int{38, 73}
}


func 小三元() {
	name := "小三元"
	score := 64
	exceptions := []int{54, 59, 73}
}

func 字一色() {
	name := "字一色"
	score := 64
	exceptions := []int{48, 55, 73}
}

func 四暗刻() {
	name := "四暗刻"
	score := 64
	exceptions := []int{48, 56, 62}
}

func 一色双龙会() {
	name := "一色双龙会"
	score := 64
	exceptions := []int{19, 22, 63, 69, 72, 76}
}

func 一色四同顺() {
	name := "一色四同顺"
	score := 48
	exceptions := []int{23, 24, 64, 69}
}

func 一色四节高() {
	name := "一色四节高"
	score := 48
	exceptions := []int{23, 24, 48}
}

func 一色四步高() {
	name := "一色四节高"
	score := 32
	exceptions := []int{30, 71, 72}
}

func 三杠() {
	name := "三杠"
	score := 32
}

func 混幺九() {
	name := "混幺九"
	score := 32
	exceptions := []int{48, 55, 73}
}

func 七对() {
	name := "七对"
	score := 24
	exceptions := []int{56, 62, 79}
}

func 七星不靠() {
	name := "七星不靠"
	score := 24
	exceptions := []int{34, 51, 56, 62}
}

func 全双刻() {
	name := "全双刻"
	score := 24
	exceptions := []int{48, 68, 76}
}

func 清一色() {
	name := "清一色"
	score := 24
	exceptions := []int{76}
}

func 一色三同顺() {
	name := "一色三同顺"
	score := 24
	exceptions := []int{24, 69}
}

func 一色三节高() {
	name := "一色三节高"
	score := 24
	exceptions := []int{23}
}

func 全大() {
	name := "全大"
	score := 24
	exceptions := []int{36, 76}
}

func 全中() {
	name := "全中"
	score := 24
	exceptions := []int{68, 76}
}

func 全小() {
	name := "全小"
	score := 24
	exceptions := []int{37, 76}
}

func 清龙() {
	name := "清龙"
	score := 16
	exceptions := []int{71, 72}
}

func 三色双龙会() {
	name := "三色双龙会"
	score := 16
	exceptions := []int{63, 72, 70, 76}
}

func 一色三步高() {
	name := "一色三步高"
	score := 16
}

func 全带五() {
	name := "全带五"
	score := 16
	exceptions := []int{68, 76}
}

func 三同刻() {
	name := "三同刻"
	score := 16
	exceptions := []int{65}
}

func 三暗刻() {
	name := "三暗刻"
	score := 16
}

func 全不靠() {
	name := "全不靠"
	score := 12
	exceptions := []int{51, 56, 62}
}

func 组合龙() {
	name := "组合龙"
	score := 12
}

func 大于五() {
	name := "大于五"
	score := 12
	exceptions := []int{76}
}


func 小于五() {
	name := "小于五"
	score := 12
	exceptions := []int{76}
}

func 三风刻() {
	name := "三风刻"
	score := 12
	exceptions := []int{73}
}


func 花龙() {
	name := "花龙"
	score := 8
}

func 推不倒() {
	name := "推不倒"
	score := 8
	exceptions := []int{75}
}

func 三色三同顺() {
	name := "三色三同顺"
	score := 8
	exceptions := []int{70}
}

func 三色三节高() {
	name := "三色三节高"
	score := 8
}

func 无番和() {
	name := "无番和"
	score := 8
}

func 妙手回春() {
	name := "妙手回春"
	score := 8
	exceptions := []int{80}
}

func 海底捞月() {
	name := "海底捞月"
	score := 8
}

func 杠上开花() {
	name := "杠上开花"
	score := 8
	exceptions := []int{80}
}

func 抢杠和() {
	name := "杠上开花"
	score := 8
	exceptions := []int{58}
}

func 碰碰和() {
	name := "杠上开花"
	score := 6
}

func 混一色() {
	name := "混一色"
	score := 6
}

func 三色三步高() {
	name := "三色三步高"
	score := 6
}

func 五门齐() {
	name := "五门齐"
	score := 6
}
func 全求人() {
	name := "全求人"
	score := 6
	exceptions := []int{79}
}

func 双暗杠() {
	name := "双暗杠"
	score := 6
}

func 双箭刻() {
	name := "双箭刻"
	score := 6
	exceptions := []int{59}
}

func 全带幺() {
	name := "全带幺"
	score := 4
}

func 不求人() {
	name := "不求人"
	score := 4
	exceptions := []int{80}
}
func 双明杠() {
	name := "双明杠"
	score := 4
}

func 和绝张() {
	name := "和绝张"
	score := 4
}

func 箭刻() {
	name := "箭刻"
	score := 2
	exceptions := []int{73}
}

func 圈风刻() {
	name := "圈风刻"
	score := 2
	exceptions := []int{73}
}

func 门风刻() {
	name := "门风刻"
	score := 2
	exceptions := []int{73}
}

func 门前清() {
	name := "门前清"
	score := 2
}

func 平和() {
	name := "平和"
	score := 2
	exceptions := []int{76}
}

func 四归一() {
	name := "四归一"
	score := 2
}

func 双同刻() {
	name := "双同刻"
	score := 2
}

func 双暗刻() {
	name := "双暗刻"
	score := 2
}

func 暗杠() {
	name := "暗杠"
	score := 2
}

func 断幺() {
	name := "断幺"
	score := 2
	exceptions := []int{76}
}

func 一般高() {
	name := "一般高"
	score := 1
}

func 喜相逢() {
	name := "喜相逢"
	score := 1
}

func 连六() {
	name := "连六"
	score := 1
}

func 老少副() {
	name := "老少副"
	score := 1
}

func 幺九刻() {
	name := "幺九刻"
	score := 1
}

func 明杠() {
	name := "明杠"
	score := 1
}

func 缺一门() {
	name := "缺一门"
	score := 1
}

func 无字() {
	name := "无字"
	score := 1
}

func 边张() {
	name := "边张"
	score := 1
}

func 坎张() {
	name := "坎张"
	score := 1
}

func 单钓() {
	name := "单钓"
	score := 1
}

func 自摸() {
	name := "自摸"
	score := 1
}

func 花牌() {
	name := "花牌"
	score := 1
}









