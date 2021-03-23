package quick

import "testing"

func TestQuick(t *testing.T) {
	l := []int{5, 6, 7, 3, 2, 5, 2, 4, 1, 5, 7, 8, 10, 32}
	t.Log(l)
	QuickSort(&l, 0, len(l)-1)

	t.Log(l)
}
