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
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}
/*
 The Countdown QuickPerm Algorithm:

   let a[] represent an arbitrary list of objects to permute
   let N equal the length of a[]
   create an integer array p[] of size N+1 to control the iteration     
   initialize p[0] to 0, p[1] to 1, p[2] to 2, ..., p[N] to N
   initialize index variable i to 1
   while (i < N) do {
      decrement p[i] by 1
      if i is odd, then let j = p[i] otherwise let j = 0
      swap(a[j], a[i])
      let i = 1
      while (p[i] is equal to 0) do {
         let p[i] = i
         increment i by 1
      } // end while (p[i] is equal to 0)
   } // end while (i < N)



a = ['a', 'b', 'c', 'd']

N = len(a)
p = range(0, N + 1)
i = 1
print a 
while i < N:
    p[i] -= 1
    if i % 2 == 1:
        j = p[i]
    else:
        j = 0
    a[j], a[i] = a[i], a[j]
    print a
    i = 1
    while p[i] == 0:
        p[i] = i
        i += 1

        */