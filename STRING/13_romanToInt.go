package main

import(
	"fmt"
)

func main (){
	fmt.Print(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    result := 0
    for i := 0; i < len(s); i++ {
        curr := m[s[i]]
        if i+1 < len(s) {
            next := m[s[i+1]]
            if curr < next {
                result += next - curr
                i++
            } else {
                result += curr
            }
        } else {
            result += curr
        }
    }
    return result
}