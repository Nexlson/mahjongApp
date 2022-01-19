package calculator

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