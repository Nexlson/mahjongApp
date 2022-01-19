package main

import (
	"fmt"
)

func checkWon(hand Hands) Hands {
	var count int 
	for _, group := range hands.grouped {
		if group.pong == 1 or group.chi == 1{
			count += 1
		}
	}
	if count == 4 {
		hands.won = true
		return hands
	}else {
		hands.won = false
		return hands
	}
}

// func checkSpecialPattern(hands Hands, outputs Output) Output{
// 	specialFuns := []func(Hands, Output)Output{九莲宝灯, 连七对, 十三幺, 七对, 七星不靠, 全不靠, 组合龙}
// 	for _, v := range specialFuns {
// 		result := v(hands, outputs)
// 		if result.score != 0 {
// 			break
// 		}
// 	}
// 	result.exceptions := append(result.exceptions, int[]{4, 6, 7, 19, 20, 34, 35}...)
// 	return result
// }

// func checkFixedPattern() {
	
// }

// func checkRepeatedPattern() {
	 
// }

func calculateScore(hands Hands) Output{
	finalResult := Output{}

	// add initial score
	finalResult.score += hands.score

	// check won
	hands := checkWon(hands)

	// loop all rules 
	for k, v := range RulesMaps{
		if !containedInt(finalResult.exceptions, k){
			finalResult := v.((func(Hands, Output) Output) (hands, finalResult))
		}
	}

	return finalResult
}

func main() {
	fmt.Println(checkSamePattern([]int{10, 11}))
}