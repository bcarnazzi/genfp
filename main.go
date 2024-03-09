// package genfp provides primitives for functional programming in go using generics.
package genfp

import (
	"cmp"
	"sort"
)

type mapFunc[E, F any] func(E) F
type filterFunc[E any] func(E) bool
type reduceFunc[E any] func(E, E) E

// Contains returns true if slice s contains v, false otherwise.
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}

// Reverse returns a reversed slice form slice s.
func Reverse[E any](s []E) []E {
	result := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

// AscendingSort returns an ascending sorted slice from slice s.
func AscendingSort[E cmp.Ordered](s []E) []E {
	result := make([]E, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// DescendingSort returns an ascending sorted slice from slice s.
func DescendingSort[E cmp.Ordered](s []E) []E {
	result := make([]E, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}

// Map returns a slice containing all s elements, with mapFunc f applied.
func Map[E, F any](s []E, f mapFunc[E, F]) []F {
	result := make([]F, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

// Filter returns a slice containing all s elements that matches filterFunc f.
func Filter[E any](s []E, f filterFunc[E]) []E {
	result := []E{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce return a single element representing the reduction of s with function f.
func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}
