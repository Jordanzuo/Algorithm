package nQueen

import (
	"testing"
)

func TestGetValidCount(t *testing.T) {
	queen := NewNQueen(4, false)
	expect := 2
	get := queen.GetValidCount()
	if get != expect {
		t.Errorf("Expect to get %d solutions, but get %d now.", expect, get)
		return
	}

	queen = NewNQueen(4, true)
	expect = 2
	get = queen.GetValidCount()
	if get != expect {
		t.Errorf("Expect to get %d solutions, but get %d now.", expect, get)
		return
	}

	queen = NewNQueen(8, false)
	expect = 92
	get = queen.GetValidCount()
	if get != expect {
		t.Errorf("Expect to get %d solutions, but get %d now.", expect, get)
		return
	}
}
