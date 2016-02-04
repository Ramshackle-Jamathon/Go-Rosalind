package tests
import (
    "testing"
    "io/ioutil"
    "github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkIEV(b *testing.B) {
        dat, err := ioutil.ReadFile("../inputs/iev.in")
        check(err)

        var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionIEV(string(dat))
        }
        result = r

}