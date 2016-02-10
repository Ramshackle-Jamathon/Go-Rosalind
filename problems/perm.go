package problems

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func PERM() {
    dat, err := ioutil.ReadFile("inputs/perm.in")
    check(err)
    result := SolutionPERM(string(dat))
    err = ioutil.WriteFile("inputs/perm.out", []byte(result), 0644)
    check(err)
}

func SolutionPERM(s string) string{
    input := strings.Split(s, " ")
    n, err := strconv.Atoi(input[0])
    check(err)
    inputs := make([]int, n) 
    for i := 1; i <= n; i++ {
        inputs[i-1] = i
    }
    result := ""
    count := 0
    for perm := range GeneratePermutations(inputs) {
        for _, i := range perm{
            result = result + strconv.Itoa(i) + " "
        }
        result = result + "\n"
        count++
    }
    result = strconv.Itoa(count) + "\n" + result
    return result
}
func GeneratePermutations(data []int) <-chan []int {  
    c := make(chan []int)
    go func(c chan []int) {
        defer close(c) 
        permutate(c, data)
    }(c)
    return c 
}
func permutate(c chan []int, inputs []int){
    output := make([]int, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int, len(inputs))
        copy(output, inputs)
        c <- inputs
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}
/*
package main

import "fmt"

func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}

func main() {
    orig := []int{1, 2, 3, 4, 5, 6, 7}
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
        fmt.Println(getPerm(orig, p))
    }
}
*/