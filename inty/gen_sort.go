package inty

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []int) []int {
	sort.Sort(IntSlice(vs))
	return vs
}
