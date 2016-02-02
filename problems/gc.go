package problems

import (
    "io/ioutil"
    "os"
    "fmt"
)

func GC() {
    result := SolutionGC("inputs/gc.in")
    fmt.Println(result)
    err := ioutil.WriteFile("inputs/gc.out", []byte(result), 0644)
    check(err)
}

func SolutionGC(filePath string) string{
    f, err := os.Open(filePath)
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    var max = 0.0
    var maxName = ""
    for dna in fasta{
        var test := ration(dna)
        if(test > max){
            max = test
            maxName = dna
        }
    }
    return maxName + " " + strconv.ParseFloat(max, 64)
