package int16y

// Code generated. DO NOT EDIT.

type Set map[int16]struct{}

// Contains returns true if Set contains v.
func (s Set) Contains(v int16) bool {
	_, exists := s[v]
	return exists
}

// Intersect returns the intersection of two Sets.
func (s Set) Intersect(s2 Set) Set {
	short, long := s, s2
	if len(s2) < len(s) {
		short, long = s2, s
	}

	m := make(map[int16]struct{}, len(short))
	for k := range short {
		if _, exists := long[k]; exists {
			m[k] = struct{}{}
		}
	}
	return m
}

// Union returns the union of two Sets.
func (s Set) Union(s2 Set) Set {
	m := make(map[int16]struct{}, len(s)+len(s2))
	for k := range s {
		m[k] = struct{}{}
	}

	for k := range s2 {
		m[k] = struct{}{}
	}

	return m
}

// Equal checks whether all
func (s Set) Equal(s2 Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for v := range s {
		if _, ok := s2[v]; !ok {
			return false
		}
	}

	return true
}

// ToSet returns a lookup map for int16.
func ToSet(vs []int16) Set {
	m := make(map[int16]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}

	return m
}
