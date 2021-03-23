package bubble

import (
	"testing"
)

func TestBubble(t *testing.T) {
	l := []int{7, 3, 5, 2, 5, 6, 3, 6, 5, 3, 2}
	Bubble(&l)
	t.Log(l)
}
