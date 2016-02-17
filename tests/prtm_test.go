package tests
import (
    "testing"
    "io/ioutil"
    "strconv"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkPRTM(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/prtm.in")
	    check(err)

	    var r float64
        for n := 0; n < b.N; n++ {
            r = problems.SolutionPRTM(string(dat))
        }
        result = strconv.FormatFloat(r, 'f', 6, 64)

}