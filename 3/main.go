package main

import (
	"fmt"
)

func main() {
	a := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(a))
}

/*
可以考虑从一个空字符串每次增加一个字符直到s结束,

当前字符为s[i],

left = max(left, last[s[i]]);获得的区间(left, i]是以s[i]结尾无重复字符的最长字串,

因为s[left]与s[i]是同一个字符,

减小left会有重复字符,

从0遍历到s.size()就取到了每个字符结尾的最长无重复字符字串,

ans记录其中的最大值,


class Solution {
    public int lengthOfLongestSubstring(String s) {
        int n = s.length(), ans = 0;
        Map<Character, Integer> map = new HashMap<>();
        for (int end = 0, start = 0; end < n; end++) {
            char alpha = s.charAt(end);
            if (map.containsKey(alpha)) {
                start = Math.max(map.get(alpha), start);
            }
            ans = Math.max(ans, end - start + 1);
            map.put(s.charAt(end), end + 1);
        }
        return ans;
    }
}


*/
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int {}
	res := 0
	for i, j := 0, 0; i < len(s); i ++ {
		m[s[i]] ++
		for m[s[i]] > 1 {
			m[s[j]] --
			j ++
		}
		res = max(res, i - j + 1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
