package main

import (
    "runtime"
    "fmt"
)

func main(){
   num := runtime.NumCPU()
   fmt.Println(num)
}
