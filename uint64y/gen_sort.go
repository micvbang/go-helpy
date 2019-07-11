package uint64y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []uint64) []uint64 {
	sort.Sort(Uint64Slice(vs))
	return vs
}
