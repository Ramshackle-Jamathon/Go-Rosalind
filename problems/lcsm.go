package problems

import (
    "io/ioutil"
    "os"
    "strings"
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
    var strings []string
    for _, entry := range fasta {
        strings = append(strings, entry.Data)
    }
    return terribleMLCS(strings)
}

/*
a complete POS solution
*/
func terribleMLCS(inputs []string) string{
    res := ""
    if len(inputs) > 0  && len(inputs[0]) > 0 {
        for i := 0; i < len(inputs[0]); i++{
            for j := 0; j < len(inputs[0])-i+1; j++{
                check := true
                for _, input := range inputs{
                    if strings.Index(input, inputs[0][i:i+j]) == -1 {
                        check = false
                        break
                    }
                }
                if j > len(res) && check {
                    res = inputs[0][i:i+j]
                }
            }
        }
    }
    return res
}

/*
//LCS ITERATIVE PseudoCode
function LCSLength(X[1..m], Y[1..n])
    C = array(0..m, 0..n)
    for i := 0..m
       C[i,0] = 0
    for j := 0..n
       C[0,j] = 0
    for i := 1..m
        for j := 1..n
            if X[i] = Y[j]
                C[i,j] := C[i-1,j-1] + 1
            else
                C[i,j] := max(C[i,j-1], C[i-1,j])
    return C[m,n]
*/