package main

/*
解题思路: 减少算法复杂度，不用两层遍历
*/

func twoSum(nums []int, target int) []int {
    indexSlice := make([]int, 0)
    mapSlice := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        temp := target - nums[i]
        if temp < 0 {
            return indexSlice
        }
        if index, ok := mapSlice[temp]; ok && i != index {
            indexSlice = append(indexSlice, i, index)
            return indexSlice
        }
        mapSlice[nums[i]] = i
    }
    return indexSlice
}

func main() {
//    fmt.Println(twoSum([]int{1, 2},1))
}
