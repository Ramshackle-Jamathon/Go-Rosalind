package problems

import (
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

func GC() {
    f, err := os.Open("inputs/gc.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionGC(fasta)
    err = ioutil.WriteFile("inputs/gc.out", []byte(result), 0644)
    check(err)
}

func SolutionGC(fasta []String) string{
    var maxRatio = 0.0
    var maxName = ""
    for _, entry := range fasta {
        var ratio = ratio(entry.Data)
        if(ratio > maxRatio){
            maxRatio = ratio
            maxName = entry.Description
        }
    }
    return maxName + " " + strconv.FormatFloat(maxRatio, 'f', -1, 64)
}

func ratio(dna string) float64{
    return float64(strings.Count(dna, "C") + strings.Count(dna, "G")) / float64(len([]rune(dna))) * 100.0
}