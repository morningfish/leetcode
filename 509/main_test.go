package main

import "testing"

func TestAdd(t *testing.T) {
	sum := fib(3)
	if sum == 2 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
	sum = fib(10)
	t.Log(sum)
}
