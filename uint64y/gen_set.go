package uint64y

// Code generated. DO NOT EDIT.

type Set map[uint64]struct{}

// Contains returns true if Set contains v.
func (s Set) Contains(v uint64) bool {
	_, exists := s[v]
	return exists
}

// Intersect returns a Set of the intersection between two Sets.
func (s Set) Intersect(s2 Set) Set {
	short, long := s, s2
	if len(s2) < len(s) {
		short, long = s2, s
	}

	m := make(map[uint64]struct{}, len(short))
	for k := range short {
		if _, exists := long[k]; exists {
			m[k] = struct{}{}
		}
	}
	return m
}

// ToSet returns a lookup map for uint64.
func ToSet(vs []uint64) Set {
	m := make(map[uint64]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}

	return m
}
