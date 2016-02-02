package problems

import (
    "strings"
    "io/ioutil"
)

func RNA() {
    dat, err := ioutil.ReadFile("inputs/rna.in")
    check(err)
    result := SolutionRNA(string(dat))
    err = ioutil.WriteFile("inputs/rna.out", []byte(result), 0644)
    check(err)

}

func SolutionRNA(dna string) string{
	return strings.Replace(dna, "T", "U", -1)
}
