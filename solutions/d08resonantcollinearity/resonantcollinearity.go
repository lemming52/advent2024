package d08resonantcollinearity

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
)

type Antenna struct {
	x, y      int
	frequency rune
}

func FindAntiNodes(lines []string) (int, int) {
	antennae := map[rune][]*Antenna{}
	for y, l := range lines {
		for x, r := range l {
			if r != '.' {
				a := &Antenna{
					x:         x,
					y:         y,
					frequency: r,
				}
				as, ok := antennae[r]
				if !ok {
					antennae[r] = []*Antenna{a}
				} else {
					antennae[r] = append(as, a)
				}
			}
		}
	}

	xMax, yMax := len(lines[0])-1, len(lines)-1
	antinodes := map[string]bool{}
	resonantAntiNodes := map[string]bool{}
	for _, frequency := range antennae {
		for i, a := range frequency[:len(frequency)-1] {
			for _, b := range frequency[i+1:] {
				coords := findAntiNodes(a, b)
				for _, xy := range coords {
					if !utils.IsValidGridCoord(xy[0], xy[1], xMax, yMax) {
						continue
					}
					antinodes[fmt.Sprintf("%d-%d", xy[0], xy[1])] = true
				}
				coords = findResonantAntiNodes(a, b, xMax, yMax)
				for _, xy := range coords {
					resonantAntiNodes[fmt.Sprintf("%d-%d", xy[0], xy[1])] = true
				}
			}
		}
	}

	return len(antinodes), len(resonantAntiNodes)
}

func findAntiNodes(a, b *Antenna) [][]int {
	dx, dy := b.x-a.x, b.y-a.y
	return [][]int{
		{b.x + dx, b.y + dy},
		{a.x - dx, a.y - dy},
	}
}

func findResonantAntiNodes(a, b *Antenna, xMax, yMax int) [][]int {
	dx, dy := b.x-a.x, b.y-a.y
	antiNodes := [][]int{
		{a.x, a.y},
		{b.x, b.y},
	}
	nx, ny := b.x+dx, b.y+dy
	for utils.IsValidGridCoord(nx, ny, xMax, yMax) {
		antiNodes = append(antiNodes, []int{nx, ny})
		nx, ny = nx+dx, ny+dy
	}
	nx, ny = a.x-dx, a.y-dy
	for utils.IsValidGridCoord(nx, ny, xMax, yMax) {
		antiNodes = append(antiNodes, []int{nx, ny})
		nx, ny = nx-dx, ny-dy
	}
	return antiNodes
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	resA, resB := FindAntiNodes(lines)
	return strconv.Itoa(resA), strconv.Itoa(resB)
}
