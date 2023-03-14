package tcp_scanner

type portNumber interface {
    uint16 | uint32 | uint64 | int64 | int32 | int
}

var (
    maxPortNumber = 65535
    minPortNumber = 1
)
