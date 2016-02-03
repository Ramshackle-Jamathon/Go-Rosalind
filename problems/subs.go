package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func SUBS() {
    dat, err := ioutil.ReadFile("inputs/subs.in")
    check(err)
    result := SolutionSUBS(string(dat))
    err = ioutil.WriteFile("inputs/subs.out", []byte(result), 0644)
    check(err)
}

func SolutionSUBS(s string) string{
    inputs := strings.Split(s, "\n")
    locs := getMotif(inputs[0], inputs[1])
    var res = ""
    for _, entry := range locs {
        res = res + strconv.Itoa(entry) + " "
    }
    return res
}

func getMotif(data, sub string) []int{
    var r []int
    for i := len(data) + 1; i > 0; i-- {
        loc := strings.LastIndex(data, sub)
        if(loc == -1){
            return r
        }
        r = append([]int{loc + 1}, r...)
        data = data[:loc + len(sub) - 1]
    }

    return r
}