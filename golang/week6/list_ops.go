package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, i := range s {
		initial = fn(initial, i)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for _, i := range s.Reverse() {
		initial = fn(i, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	final := IntList{} // this feels dicey
	for _, i := range s {
		if fn(i) {
			final = append(final, i)
		}
	}
	return final
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	final := make(IntList, len(s))
	for ix, i := range s {
		final[ix] = fn(i)
	}
	return final
}

func (s IntList) Reverse() IntList {
	l := len(s)
	final := make(IntList, len(s))
	for i := 0; i < len(s); i++ {
		final[i] = s[l-i-1]
	}
	return final
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	final := s // copy in
	for _, list := range lists {
		final = append(final, list...)
	}
	return final
}
