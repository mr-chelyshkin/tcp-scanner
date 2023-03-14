package main

import (
    "fmt"
    "flag"
    "os"
    "strconv"

    scanner "github.com/mr-chelyshkin/tcp-scanner"
)

func main() {
    threads := flag.Int("threads", 50, "an int")
    ports := flag.String("ports", "1024...65535", "a string")
    host := flag.String("host", "", "a string")
    flag.Parse()

    sc, err := scanner.Scan(*host, *ports, *threads)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for host, status := range sc.All() {
        fmt.Println(host + ": " + strconv.FormatBool(status))
    }
    os.Exit(0)
}