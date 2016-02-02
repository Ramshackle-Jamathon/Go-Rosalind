package tests
import (
    "testing"
    "os"
	"github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkGC(b *testing.B) {
        f, err := os.Open("../inputs/gc.in")
        check(err)

        fasta, err := problems.ReadAllFasta(f)
        check(err)

	    var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionGC(fasta)
        }
        result = r
}