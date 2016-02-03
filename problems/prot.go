package problems

import (
    "io/ioutil"
)

func PROT() {
    dat, err := ioutil.ReadFile("inputs/prot.in")
    check(err)
    result := SolutionPROT(string(dat))
    err = ioutil.WriteFile("inputs/prot.out", []byte(result), 0644)
    check(err)
}

func SolutionPROT(s string) string{
    return getProtein(s)
}

func getProtein(rna string) string{
    var protein = ""
    var res = ""
    for i, c := range rna {
        res = res + string(c)
        if i > 0 && (i+1)%3 == 0 {
            i, ok := proteinMap[res]
            if(i == "Stop"){
                return protein
            }
            if(ok){
                protein = protein + i
            }
            res = ""
        }
    }
    return protein
}