package uinty

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []uint) []uint {
	sort.Sort(UintSlice(vs))
	return vs
}
