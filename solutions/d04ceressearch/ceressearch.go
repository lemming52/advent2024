package d04ceressearch

import (
	"advent/solutions/utils"
	"strconv"
)

const xmas = "XMAS"

func WordSearch(text []string) (int, int) {
	xMax, yMax := len(text[0])-1, len(text)-1

	total, crosses := 0, 0
	for y, line := range text {
		for x, r := range line {
			if r == 'X' {
				neighbours := utils.All2DNeighbours(x, y, xMax, yMax)
				for _, n := range neighbours {
					if checkNeighbours(x, y, xMax, yMax, n, text) {
						total += 1
					}
				}
				continue
			}
			if r == 'A' {
				if checkCross(x, y, xMax, yMax, text) {
					crosses += 1
				}
			}
		}
	}
	return total, crosses
}

func checkNeighbours(x, y, xMax, yMax int, p []int, text []string) bool {
	if text[p[1]][p[0]] != 'M' {
		return false
	}
	position := []int{p[0], p[1]}
	direction := []int{p[0] - x, p[1] - y}
	for _, r := range xmas[2:] {
		position[0], position[1] = position[0]+direction[0], position[1]+direction[1]
		if !isValidCoord(position[0], position[1], xMax, yMax) {
			return false
		}
		if rune(text[position[1]][position[0]]) != r {
			return false
		}
	}
	return true
}

func checkCross(x, y, xMax, yMax int, text []string) bool {
	neighbours := getCorner2DNeighbours(x, y, xMax, yMax)
	if len(neighbours) != 4 {
		return false
	}
	rTL := rune(text[neighbours[0][1]][neighbours[0][0]])
	rTR := rune(text[neighbours[1][1]][neighbours[1][0]])
	rBL := rune(text[neighbours[2][1]][neighbours[2][0]])
	rBR := rune(text[neighbours[3][1]][neighbours[3][0]])
	return isValidPair(rTL, rBR) && isValidPair(rTR, rBL)
}

func getCorner2DNeighbours(x, y, xMax, yMax int) [][]int {
	coords := [][]int{
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}
	correct := [][]int{}
	for _, xy := range coords {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] <= xMax && xy[1] <= yMax {
			correct = append(correct, xy)
		}
	}
	return correct
}

func isValidPair(r1, r2 rune) bool {
	if r1 == 'M' {
		return r2 == 'S'
	}
	if r1 == 'S' {
		return r2 == 'M'
	}
	return false
}

func isValidCoord(x, y, xMax, yMax int) bool {
	return x >= 0 && y >= 0 && x <= xMax && y <= yMax
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	resA, resB := WordSearch(lines)
	return strconv.Itoa(resA), strconv.Itoa(resB)
}
