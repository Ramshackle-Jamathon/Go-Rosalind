package problems

import (
    "io/ioutil"
    "math/big"
    "fmt"
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
	posRNA.Mul(posRNA, big.NewInt(int64(len(reverseProteinMap["Stop"]))))
	for _, protein := range proteins{
		fmt.Println(posRNA)
		posRNA = posRNA.Mod(posRNA.Mul(posRNA, big.NewInt(int64(len(reverseProteinMap[string(protein)])))), big.NewInt(1000000))
	}
	fmt.Println("nice")
	fmt.Println(len(reverseProteinMap["Stop"]))
	//Stop Codon
	fmt.Println(posRNA)
    return posRNA
}