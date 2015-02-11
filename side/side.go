package side

import (
    "fmt"
    "github.com/rekby/multiex"
)

func init(){
    multiex.Register(multiex.ExecutorDescribe{Name: "side", Function: Main})
}

func Main(){
    fmt.Println("Main function, imported as side effect from init func")
}