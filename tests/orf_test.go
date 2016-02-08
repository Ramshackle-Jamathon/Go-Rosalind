package tests
import (
    "testing"
    "os"
    "github.com/Ramshackle-Jamathon/Rosalind/Problems"
)


func BenchmarkORF(b *testing.B) {
        f, err := os.Open("../inputs/orf.in")
        check(err)

        fasta, err := problems.ReadAllFasta(f)
        check(err)

        var r string
        for n := 0; n < b.N; n++ {
            r = problems.SolutionORF(fasta)
        }
        result = r
}