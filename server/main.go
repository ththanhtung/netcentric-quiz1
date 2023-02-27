package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// list of user name and password
var users = map[string]string{
	"student": "imastudent",
}

func main() {
	// listen on all interfaces
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 9999")

	// accept connections from outside
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Connected by:", conn.RemoteAddr())

		// ask for username and password
		conn.Write([]byte("Please enter your username: "))
		username, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		username = strings.TrimSpace(username)
		conn.Write([]byte("Please enter your password: "))
		password, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		password = strings.TrimSpace(password)

		// verify the username and password
		if p, ok := users[username]; ok && p == password {
			conn.Write([]byte("Access granted!\n"))

			// start receiving data from the client
			for {
				data, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					break
				}
				data = strings.TrimSpace(data)
				if data == "EOF" {
					fmt.Println("Connection closed by:", conn.RemoteAddr())
					break
				} else {
					conn.Write([]byte(strings.ToUpper(data) + "\n"))
				}
			}
		} else {
			conn.Write([]byte("Access denied!\n"))
		}

		// close the connection
		conn.Close()
	}
}
