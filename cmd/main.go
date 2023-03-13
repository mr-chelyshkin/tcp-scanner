package main

import (
    "fmt"
    "flag"
    "os"

    scanner "github.com/mr-chelyshkin/tcp-scanner"
)

func main() {
    ports := flag.String("ports", "1024...65535", "a string")
    host := flag.String("host", "", "a string")
    flag.Parse()

    if _, err := scanner.Scan(*host, *ports); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
}