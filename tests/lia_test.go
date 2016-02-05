package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkLIA(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/lia.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionLIA(string(dat))
        }
        result = r

}