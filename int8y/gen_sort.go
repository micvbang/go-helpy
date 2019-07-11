package int8y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Int8Slice []int8

func (p Int8Slice) Len() int           { return len(p) }
func (p Int8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []int8) []int8 {
	sort.Sort(Int8Slice(vs))
	return vs
}
