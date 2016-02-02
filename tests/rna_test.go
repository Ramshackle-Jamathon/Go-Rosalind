package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkRNA(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/rna.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionRNA(string(dat))
        }
        result = r

}