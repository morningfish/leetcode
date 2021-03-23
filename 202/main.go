package main

func main() {
}

func isHappy(n int) bool {
	for n > 10 {
	}
	if n == 1 {
		return true
	}
	return false
}

func sum(n int) int {
	res := 0
	for n/10 != 0 {
		g := n % 10
		n -= g
		n = n / 10
	}
}
