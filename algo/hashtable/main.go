package main

type Node struct {
	Next  *Node
	Value int
}

type HashTable struct {
	Table map[int]*Node
}

func HashFunction(v int) int {
	return v % 2
}

func insert(v int, t *HashTable) *HashTable {
	hash := HashFunction(v)
	temp := t.Table[hash]
	lastInsert := &Node{
		Next:  temp,
		Value: v,
	}
	t.Table[hash] = lastInsert
	return t
}

func create() *HashTable {
	h := &HashTable{
		Table: make(map[int]*Node),
	}
	for n := 0; n < 100; n++ {
		h = insert(n, h)
	}
	return h
}

func lookup(v int, h *HashTable) bool {
	index := HashFunction(v)
	current := h.Table[index]
	for {
		if current.Value == v {
			return true
		}
		if current.Next == nil {
			return false
		}
		current = current.Next
	}
}

func main() {

	h := create()
	insert(435, h)
	lookup(2, h)
}
