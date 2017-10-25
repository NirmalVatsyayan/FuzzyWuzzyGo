package fuzzymatch

import (
	"strings"
	"unicode/utf8"
)

//Ratio allow you to calculate the percentage of variance between two strings
//if the two strings are equals the function returns 1.
func Ratio(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	len := utf8.RuneCountInString(s1) + utf8.RuneCountInString(s2)
	dist := LevenshteinDistance(processString(s1), processString(s2))
	return int((1 - (float32(dist) / float32(len))) * 100)
}

//PartialRatio allow you to calculate the "best partial" ratio. It takes
//the smaller string and we compare the smaller with a partial string from
//the bigger one. Could be useful if you have to compare two strings with
//very different length
func PartialRatio(s1, s2 string) int {
	min, max := minMax(s1, s2)
	var bestRatio int
	for i := 0; i < len(max)-len(min)+1; i++ {
		Ratio := Ratio(min, max[i:i+len(min)])
		if Ratio > bestRatio {
			bestRatio = Ratio
		}
	}
	return bestRatio
}

//TokenSortRatio allow you to compare two strings "ordered" alphabetically
//so if you have two strings not ordered. This function could be useful.
func TokenSortRatio(s1, s2 string) int {
	s1Sort, s2Sort := sortString(strings.Split(s1, " ")), sortString(strings.Split(s2, " "))
	return Ratio(s1Sort, s2Sort)
}

//TokenSetRatio splits the strings in two groups : intersection and remainder
//and then we compare the group with each other.
func TokenSetRatio(s1, s2 string) int {
	t0, t1, t2 := sortListTokenSetRatio(strings.Split(s1, " "), strings.Split(s2, " "))
	t1, t2 = sumSlices(t0, t1), sumSlices(t0, t2)
	t0str, t1str := strings.Join(t0, " "), strings.Join(t1, " ")
	t2str := strings.Join(t2, " ")
	return minimum(Ratio(t0str, t1str), Ratio(t0str, t2str), Ratio(t1str, t2str))
}
