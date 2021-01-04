package main

import (
    "errors"
)

func factorize(num int, primes []int) (map[int]int, error) {
    result := make(map[int]int)

    for _, p := range primes {
        if num == 1 {
            return result, nil
        }
        if p*p > num {
            result[num] = 1
            return result, nil
        }
        i := 0
        for {
            if num%p == 0 {
                i++
                num = num / p
            } else {
                break
            }
        }
        if i > 0 {
            result[p] = i
        }
    }
    return nil, errors.New("The number is too large to factorize")
}

func primesUnder(N int) []int {
    isComposite := make([]bool, 0)

    for i := 0; i < N; i++ {
        isComposite = append(isComposite, false)
    }

    for n := 2; n < N/2; n++ {
        for i := 2; i < N/n; i++ {
            isComposite[i*n] = true
        }
    }

    result := make([]int, 0)

    for n := 2; n < N; n++ {
        if !isComposite[n] {
            result = append(result, n)
        }
    }

    return result
}
