package calculator

import "testing"

func Test大四喜(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{28, 28, 28}, false, true, false, false, false, 5},
		{[]int{29, 29, 29}, false, true, false, false, false, 5},
		{[]int{30, 30, 30}, false, true, false, false, false, 5},
		{[]int{31, 31, 31}, false, true, false, false, false, 5},
		{[]int{1, 1}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28, 28, 28, 29, 29, 29, 30, 30, 30, 31, 31, 31, 1, 1},
		won: true,
	}
	result := Output{}

	// test function 
	got := 大四喜(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test大三元(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{32, 32, 32}, false, true, false, false, false, 5},
		{[]int{33, 33, 33}, false, true, false, false, false, 5},
		{[]int{34, 34, 34}, false, true, false, false, false, 5},
		{[]int{31, 31, 31}, false, true, false, false, false, 5},
		{[]int{1, 1}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32, 32, 32, 33, 33, 33, 34, 34, 34, 31, 31, 31, 1, 1},
		won: true,
	}
	result := Output{}

	// test function 
	got := 大三元(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test绿一色(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{11, 11, 11}, false, true, false, false, false, 2},
		{[]int{12, 12, 12}, false, true, false, false, false, 2},
		{[]int{15, 15, 15}, false, true, false, false, false, 2},
		{[]int{33, 33, 33}, false, true, false, false, false, 5},
		{[]int{17, 17}, false, false, false, false, true, 2},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{11,11,11,12,12,12,15,15,15,33,33,33,17,17},
		won: true,
	}
	result := Output{}

	// test function 
	got := 绿一色(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test九莲宝灯(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{1, 1, 1}, false, true, false, false, false, 1},
		{[]int{9, 9, 9}, false, true, false, false, false, 1},
		{[]int{2, 3, 4}, false, false, false, true, false, 1},
		{[]int{5, 6, 7}, false, false, false, true, false, 1},
		{[]int{8, 8}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,9,9,9,2,3,4,5,6,7,8,8},
		won: true,
	}
	result := Output{}

	// test function 
	got := 九莲宝灯(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test四杠(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{1, 1, 1}, false, false, true, false, false, 1},
		{[]int{9, 9, 9}, false, false, true, false, false, 1},
		{[]int{2, 2, 2}, false, false, true, false, false, 1},
		{[]int{3, 3, 3}, false, false, true, false, false, 1},
		{[]int{8, 8}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,9,9,9,2,2,2,3,3,3,8,8},
		won: true,
	}
	result := Output{}

	// test function 
	got := 四杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test连七对(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{1, 1, 2}, false, false, false, false, true, 1},
		{[]int{2, 3, 3}, false, false, false, false, true, 1},
		{[]int{4, 4, 5}, false, false, false, false, true, 1},
		{[]int{5, 6, 6}, false, false, false, false, true, 1},
		{[]int{8, 8}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,2,2,3,3,4,4,5,5,6,6,8,8},
		won: true,
	}
	result := Output{}

	// test function 
	got := 连七对(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test十三幺(t *testing.T) {
	// const
	want := 88
	groupedHand := []TileGroup{
		{[]int{1, 9, 10}, false, false, false, false, false, 5},
		{[]int{18, 19, 27}, false, false, false, false, false, 5},
		{[]int{28, 29, 30}, false, false, false, false, false, 5},
		{[]int{31, 32, 33}, false, false, false, false, false, 5},
		{[]int{34, 34}, false, false, false, false, true, 5},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1, 9, 10, 18, 19, 27, 28, 29, 30, 31, 32, 33, 34, 34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 十三幺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

// func Test清幺九(t *testing.T) {
// 	// const
// 	want := 64
// 	groupedHand := []TileGroup{
// 		{[]int{1, 1, 1}, false, true, false, false, false, 1},
// 		{[]int{9, 9, 9}, false, true, false, false, false, 1},
// 		{[]int{10, 10, 10}, false, true, false, false, false, 2},
// 		{[]int{18, 18, 18}, false, true, false, false, false, 2},
// 		{[]int{19, 19}, false, false, false, false, true, 3},
// 	}
// 	hands := Hands{
// 		grouped: groupedHand,
// 		ungrouped: []int{1,1,1,9,9,9,10,10,10,18,18,18,19,19},
// 	}
// 	result := Output{}

// 	// test function 
// 	got = 清幺九(hands, result) 

// 	// check results
// 	if got.score != want {
// 		t.Errorf("Test output %d is not equal to expected output %d", got, want)
// 	}
// }

func Test小四喜(t *testing.T) {
	// const
	want := 64
	groupedHand := []TileGroup{
		{[]int{28, 28, 28}, false, true, false, false, false, 4},
		{[]int{29, 29, 29}, false, true, false, false, false, 4},
		{[]int{30, 30, 30}, false, true, false, false, false, 4},
		{[]int{18, 18, 18}, false, true, false, false, false, 2},
		{[]int{31, 31}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28,28,28,29,29,29,30,30,30,18,18,18,31,31},
		won: true,
	}
	result := Output{}

	// test function 
	got := 小四喜(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test小三元(t *testing.T) {
	// const
	want := 64
	groupedHand := []TileGroup{
		{[]int{32, 32, 32}, false, true, false, false, false, 4},
		{[]int{33, 33, 33}, false, true, false, false, false, 4},
		{[]int{30, 30, 30}, false, true, false, false, false, 4},
		{[]int{18, 18, 18}, false, true, false, false, false, 2},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,32,33,33,33,30,30,30,18,18,18,34,34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 小三元(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test字一色(t *testing.T) {
	// const
	want := 64
	groupedHand := []TileGroup{
		{[]int{32, 32, 32}, false, true, false, false, false, 4},
		{[]int{33, 33, 33}, false, true, false, false, false, 4},
		{[]int{30, 30, 30}, false, true, false, false, false, 4},
		{[]int{29, 29, 29}, false, true, false, false, false, 4},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,32,33,33,33,30,30,30,29,29,29,34,34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 字一色(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test四暗刻(t *testing.T) {
	// const
	want := 64
	groupedHand := []TileGroup{
		{[]int{32, 32, 32}, false, true, false, false, false, 4},
		{[]int{33, 33, 33}, false, true, false, false, false, 4},
		{[]int{30, 30, 30}, false, true, false, false, false, 4},
		{[]int{29, 29, 29}, false, true, false, false, false, 4},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,32,33,33,33,30,30,30,29,29,29,34,34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 四暗刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色双龙会(t *testing.T) {
	want := 64
	groupedHand := []TileGroup{
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{7, 8, 9}, false, false, false, true, false, 1},
		{[]int{7, 8, 9}, false, false, false, true, false, 1},
		{[]int{5, 5}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,1,2,3,7,8,9,7,8,9,5,5},
	}
	result := Output{}

	// test function 
	got := 一色双龙会(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色四同顺(t *testing.T) {
	// const
	want := 48
	groupedHand := []TileGroup{
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,1,2,3,1,2,3,1,2,3,34,34},
	}
	result := Output{}

	// test function 
	got := 一色四同顺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色四节高(t *testing.T) {
	// const
	want := 48
	groupedHand := []TileGroup{
		{[]int{1, 1, 1}, false, true, false, false, false, 1},
		{[]int{2, 2, 2}, false, true, false, false, false, 1},
		{[]int{3, 3, 3}, false, true, false, false, false, 1},
		{[]int{4, 4, 4}, false, true, false, false, false, 1},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,2,2,2,3,3,3,4,4,4,34,34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 一色四节高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色四步高(t *testing.T) {
	// const
	want := 32
	groupedHand := []TileGroup{
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{2, 3, 4}, false, false, false, true, false, 1},
		{[]int{3, 4, 5}, false, false, false, true, false, 1},
		{[]int{4, 5, 6}, false, false, false, true, false, 1},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,2,3,4,3,4,5,4,5,6,34,34},
		won : true,
	}
	result := Output{}

	// test function 
	got := 一色四步高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三杠(t *testing.T) {
	// const
	want := 32
	groupedHand := []TileGroup{
		{[]int{32, 32, 32}, false, false, true, false, false, 4},
		{[]int{33, 33, 33}, false, false, true, false, false, 4},
		{[]int{30, 30, 30}, false, false, true, false, false, 4},
		{[]int{29, 29, 29}, false, true, false, false, false, 4},
		{[]int{34, 34}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,32,33,33,33,30,30,30,29,29,29,34,34},
		won: true,
	}
	result := Output{}

	// test function 
	got := 三杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got, want)
	}
}

func Test混幺九(t *testing.T) {
	// const
	want := 32
	groupedHand := []TileGroup{
		{[]int{1, 1, 1}, false, true, false, false, false, 1},
		{[]int{9, 9, 9}, false, true, false, false, false, 1},
		{[]int{10, 10, 10}, false, true, false, false, false, 2},
		{[]int{29, 29, 29}, false, true, false, false, false, 4},
		{[]int{30, 30}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,9,9,9,10,10,10,29,29,29,34,34},
	}
	result := Output{}

	// test function 
	got := 混幺九(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test七对(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{32, 32, 33}, false, false, true, false, true, 4},
		{[]int{33, 34, 34}, false, false, true, false, true, 4},
		{[]int{30, 30, 9}, false, false, true, false, true, 4},
		{[]int{9, 29, 29}, false, true, false, false, true, 4},
		{[]int{1, 1}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,33,33,34,34,30,30,9,9,29,29,1,1},
		won: true,
	}
	result := Output{}

	// test function 
	got := 七对(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test七星不靠(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{28, 29, 30}, false, false, true, false, true, 4},
		{[]int{31, 32, 33}, false, false, true, false, true, 4},
		{[]int{34, 1, 4}, false, false, true, false, true, 4},
		{[]int{7, 11, 14}, false, true, false, false, true, 4},
		{[]int{17, 21}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28, 29, 30, 31, 32, 33, 34, 1, 4, 7, 11, 14, 17, 21},
	}
	result := Output{}

	// test function 
	got := 七星不靠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全双刻(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{2, 2, 2}, false, true, false, false, false, 1},
		{[]int{4, 4, 4}, false, true, false, false, false, 1},
		{[]int{6, 6, 6}, false, true, false, false, false, 1},
		{[]int{8, 8, 8}, false, true, false, false, false, 1},
		{[]int{11, 11}, false, false, false, false, true, 2},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{2,2,2,4,4,4,6,6,6,8,8,8,11,11},
		won: true,
	}
	result := Output{}

	// test function 
	got := 全双刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test清一色(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{1, 2, 3}, false, false, true, false, false, 1},
		{[]int{4, 5, 6}, false, false, true, false, false, 1},
		{[]int{7, 8, 9}, false, false, true, false, false, 1},
		{[]int{2, 3, 4}, false, false, true, false, false, 1},
		{[]int{3, 3}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,4,5,6,7,8,9,2,3,4,3,3},
		won: true,
	}
	result := Output{}

	// test function 
	got := 清一色(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色三同顺(t *testing.T) {
	want := 24
	groupedHand := []TileGroup{
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{1, 2, 3}, false, false, false, true, false, 1},
		{[]int{2, 3, 4}, false, false, false, true, false, 1},
		{[]int{3, 3}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,1,2,3,1,2,3,2,3,4,3,3},
		won : true,
	}
	result := Output{}

	// test function 
	got := 一色三同顺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色三节高(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{5, 5, 5}, false, true, false, false, false, 1},
		{[]int{6, 6, 6}, false, true, false, false, false, 1},
		{[]int{7, 7, 7}, false, true, false, false, false, 1},
		{[]int{2, 3, 4}, false, false, false, true, false, 1},
		{[]int{3, 3}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{5,5,5,6,6,6,7,7,7,2,3,4,3,3},
		won : true,
	}
	result := Output{}

	// test function 
	got := 一色三节高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全大(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{7, 8, 9}, false, false, true, false, false, 1},
		{[]int{16, 17, 18}, false, false, true, false, false, 2},
		{[]int{25, 26, 27}, false, false, true, false, false, 3},
		{[]int{7, 8, 9}, false, false, true, false, false, 1},
		{[]int{18, 18}, false, false, false, false, true, 2},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{7,8,9,16,17,18,25,26,27,7,8,9,18,18},
		won: true,
	}
	result := Output{}

	// test function 
	got := 全大(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全中(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{22,23,24}, false, false, true, false, false, 3},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{14,14}, false, false, false, false, true, 2},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{4,5,6,13,14,15,22,23,24,4,5,6,14,14},
		won: true,
	}
	result := Output{}

	// test function 
	got := 全中(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全小(t *testing.T) {
	// const
	want := 24
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, true, false, false, 1},
		{[]int{10,11,12}, false, false, true, false, false, 2},
		{[]int{19,20,21}, false, false, true, false, false, 3},
		{[]int{1,2,3}, false, false, true, false, false, 1},
		{[]int{20,20}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,10,11,12,19,20,21,1,2,3,20,20},
		won: true,
	}
	result := Output{}

	// test function 
	got := 全小(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test清龙(t *testing.T) {
	// const
	want := 16
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{7,8,9}, false, false, true, false, false, 1},
		{[]int{1,2,3}, false, false, true, false, false, 1},
		{[]int{20,20}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,4,5,6,7,8,9,1,2,3,20,20},
		won: true,
	}
	result := Output{}

	// test function 
	got := 清龙(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三色双龙会(t *testing.T) {
	want := 16
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{7,8,9}, false, false, false, true, false, 1},
		{[]int{10,11,12}, false, false, false, true, false, 2},
		{[]int{16,17,18}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,7,8,9,10,11,12,16,17,18,23,23},
	}
	result := Output{}

	// test function 
	got := 三色双龙会(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一色三步高(t *testing.T) {
	want := 16
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{2,3,4}, false, false, false, true, false, 1},
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{10,10,10}, false, true, false, false, false, 2},
		{[]int{20,20}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,2,3,4,3,4,5,10,10,10,20,20},
		won : true,
	}
	result := Output{}

	// test function 
	got := 一色三步高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全带五(t *testing.T) {
	// const
	want := 16
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won: true,
	}
	result := Output{}

	// test function 
	got := 全带五(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三同刻(t *testing.T) {
	want := 16
	groupedHand := []TileGroup{
		{[]int{1,1,1}, false, true, false, false, false, 1},
		{[]int{10,10,10}, false, true, false, false, false, 2},
		{[]int{19,19,19}, false, true, false, false, false, 3},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,10,10,10,19,19,19,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三同刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三暗刻(t *testing.T) {
	// const
	want := 16
	groupedHand := []TileGroup{
		{[]int{3,3,3}, false, true, false, false, false, 1},
		{[]int{4,4,4}, false, true, false, false, false, 1},
		{[]int{10,10,10}, false, true, false, false, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,3,3,4,4,4,10,10,10,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三暗刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全不靠(t *testing.T) {
	// const
	want := 12
	groupedHand := []TileGroup{
		{[]int{1,4,7}, false, false, false, false, false, 1},
		{[]int{11,14,17}, false, false, false, false, false, 2},
		{[]int{21,24,27}, false, false, false, false, false, 3},
		{[]int{29,30,31}, false, false, false, false, false, 4},
		{[]int{28,32}, false, false, false, false, false, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,4,7,11,14,17,21,24,27,29,30,31,28,32},
	}
	result := Output{}

	// test function 
	got := 全不靠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test组合龙(t *testing.T) {
	// const
	want := 12
	groupedHand := []TileGroup{
		{[]int{1,4,7}, false, false, false, false, false, 1},
		{[]int{11,14,17}, false, false, false, false, false, 2},
		{[]int{21,24,27}, false, false, false, false, false, 3},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,4,7,11,14,17,21,24,27,14,15,16,23,23},
	}
	result := Output{}

	// test function 
	got := 组合龙(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test大于五(t *testing.T) {
	// const
	want := 12
	groupedHand := []TileGroup{
		{[]int{6,7,8}, false, false, false, true, false, 1},
		{[]int{9,9,9}, false, true, false, true, false, 1},
		{[]int{15,16,17}, false, false, false, true, false, 2},
		{[]int{16,17,18}, false, false, false, true, false, 2},
		{[]int{26,26}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{6,7,8,9,9,9,15,16,17,16,17,18,26,26},
		won: true,
	}
	result := Output{}

	// test function 
	got := 大于五(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test小于五(t *testing.T) {
	// const
	want := 12
	groupedHand := []TileGroup{
		{[]int{2,3,4}, false, false, false, true, false, 1},
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{10,10,10}, false, true, false, false, false, 2},
		{[]int{11,12,13}, false, false, false, true, false, 2},
		{[]int{22,22}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{2,3,4,1,2,3,10,10,10,11,12,13,22,22},
		won : true,
	}
	result := Output{}

	// test function 
	got := 小于五(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三风刻(t *testing.T) {
	// const
	want := 12
	groupedHand := []TileGroup{
		{[]int{28,28,28}, false, true, false, false, false, 4},
		{[]int{29,29,29}, false, true, false, false, false, 4},
		{[]int{30,30,30}, false, true, false, false, false, 4},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28,28,28,29,29,29,30,30,30,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三风刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test花龙(t *testing.T) {
	// const
	want := 8
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{25,26,27}, false, false, false, true, false, 3},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{30,30,30}, false, true, false, false, false, 4},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,25,26,27,13,14,15,30,30,30,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 花龙(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test推不倒(t *testing.T) {
	// const
	want := 8
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{4,4,4}, false, true, false, false, false, 1},
		{[]int{13,14, 15}, false, false, false, true, false, 2},
		{[]int{8,8,8}, false, true, false, false, false, 1},
		{[]int{34,34}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,4,4,4,13,14,15,8,8,8,34,34},
		won : true,
	}
	result := Output{}

	// test function 
	got := 推不倒(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三色三同顺(t *testing.T) {
	want := 8
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{10,11,12}, false, false, false, true, false, 2},
		{[]int{19,20,21}, false, false, false, true, false, 3},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,10,11,12,19,20,21,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三色三同顺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三色三节高(t *testing.T) {
	want := 8
	groupedHand := []TileGroup{
		{[]int{1,1,1}, false, true, false, false, false, 1},
		{[]int{11,11,11}, false, true, false, false, false, 2},
		{[]int{21,21,21}, false, true, false, false, false, 3},
		{[]int{2,3,4}, false, false, true, false, false, 1},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,11,11,11,21,21,21,2,3,4,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三色三节高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test无番和(t *testing.T) {
	// const
	want := 8
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won: true,
	}
	result := Output{}

	// test function 
	got := 无番和(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test碰碰和(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,3,3}, false, true, false, false, false, 1},
		{[]int{2,2,2}, false, true, false, false, false, 1},
		{[]int{4,4,4}, false, true, false, false, false, 1},
		{[]int{5,5,5}, false, true, false, false, false, 1},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,3,3,2,2,2,4,4,4,5,5,5,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 碰碰和(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test混一色(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, true, false, 1},
		{[]int{4,5,6}, false, false, true, true, false, 1},
		{[]int{1,2,3}, false, false, true, true, false, 1},
		{[]int{30,30,30}, false, false, true, false, false, 4},
		{[]int{33,33}, false, false, false, false, true, 5},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,1,2,3,30,30,30,33,33},
		won : true,
	}
	result := Output{}

	// test function 
	got := 混一色(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test三色三步高(t *testing.T) {
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{23,24,25}, false, false, false, true, false, 3},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,13,14,15,23,24,25,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 三色三步高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %d is not equal to expected output %d", got.score, want)
	}
}

func Test五门齐(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{19,20,21}, false, false, true, false, false, 3},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{30,30,30}, false, false, true, false, false, 4},
		{[]int{32,32}, false, false, false, false, true, 5},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,19,20,21,13,14,15,30,30,30,32,32},
		won : true,
	}
	result := Output{}

	// test function 
	got := 五门齐(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全求人(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,4,5}, true, false, false, true, false, 1},
		{[]int{4,5,6}, true, false, false, true, false, 1},
		{[]int{13,14,15}, true, false, false, true, false, 2},
		{[]int{14,15,16}, true, false, false, true, false, 2},
		{[]int{23,23}, true, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 全求人(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test双暗杠(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{3,3,3}, false, false, true, false, false, 1},
		{[]int{4,4,4}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,3,3,4,4,4,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 双暗杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test双箭刻(t *testing.T) {
	// const
	want := 6
	groupedHand := []TileGroup{
		{[]int{32,32,32}, false, true, false, false, false, 4},
		{[]int{33,33,33}, false, true, false, false, false, 4},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{32,32,32,33,33,33,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 双箭刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test全带幺(t *testing.T) {
	// const
	want := 4
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{7,8,9}, false, false, false, true, false, 1},
		{[]int{10,11,12}, false, false, false, true, false, 2},
		{[]int{30,30,30}, false, true, false, false, false, 4},
		{[]int{31,31}, false, false, false, false, true, 4},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 全带幺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test不求人(t *testing.T) {
	// const
	want := 4
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 不求人(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test双明杠(t *testing.T) {
	// const
	want := 4
	groupedHand := []TileGroup{
		{[]int{3,3,3}, true, false, true, false, false, 1},
		{[]int{4,4,4}, true, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,3,3,4,4,4,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 双明杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test箭刻(t *testing.T) {
	want := 2
	groupedHand := []TileGroup{
		{[]int{33,33,33}, false, true, false, false, false, 4},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{33,33,33,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 箭刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test圈风刻(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{28,28,28}, false, true, false, false, false, 4},
		{[]int{30,30,30}, false, true, false, false, false, 4},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28,28,28,30,30,30,13,14,15,14,15,16,23,23},
		round: 30,
		won : true,
	}
	result := Output{}

	// test function 
	got := 圈风刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test门风刻(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{28,28,28}, false, true, false, false, false, 4},
		{[]int{30,30,30}, false, true, false, false, false, 4},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{28,28,28,30,30,30,13,14,15,14,15,16,23,23},
		turn: 28,
		won : true,
	}
	result := Output{}

	// test function 
	got := 门风刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test门前清(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, true, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 门前清(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test平和(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 平和(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test四归一(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{4,4}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,4,4},
		won : true,
	}
	result := Output{}

	// test function 
	got := 四归一(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test双同刻(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{1,1,1}, false, true, false, false, false, 1},
		{[]int{10,10,10}, false, true, false, false, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,10,10,10,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 双同刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test双暗刻(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{1,1,1}, false, true, false, false, false, 1},
		{[]int{2,2,2}, false, true, false, false, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,2,2,2,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 双暗刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test暗杠(t *testing.T) {
	// const
	want := 2
	groupedHand := []TileGroup{
		{[]int{3,3,3}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,3,3,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 暗杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test断幺(t *testing.T) {
	want := 2
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{2,2}, false, false, false, false, true, 1},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,2,2},
		won : true,
	}
	result := Output{}

	// test function 
	got := 断幺(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test一般高(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,3,4,5,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 一般高(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test喜相逢(t *testing.T) {
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{12,13,14}, false, false, false, true, false, 2},
		{[]int{1,1,1}, false, false, true, false, false, 1},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,12,13,14,1,1,1,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 喜相逢(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %d is not equal to expected output %d", got.score, want)
	}
}

func Test连六(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{6,7,8}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,6,7,8,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 连六(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test老少副(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{1,2,3}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,2,3,1,2,3,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 老少副(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got, want)
	}
}

func Test幺九刻(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{1,1,1}, false, true, false, false, false, 1},
		{[]int{9,9,9}, false, true, false, true, false, 1},
		{[]int{7,8,9}, false, false, false, true, false, 4},
		{[]int{12,13,14}, false, false, false, true, false, 4},
		{[]int{10,10}, false, false, false, false, true, 2},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{1,1,1,9,9,9,7,8,9,12,13,14,10,10},
		won : true,
	}
	result := Output{}

	// test function 
	got := 幺九刻(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test明杠(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{1,1,1}, true, false, true, false, false, 1},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,1,1,1,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 明杠(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test缺一门(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{1,1}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,1,1},
		won : true,
	}
	result := Output{}

	// test function 
	got := 缺一门(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test无字(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, true, false, false, 1},
		{[]int{4,5,6}, false, false, true, false, false, 1},
		{[]int{13,14,15}, false, false, true, false, false, 2},
		{[]int{14,15,16}, false, false, true, false, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		won : true,
	}
	result := Output{}

	// test function 
	got := 无字(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got, want)
	}
}

func Test边张(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		winnngTile: 2,
		won : true,
	}
	result := Output{}

	// test function 
	got := 边张(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %d is not equal to expected output %d", got.score, want)
	}
}

func Test坎张(t *testing.T) {
	// const
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		winnngTile: 7,
		won : true,
	}
	result := Output{}

	// test function 
	got := 坎张(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}

func Test单钓(t *testing.T) {
	want := 1
	groupedHand := []TileGroup{
		{[]int{3,4,5}, false, false, false, true, false, 1},
		{[]int{4,5,6}, false, false, false, true, false, 1},
		{[]int{13,14,15}, false, false, false, true, false, 2},
		{[]int{14,15,16}, false, false, false, true, false, 2},
		{[]int{23,23}, false, false, false, false, true, 3},
	}
	hands := Hands{
		grouped: groupedHand,
		ungrouped: []int{3,4,5,4,5,6,13,14,15,14,15,16,23,23},
		winnngTile: 13,
		won : true,
	}
	result := Output{}

	// test function 
	got := 单钓(hands, result) 

	// check results
	if got.score != want {
		t.Errorf("Test output %v is not equal to expected output %v", got.score, want)
	}
}