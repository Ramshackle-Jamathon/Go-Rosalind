package problems

import (
    "io/ioutil"
    "os"
    "strconv"
)

func CONS() {
    f, err := os.Open("inputs/cons.in")
    check(err)
    fasta, err := ReadAllFasta(f)
    check(err)
    result := SolutionCONS(fasta)
    err = ioutil.WriteFile("inputs/cons.out", []byte(result), 0644)
    check(err)
}

func SolutionCONS(fasta []String) string{
    /*
        initializing arrays
    */
    a := make([]int, 0, len(fasta[0].Data))
    c := make([]int, 0, len(fasta[0].Data))
    g := make([]int, 0, len(fasta[0].Data))
    t := make([]int, 0, len(fasta[0].Data))
    for _, _ = range fasta[0].Data {
        a = append(a, 0)
        c = append(c, 0)
        g = append(g, 0)
        t = append(t, 0)
    }

    /*
        getting nt count
    */
    for _, entry := range fasta {
        for i, char := range entry.Data{
            switch char {
                case 'A':
                    a[i]++
                case 'C':
                    c[i]++
                case 'G':
                    g[i]++
                case 'T':
                    t[i]++
            }
        }
    }

    /*
        Generating consensus string
    */
    var consensus string
    for i := 0; i < len(a); i++ {
            switch {
                case a[i] >= c[i] && a[i] >= t[i] && a[i] >= g[i]:
                    consensus = consensus + "A"
                case c[i] >= a[i] && c[i] >= t[i] && c[i] >= g[i]:
                    consensus = consensus + "C"
                case g[i] >= c[i] && g[i] >= t[i] && g[i] >= a[i]:
                    consensus = consensus + "G"
                case t[i] >= c[i] && t[i] >= a[i] && t[i] >= g[i]:
                    consensus = consensus + "T"
            }
    }

    /*
        Preparing return
    */
    var result = consensus + "\nA:"
    for _, entry := range a {
        result = result + " " + strconv.Itoa(entry)
    }
    result = result + "\nC:"
    for _, entry := range c {
        result = result + " " + strconv.Itoa(entry)
    }
    result = result + "\nG:"
    for _, entry := range g {
        result = result + " " + strconv.Itoa(entry)
    }
    result = result + "\nT:"
    for _, entry := range t {
        result = result + " " + strconv.Itoa(entry)
    }
    return result
}

