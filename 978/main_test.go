package main

import "testing"

func TestmaxTurbulenceSize(t *testing.T) {
	test1 := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	res1 := maxTurbulenceSize(test1)
	if res1 == 5 {
		t.Logf("successful")
	} else {
		t.Fatal("error")
	}
}
