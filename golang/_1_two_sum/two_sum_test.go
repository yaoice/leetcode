package main

import (
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {

    varStrSlice := []struct{
        intSlice []int
        target int
        res []int
        res2 []int
    }{
        {
            []int{2,7,11,15},
            9,
            []int{0, 1},
            []int{1, 0},
        },
        {
            []int{1,2},
            2,
            []int{0, 1},
            []int{1, 0},
        },
        {
            []int{2},
            1,
            []int{0},
            []int{1},
        },
    }

    for _, str := range varStrSlice {
        result := twoSum(str.intSlice, str.target)
        if reflect.DeepEqual(result, str.res) {
            t.Logf("success")
        } else if reflect.DeepEqual(result, str.res2) {
            t.Logf("success")
        } else {
            t.Logf("failed")
        }
    }
}