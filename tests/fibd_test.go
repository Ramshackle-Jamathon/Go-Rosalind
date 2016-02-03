package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkFIBD(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/fibd.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionFIBD(string(dat))
        }
        result = r

}