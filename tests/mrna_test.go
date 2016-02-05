package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkMRNA(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/mrna.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionMRNA(string(dat))
        }
        result = r

}