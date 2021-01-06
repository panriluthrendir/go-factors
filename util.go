package main

import (
    "sort"
    "strconv"
)

func sortedKeys(m map[int]int) []int {
    result := make([]int, 0)
    
    for k := range m {
        result = append(result, k)
    }
    
    sort.Ints(result)
    
    return result
}

func pprint(n int, sep string) string {
    digits := strconv.FormatInt(int64(n), 10)
    
    if  len(digits) < 4 {
        return digits
    }
    
    overFlow := len(digits)%3
    result := digits[:overFlow]
    
    for i := 0; i < len(digits)/3; i++ {
        if len(result) > 0 {
            result += sep
        }
        result += digits[3*i + overFlow : 3*(i+1) + overFlow]
    }
    return result
}
