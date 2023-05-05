package stack

type node[t any] struct {
	value t
	prev  *node[t]
}

type Stack[t any] struct {
	head   *node[t]
	length int
}

func (s *Stack[t]) Push(elem t) {
	if s.head == nil {
		s.head = &node[t]{
			value: elem,
		}

		s.length++
		return
	}

	n := &node[t]{
		value: elem,
		prev:  s.head,
	}

	s.head = n
	s.length++
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
