package main

import "testing"

func Test(t *testing.T) {
	exmp1 := []int{1, 2, 3, 4}
	res := numberOfArithmeticSlices(exmp1)

	if res == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
