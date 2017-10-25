package Helpers

import (
	fuzz "github.com/NirmalVatsyayan/FuzzyWuzzyGo/FuzzyMatch"
	"math"
	//"log"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}


func CalculateTokenRatio(data1, data2 string) (int, int, int){

	token_set_ratio := fuzz.TokenSetRatio(data1, data2)
	token_sort_ratio := fuzz.TokenSortRatio(data1, data2)
	partial_ratio := fuzz.PartialRatio(data1, data2)

        return token_set_ratio, token_sort_ratio, partial_ratio
}
