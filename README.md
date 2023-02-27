
# Net Centric - Quiz 1



## Requirements

 Write a client/server program to support the following interactions:

Server maintains a list of simple user name and password (plain text, no encryption, hardcode in your program). Server will ask for user name and password whenever client connects to and verify
Client connects to server using username and password. Communication is granted only when username/password is matched.
After successfully establishing the connection, client can send data to server, server echos back the data in all uppercase format
Connection closes when client sends “EOF” 

Example:

client log in with account: student/imastudent
server verifies with its database
if matched, server grants access
client sends “Hello friend”
server echos back “HELLO FRIEND”
clent sends “EOF” to terminate connection

## Note

In order to run this program your computer have to install Go Programming language

## Run Locally

Start the server first

```bash
  go run server/main.go
```

Start the client

```bash
  go run client/main.go
```
