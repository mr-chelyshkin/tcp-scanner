package tcp_scanner

import(
    "fmt"
    "strings"
    "strconv"
)

var (
    maxPortNumber = 65535
    minPortNumber = 1
)

func isHost(host string) bool {
    return true
}

func parsePortsForScan(ports string) (scanPorts []int, err error) {    
    for _, i := range strings.Split(ports, ",") {
        port, err := strconv.Atoi(i)
        if err != nil {
            area, err := parseArea(i)
            if err != nil {
                return nil, err
            }
            scanPorts = append(scanPorts, area...)
            continue
        }
        scanPorts = append(scanPorts, port)
    }
    return
}

func parseArea(item string) (scanPorts []int, err error) {
    area := strings.Split(item, "...")
    if len(area) != 2 {
        err = fmt.Errorf("port area '%s' has incorrect format, should be '[start...stop]'", item)
        return
    }
    start, err := strconv.Atoi(area[0])
    if err != nil {
        err = fmt.Errorf("can't parse port from '%s'", area[0])
        return
    }
    end, err := strconv.Atoi(area[1])
    if err != nil {
        err = fmt.Errorf("can't parse port from '%s'", area[1])
        return
    }
    if start >= end {
        err = fmt.Errorf("area is incorrect, start is more than end")
        return
    }
    if end > maxPortNumber {
        err = fmt.Errorf("incorrect port number, '%d' more than %d", end, maxPortNumber)
        return
    }
    if start < minPortNumber {
        err = fmt.Errorf("incorrect port number, '%d' less than %d", start, minPortNumber)
    }

    for i:=start; i<=end; i++ {
        scanPorts = append(scanPorts, i)
    }
    return
}