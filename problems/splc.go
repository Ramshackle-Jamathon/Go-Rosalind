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
    result := SolutionSPLC(fasta)
    err = ioutil.WriteFile("inputs/splc.out", []byte(result), 0644)
    check(err)
}

func SolutionSPLC(fastaInput []String) string{
    dna := fastaInput[0].Data
    for _, intron := range fastaInput[1:]{
        for _,i := range getMotif(dna, intron.Data){
            dna = dna[:i-1] + dna[i -1 + len(intron.Data):]
        }
    }
    result := SolutionPROT(SolutionRNA(dna))
    return result
}