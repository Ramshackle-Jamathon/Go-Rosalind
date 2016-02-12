package problems

import (
    "io/ioutil"
    "strconv"
)

func PRTM() {
    dat, err := ioutil.ReadFile("inputs/prtm.in")
    check(err)
    result := SolutionPRTM(string(dat))
    err = ioutil.WriteFile("inputs/prtm.out", []byte(strconv.FormatFloat(result, 'f', 6, 64)), 0644)
    check(err)
}

func SolutionPRTM(s string) float64{
    return getProteinMass(s)
}

func getProteinMass(proteins string) float64{
    mass := 0.0
    for _, c := range proteins {
        a, ok := monoisotopicMassMap[string(c)]
        if(ok){
            mass = mass + a
        }
    }
    return mass
}