package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func IEV() {
    dat, err := ioutil.ReadFile("inputs/iev.in")
    check(err)
    result := SolutionIEV(string(dat))
    err = ioutil.WriteFile("inputs/iev.out", []byte(result), 0644)
    check(err)
}

func SolutionIEV(s string) string{
    input := strings.Split(s, " ")
    n0, err := strconv.ParseFloat(input[0], 64)
    check(err)
    n1, err := strconv.ParseFloat(input[1], 64)
    check(err)
    n2, err := strconv.ParseFloat(input[2], 64)
    check(err)
    n3, err := strconv.ParseFloat(input[3], 64)
    check(err)
    n4, err := strconv.ParseFloat(input[4], 64)
    check(err)
    n5, err := strconv.ParseFloat(input[5], 64)
    check(err)
    return strconv.FormatFloat(calcExpectedOffspring(n0, n1, n2, n3, n4, n5), 'f', -1, 64)
}

func calcExpectedOffspring(n0, n1, n2, n3, n4, n5 float64) float64{
    return 2.0 * (n0 + n1 + n2) + 1.5 * n3 + n4
}

