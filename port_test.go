package tcp_scanner

import (
	"net"
    "strconv"
	"testing"
)

func TestNewDial(t *testing.T) {
	host := "example.com"
	port := 80
	expectedPort := uint16(port)
	expectedIsOpen := true

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		t.Fatalf("Failed to start fake TCP server: %v", err)
	}
	defer ln.Close()


	p := NewDial(host, port)
	if p.port != expectedPort {
		t.Errorf("NewDial(%q, %d) returned port %d, expected %d", host, port, p.port, expectedPort)
	}
	if p.isOpen != expectedIsOpen {
		t.Errorf("NewDial(%q, %d) returned isOpen %t, expected %t", host, port, p.isOpen, expectedIsOpen)
	}
}
