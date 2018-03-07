package main

import "fmt"
import "os"

func main() {
    //fmt.Printf("hello, world\n")
    var PORT string = os.Getenv("PORT")
    var IP string = os.Getenv("IP")
    
    fmt.Printf("%s\n%s\n", PORT, IP)
}