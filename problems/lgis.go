package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func LGIS() {
    dat, err := ioutil.ReadFile("inputs/lgis.in")
    check(err)
    result := SolutionLGIS(string(dat))
    err = ioutil.WriteFile("inputs/lgis.out", []byte(result), 0644)
    check(err)
}

func SolutionLGIS(s string) string{
    lines := strings.Split(s, "\n")
    perm := strings.Split(lines[1], " ")
    input := make([]int, len(perm))
    var err error
    for i,e := range perm {
        input[i], err = strconv.Atoi(e)
        check(err)
    }
    result := ""
    for _, e := range LIS(input){
        result = result + strconv.Itoa(e) + " "
    }
    result = result + "\n"
    for _, e := range LDS(input){
        result = result + strconv.Itoa(e) + " "
    }
    return result
}

func LIS (s []int) []int{
    l := 0;
    P := make([]int, len(s));
    M := make([]int, len(s) + 1);
    for i := 0; i < len(s); i++ {
        lo := 1;
        hi := l;
        for lo <= hi {
            mid := int(math.Ceil(float64((lo+hi)/2)))
            if(s[M[mid]] < s[i]){
                lo = mid + 1
            } else {
                hi = mid-1
            }
        }
        newL := lo;

        P[i] = M[newL-1];
        M[newL] = i;

        if newL > l{
            l = newL
        }
    }
    S := make([]int, l);
    k := M[l]
    for i := l-1; i >= 0; i--{
        S[i] = s[k]
        k = P[k]
    } 
    return S
}
func LDS (s []int) []int{
    l := 0;
    P := make([]int, len(s));
    M := make([]int, len(s) + 1);
    for i := 0; i < len(s); i++ {
        lo := 1;
        hi := l;
        for lo <= hi {
            mid := int(math.Ceil(float64((lo+hi)/2)))
            if(s[M[mid]] > s[i]){
                lo = mid + 1
            } else {
                hi = mid-1
            }
        }
        newL := lo;

        P[i] = M[newL-1];
        M[newL] = i;

        if newL > l{
            l = newL
        }
    }
    S := make([]int, l);
    k := M[l]
    for i := l-1; i >= 0; i--{
        S[i] = s[k]
        k = P[k]
    } 
    return S
}
