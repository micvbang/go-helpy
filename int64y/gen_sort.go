package int64y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []int64) []int64 {
	sort.Sort(Int64Slice(vs))
	return vs
}
