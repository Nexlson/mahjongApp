package calculator

type Hands struct {
	Grouped []TileGroup 
	Ungrouped []int
	Won bool
	Score int
	WiningTile int // position of winningTile
	Round int // 圈风
	Turn int // 门风
}

type TileGroup struct {
	Tiles []int
	Open bool
	Pong bool
	Kong bool 
	Chi bool
	Pair bool
	Pattern int // 1 - 筒, 2 - 条, 3 - 万 4 - 风 5 - 箭
}

type Output struct {
	Names []string
	Score int
	Exceptions []int
}

var PatternLib = map[int][]int {
	// 筒
	1: {1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 条
	2: {10, 11, 12, 13, 14, 15, 16, 17, 18},
	// 万
	3: {19, 20, 21, 22, 23, 24, 25, 26, 27},
	// 风
	4: {28, 29, 30, 31, 32, 33, 34},
	// 箭
	5: {32, 33, 34},
}

// map tiles to id
var TilesMaps = map[int]string{
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

func generateFunctions() []func(Hands, Output)Output {
	var rulesSlice []func(Hands, Output)Output
	rulesSlice = append(rulesSlice, 大四喜)
	rulesSlice = append(rulesSlice, 大三元)
	rulesSlice = append(rulesSlice, 绿一色)
	rulesSlice = append(rulesSlice, 九莲宝灯)
	rulesSlice = append(rulesSlice, 四杠)
	rulesSlice = append(rulesSlice, 连七对)
	rulesSlice = append(rulesSlice, 十三幺)
	rulesSlice = append(rulesSlice, 清幺九)
	rulesSlice = append(rulesSlice, 小四喜)
	rulesSlice = append(rulesSlice, 小三元)
	rulesSlice = append(rulesSlice, 字一色)
	rulesSlice = append(rulesSlice, 四暗刻)
	rulesSlice = append(rulesSlice, 一色双龙会)
	rulesSlice = append(rulesSlice, 一色四同顺)
	rulesSlice = append(rulesSlice, 一色四节高)
	rulesSlice = append(rulesSlice, 一色四步高)
	rulesSlice = append(rulesSlice, 三杠)
	rulesSlice = append(rulesSlice, 混幺九)
	rulesSlice = append(rulesSlice, 七对)
	rulesSlice = append(rulesSlice, 七星不靠)
	rulesSlice = append(rulesSlice, 全双刻)
	rulesSlice = append(rulesSlice, 清一色)
	rulesSlice = append(rulesSlice, 一色三同顺)
	rulesSlice = append(rulesSlice, 一色三节高)
	rulesSlice = append(rulesSlice, 全大)
	rulesSlice = append(rulesSlice, 全中)
	rulesSlice = append(rulesSlice, 全小)
	rulesSlice = append(rulesSlice, 清龙)
	rulesSlice = append(rulesSlice, 三色双龙会)
	rulesSlice = append(rulesSlice, 一色三步高)
	rulesSlice = append(rulesSlice, 全带五)
	rulesSlice = append(rulesSlice, 三同刻)
	rulesSlice = append(rulesSlice, 三暗刻)
	rulesSlice = append(rulesSlice, 全不靠)
	rulesSlice = append(rulesSlice, 组合龙)
	rulesSlice = append(rulesSlice, 大于五)
	rulesSlice = append(rulesSlice, 小于五)
	rulesSlice = append(rulesSlice, 三风刻)
	rulesSlice = append(rulesSlice, 花龙)
	rulesSlice = append(rulesSlice, 推不倒)
	rulesSlice = append(rulesSlice, 三色三同顺)
	rulesSlice = append(rulesSlice, 三色三节高)
	rulesSlice = append(rulesSlice, 碰碰和)
	rulesSlice = append(rulesSlice, 混一色)
	rulesSlice = append(rulesSlice, 三色三步高)
	rulesSlice = append(rulesSlice, 五门齐)
	rulesSlice = append(rulesSlice, 全求人)
	rulesSlice = append(rulesSlice, 双暗杠)
	rulesSlice = append(rulesSlice, 双箭刻)
	rulesSlice = append(rulesSlice, 全带幺)
	rulesSlice = append(rulesSlice, 不求人)
	rulesSlice = append(rulesSlice, 双明杠)
	rulesSlice = append(rulesSlice, 箭刻)
	rulesSlice = append(rulesSlice, 圈风刻)
	rulesSlice = append(rulesSlice, 门风刻)
	rulesSlice = append(rulesSlice, 门前清)
	rulesSlice = append(rulesSlice, 平和)
	rulesSlice = append(rulesSlice, 四归一)
	rulesSlice = append(rulesSlice, 双同刻)
	rulesSlice = append(rulesSlice, 双暗刻)
	rulesSlice = append(rulesSlice, 暗杠)
	rulesSlice = append(rulesSlice, 断幺)
	rulesSlice = append(rulesSlice, 一般高)
	rulesSlice = append(rulesSlice, 喜相逢)
	rulesSlice = append(rulesSlice, 连六)
	rulesSlice = append(rulesSlice, 老少副)
	rulesSlice = append(rulesSlice, 幺九刻)
	rulesSlice = append(rulesSlice, 明杠)
	rulesSlice = append(rulesSlice, 缺一门)
	rulesSlice = append(rulesSlice, 无字)
	rulesSlice = append(rulesSlice, 边张)
	rulesSlice = append(rulesSlice, 坎张)
	rulesSlice = append(rulesSlice, 单钓)
	rulesSlice = append(rulesSlice, 无番和)
	return rulesSlice
}


// map ruleFunc to id
var RulesMaps = map[int]interface{}{
	0: 大四喜,
	1: 大三元,
	2: 绿一色,
	3: 九莲宝灯,
	4: 四杠,
	5: 连七对,
	6: 十三幺,
	7: 清幺九,
	8: 小四喜,
	9: 小三元,
	10: 字一色,
	11: 四暗刻,
	12: 一色双龙会,
	13: 一色四同顺,
	14: 一色四节高,
	15: 一色四步高,
	16: 三杠,
	17: 混幺九,
	18: 七对,
	19: 七星不靠,
	20: 全双刻,
	21: 清一色,
	22: 一色三同顺,
	23: 一色三节高,
	24: 全大,
	25: 全中,
	26: 全小,
	27: 清龙,
	28: 三色双龙会,
	29: 一色三步高,
	30: 全带五,
	31: 三同刻,
	32: 三暗刻,
	33: 全不靠,
	34: 组合龙,
	35: 大于五,
	36: 小于五,
	37: 三风刻,
	38: 花龙,
	39: 推不倒,
	40: 三色三同顺,
	41: 三色三节高,
	42: 碰碰和,
	43: 混一色,
	44: 三色三步高,
	45: 五门齐,
	46: 全求人,
	47: 双暗杠,
	48: 双箭刻,
	49: 全带幺,
	50: 不求人,
	51: 双明杠,
	52: 箭刻,
	53: 圈风刻,
	54: 门风刻,
	55: 门前清,
	56: 平和,
	57: 四归一,
	58: 双同刻,
	59: 双暗刻,
	60: 暗杠,
	61: 断幺,
	62: 一般高,
	63: 喜相逢,
	64: 连六,
	65: 老少副,
	66: 幺九刻,
	67: 明杠,
	68: 缺一门,
	69: 无字,
	70: 边张,
	71: 坎张,
	72: 单钓,
	73: 无番和,
}