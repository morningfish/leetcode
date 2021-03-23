package main

import (
	"fmt"
	"strings"
)

func main() {
	l := []string{"a==b", "b==c", "c==a"}
	res := equationsPossible(l)
	fmt.Println(res)
}

func equationsPossible(equations []string) bool {
	equal := make(map[string][]string)
	unequal := make(map[string][]string)
	for _, index := range equations {
		res := strings.Split(index, "=")
		if len(res) == 2 {
			if _, ok := unequal[res[0]]; ok {
				unequal[res[0]] = append(unequal[res[0]], strings.Split(res[1], "!")[1])
			}else{
			unequal[res[0]] = []string{strings.Split(res[1], "!")[1]}}
		}
		else{
			if _, ok := equal[res[0]]; ok {
				equal[res[0]] = append(equal[res[0]], strings.Split(res[1], "!")[1]))
			}else{
			equal[res[0]] = []string{strings.Split(res[1], "!")[1]}
		}
	}
}
	return true
}
