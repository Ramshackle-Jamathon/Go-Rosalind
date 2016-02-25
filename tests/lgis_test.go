package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkLGIS(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/lgis.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionLGIS(string(dat))
        }
        result = r
}