package tcp_scanner

import (
    "fmt"
    "net"
)

type Dump struct {
    items []*item
}

func (d *Dump) Opened() (ports []int) {
    for _, item := range d.items{
        if item.isOpen {
            ports = append(ports, item.port)
        }
    }
    return
}


type item struct {
    isOpen bool
    port   int
}

// Scan ...
func Scan(host, ports string) (*Dump, error) {
    if !isHost(host) {
        return nil, fmt.Errorf("host '%s' is invalid", host)
    }
    scanPorts, err := parsePortsForScan(ports)
    if err != nil {
        return nil, fmt.Errorf("cannot parse ports for scan from '%s', got %s", ports, err.Error())
    }

    dump := &Dump{}
    for _, port := range scanPorts {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
        if err != nil {
            dump.items = append(dump.items, &item{
                isOpen: false,
                port:   port,
            })
            conn.Close()
            continue
        }
        dump.items = append(dump.items, &item{
            isOpen: true,
            port:   port,
        })
        conn.Close()
    }
    return dump, nil
}
