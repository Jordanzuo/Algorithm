package maxRegionSum

import (
	"math/rand"
	"testing"
	"time"
)

// go test -bench .
// goos: linux
// goarch: amd64
// pkg: maxRegionSum
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkGetMaxRegionSum1-12                  36          32362822 ns/op
// BenchmarkGetMaxRegionSum2-12               19452             63265 ns/op
// PASS
// ok      maxRegionSum    5.205s

type Rand struct {
	*rand.Rand
}

// 获得Rand对象(如果是循环生成多个数据，则只需要调用本方法一次，而不能调用多次，否则会得到相同的随机值)
func GetRand() *Rand {
	return &Rand{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func TestGetMaxRegionSum(t *testing.T) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = -1*count + GetRand().Intn(2*count)
	}

	lowerIndex1, upperIndex1 := GetMaxRegionSum1(list)
	lowerIndex2, upperIndex2 := GetMaxRegionSum2(list)
	if lowerIndex1 != lowerIndex2 || upperIndex1 != upperIndex2 {
		t.Errorf("(%d, %d) != (%d, %d)", lowerIndex1, upperIndex1, lowerIndex2, upperIndex2)
		return
	}
}

func TestGetMaxRegionSum1(t *testing.T) {
	list := []int{-4, -2, -3, -2, -1}
	expectLowerIndex, expectUpperIndex := 4, 4
	getLowerIndex, getUpperIndex := GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 4
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -4, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 5
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -6, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 3
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 3
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 5, 6, -1}
	expectLowerIndex, expectUpperIndex = 5, 6
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 1, 8, -1}
	expectLowerIndex, expectUpperIndex = 5, 6
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, -20, 1, 8, -1}
	expectLowerIndex, expectUpperIndex = 6, 7
	getLowerIndex, getUpperIndex = GetMaxRegionSum1(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
}

func TestGetMaxRegionSum2(t *testing.T) {
	list := []int{-4, -2, -3, -2, -1}
	expectLowerIndex, expectUpperIndex := 4, 4
	getLowerIndex, getUpperIndex := GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get upper index: %d, got %d", expectUpperIndex, getUpperIndex)
		return
	}

	list = []int{-4, -2, 3, 4, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 4
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get upper index: %d, got %d", expectUpperIndex, getUpperIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -4, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 5
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -6, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 3
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 5, -1}
	expectLowerIndex, expectUpperIndex = 2, 3
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 5, 6, -1}
	expectLowerIndex, expectUpperIndex = 5, 6
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, 1, 8, -1}
	expectLowerIndex, expectUpperIndex = 5, 6
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}

	list = []int{-4, -2, 3, 4, -10, -20, 1, 8, -1}
	expectLowerIndex, expectUpperIndex = 6, 7
	getLowerIndex, getUpperIndex = GetMaxRegionSum2(list)
	if getLowerIndex != expectLowerIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
	if getUpperIndex != expectUpperIndex {
		t.Errorf("Expect to get lower index: %d, got %d", expectLowerIndex, getLowerIndex)
		return
	}
}

func BenchmarkGetMaxRegionSum1(b *testing.B) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = -1*count + GetRand().Intn(2*count)
	}

	for i := 0; i < b.N; i++ {
		GetMaxRegionSum1(list)
	}
}

func BenchmarkGetMaxRegionSum2(b *testing.B) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = -1*count + GetRand().Intn(2*count)
	}

	for i := 0; i < b.N; i++ {
		GetMaxRegionSum2(list)
	}
}
