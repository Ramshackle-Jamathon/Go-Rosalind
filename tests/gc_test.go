package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkGC(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/gc.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionGC(string(dat))
        }
        result = r
}