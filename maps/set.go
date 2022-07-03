package maps

import "litu/cmp"

type Set[K comparable] map[K]struct{}

func (s Set[K]) Contains(k K) bool {
	_, ok := s[k]
	return ok
}

func (s Set[K]) Set(k K) {
	s[k] = struct{}{}
}

func (s Set[K]) TrySet(k K) bool {
	if s.Contains(k) {
		return true
	} else {
		s.Set(k)
		return false
	}
}

func (s Set[K]) Union(other Set[K]) Set[K] {
	r := make(Set[K], len(s)+len(other))
	for k := range s {
		r.Set(k)
	}
	for k := range other {
		r.Set(k)
	}
	return r
}

func (s Set[K]) Intersect(other Set[K]) Set[K] {
	r := make(Set[K], cmp.Min(len(s), len(other)))
	for k := range s {
		if other.Contains(k) {
			r.Set(k)
		}
	}
	return r
}

func (s Set[K]) Diff(other Set[K]) Set[K] {
	r := make(Set[K], len(s))
	for k := range s {
		if !other.Contains(k) {
			r.Set(k)
		}
	}
	return r
}

func (s Set[K]) IsDisjoint(other Set[K]) bool {
	for k := range s {
		if other.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[K]) IsSubset(other Set[K]) bool {
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}
