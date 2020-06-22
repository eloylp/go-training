package main

import (
	"testing"
)

func TestRemove(t *testing.T) {

	var linkedList *Node = nil
	linkedList = insert(linkedList, "start")
	hello := insert(linkedList, "hello")
	insert(linkedList, "bye")
	insert(linkedList, "cat")

	remove(linkedList, "bye")

	if hello.Next.Value != "cat" {
		t.Error("Expected cat as next element")
	}

	if hello.Next.Next.Value != "bye" {
		t.Error("Expected bye as next element")
	}

}
