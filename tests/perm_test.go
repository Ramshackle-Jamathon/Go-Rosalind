package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkPERM(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/perm.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionPERM(string(dat))
        }
        result = r
}