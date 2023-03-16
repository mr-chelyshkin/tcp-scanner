package tcp_scanner

import(
    "reflect"
    "testing"
)

func TestScan(t *testing.T) {
    host := "example.com"
    ports := "80...82"

    dump, err := Scan(host, ports, 1)
    if err != nil {
        t.Errorf("unexpected error: %s", err.Error())
    }
    if dump.host != host {
        t.Errorf("expected host %s, got %s", host, dump.host)
    }

    expectedNumPorts := 3
    if len(dump.ports) != expectedNumPorts {
        t.Errorf("expected %d ports, got %d", expectedNumPorts, len(dump.ports))
    }
    for i, port := range dump.ports {
        expectedPort := uint16(80 + i)
        if port == nil {
            t.Errorf("port %d is nil", expectedPort)
        } else if port.port != expectedPort {
            t.Errorf("expected port %d, got %d", expectedPort, port.port)
        }
    }
}


func TestParseArea(t *testing.T) {
    tests := []struct{
        name string
        input string
        want []uint16
        err string
    }{
        {
            name:  "valid input",
            input: "80...90",
            want:  []uint16{80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
            err:   "",
        },
        {
            name:  "invalid input - incorrect format",
            input: "80...90...",
            want:  nil,
            err:   "port area '80...90...' has incorrect format, should be '[start...stop]'",
        },
        {
            name:  "invalid input - start is greater than end",
            input: "90...80",
            want:  nil,
            err:   "area is incorrect, start is more than end",
        },
        {
            name:  "invalid input - end is more than maxPortNumber",
            input: "80...65536",
            want:  nil,
            err:   "incorrect port number, '65536' more than 65535",
        },
        {
            name:  "invalid input - start is less than minPortNumber",
            input: "-1...80",
            want:  nil,
            err:   "incorrect port number, '-1' less than 1",
        },
        {
            name:  "invalid input - cannot parse start",
            input: "abc...80",
            want:  nil,
            err:   "can't parse port from 'abc'",
        },
        {
            name:  "invalid input - cannot parse end",
            input: "80...def",
            want:  nil,
            err:   "can't parse port from 'def'",
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got, err := parseArea(test.input)
            if err != nil {
                if err.Error() != test.err {
                    t.Fatalf("unexpected error: got %v, want %v", err.Error(), test.err)
                }
                return
            }

            if !reflect.DeepEqual(got, test.want) {
                t.Fatalf("unexpected result: got %v, want %v", got, test.want)
            }
        })
    }
}

func TestParsePortsForScan(t *testing.T) {
    tests := []struct{
        ports       string
        expected    []uint16
        expectError bool
    }{
        {"80,443,8080", []uint16{80, 443, 8080}, false},
        {"1...10", []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
        {"22,25,abc", nil, true},
    }

    for _, test := range tests {
        scanPorts, err := parsePortsForScan(test.ports)

        if test.expectError {
            if err == nil {
                t.Errorf("Expected error for input %s but got none", test.ports)
            }
            continue
        }
        if err != nil {
            t.Errorf("Unexpected error for input %s: %s", test.ports, err.Error())
        }

        if !reflect.DeepEqual(scanPorts, test.expected) {
            t.Errorf("Expected %v for input %s but got %v", test.expected, test.ports, scanPorts)
        }
        
    }
}