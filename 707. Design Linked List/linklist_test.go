package ds

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := NewLinkList()
	// obj.Get(1)
	// obj.AddAtHead(1)
	// obj.AddAtTail(3)
	// obj.AddAtIndex(1, 2)
	// obj.Get(1)
	// obj.DeleteAtIndex(1)
	// obj.Get(1)

	// obj.AddAtIndex(0, 10)
	// obj.AddAtIndex(0, 20)
	// obj.AddAtIndex(1, 30)
	// obj.Get(0)

	// obj.AddAtHead(2)
	// obj.DeleteAtIndex(1)
	// obj.AddAtHead(2)
	// obj.AddAtHead(7)
	// obj.AddAtHead(3)
	// obj.AddAtHead(2)
	// obj.AddAtHead(5)
	// obj.AddAtTail(5)
	// obj.Get(5)
	// obj.DeleteAtIndex(6)
	// obj.DeleteAtIndex(4)

	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	obj.Get(5)
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)

}
