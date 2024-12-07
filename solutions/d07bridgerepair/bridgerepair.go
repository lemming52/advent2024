package d07bridgerepair

import (
	"advent/solutions/utils"
	"math"
	"strconv"
	"strings"
)

func Operate(lines []string) (int, int) {
	total, concat := 0, 0
	for _, l := range lines {
		components := strings.Split(l, ":")
		target := utils.Stoi(components[0])
		values := strings.Split(strings.TrimSpace(components[1]), " ")
		current := utils.Stoi(values[0])
		if operate(current, target, values[1:], false) {
			total += target
			concat += target
			continue
		}
		if operate(current, target, values[1:], true) {
			concat += target
		}
	}
	return total, concat
}

func operate(current, target int, values []string, useConcat bool) bool {
	if len(values) == 0 {
		return current == target
	}
	val := utils.Stoi(values[0])
	newVal := current + val
	if operate(newVal, target, values[1:], useConcat) {
		return true
	}
	newVal = current * val
	if operate(newVal, target, values[1:], useConcat) {
		return true
	}
	if !useConcat {
		return false
	}
	newVal = current*int(math.Pow(10.0, float64(len(values[0])))) + val
	return operate(newVal, target, values[1:], useConcat)
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	resA, resB := Operate(lines)
	return strconv.Itoa(resA), strconv.Itoa(resB)
}
