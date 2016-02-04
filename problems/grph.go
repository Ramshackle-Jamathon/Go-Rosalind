package problems

import (
    "io/ioutil"
    "os"
)

type Edge struct {
    Tail, Head string
}

func GRPH() {
    f, err := os.Open("inputs/grph.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionGRPH(fasta)
    err = ioutil.WriteFile("inputs/grph.out", []byte(result), 0644)
    check(err)
}

func SolutionGRPH(fasta []String) string{
    edges := kEdges(fasta, 3)
    result := ""
    for _, entry := range edges {
        result = result + entry.Tail + " " + entry.Head + "\n" 
    }
    return result
}

func kEdges(data []String, k int) []Edge{
    edges := []Edge{} 
    for dnaPair := range GenerateCombinations(data, 2) {
        //fmt.Println(dnaPair)
        if(dnaPair[0].Data != dnaPair[1].Data){
            if kOverlap(dnaPair[0].Data, dnaPair[1].Data, k) {
                edges = append(edges, Edge{Tail: dnaPair[0].Description, Head: dnaPair[1].Description})
            }
        }
    }
    return edges
}

func kOverlap(s1 string, s2 string, k int) bool{
    return s1[len(s1)-k:] == s2[0:k]
}

func GenerateCombinations(data []String, length int) <-chan []String {
    c := make(chan []String)
    // Starting a separate goroutine that will create all the combinations,
    // feeding them to the channel c
    go func(c chan []String) {
        defer close(c) // Once the iteration function is finished, we close the channel
        Combos(c, []String{}, data, length) // We start by feeding it an empty string
    }(c)
    return c // Return the channel to the calling function
}
// Combos adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func Combos(c chan []String, combo []String, data []String, length int) {
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []String
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets below k
        if(length == 1){
            c <- newCombo

        }
        Combos(c, newCombo, data, length-1)
    }
}