package tcp_scanner

import (
    "fmt"
    "net"
)

type Dump struct {
    items []*item
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

    var dump *Dump
    for _, port := range scanPorts {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
        if err != nil {
            return nil, err
        }
        dump.items = append(dump.items, &item{port: port})
        conn.Close()
    }
    return dump, nil
}
