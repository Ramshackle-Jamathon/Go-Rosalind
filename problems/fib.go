package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func FIB() {
    dat, err := ioutil.ReadFile("inputs/fib.in")
    check(err)
    result := SolutionFIB(string(dat))
    err = ioutil.WriteFile("inputs/fib.out", []byte(result), 0644)
    check(err)
}

func SolutionFIB(s string) string{
    input := strings.Split(s, " ")
    n, err := strconv.Atoi(input[0])
    check(err)
    k, err := strconv.Atoi(input[1])
    check(err)
    return strconv.Itoa(calcRabbits(n, k))
}

func calcRabbits(n int, k int) int{
    if (n<=2){
        return 1
    }
    return calcRabbits(n-1, k) + (k * calcRabbits(n-2, k))
}