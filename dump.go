package tcp_scanner

import(
    "strconv"
)

// Dump scan result object.
type Dump struct {
    host  string
    ports []*Port
}

// Host return hostname for scan.
func (d *Dump) Host() string {
    return d.host
}

// OpenPorts return list of ports thats are open.
func (d *Dump) OpenPorts() (ports []int) {
    for _, port := range d.ports{
        if port.isOpen {
            ports = append(ports, int(port.port))
        }
    }
    return
}

// OpenAddrs return full addr thats are open.
func (d *Dump) OpenAddrs() (addr []string) {
    for _, port := range d.ports{
        if port.isOpen{
            addr = append(addr, d.host + ":" + strconv.FormatUint(uint64(port.port), 10))
        }
    }
    return
}

// All return all dump.
func (d *Dump) All() map[string]bool {
    r := make(map[string]bool, len(d.ports))

    for _, port := range d.ports{
        r[d.host + ":" + strconv.FormatUint(uint64(port.port), 10)] = port.isOpen
    }
    return r
}                           
