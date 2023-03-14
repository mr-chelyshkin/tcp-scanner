package tcp_scanner
import ("fmt")
// Dump ...
type Dump struct {
    host  string
    ports []*Port
}

// Host return hostname for scan.
func (d *Dump) Host() string {
    return d.host
}

// OpenPorts return list of ports that are open.
func (d *Dump) OpenPorts() (ports []int) {
    fmt.Println(d.ports)
    for _, port := range d.ports{
        if port.isOpen {
            ports = append(ports, int(port.port))
        }
    }
    return
}
