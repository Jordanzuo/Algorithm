package nDiagonal

import (
	"testing"
)

func TestGetValidCount(t *testing.T) {
	diagonal := NewDiagonal(2, 3, true)
	expect := 3
	get := diagonal.GetValidCount()
	if get != expect {
		t.Errorf("Expect to get %d solutions, but get %d instead.", expect, get)
		return
	}

	diagonal = NewDiagonal(5, 16, false)
	expect = 2
	get = diagonal.GetValidCount()
	if get != expect {
		t.Errorf("Expect to get %d solutions, but get %d instead.", expect, get)
		return
	}
}
