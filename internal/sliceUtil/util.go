package sliceutil

func Last[S ~[]E, E any](slice S) E {
	return slice[len(slice)-1]
}
