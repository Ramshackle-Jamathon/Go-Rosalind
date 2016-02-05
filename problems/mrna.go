package problems

import (
    "io/ioutil"
    "math/big"
)

func MRNA() {
    dat, err := ioutil.ReadFile("inputs/mrna.in")
    check(err)
    result := SolutionMRNA(string(dat))
    err = ioutil.WriteFile("inputs/mrna.out", []byte(result), 0644)
    check(err)
}

func SolutionMRNA(protein string) string{
	return inferRna(protein).String()
}

func inferRna(proteins string) *big.Int {
	var posRNA = big.NewInt(1)
	reverseProteinMap := reverseMap(proteinMap)
	for _, protein := range proteins{
		posRNA.Mul(posRNA, big.NewInt(int64(len(reverseProteinMap[string(protein)]))))
	}
	posRNA.Mul(posRNA, big.NewInt(int64(len(reverseProteinMap["Stop"]))))
	posRNA.Mod(posRNA, big.NewInt(1000000))
    return posRNA
}