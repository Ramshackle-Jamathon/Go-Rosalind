package problems

import (
    "io/ioutil"
    "os"
    "strconv"
)

func LCSM() {
    f, err := os.Open("inputs/lcsm.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionLCSM(fasta)
    err = ioutil.WriteFile("inputs/lcsm.out", []byte(result), 0644)
    check(err)
}

func SolutionLCSM(fasta []String) string{
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
