package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func IPRB() {
    dat, err := ioutil.ReadFile("inputs/iprb.in")
    check(err)
    result := SolutionIPRB(string(dat))
    err = ioutil.WriteFile("inputs/iprb.out", []byte(result), 0644)
    check(err)
}

func SolutionIPRB(s string) string{
    inputs := strings.Split(s, " ")
    k, err := strconv.ParseFloat(inputs[0], 64)
    check(err)
    m, err := strconv.ParseFloat(inputs[1], 64)
    check(err)
    n, err := strconv.ParseFloat(inputs[2], 64)
    check(err)
    p := mendelsFirst(k, m, n)
    return strconv.FormatFloat(p, 'f', -1, 64)
}

func mendelsFirst(k, m, n float64) float64 {
    return ((k*k - k) + 2*(k*m) + 2*(k*n) + (.75*(m*m - m)) + 2*(.5*m*n))/((k + m + n)*(k + m + n -1)); 
}