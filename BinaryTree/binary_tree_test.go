package binary_tree

import (
	"testing"
)

func TestContainBinaryTree(t *testing.T) {
	a := BinaryTree{}
	a.add(5)

	if a.contains(5) != true {
		t.Error("Contain root value not work")
	}

	a.add(10)
	a.add(1)

	if a.contains(10) != true {
		t.Error("Contain value in left not work")
	}

	if a.contains(1) != true {
		t.Error("Contain in right not work")
	}
}

func TestRemoveBinaryTree(t *testing.T)  {
	// Create Bt
	a := BinaryTree{}

	// Add values
	a.add(5)
	a.add(1)
	a.add(10)
	a.add(11)

	//Remove value on one side
	a.remove(11)

	if a.contains(11) == true {
		t.Error("Remove element not work")
	}

	//Remove another
	a.remove(5)

	if a.contains(5) == true {
		t.Error("Remove element not work")
	}

	//Remove one more
	a.remove(1)

	if a.contains(1) == true {
		t.Error("Remove element not work")
	}
}

