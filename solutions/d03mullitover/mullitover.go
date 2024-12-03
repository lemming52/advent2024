package d03mullitover

import (
	"advent/solutions/utils"
	"regexp"
	"strconv"
	"strings"
)

const instructions = `(?:mul\((\d{1,3})\,(\d{1,3})\))|(?:do\(\))|(?:don't\(\))`

func SumMultiply(line string) (int, int) {
	total, controlledTotal := 0, 0
	enabled := true
	pattern := regexp.MustCompile(instructions)
	matches := pattern.FindAllStringSubmatch(line, -1)
	for _, m := range matches {
		if strings.HasPrefix(m[0], "mul") {
			val := utils.Stoi(m[1]) * utils.Stoi(m[2])
			total += val
			if enabled {
				controlledTotal += val
			}
		} else {
			enabled = !strings.HasPrefix(m[0], "don't")
		}
	}
	return total, controlledTotal
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	line := strings.Join(lines, "")
	resA, resB := SumMultiply(line)
	return strconv.Itoa(resA), strconv.Itoa(resB)
}
