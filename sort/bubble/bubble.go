package bubble

func Bubble(l *[]int) {
	for x := 1; x < len(*l); x++ {
		for c := 0; c < len(*l)-x; c++ {
			if (*l)[c] > (*l)[c+1] {
				temp := (*l)[c]
				(*l)[c] = (*l)[c+1]
				(*l)[c+1] = temp
			}
		}
	}

}
