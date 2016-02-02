package problems

import (
    "io/ioutil"
)

func REVC() {
    dat, err := ioutil.ReadFile("inputs/revc.in")
    check(err)
    result := SolutionREVC(string(dat))
    err = ioutil.WriteFile("inputs/revc.out", []byte(result), 0644)
    check(err)

}

func SolutionREVC(dna string) string{
	runes := []rune(dna)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = getCompliment(runes[j]), getCompliment(runes[i])
    }
    return string(runes)
}

func getCompliment(r rune) rune{
	switch r {
    case 'A':
        return 'T'
    case 'C':
        return 'G'
    case 'T':
        return 'A'
    case 'G':
        return 'C'
    }
    return r
}