package main

import "fmt"

func main() {
	a := []string{"flower", "flow", "flight"}
	b := []string{"dog", "racecar", "car"}
	c := []string{"aa", "ab"}
	fmt.Println(longestCommonPrefix(a))
	fmt.Println(longestCommonPrefix(b))
	fmt.Println(longestCommonPrefix(c))
}

/*
先获取前两个的相同字符串，再把这个字符串和后面的对比
时间复杂度：0(n)
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) < 1 {
		return ""
	}
	defaultStr1 := strs[0]
	defaultStr2 := strs[1]
	strs = strs[2:]
	defaultStr := getString(defaultStr1, defaultStr2)
	for _, value := range strs {
		defaultStr = getString(defaultStr, string(value))
	}
	return defaultStr
}
func getString(str1, str2 string) string {
	result := ""
	rune1 := []rune(str1)
	rune2 := []rune(str2)
	length1 := len(rune1)
	length2 := len(rune2)
	if length1 > length2 {
		for count, value := range str2 {
			if rune1[count] == value {
				result += string(value)
			} else {
				break
			}
		}
	} else {
		for count, value := range str1 {
			if rune2[count] == value {
				result += string(value)
			} else {
				break
			}
		}
	}

	return result
}

func longestCommonPrefix1(strs []string) string {
	// 长度为0的时候直接返回空字符串
	if len(strs) == 0 {
		return ""
	}
	// 以列表第一个字符串为基准字符串， 逐一判断每一个字符
	for index, char := range strs[0] {
		// 根据基准字符串对比列表中其他字符串相同位置字符是否相等
		for _, v := range strs {
			// 如果index大于当前字符串长度，直接返回当前字符串。不可能还有更长的公共前缀
			if index >= len(v) {
				return v
			}
			// 如果当前字符串不相等，说明前面的是已经匹配上的
			if v[index] != byte(char) {
				return v[:index]
			}
		}
	}
	return strs[0]
}
