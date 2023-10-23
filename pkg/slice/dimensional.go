package slice

func New2D[T any](n, m int) [][]T {
	s := make([][]T, n)
	for i := range s {
		s[i] = make([]T, m)
	}
	return s
}
