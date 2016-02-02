package tests
import (
    "testing"
    "io/ioutil"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkDNA(b *testing.B) {
	    dat, err := ioutil.ReadFile("../inputs/dna.in")
	    check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionDNA(string(dat))
        }
        result = r

}