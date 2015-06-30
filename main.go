package main

import (
	"fmt"
	"github.com/rekby/multiex"
	"os"

    _ "github.com/rekby/multiex-example/side"
)

func f1() {
	fmt.Println("Func 1")
	fmt.Println(os.Args)
}

func f2() {
	fmt.Println("Func 2")
	fmt.Println(os.Args)
}

func main() {
	multiex.Register(multiex.ExecutorDescribe{Name: "test-f1", Function: f1, Describtion: "Test function 1"})
	multiex.Register(multiex.ExecutorDescribe{Name: "test-f2", Function: f2, Describtion: "Test function 2"})
	multiex.Main()
}
