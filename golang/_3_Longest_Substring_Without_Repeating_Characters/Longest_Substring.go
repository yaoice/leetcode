package main

/*
str: abcdea，统计两个重复字符之间的距离，对比取最大值

解题思路: 动态规划-滑动窗口
*/

func lengthOfLongestSubstring(s string) int {
    flag := [256]int{}
    beg := 0
    max := 0
    for i := 0; i < len(s); i++ {
        if flag[s[i]] > 0 && flag[s[i]] > beg {
            beg = flag[s[i]]
        }
        flag[s[i]] = i+1
        max = maxnum(max, i - beg + 1)
    }
    return max
}

func maxnum(a, b int) int {
    if a > b {
        return a
    }
    return b
}


/*
func main() {
    fmt.Println(lengthOfLongestSubstring("abcdaed"))
}
 */




