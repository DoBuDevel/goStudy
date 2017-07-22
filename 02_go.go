package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 100
    fmt.Println("초기 값 : ", i)

    zeroval(i)
    fmt.Println("zeroval : ", i)

    fmt.Println("%i : ", &i)

    zeroptr(&i)
    fmt.Println("zeroptr : ", i)

    fmt.Println("포인터 :", &i)

    fmt.Println("포인터 :", i)
    //fmt.Println(*i)
}
