package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkHAMM(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/hamm.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionHAMM(string(dat))
        }
        result = r

}