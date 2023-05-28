package linkedlist

func Create[t comparable]() DoublyLinkedList[t] {
	var result DoublyLinkedList[t]
	return result
}

type node[t comparable] struct {
	value t
	next  *node[t]
	prev  *node[t]
}

type DoublyLinkedList[t comparable] struct {
	head   *node[t]
	tail   *node[t]
	length int
}

func (d *DoublyLinkedList[t]) insertWhenEmpty(item t) {
	d.head = &node[t]{
		value: item,
		next:  nil,
		prev:  nil,
	}
	d.tail = d.head
	d.length++
}

func (d *DoublyLinkedList[t]) Append(item t) {
	if d.head == nil {
		d.insertWhenEmpty(item)
		return
	}

	n := &node[t]{
		value: item,
		next:  nil,
		prev:  d.tail,
	}
	d.tail.next = n
	d.tail = n
	d.length++
}

func (d *DoublyLinkedList[t]) Prepend(item t) {
	if d.head == nil {
		d.insertWhenEmpty(item)
		return
	}

	n := &node[t]{
		value: item,
		next:  d.head,
		prev:  nil,
	}
	d.head.prev = n
	d.head = n
	d.length++
}

func (d *DoublyLinkedList[t]) InsertAt(item t, index int) {
	if index > d.length {
		panic("OUT OF BOUND ERROR")
	}
	if index == d.length {
		d.Append(item)
		return
	}
	if index == 0 {
		d.Prepend(item)
		return
	}

	cur := d.getNode(index)
	n := &node[t]{
		value: item,
		next:  cur,
		prev:  cur.prev,
	}
	cur.prev.next = n
	cur.prev = n
	d.length++
}

func (d *DoublyLinkedList[t]) RemoveAt(index int) *t {
	cur := d.getNode(index)
	if cur != nil {
		if d.head == d.tail {
			d.length = 0
			d.head = nil
			d.tail = nil
			defer func() { cur = nil }()
			return &cur.value
		}
		if cur.next == nil {
			cur.prev.next = nil
			d.tail = cur.prev
		} else if cur.prev == nil {
			cur.next.prev = nil
			d.head = cur.next
		} else {
			n := cur.next
			p := cur.prev
			cur.prev.next = n
			cur.next.prev = p
		}
		defer func() { cur = nil }()
	} else {
		return nil
	}

	d.length--
	return &cur.value
}

// removes first occurence of given item
func (d *DoublyLinkedList[t]) Remove(item t) *t {
	index := 0
	cur := d.head
	for cur != nil {
		if cur.value == item {
			break
		}
		cur = cur.next
		index++
	}

	res := d.RemoveAt(index)
	return res
}

func (d DoublyLinkedList[t]) ElementAt(index int) *t {
	n := d.getNode(index)
	if n != nil {
		return &n.value
	}
	return nil
}

func (d DoublyLinkedList[t]) getNode(index int) *node[t] {
	cur := d.head
	i := 0
	for cur != nil && i < index {
		cur = cur.next
		i++
	}
	if i == index && cur != nil {
		return cur
	}
	return nil
}

func (d DoublyLinkedList[t]) Length() int {
	return d.length
}
