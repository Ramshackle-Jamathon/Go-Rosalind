package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkREVC(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/revc.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionREVC(string(dat))
        }
        result = r

}