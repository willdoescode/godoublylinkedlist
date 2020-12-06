package DubList

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type node struct {
	elem int
	prev *node
	next *node
}

func (n node) Slicify() []int {
	res := make([]int, 0)
	res = append(res, n.elem)
	for n.next != nil {
		res = append(res, n.next.elem)
		n = *n.next
	}
	return res
}

func (n node) Stringify() string {
	res := make([]string, 0)
	for _, v := range n.Slicify() {
		res = append(res, fmt.Sprintf("%d", v))
	}
	return strings.Join(res, " ")
}

func (n node) Reverse() node {
	for n.next != nil {
		n = *n.next
	}

	h := New(n.elem)
	for n.prev != nil {
		h.Append(n.prev.elem)
		n = *n.prev
	}
	return h
}

func (n *node) Append(elem int) {
	if n.next == nil {
		n.next = &node{
			elem: elem,
			prev: n,
			next: nil,
		}
	} else {
		n.next.Append(elem)
	}
}

func (n node) AppendFront(elem int) node {
	return node{
		elem: elem,
		prev: nil,
		next: &n,
	}
}

func (n node) Search(elem ...int) int {
	var amount int
	if len(elem) == 1 {
		amount = 0
	} else {
		amount = elem[1]
		amount++
	}
	if n.elem == elem[0] {
		return amount
	} else if n.next == nil {
		return 0
	} else {
		return n.next.Search(elem[0], amount)
	}
}

func (n *node) Remove(index int) {
	if index == 0 {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = n.prev
		}
		return
	}
	index--
	n.next.Remove(index)
}

func (n *node) RemoveInt(a int) {
	pos := n.Search(a)
	n.Remove(pos)
}

func (n node) Len() int {
	var count int
	for n.next != nil {
		count++
		n = *n.next
	}
	return count
}

func (n node) Display() []int {
	res := make([]int, 0)
	res = append(res, n.elem)
	if n.next != nil {
		res = append(res, n.next.Display()...)
	}
	return res
}

func (n node) String() string {
	res := make([]string, 0)
	for i, v := range n.Display() {
		res = append(res, fmt.Sprintf("%d", v))
		if i != n.Len() {
			res = append(res, "->")
		}
	}
	return strings.Join(res, " ")
}

func New(pass interface{}) node {
	switch pass.(type) {
	case []int:
		h := node{
			elem: pass.([]int)[0],
			prev: nil,
			next: nil,
		}
		for _, v := range pass.([]int)[1:] {
			h.Append(v)
		}
		return h
	case int:
		return node{
			elem: pass.(int),
			prev: nil,
			next: nil,
		}
	case string:
		elem, err := strconv.Atoi(pass.(string))
		if err != nil {
			log.Fatal(err)
		}
		return node{
			elem: elem,
			prev: nil,
			next: nil,
		}
	case []string:
		elem, err := strconv.Atoi(pass.([]string)[0])
		if err != nil {
			log.Fatal(err)
		}
		h := node{
			elem: elem,
			prev: nil,
			next: nil,
		}
		for _, v := range pass.([]string)[1:] {
			elem, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			h.Append(elem)
		}
		return h
	}
	panic(fmt.Sprintf("Cannot handle %T", pass))
}
