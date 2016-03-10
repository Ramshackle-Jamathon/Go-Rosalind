package problems

import (
    "io/ioutil"
    "os"
    "strings"
)

func LONG() {
    f, err := os.Open("inputs/long.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionLONG(fasta)
    err = ioutil.WriteFile("inputs/long.out", []byte(result), 0644)
    check(err)
}

func SolutionLONG(fasta []String) string{
    var strings []string
    for _, entry := range fasta {
        strings = append(strings, entry.Data)
    }
    return ""
}