package main

import (
	"fmt"
    "os"
    "github.com/rekby/multiex"
)

func f1(){
    fmt.Println("Func 1")
    fmt.Println(os.Args)
}

func f2(){
    fmt.Println("Func 2")
    fmt.Println(os.Args)
}

func main() {
    multiex.Register(multiex.ExecutorDescribe{Name: "test-f1", Function: f1})
    multiex.Register(multiex.ExecutorDescribe{Name: "test-f2", Function: f2})
    multiex.Main()
}
