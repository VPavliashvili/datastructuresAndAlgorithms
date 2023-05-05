package stack

type node[t any] struct {
	value t
	prev  *node[t]
}

func newnode[t any](v t, previous *node[t]) *node[t] {
	return &node[t]{
		value: v,
		prev:  previous,
	}
}

type Stack[t any] struct {
	head   *node[t]
	length int
}

func (s *Stack[t]) Push(elem t) {
	s.length++

	if s.head == nil {
		s.head = newnode(elem, nil)
		return
	}

	n := newnode(elem, s.head)
	s.head = n
}

func (s *Stack[t]) Pop() *t {
	if s.head == nil {
		return nil
	}

	popped := *s.head
	s.head = s.head.prev

	val := popped.value

	s.length--

	return &val
}

func (s Stack[t]) Peek() *t {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}

func (s Stack[t]) Length() int {
	return s.length
}
