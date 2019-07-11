package int16y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Int16Slice []int16

func (p Int16Slice) Len() int           { return len(p) }
func (p Int16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []int16) []int16 {
	sort.Sort(Int16Slice(vs))
	return vs
}
