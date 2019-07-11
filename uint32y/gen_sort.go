package uint32y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []uint32) []uint32 {
	sort.Sort(Uint32Slice(vs))
	return vs
}
