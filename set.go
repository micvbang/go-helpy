package helpy

// Code generated. DO NOT EDIT.

type Set[T comparable] map[T]struct{}

// Contains returns true if Set contains v.
func (s Set[T]) Contains(v T) bool {
	_, exists := s[v]
	return exists
}

// Intersect returns the intersection of two Sets.
func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	short, long := s, s2
	if len(s2) < len(s) {
		short, long = s2, s
	}

	m := make(map[T]struct{}, len(short))
	for k := range short {
		if _, exists := long[k]; exists {
			m[k] = struct{}{}
		}
	}
	return m
}

// Union returns the union of two Sets.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	m := make(map[T]struct{}, len(s) + len(s2))
	for k := range s {
		m[k] = struct{}{}
	}

	for k := range s2 {
		m[k] = struct{}{}
	}

	return m
}


// Subtract returns a new Set containing all values from s that are not in
// s2.
func (s Set[T]) Subtract(s2 Set[T]) Set[T] {
	output := Set[T]{}
	for v1 := range s {
		if _, ok := s2[v1]; !ok{
			output[v1] = struct{}{}
		}
	}

	return output
}


// Add adds v to its elements
func (s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

// Equal returns true if s and s2 are of equal length and all elements 
// contained in one set are contained in the other.
func (s Set[T]) Equal(s2 Set[T]) bool {
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

// ToSet returns a lookup map for T.
func ToSet[T comparable](vs []T) Set[T] {
	m := make(map[T]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}

	return m
}

// MakeSet returns a lookup map for T
func MakeSet[T comparable](vs ...T) Set[T] {
	return ToSet(vs)
}