package tcp_scanner

import(
    "fmt"
    "strings"
    "strconv"
    "sync"
)

// Scan ports and return Dump oject with results.
func Scan(host, ports string) (*Dump, error) {
    if !isHost(host) {
        return nil, fmt.Errorf("host '%s' is invalid", host)
    }
    scanPorts, err := parsePortsForScan(ports)
    if err != nil {
        return nil, fmt.Errorf("cannot parse ports for scan from '%s', got %s", ports, err.Error())
    }
    dumpPorts := make([]*Port, len(scanPorts))

    wg := &sync.WaitGroup{}
    sem := make(chan struct{}, 5)
    for i, port := range scanPorts {
        sem <- struct{}{}
        wg.Add(1)
        i := i
        go func(h string, p uint16) {
            defer func() { 
                <-sem 
                wg.Done()
            }()
            dumpPorts[i] = NewDial(h, p)
        }(host, port)
    }
    wg.Wait()

    return &Dump{
        host:  host,
        ports: dumpPorts,
    }, nil
}

func isHost(host string) bool {
    return true
}

func parsePortsForScan(ports string) (scanPorts []uint16, err error) {    
    for _, i := range strings.Split(ports, ",") {
        port, err := strconv.Atoi(i)
        if err != nil {
            area, err := parseArea(i)
            if err != nil {
                return nil, fmt.Errorf("parse ports from '%s' failed, got %s", ports, err.Error())
            }
            scanPorts = append(scanPorts, area...)
            continue
        }
        scanPorts = append(scanPorts, uint16(port))
    }
    return
}

func parseArea(item string) ([]uint16, error) {
    area := strings.Split(item, "...")
    if len(area) != 2 {
        return nil, fmt.Errorf("port area '%s' has incorrect format, should be '[start...stop]'", item)
    }
    
    start, err := strconv.Atoi(area[0])
    if err != nil {
        return nil, fmt.Errorf("can't parse port from '%s'", area[0])
    }
    end, err := strconv.Atoi(area[1])
    if err != nil {
        return nil, fmt.Errorf("can't parse port from '%s'", area[1])
    }
    
    if start >= end {
        return nil, fmt.Errorf("area is incorrect, start is more than end")
    }
    if end > maxPortNumber {
        return nil, fmt.Errorf("incorrect port number, '%d' more than %d", end, maxPortNumber)
    }
    if start < minPortNumber {
        return nil, fmt.Errorf("incorrect port number, '%d' less than %d", start, minPortNumber)
    }

    scanPorts := make([]uint16, end-start+1)
    for i,j := start,0; i<=end; i,j = i+1, j+1{
        scanPorts[j] = uint16(i)
    }
    return scanPorts, nil
}