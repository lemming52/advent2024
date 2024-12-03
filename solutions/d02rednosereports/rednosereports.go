package d02rednosereports

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

const deltaThreshold = 3

func CheckReports(lines []string) (int, int) {
	safe, damped := 0, 0
	for _, l := range lines {
		vals := parseLine(l)
		if checkReport(vals) {
			safe += 1
			damped += 1
			continue
		}
		if checkReportDamped(vals) {
			damped += 1
		}
	}
	return safe, damped
}

func parseLine(s string) []int {
	components := strings.Split(s, " ")
	vals := make([]int, len(components))
	for i, c := range components {
		vals[i] = utils.Stoi(c)
	}
	return vals
}

func checkReport(vals []int) bool {
	direction := 1
	delta := vals[1] - vals[0]
	if delta < 0 {
		direction = -1
	}
	if utils.Abs(delta) > deltaThreshold || delta == 0 {
		return false
	}

	for i := range vals[1:] {
		delta = vals[i+1] - vals[i]
		if utils.Abs(delta) > deltaThreshold || delta == 0 {
			return false
		}
		if delta*direction < 0 {
			return false
		}
	}
	return true
}

func checkReportDamped(vals []int) bool {
	for i := range vals {
		newArr := []int{}
		newArr = append(newArr, vals[:i]...)
		newArr = append(newArr, vals[i+1:]...)
		safe := checkReport(newArr)
		if safe {
			return true
		}
	}
	return false
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := CheckReports(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
