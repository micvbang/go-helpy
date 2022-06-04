package uint16y

import "github.com/micvbang/go-helpy"

// Code generated. DO NOT EDIT.

// NOTE: this type is deprecated. Use helpy.Set instead
type Set helpy.Set[uint16]

func (s Set) t() helpy.Set[uint16] {
	return helpy.Set[uint16](s)
}

func (s Set) Intersect(s2 Set) Set {
	return Set(s.t().Intersect(s2.t()))
}

func (s Set) Union(s2 Set) Set {
	return Set(s.t().Union(s2.t()))
}

func (s Set) Contains(v uint16) bool {
	return s.t().Contains(v)
}

func (s Set) Equal(s2 Set) bool {
	return s.t().Equal(s2.t())
}

// ToSet returns a lookup map for uint16.
// NOTE: this function is deprecated. Use helpy.ToSet instead.
func ToSet(vs []uint16) Set {
	return Set(helpy.ToSet(vs))
}

// MakeSet returns a lookup map for uint16
// NOTE: this function is deprecated. Use helpy.MakeSet instead.
func MakeSet(vs ...uint16) Set {
	return Set(helpy.ToSet(vs))
}
