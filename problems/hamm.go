package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func HAMM() {
    dat, err := ioutil.ReadFile("inputs/hamm.in")
    check(err)
    result := SolutionHAMM(string(dat))
    err = ioutil.WriteFile("inputs/hamm.out", []byte(result), 0644)
    check(err)
}

func SolutionHAMM(s string) string{
    inputs := strings.Split(s, "\n")
    count := uncommonRunes(inputs[0], inputs[1])
    return strconv.Itoa(count)
}

func uncommonRunes(s, t string) int {
    if len(s) > len(t) {
        s, t = t, s
    }
    count := 0
    runes := []rune(t)
    for i, entry := range s {
        if entry != runes[i] {
            count++
        }
    }
    return count
}