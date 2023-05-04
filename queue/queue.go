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
	Length int
}

func (q *Queue[t]) Enqueue(elem t) {
	if q.Length == 0 {
		n := newnode(elem, nil)

		q.head = n
		q.tail = n
		q.Length++
		return
	}

	n := newnode(elem, q.tail)

	q.tail.next = n
	q.tail = n
	q.Length++
}

func (q *Queue[t]) Dequeue() *t {
	if q.head == nil {
		return nil
	}

	if q.Length == 1 {
		res := q.head
		q.tail = nil
		q.head = nil

		q.Length--
		return &res.value
	}

    h := q.head
	q.head = q.head.next
	q.Length--

    h.next = nil
	return &h.value
}

func (q Queue[t]) Peek() *t {
	if q.Length == 0 {
		return nil
	}

	return &q.head.value
}
