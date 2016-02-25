package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func LEXF() {
    dat, err := ioutil.ReadFile("inputs/lexf.in")
    check(err)
    result := SolutionLEXF(string(dat))
    err = ioutil.WriteFile("inputs/lexf.out", []byte(result), 0644)
    check(err)
}

func SolutionLEXF(s string) string{
    lines := strings.Split(s, "\n")
    alphabet := strings.Split(lines[0], " ")
    k, err := strconv.Atoi(lines[1])
    check(err)
    result := ""
    for combo := range generateCombinations(alphabet, k) {
        result = result + strings.Join(combo, "") + "\n"
    }
    return result
}

func generateCombinations(data []string, length int) <-chan []string {  
    c := make(chan []string)
    // Starting a separate goroutine that will create all the combinations,
    // feeding them to the channel c
    go func(c chan []string) {
        defer close(c) // Once the iteration function is finished, we close the channel
        combos(c, []string{}, data, length) // We start by feeding it an empty string
    }(c)
    return c // Return the channel to the calling function
}
// Combos adds a element to the combination to create a new combination.
// This new combination is passed on to the channel before we call Combos once again
// to add yet another element to the new combination in case length allows it
func combos(c chan []string, combo []string, data []string, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []string
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]string, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combos(c, newCombo, data, length-1)
    }
}