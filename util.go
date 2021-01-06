package main

import "sort"

func sortedKeys(m map[int]int) []int {
    result := make([]int, 0)
    
    for k := range m {
        result = append(result, k)
    }
    
    sort.Ints(result)
    
    return result
}
