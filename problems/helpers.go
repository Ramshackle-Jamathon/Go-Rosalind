package problems

import(
	"bufio"
	"fmt"
	"io"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func binomial(n, k float64) float64{
	if k > n - k {
		k = n - k
	}
	accum := 1.0
	for i := 1.0; i < k + 1.0; i++{
		accum = accum * (n - (k - i))
		accum = accum / i
	}
	return accum
}

/*
Check if slice variable list countains string variable a
*/
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
/*
handles duplicates like a champ
*/
func reverseMap(m map[string]string) map[string][]string {
    n := make(map[string][]string)
    for k, v := range m {
    	if val, ok := n[v]; ok {
    		n[v] = append(val, k)
    	} else {
        	n[v] = []string{k}
    	}
    }
    return n
}

/*
PROTEIN -> Mass
*/
var monoisotopicMassMap = map[string]float64{
	"A": 71.03711,
	"C": 103.00919,
	"D": 115.02694,
	"E": 129.04259,
	"F": 147.06841,
	"G": 57.02146,
	"H": 137.05891,
	"I": 113.08406,
	"K": 128.09496,
	"L": 113.08406,
	"M": 131.04049,
	"N": 114.04293,
	"P": 97.05276,
	"Q": 128.05858,
	"R": 156.10111,
	"S": 87.03203,
	"T": 101.04768,
	"V": 99.06841,
	"W": 186.07931,
	"Y": 163.06333,
}
/*
RNA -> PROTEIN ENCODING MAP
*/
var proteinMap = map[string]string{
	"UUU": "F",      
	"CUU": "L",      
	"AUU": "I",      
	"GUU": "V",
	"UUC": "F",      
	"CUC": "L",      
	"AUC": "I",      
	"GUC": "V",
	"UUA": "L",      
	"CUA": "L",      
	"AUA": "I",      
	"GUA": "V",
	"UUG": "L",      
	"CUG": "L",      
	"AUG": "M",      
	"GUG": "V",
	"UCU": "S",      
	"CCU": "P",      
	"ACU": "T",      
	"GCU": "A",
	"UCC": "S",      
	"CCC": "P",      
	"ACC": "T",      
	"GCC": "A",
	"UCA": "S",      
	"CCA": "P",      
	"ACA": "T",      
	"GCA": "A",
	"UCG": "S",      
	"CCG": "P",      
	"ACG": "T",      
	"GCG": "A",
	"UAU": "Y",      
	"CAU": "H",      
	"AAU": "N",      
	"GAU": "D",
	"UAC": "Y",      
	"CAC": "H",      
	"AAC": "N",      
	"GAC": "D",
	"UAA": "Stop",   
	"CAA": "Q",      
	"AAA": "K",      
	"GAA": "E",
	"UAG": "Stop",   
	"CAG": "Q",      
	"AAG": "K",      
	"GAG": "E",
	"UGU": "C",      
	"CGU": "R",      
	"AGU": "S",      
	"GGU": "G",
	"UGC": "C",      
	"CGC": "R",      
	"AGC": "S",      
	"GGC": "G",
	"UGA": "Stop",   
	"CGA": "R",      
	"AGA": "R",      
	"GGA": "G",
	"UGG": "W",      
	"CGG": "R",      
	"AGG": "R",      
	"GGG": "G",
}


/*
FASTA
*/
type String struct {
	Description string
	Data string
}

type FastaError string
func (f FastaError) Error() string {
	return string(f)
}

type lineReader struct {
	*bufio.Reader
	Lineno int
}

func (r *lineReader) ReadLine() (string, error) {
	line, err := r.ReadString('\n')
	r.Lineno++
	return line, err
}

func (lr *lineReader) ReadFasta() (ret String, retErr error) {
	newFastaError := func(s string) FastaError {
		return FastaError(fmt.Sprintf("Line %v: %s", lr.Lineno, s))
	}

	line, err := lr.ReadLine()
	if err == io.EOF {
		retErr = FastaError("EOF before description")
		return
	}
	if err != nil {
		retErr = newFastaError(err.Error())
		return
	}

	line = strings.TrimRight(line, "\n")

	if line[0] != '>' {
		retErr = newFastaError("Missing > before description")
		return
	}

	desc := line[1:]
	data := []byte{}
	for {
		buf, err := lr.Reader.Peek(1)
		done := len(buf) == 0 && err == io.EOF
		if len(buf) > 0 && (buf[0] == '\n' || buf[0] == '>') {
			done = true
		}
		if done {
			return String{Description: desc, Data: string(data)}, err
		}

		line, err := lr.ReadLine()
		if err != nil {
			return String{Description: desc, Data: string(data)},
			       newFastaError(err.Error())
		}

		line = strings.TrimRight(line, "\n")
		data = append(data, line...)
	}
	return
}

// ReadFasta reads a single String and assumes that nothing else follows it
// in the reader's data.
func ReadFasta(r io.Reader) (String, error) {
	lr := lineReader{bufio.NewReader(r), 0}
	return lr.ReadFasta()
}

func ReadAllFasta(r io.Reader) ([]String, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	lr := lineReader{br, 0}
	a := []String{}
	for {
		s, err := lr.ReadFasta()
		a = append(a, s)
		if err == io.EOF {
			return a, nil
		}
		if err != nil {
			return a, err
		}
	}
	return nil, nil
}

func WriteFasta(w io.Writer, s String) error {
	_, err := fmt.Fprintf(w, ">%s\n", s.Description)
	if err != nil {
		return err
	}
	for i := 0; i < len(s.Data); i += 80 {
		t := s.Data[i:]
		if len(t) > 80 {
			t = t[:80]
		}
		_, err = fmt.Fprintln(w, t)
		if err != nil {
			return err
		}
	}
	return nil
}