package tests

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var result string