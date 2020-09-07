package main

import (
    "fmt"
    "github.com/itchin/bytes-packet"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:9000")
    if err != nil {
        fmt.Printf("connect failed, err : %v\n", err.Error())
        return
    }
    defer conn.Close()

    msg := "hello world"

    conn.Write(bpacket.Packet([]byte(msg)))
    time.Sleep(time.Second / 10)
}