package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkLEXF(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/lexf.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionLEXF(string(dat))
        }
        result = r
}