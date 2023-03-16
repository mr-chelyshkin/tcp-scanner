package tcp_scanner

import (
    "testing"
)

func TestDump_Host(t *testing.T) {
    d := &Dump{host: "example.com", ports: []*Port{}}
    if d.Host() != "example.com" {
        t.Errorf("Unexpected Host(): %s", d.Host())
    }
}

func TestDump_OpenPorts(t *testing.T) {
    d := &Dump{host: "example.com", ports: []*Port{
        {port: 80, isOpen: true},
        {port: 443, isOpen: false},
        {port: 8080, isOpen: true},
    }}
    expected := []int{80, 8080}
    actual := d.OpenPorts()
    if len(actual) != len(expected) {
        t.Errorf("Unexpected OpenPorts(): %v", actual)
    }
    for i := range actual {
        if actual[i] != expected[i] {
            t.Errorf("Unexpected OpenPorts(): %v", actual)
        }
    }
}

func TestDump_OpenAddrs(t *testing.T) {
    d := &Dump{host: "example.com", ports: []*Port{
        {port: 80, isOpen: true},
        {port: 443, isOpen: false},
        {port: 8080, isOpen: true},
    }}
    expected := []string{"example.com:80", "example.com:8080"}
    actual := d.OpenAddrs()
    if len(actual) != len(expected) {
        t.Errorf("Unexpected OpenAddrs(): %v", actual)
    }
    for i := range actual {
        if actual[i] != expected[i] {
            t.Errorf("Unexpected OpenAddrs(): %v", actual)
        }
    }
}

func TestDump_All(t *testing.T) {
    d := &Dump{host: "example.com", ports: []*Port{
        {port: 80, isOpen: true},
        {port: 443, isOpen: false},
        {port: 8080, isOpen: true},
    }}
    expected := map[string]bool{
        "example.com:80": true,
        "example.com:443": false,
        "example.com:8080": true,
    }
    actual := d.All()
    if len(actual) != len(expected) {
        t.Errorf("Unexpected All(): %v", actual)
    }
    for k, v := range actual {
        if expected[k] != v {
            t.Errorf("Unexpected All(): %v", actual)
        }
    }
}

