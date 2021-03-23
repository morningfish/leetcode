package main

func main(){
	generate(5)
}

// 直接
/*
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    for i := range ans {
        ans[i] = make([]int, i+1)
        ans[i][0] = 1
        ans[i][i] = 1
        for j := 1; j < i; j++ {
            ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
        }
    }
    return ans
}
*/

// 间接 移位相加
func generate(numRows int) [][]int{
	if numRows==0{
		return [][]int{}
	}
 ans:= [][]int{
	 []int{1},
 }
 for i:=0;i<numRows-1;i++{
	 currentx:=append(ans[i],0)
	 currenty:=append([]int{0},ans[i]...)
	 current:= []int{}
	 for x:=0;x<len(currentx);x++{
		 current=append(current,currentx[x]+currenty[x])
	 }

	 ans = append(ans,current)
 }
 return ans 
}
