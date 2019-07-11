package uint16y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Uint16Slice []uint16

func (p Uint16Slice) Len() int           { return len(p) }
func (p Uint16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []uint16) []uint16 {
	sort.Sort(Uint16Slice(vs))
	return vs
}
