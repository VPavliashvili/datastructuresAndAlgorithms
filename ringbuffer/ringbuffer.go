package ringbuffer

var buffersize = 3

// AKA circular queue
func Create[t comparable]() ringbuff[t] {
	return ringbuff[t]{
		buffer:      make([]*t, buffersize),
		head:        0,
		tail:        -1,
		length:      0,
	}
}

type ringbuff[t comparable] struct {
	buffer      []*t
	head        int
	tail        int
	length      int
}

var tailGottaLeap bool

func (rb *ringbuff[t]) Enqueue(elem t) {
	idx := (rb.tail + 1) % len(rb.buffer)
	if rb.tail+1 == len(rb.buffer) {
		tailGottaLeap = true
	}

	if idx == rb.head && tailGottaLeap {
		resize(rb)
	}

	rb.buffer[idx] = &elem
	rb.tail = idx
	rb.length++
	tailGottaLeap = false
}

func resize[t comparable](rb *ringbuff[t]) {
	toend := rb.buffer[rb.head:len(rb.buffer)]
	fromstart := rb.buffer[0:rb.head]
	rb.buffer = append(toend, fromstart...)
	rb.buffer = append(rb.buffer, make([]*t, buffersize)...)
	rb.head = 0
	rb.tail = rb.length
}

func tailHasDoneLap[t comparable](rb ringbuff[t]) bool {
	return rb.head > 0 && rb.buffer[rb.head-1] != nil
}

func (rb *ringbuff[t]) Dequeue() *t {
	if rb.length == 0 {
		return nil
	}

	idx := rb.head % len(rb.buffer)
	res := rb.buffer[idx]
	rb.buffer[idx] = nil
	rb.head = (idx + 1) % len(rb.buffer)
	rb.length--

	return res
}

func (rb ringbuff[t]) Get(index int) *t {
	if index > rb.length-1 {
		return nil
	}

	if rb.tail > rb.head {
		return rb.buffer[rb.head : rb.tail+1][index]
	} else if rb.tail < rb.head {
		toend := rb.buffer[rb.head:len(rb.buffer)]
		fromstart := rb.buffer[0 : rb.tail+1]
		return append(fromstart, toend...)[index]
	} else {
		if index == rb.head {
			return rb.buffer[index]
		} else {
			return nil
		}
	}
}
