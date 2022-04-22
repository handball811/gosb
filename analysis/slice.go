package analysis

func Map[T any, S any](s []T, f func(T) (S, bool)) []S {
	if s == nil {
		return []S{}
	}
	ret := make([]S, 0, len(s))
	for _, t := range s {
		n, ok := f(t)
		if !ok {
			continue
		}
		ret = append(ret, n)
	}
	return ret
}

type StringSlice []string

func (s StringSlice) Where(f func(string) bool) []string {
	ret := make([]string, 0, len(s))
	for _, t := range s {
		if f(t) {
			ret = append(ret, t)
		}
	}
	return ret
}

func Keys[T comparable, S any](mp map[T]S) []T {
	ret := make([]T, 0, len(mp))
	for k := range mp {
		ret = append(ret, k)
	}
	return ret
}
