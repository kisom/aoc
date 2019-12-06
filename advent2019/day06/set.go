package main

// A QueueSet is a queue with fast lookups. It maintains both a hash
// table and an array for the objects in the set. The QueueSet keeps
// history; items are never removed from the set.
//
// The motivation was that I want to be able to basically range over
// an ordered dictionary while adding and removing elements as needed.
// That is, you can do something like
//
// for !qs.Empty() {
//     // add or remove from qs
// }
type QueueSet struct {
	m map[string]bool
	s []string
}

func NewSet(initial string) *QueueSet {
	return &QueueSet{
		m: map[string]bool{initial: true},
	}
}

// Has returns true if v is present in the QueueSet.
func (s *QueueSet) Has(v string) bool {
	return s.m[v]
}

// Empty returns true if the QueueSet has no elements.
func (s *QueueSet) Empty() bool {
	return len(s.s) == 0
}

// Pop removes the first element in the queue and returns it.
func (s *QueueSet) Pop() string {
	v := s.s[0]
	s.s = s.s[1:]
	return v
}

// Add enqueues x if it is not present in the queueset.
func (s *QueueSet) Add(x string) {
	if !s.m[x] {
		s.m[x] = true
		s.s = append(s.s, x)
	}
}

// AddList adds
func (s *QueueSet) AddList(xs []string) {
	for _, x := range xs {
		s.Add(x)
	}
}
