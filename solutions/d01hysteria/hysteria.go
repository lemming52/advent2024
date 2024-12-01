package d01hysteria

import (
	"advent/solutions/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func CompareIds(lines []string) (int, int) {
	setA, setB := make([]int, len(lines)), make([]int, len(lines))
	counts := map[int]int{}
	for _, s := range lines {
		a, b := parseLine(s)
		ia, _ := slices.BinarySearch(setA, a)
		setA = slices.Insert(setA, ia, a)
		ib, _ := slices.BinarySearch(setB, b)
		setB = slices.Insert(setB, ib, b)
		_, ok := counts[b]
		if !ok {
			counts[b] = 1
		} else {
			counts[b] += 1
		}
	}
	distance, similarity := 0, 0
	for i, a := range setA {
		distance += int(math.Abs(float64(setB[i] - a)))
		count, ok := counts[a]
		if !ok {
			count = 0
		}
		similarity += a * count
	}
	return distance, similarity
}

func parseLine(s string) (int, int) {
	components := strings.Split(s, "   ")
	a := utils.Stoi(components[0])
	b := utils.Stoi(components[1])
	return a, b
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := CompareIds(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
