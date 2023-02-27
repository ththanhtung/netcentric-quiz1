package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    // connect to the server
    conn, err := net.Dial("tcp", "localhost:9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    // send username and password
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your username: ")
    username, _ := reader.ReadString('\n')
    fmt.Print("Enter your password: ")
    password, _ := reader.ReadString('\n')
    fmt.Fprintf(conn, strings.TrimSpace(username)+"\n")
    fmt.Fprintf(conn, strings.TrimSpace(password)+"\n")

    // receive the access status
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Print(status)

    // start sending data to the server
    for {
        fmt.Print("Enter data to send: ")
        data, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, strings.TrimSpace(data)+"\n")
        response, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }
        fmt.Print("Server response:", response)
        if strings.TrimSpace(data) == "EOF" {
            break
        }
    }
}
