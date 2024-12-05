package d05printqueue

import (
	"advent/solutions/utils"
	"sort"
	"strconv"
	"strings"
)

type Page struct {
	number   string
	isAfter  map[string]*Page
	isBefore map[string]*Page
}

func (p *Page) AddPageAfter(q *Page) {
	p.isBefore[q.number] = q
}

func (p *Page) AddPageBefore(q *Page) {
	p.isAfter[q.number] = q
}

func newPage(number string) *Page {
	return &Page{
		number:   number,
		isAfter:  map[string]*Page{},
		isBefore: map[string]*Page{},
	}
}

func CheckPageOrdering(input []string) (int, int) {
	cutoff := 0
	rules := []string{}
	for i, l := range input {
		if l == "" {
			cutoff = i
			break
		}
		rules = append(rules, l)
	}
	pages := constructRules(rules)
	correct, incorrect := 0, 0
	for _, l := range input[cutoff+1:] {
		middle, ok := checkOrdering(l, pages)
		if ok {
			correct += utils.Stoi(middle)
			continue
		}
		incorrect += utils.Stoi(correctOrdering(l, pages))
	}
	return correct, incorrect
}

func constructRules(rules []string) *map[string]*Page {
	pages := map[string]*Page{}
	for _, r := range rules {
		components := strings.Split(r, "|")
		f, l := strings.TrimSpace(components[0]), strings.TrimSpace(components[1])
		first, ok := pages[f]
		if !ok {
			pages[f] = newPage(f)
			first = pages[f]
		}
		last, ok := pages[l]
		if !ok {
			pages[l] = newPage(l)
			last = pages[l]
		}
		first.AddPageAfter(last)
		last.AddPageBefore(first)
	}
	return &pages
}

func checkOrdering(line string, pages *map[string]*Page) (string, bool) {
	entries := strings.Split(line, ",")
	prior := []*Page{}
	for _, e := range entries {
		page, ok := (*pages)[e]
		if !ok {
			continue
		}
		for _, p := range prior {
			_, ok := page.isAfter[p.number]
			if !ok {
				return "", false
			}
		}
		prior = append(prior, page)
	}
	return entries[len(entries)/2], true
}

func correctOrdering(line string, pages *map[string]*Page) string {
	entries := strings.Split(line, ",")
	sort.Slice(entries, func(i, j int) bool {
		p, ok := (*pages)[entries[i]]
		if !ok {
			return true
		}
		q, ok := (*pages)[entries[j]]
		if !ok {
			return false
		}
		_, ok = p.isBefore[q.number]
		return ok
	})
	return entries[len(entries)/2]

}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	resA, resB := CheckPageOrdering(lines)
	return strconv.Itoa(resA), strconv.Itoa(resB)
}
