package int32y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []int32) []int32 {
	sort.Sort(Int32Slice(vs))
	return vs
}
