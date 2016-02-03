package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkPROT(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/prot.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionPROT(string(dat))
        }
        result = r

}