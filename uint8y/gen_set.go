package uint8y

import "github.com/micvbang/go-helpy"

// Code generated. DO NOT EDIT.

// NOTE: this type is deprecated. Use helpy.Set instead
type Set helpy.Set[uint8]

func (s Set) t() helpy.Set[uint8] {
	return helpy.Set[uint8](s)
}

func (s Set) Intersect(s2 Set) Set {
	return Set(s.t().Intersect(s2.t()))
}

func (s Set) Union(s2 Set) Set {
	return Set(s.t().Union(s2.t()))
}

func (s Set) Contains(v uint8) bool {
	return s.t().Contains(v)
}

func (s Set) Equal(s2 Set) bool {
	return s.t().Equal(s2.t())
}

// ToSet returns a lookup map for uint8.
// NOTE: this function is deprecated. Use helpy.ToSet instead.
func ToSet(vs []uint8) Set {
	return Set(helpy.ToSet(vs))
}

// MakeSet returns a lookup map for uint8
// NOTE: this function is deprecated. Use helpy.MakeSet instead.
func MakeSet(vs ...uint8) Set {
	return Set(helpy.ToSet(vs))
}
