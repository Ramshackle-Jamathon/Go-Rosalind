package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkSUBS(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/subs.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionSUBS(string(dat))
        }
        result = r

}