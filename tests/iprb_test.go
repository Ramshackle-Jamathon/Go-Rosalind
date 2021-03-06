package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkIPRB(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/iprb.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionIPRB(string(dat))
        }
        result = r

}