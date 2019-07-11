package uint8y

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type Uint8Slice []uint8

func (p Uint8Slice) Len() int           { return len(p) }
func (p Uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []uint8) []uint8 {
	sort.Sort(Uint8Slice(vs))
	return vs
}
