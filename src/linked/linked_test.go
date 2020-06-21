package linked_test

import (
	"testing"

	"../linked"
)

func TestLinkedList(t *testing.T) {

}

// -------------------Skip List Test-------------------
func TestSkipListInsert(t *testing.T) {
	list := linked.NewSkipList()
	list.Insert(3)
	list.Insert(5)
	list.Insert(7)
	list.Insert(4)
	list.Insert(2)

	list.Print()
	if list.Size() != 5 {
		t.Error("Unexpected size ", list.Size())
	}
}

func TestSkipListFind(t *testing.T) {
	list := linked.NewSkipList()
	list.Insert(3)
	list.Insert(5)
	list.Insert(7)
	list.Insert(4)
	list.Insert(2)

	node := list.Find(3)
	if node.Value() != 3 {
		t.Error("Unexpected value ", node.Value())
	}
}

func TestSkipListDelete(t *testing.T) {
	list := linked.NewSkipList()
	list.Insert(3)
	list.Insert(5)
	list.Insert(7)
	list.Insert(4)
	list.Insert(2)

	if !list.Delete(3) {
		t.Error("Delete value ", 3, " result is unexpected")
	}
	if list.Delete(6) {
		t.Error("Delete value ", 6, " result is unexpected")
	}
	if list.Size() != 4 {
		t.Error("Unexpected size ", list.Size())
	}
	list.Print()
}

// ---------------------Cache Skip List Test-------------------
func TestCacheSkipListInsert(t *testing.T) {
	list := linked.NewCacheSkipList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(5)
	list.Insert(6)
	list.Insert(8)
	list.Insert(4)

	list.Print()
	if list.Size() != 6 {
		t.Error("Unexpected size ", list.Size())
	}
}

func TestCacheSkipListFind(t *testing.T) {
	list := linked.NewCacheSkipList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(5)
	list.Insert(6)
	list.Insert(8)
	list.Insert(4)

	node := list.Find(5)
	if node.Value() != 5 {
		t.Error("Unexpected value ", node.Value())
	}
}

func TestCacheSkipListDelete(t *testing.T) {
	list := linked.NewCacheSkipList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(5)
	list.Insert(6)
	list.Insert(8)
	list.Insert(4)

	if !list.Delete(3) {
		t.Error("Delete value ", 3, " result is unexpected")
	}
	if list.Delete(7) {
		t.Error("Delete value ", 7, " result is unexpected")
	}
	if list.Size() != 5 {
		t.Error("Unexpected size ", list.Size())
	}
	list.Print()
}
