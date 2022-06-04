package uinty

import "github.com/micvbang/go-helpy"

// Code generated. DO NOT EDIT.

// NOTE: this type is deprecated. Use helpy.Set instead
type Set helpy.Set[uint]

func (s Set) t() helpy.Set[uint] {
	return helpy.Set[uint](s)
}

func (s Set) Intersect(s2 Set) Set {
	return Set(s.t().Intersect(s2.t()))
}

func (s Set) Union(s2 Set) Set {
	return Set(s.t().Union(s2.t()))
}

func (s Set) Contains(v uint) bool {
	return s.t().Contains(v)
}

func (s Set) Equal(s2 Set) bool {
	return s.t().Equal(s2.t())
}

// ToSet returns a lookup map for uint.
// NOTE: this function is deprecated. Use helpy.ToSet instead.
func ToSet(vs []uint) Set {
	return Set(helpy.ToSet(vs))
}

// MakeSet returns a lookup map for uint
// NOTE: this function is deprecated. Use helpy.MakeSet instead.
func MakeSet(vs ...uint) Set {
	return Set(helpy.ToSet(vs))
}
