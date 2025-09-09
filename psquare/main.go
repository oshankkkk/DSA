package main

import "fmt"
func main() {
    num := 12
    array := []int{}
    for i := 1; i*i <= num; i++ {
        array = append(array, i*i) // squares
    }
    a := psqrec(num, array)
    fmt.Println(a)
}

