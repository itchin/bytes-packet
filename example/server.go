package main

import (
    "fmt"
    "io"
    "net"
    "github.com/itchin/bytes-packet"
)

func main() {
    server, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }
    defer server.Close()

    for {
        conn, err := server.Accept()
        if err != nil {
            continue
        }
        defer conn.Close()
        go run(conn)
    }
}

func run(conn net.Conn) {
    // 每次读取的字节流
    buf := make([]byte, 1)

    for {
        n, err := conn.Read(buf)

        if err != nil || err == io.EOF {
            conn.Close()
            break
        }

        packet := bpacket.Unpack(buf[:n])
        if len(packet) > 0 {
            fmt.Println(string(packet))
        }
    }
}
