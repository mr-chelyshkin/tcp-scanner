package tcp_scanner

import(
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
    p = &Port{
        port:   uint16(port),
        isOpen: true,
    }
    
    conn, err := net.Dial("tcp", host + ":" + strconv.Itoa(int(port)))
    if err != nil {
        p.isOpen = false
        return
    }
    conn.Close()
    return
}