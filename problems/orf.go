package problems

import (
    "io/ioutil"
    "os"
    "strings"
)

func ORF() {
    f, err := os.Open("inputs/orf.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionORF(fasta)
    err = ioutil.WriteFile("inputs/orf.out", []byte(result), 0644)
    check(err)
}

func SolutionORF(fasta []String) string{
    frames := getORFS(fasta[0].Data)
    result := ""
    for _, frame := range frames{
        result = result + frame + "\n"
    }
    return result
}

func getORFS(input string) []string{
    inputRComp := SolutionREVC(input)
    frames := []string{}
    frames = append(frames, strings.Replace(input, "T", "U", -1))
    frames = append(frames, strings.Replace(input[1:], "T", "U", -1))
    frames = append(frames, strings.Replace(input[2:], "T", "U", -1))
    frames = append(frames, strings.Replace(inputRComp, "T", "U", -1))
    frames = append(frames, strings.Replace(inputRComp[1:], "T", "U", -1))
    frames = append(frames, strings.Replace(inputRComp[2:], "T", "U", -1))
    result := []string{}
    for _, frame := range frames {
        for i := 0; i < len(frame) - 2; i = i + 3{
            if frame[i:i+3] == "AUG" {
                p := getProtein(frame[i:])
                if !stringInSlice(p, result) && p != "" {
                    result = append(result, p)
                }
            }
        }   
    }
    return result
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}