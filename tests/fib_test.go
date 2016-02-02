package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkFIB(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/fib.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionFIB(string(dat))
        }
        result = r

}