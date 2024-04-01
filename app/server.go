package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

const (
	httpVersion = "HTTP/1.1"
	CRLF        = "\r\n"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	_ = readConn(conn)

	responseData := response(nil)
	_, err = conn.Write(responseData)
	if err != nil {
		fmt.Println("Failed to write response data: ", err.Error())
		os.Exit(1)
	}
}

// readConn read the connection data
func readConn(conn net.Conn) []byte {
	return nil
}

// response send the response to the connection
func response(reqData []byte) []byte {
	// process the request data

	// response data
	respHeader := []byte(fmt.Sprintf("%s %d%s%s", httpVersion, http.StatusOK, CRLF, CRLF))

	return respHeader
}
