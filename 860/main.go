package main

import "fmt"


func main(){
	a:=[]int{5,5,5,10,10,10}
	fmt.Println(lemonadeChange(a))
}


func lemonadeChange(bills []int) bool {
	five,ten:=0,0
	for _, bill := range bills{
		switch bill{
		case 5:
			five++
		case 10:
			if five >0{
				five--
				ten++
			}else{
				return false
			}
		case 20:
			if five>0 && ten>0{
				five--
				ten--
			}else if five>2{
				five-=3
			}else{
				return false
			}
			}
		}
		return true
	}
