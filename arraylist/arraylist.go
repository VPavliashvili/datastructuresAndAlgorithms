package arraylist

import (
	"fmt"
)

var memsize = 2

type ArrayList[t comparable] struct {
	memory []t
	length int
}

func Create[t comparable]() ArrayList[t] {
	return ArrayList[t]{
		memory: make([]t, memsize),
		length: 0,
	}
}

func tryMemincrease[t comparable](al *ArrayList[t]) {
    if al.length >= len(al.memory) {
        newmemory := make([]t, memsize)
        al.memory = append(al.memory, newmemory...)
    }
}

func (al *ArrayList[t]) Append(elem t) {
    tryMemincrease(al)

	al.memory[al.length] = elem
	al.length++
}

func (al *ArrayList[t]) Prepend(elem t) {
    tryMemincrease(al)

	al.length++
	al.memory = append([]t{elem}, al.memory...)
}

func (al ArrayList[t]) ElementAt(index int) t {
	if index >= al.length {
		msg := fmt.Sprintf("array out of bound. index %v does not exist", index)
		panic(msg)
	}

	res := al.memory[index]
	return res
}

func (al *ArrayList[t]) RemoveAt(index int) t {
	if index >= al.length {
		msg := fmt.Sprintf("array out of bound. index %v does not exist", index)
		panic(msg)
	}
	al.length--
	res := al.memory[index]
	al.memory = append(al.memory[:index], al.memory[index+1:]...)

	return res
}

func (al *ArrayList[t]) Remove(elem t) *t {
	var res *t

	for i, item := range al.memory {
		if item == elem {
			res = &elem
			al.memory = append(al.memory[:i], al.memory[i+1:]...)
			al.length--
			break
		}
	}

	return res
}

func (al ArrayList[t]) Length() int {
	return al.length
}
