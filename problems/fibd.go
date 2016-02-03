package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func FIBD() {
    dat, err := ioutil.ReadFile("inputs/fibd.in")
    check(err)
    result := SolutionFIBD(string(dat))
    err = ioutil.WriteFile("inputs/fibd.out", []byte(result), 0644)
    check(err)
}

func SolutionFIBD(s string) string{
    input := strings.Split(s, " ")
    n, err := strconv.Atoi(input[0])
    check(err)
    k, err := strconv.Atoi(input[1])
    check(err)
    return strconv.Itoa(calcMortalRabbits(n, k))
}

func calcMortalRabbits(n int, m int) int{
    if (n<=2){
        return 1
    }
    return calcMortalRabbits(n-1, 1) + (1 * calcMortalRabbits(n-2, 1))
}