package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func LIA() {
    dat, err := ioutil.ReadFile("inputs/lia.in")
    check(err)
    result := SolutionLIA(string(dat))
    err = ioutil.WriteFile("inputs/lia.out", []byte(result), 0644)
    check(err)
}

func SolutionLIA(s string) string{
    inputs := strings.Split(s, " ")
    n0, err := strconv.ParseFloat(inputs[0], 64)
    check(err)
    n1, err := strconv.ParseFloat(inputs[1], 64)
    check(err)
    return strconv.FormatFloat(calcIndependentAlleles(n0, n1), 'f', -1, 64)
}

func calcIndependentAlleles(k, n float64) float64{
    sum := 0.0
    for i := 0.0; i < n; i++{
        sum = sum + binomial(math.Pow(2.0,k), i) * math.Pow(0.25,i) * math.Pow(0.75,(math.Pow(2.0,k) - i))
    }
    return 1.0 - sum
}
// return binomial(2**k, n) * 0.25**n * 0.75**(2**k - n)
// return 1 - sum([P(n, k) for n in range(N)])