package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	res "github.com/codecrafters-io/http-server-starter-go/app/response"
	"net"
	"net/http"
	"os"
	"strings"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

const (
	echo      = "/echo/"
	userAgent = "/user-agent"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
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

	defer func(conn net.Conn) {
		fmt.Println("Closing Connection")
		err = conn.Close()
		if err != nil {
			fmt.Println("failed to close connection ", err.Error())
			return
		}
	}(conn)

	req, err := getReq(conn)
	if err != nil {
		fmt.Println("failed to read conn", err.Error())
		os.Exit(1)
	}

	fmt.Println("Processing Request")
	resp, err := processRequest(req)
	if err != nil {
		fmt.Println("failed process request", err.Error())
		os.Exit(1)
	}

	fmt.Println("Writing Response")
	err = resp.WriteResponse(conn)
	if err != nil {
		fmt.Println("failed to write response", err.Error())
		os.Exit(1)
	}

}

func processRequest(req *request.Request) (*res.Response, error) {
	var (
		data string
	)

	path := req.GetHeader().GetPath()
	switch true {
	case strings.EqualFold(path, "/"):
		data = ""
	case strings.HasPrefix(path, echo):
		data, _ = strings.CutPrefix(path, echo)
	case strings.HasPrefix(path, userAgent):
		data = req.GetHeader().GetHeaderVal("user-agent")
	default:
		return res.NewResponse(req.GetHeader().GetProtocolVersion(), http.StatusNotFound, "")
	}

	return res.NewResponse(req.GetHeader().GetProtocolVersion(), http.StatusOK, data)
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
