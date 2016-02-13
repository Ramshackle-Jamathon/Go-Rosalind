package problems

import (
    "io/ioutil"
    "os"
)

func SPLC() {
    f, err := os.Open("inputs/splc.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionSPLC(fasta[0])
    err = ioutil.WriteFile("inputs/splc.out", []byte(result), 0644)
    check(err)
}

func SolutionSPLC(fastaInput String) string{
    return ""
}