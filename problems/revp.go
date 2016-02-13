package problems

import (
    "io/ioutil"
    "os"
    "strconv"
)

func REVP() {
    f, err := os.Open("inputs/revp.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionREVP(fasta[0])
    err = ioutil.WriteFile("inputs/revp.out", []byte(result), 0644)
    check(err)
}

func SolutionREVP(fastaInput String) string{
    revpData := getReverse(fastaInput.Data, 4, 12)
    result := ""
    for _, i := range revpData {
        result =  result + strconv.Itoa(i[0]) + " " + strconv.Itoa(i[1]) + "\n"
    }
    return result
}


func getReverse(data string, min int, max int) [][]int{
    var maxRatio [][]int
    substring := ""
    for i := 0; i < len(data); i++{
        for a := i + min; a <= i + max && a <= len(data); a++ {
            substring = data[i:a]
            if substring == SolutionREVC(substring) {
                maxRatio = append(maxRatio, []int{i + 1, a - i})
            }
        }
    }
    return maxRatio
}
