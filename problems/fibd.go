package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
    "math/big"
)

func FIBD() {
    dat, err := ioutil.ReadFile("inputs/fibd.in")
    check(err)
    result := SolutionFIBD(string(dat))
    err = ioutil.WriteFile("inputs/fibd.out", []byte(result), 0644)
    check(err)
}

func SolutionFIBD(s string) string{
    input := strings.Split(s, " ")
    n, err := strconv.Atoi(input[0])
    check(err)
    m, err := strconv.Atoi(input[1])
    check(err)
    return calcMortalRabbits(n, m).String()
}
func calcMortalRabbits(n, m int) *big.Int {
    r := make([]*big.Int, 0, n)
    r = append(r, big.NewInt(1), big.NewInt(1)) 
    for i := 2; i < n; i++ {
        temp := big.NewInt(0)
        if i < m {
            temp.Add(r[i-1], r[i-2])
            r = append(r,temp) 
        } else if i == m || i == m+1 {
            temp.Add(r[i-1], r[i-2])
            temp.Sub(temp, big.NewInt(1))
            r = append(r,temp) 
        } else {
            temp.Add(r[i-1], r[i-2])
            temp.Sub(temp, r[i-m-1])
            r = append(r,temp) 
        }
    }
    return r[n-1]
}


