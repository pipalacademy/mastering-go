package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net"
)

func handleConn(c net.Conn) {
    defer c.Close()

    input := bufio.NewScanner(c)
    for input.Scan() {
        echo(c, input.Text())
    }
}

func echo(c net.Conn, msg string) {
    _, err := io.WriteString(c, msg+"\n")
    if err != nil {
        log.Print("err") // connection failed
    }
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
