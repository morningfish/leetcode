package main

func numberOfArithmeticSlices(A []int) int {
	if len(A)== 0{
		return 0
	}

	res :=0
	add :=0
	for i:=2;i<len(A);i++{
		if A[i-1]-A[i]==A[i-2]-A[i-1]{
			add+=1
			res+=add
		}else{
			add=0
		}
	}
	return res
}
