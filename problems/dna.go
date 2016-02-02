package problems

import (
    "strings"
    "io/ioutil"
    "strconv"
)

func DNA() {
    dat, err := ioutil.ReadFile("inputs/dna.in")
    check(err)
    result := SolutionDNA(string(dat))
    err = ioutil.WriteFile("inputs/dna.out", []byte(result), 0644)
    check(err)

}

func SolutionDNA(dna string) string{
	return strconv.Itoa(strings.Count(dna, "A")) + " " + strconv.Itoa(strings.Count(dna, "C")) + " " + strconv.Itoa(strings.Count(dna, "G")) + " " + strconv.Itoa(strings.Count(dna, "T"))

}
