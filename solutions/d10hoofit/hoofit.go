package d10hoofit

import (
	"advent/solutions/utils"
	"fmt"
	"maps"
	"strconv"
)

const runeOffset = 48

type HeightNode struct {
	xy       []int
	xMax     int
	yMax     int
	height   int
	terminus map[string]bool
	distinct int
	explored bool
}

func newNode(x, y, xMax, yMax, height int) *HeightNode {
	return &HeightNode{
		xy:       []int{x, y},
		xMax:     xMax,
		yMax:     yMax,
		height:   height,
		terminus: map[string]bool{},
		distinct: 0,
		explored: false,
	}
}

func (n *HeightNode) Explore(nodes [][]*HeightNode) (map[string]bool, int) {
	if n.explored {
		return n.terminus, n.distinct
	}
	if n.height == 9 {
		return map[string]bool{fmt.Sprintf("%d-%d", n.xy[0], n.xy[1]): true}, 1
	}
	coords := utils.NESW2DNeighbours(n.xy[0], n.xy[1], n.xMax, n.yMax)
	for _, xy := range coords {
		neighbour := nodes[xy[1]][xy[0]]
		if neighbour.height != n.height+1 {
			continue
		}
		terminal, count := neighbour.Explore(nodes)
		maps.Copy(n.terminus, terminal)
		n.distinct += count
	}
	n.explored = true
	return n.terminus, n.distinct
}

func FindTrailHeads(lines []string) (int, int) {
	xMax, yMax := len(lines[0])-1, len(lines)-1
	hill := make([][]*HeightNode, len(lines))
	zeroes := [][]int{}
	for y, l := range lines {
		height := make([]*HeightNode, len(l))
		for x, r := range l {
			val := int(r - runeOffset)
			if val == 0 {
				zeroes = append(zeroes, []int{x, y})

			}
			height[x] = newNode(x, y, xMax, yMax, val)
		}
		hill[y] = height
	}
	total, c := 0, 0
	for _, xy := range zeroes {
		root := hill[xy[1]][xy[0]]
		terminus, count := root.Explore(hill)
		total += len(terminus)
		c += count
	}
	return total, c
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := FindTrailHeads(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
