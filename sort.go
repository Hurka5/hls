package main

import (
	"sort"
	"strings"
)

//TODO: More sorting options

type ByName []Item

func (fi ByName) Len() int      { return len(fi) }
func (fi ByName) Swap(i, j int) { fi[i], fi[j] = fi[j], fi[i] }
func (fi ByName) Less(i, j int) bool {
	return strings.ToLower(fi[i].info.Name()) < strings.ToLower(fi[j].info.Name())
}

func reverseItems(items []Item) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func sortItems(items []Item) []Item {
	sort.Sort(ByName(items))

	if args.Reverse {
		reverseItems(items)
	}

	return items
}
