package tcp_scanner

import(
    "fmt"
    "net"
    "strconv"
)

// Port ...
type Port struct {
    port   uint16
    isOpen bool
}

// NewDial create dial and return result as Port object.
func NewDial[P portNumber](host string, port P) (p *Port) {
    fmt.Println("port", port)
    p = &Port{
        port:   uint16(port),
        isOpen: true,
    }
    
    conn, err := net.Dial("tcp", host + ":" + strconv.Itoa(int(port)))
    defer conn.Close()
    if err != nil {
        p.isOpen = false
    }
    return
}