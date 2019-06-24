package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "time"
)

func handleConn(c net.Conn) error {
    defer c.Close()
    _, err := io.WriteString(c, "Hello from Server!\n")
    if err != nil {
        return err
    }
    time.Sleep(5 * time.Second)
    _, err = io.WriteString(c, "Bye from Server!\n")
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // start listening on port 8000
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("started server on localhost:8000")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print("err") // connection failed
            continue
        }
        go handleConn(conn)
    }
}
