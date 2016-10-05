package store

import (
	"strconv"
	"testing"
)

func TestCreateLeafs(t *testing.T) {
	p := NewPage()
	p.createLeafs()

	for i, leaf := range p.leafs {
		if leaf == nil {
			t.Error("Leaf", i, "was not created!")
		}
	}
}

func TestRemoveLeafs(t *testing.T) {
	p := NewPage()
	p.createLeafs()
	p.removeLeafs()

	for i, leaf := range p.leafs {
		if leaf != nil {
			t.Error("Leaf", i, "was not removed!")
		}
	}
}

func TestMoveKeys(t *testing.T) {
	p := NewPage()
	p.createLeafs()

	// Populate page keys
	for i := 1; i <= int(*pageKeysSize); i++ {
		p.add(NewStringKey(strconv.Itoa(i)))
	}

	p.moveKeysToLeafs()
	p.leaf = false
	isKeys(p, t)

	p.moveKeysFromLeafs()
	p.leaf = true
	isKeys(p, t)
}

func isKeys(p *Page, t *testing.T) {
	for i := 1; i <= int(*pageKeysSize); i++ {
		_, err := p.load(strconv.Itoa(i))

		if err != nil {
			t.Error("Received error on key load: ", err)
		}
	}
}
