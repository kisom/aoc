package main

type Set struct {
	m map[string]bool
	s []string
}

func NewSet() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

func (s *Set) Has(v string) bool {
	return s.m[v]
}

func (s *Set) Empty() bool {
	return len(s.s) == 0
}

func (s *Set) Pop() string {
	v := s.s[0]
	s.s = s.s[1:]
	delete(s.m, v)
	return v
}

func (s *Set) Drop() {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	delete(s.m, v)
}

func (s *Set) Head() string {
	return s.s[0]
}

func (s *Set) Tail() string {
	return s.s[len(s.s)-1]
}

func (s *Set) Add(x string) {
	if !s.m[x] {
		s.m[x] = true
		s.s = append(s.s, x)
	}
}

func (s *Set) AddList(xs []string) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *Set) AddMap(m map[string]bool) {
	for k := range m {
		s.Add(k)
	}
}
