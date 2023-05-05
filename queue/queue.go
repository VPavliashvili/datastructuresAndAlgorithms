package queue

type node[t any] struct {
	value t
	next  *node[t]
}

func newnode[t any](v t, n *node[t]) *node[t] {
	return &node[t]{
		value: v,
		next:  n,
	}
}

type Queue[t any] struct {
	head   *node[t]
	tail   *node[t]
	length int
}

func (q *Queue[t]) Enqueue(elem t) {
	if q.length == 0 {
		n := newnode(elem, nil)

		q.head = n
		q.tail = n
		q.length++
		return
	}

	n := newnode(elem, q.tail)

	q.tail.next = n
	q.tail = n
	q.length++
}

func (q *Queue[t]) Dequeue() *t {
	if q.head == nil {
		return nil
	}

	if q.length == 1 {
		res := q.head
		q.tail = nil
		q.head = nil

		q.length--
		return &res.value
	}

    h := q.head
	q.head = q.head.next
	q.length--

    h.next = nil
	return &h.value
}

func (q Queue[t]) Peek() *t {
	if q.length == 0 {
		return nil
	}

	return &q.head.value
}

func (q Queue[t]) Length() int {
    return q.length
}
