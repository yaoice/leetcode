package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
    varStrSlice := []struct{
        s string
        res int
    }{
        {
            "abcabcbb",
            3,
        },
        {
            "bbb",
            1,
        },
        {
          " ",
          1,
        },
        {
          "abcde",
          5,
        },
        {
            "pwwkew",
            3,
        },
        {
          "",
          0,
        },
        {
            "aab",
            2,
        },
        {
            "cdd",
            2,
        },
    }

    for _, str := range varStrSlice {
        if lengthOfLongestSubstring(str.s) == str.res {
            t.Log("success")
        } else {
            t.Log("failed")
        }
    }
}