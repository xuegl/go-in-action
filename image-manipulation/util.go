package image_manipulation

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](es ...T) T {
	var m T
	if len(es) == 0 {
		return m
	}

	m = es[0]
	for _, e := range es[1:] {
		if e > m {
			m = e
		}
	}

	return m
}

func Min[T constraints.Ordered](es ...T) T {
	var m T
	if len(es) == 0 {
		return m
	}

	m = es[0]
	for _, e := range es[1:] {
		if m > e {
			m = e
		}
	}

	return m
}
