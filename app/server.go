package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	res "github.com/codecrafters-io/http-server-starter-go/app/response"
	"net"
	"net/http"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

const (
	httpVersion = "HTTP/1.1"
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

	req, err := getReq(conn)
	if err != nil {
		fmt.Println("failed to read conn", err.Error())
		os.Exit(1)
	}

	err = processRequest(req).WriteResponse(conn)
	if err != nil {
		fmt.Println("failed process request", err.Error())
		os.Exit(1)
	}
}

func processRequest(req *request.Request) *res.Response {
	path := req.GetHeader().GetPath()
	var statusCode int
	switch path == "/" {
	case true:
		statusCode = http.StatusOK
	default:
		statusCode = http.StatusNotFound
	}

	return res.NewResponse(req.GetHeader().GetProtocolVersion(), statusCode)
}

// readReq read the connection data
func getReq(conn net.Conn) (*request.Request, error) {
	data := make([]byte, 128)
	_, err := conn.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read request data: %s", err.Error())
	}

	return request.NewRequest(data)
}
