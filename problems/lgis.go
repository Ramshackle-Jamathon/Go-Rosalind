package problems

import (
    "io/ioutil"
    //"strings"
    //"strconv"
)

func LGIS() {
    dat, err := ioutil.ReadFile("inputs/lgis.in")
    check(err)
    result := SolutionLGIS(string(dat))
    err = ioutil.WriteFile("inputs/lgis.out", []byte(result), 0644)
    check(err)
}

func SolutionLGIS(s string) string{
    //lines := strings.Split(s, "\n")
    //perm := strings.Split(lines[1], " ")
    //n, err := strconv.Atoi(lines[0])
    //check(err)
    result := ""
    return result
}