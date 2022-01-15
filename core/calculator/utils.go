package calculator

import (
	"math/rand"
)

func checkPairs(hands []int) bool {
	return len(removeDuplicateInt(hands)) == 7
}

func removeIntFromSlice(slice []int, hand []int) []int {
	output := []int{}
	for _, h := range hand {
		for , s := range slice {
			if h != s {
				output = append(output,s)
				break
			}
		}
	}
	return output
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func extractData(hands []TileGroup, mode string) []int{
	handTiles := int[]{}
	for _, g := range hands {
		if mode == "pong"{
			if g.pong == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "kong"{
			if g.kong == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "pair"{
			if g.pair == 1 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}else if mode == "close"{
			if g.open == 0 {
				handTiles = append(handTiles, g.tiles[0]...)
			}
		}
	}
	sort.Ints(handTiles)
	return handTiles
}

func containedInt(slice []int, tile int) bool {
	for _, i := range slice {
		if tile == i {
			return true
		}
	}
	return false
}
