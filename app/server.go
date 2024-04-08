package main

import (
	"flag"
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

	directoryFlag = "directory"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	err := loadFlagsAsENV()
	if err != nil {
		fmt.Println("failed to load flags", err.Error())
		os.Exit(1)
	}

	fmt.Println("Starting listener")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	var (
		conn    net.Conn
		counter int
	)

	defer l.Close()

	fmt.Println("Accepting connections")

	for conn, err = l.Accept(); err == nil; conn, err = l.Accept() {
		counter++
		fmt.Println("received connection ", counter, " from ", conn.RemoteAddr().String())
		go func(conn net.Conn) {
			err := handleConnection(conn)
			if err != nil {
				fmt.Println("failed to handle connection ", err.Error())
				os.Exit(1)
			}

			conn.Close()
		}(conn)
	}

	fmt.Println("Failed to accept connection")
}

func loadFlagsAsENV() error {
	// if we are running binary for other stages then no flags will be available
	if len(os.Args) < 2 {
		return nil
	}

	fmt.Println("loading flags")

	var directory string
	flag.StringVar(&directory, directoryFlag, "", "directory to serve")

	flag.Parse()

	if directory == "" {
		return fmt.Errorf("empty directory provided")
	}

	return os.Setenv(directoryFlag, directory)
}

func handleConnection(conn net.Conn) error {
	req, err := getReq(conn)
	if err != nil {
		return fmt.Errorf("failed to read conn: %w", err)
	}

	fmt.Println("Processing Request")
	resp, err := processRequest(req)
	if err != nil {
		return fmt.Errorf("failed process request: %w", err)
	}

	fmt.Println("Writing Response")
	err = resp.WriteResponse(conn)
	if err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

func processRequest(req *request.Request) (*res.Response, error) {
	var (
		data       = &res.Content{}
		statusCode = http.StatusOK
	)

	path := req.GetHeader().GetPath()
	switch true {
	case strings.EqualFold(path, "/"):
		data = res.CreateContent([]byte(""), res.PlainText)
	case strings.HasPrefix(path, echo):
		reqData, _ := strings.CutPrefix(path, echo)
		data = res.CreateContent([]byte(reqData), res.PlainText)
	case strings.HasPrefix(path, userAgent):
		data = res.CreateContent([]byte(req.GetHeader().GetHeaderVal("user-agent")), res.PlainText)
	case strings.HasPrefix(path, "/files/"):
		requestedFileName := strings.TrimPrefix(path, "/files/")
		directoryPath := os.Getenv(directoryFlag)
		// check if the requested data exist
		// if it doesn't exist return not found
		file, err := os.DirFS(directoryPath).Open(requestedFileName)
		if err != nil {
			statusCode = http.StatusNotFound
			break
		}

		// read data
		info, _ := file.Stat()
		var dataByte = make([]byte, info.Size())
		_, err = file.Read(dataByte)
		if err != nil {
			return nil, fmt.Errorf("failed to read content of the file: %w", err)
		}

		data = res.CreateContent(dataByte, res.FileData)

	default:
		statusCode = http.StatusNotFound
	}

	response := res.NewResponse(req.GetHeader().GetProtocolVersion(), statusCode, data)
	return response, nil
}

// readReq read the connection data
func getReq(conn net.Conn) (*request.Request, error) {
	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read request data: %s", err.Error())
	}

	return request.NewRequest(data[:n])
}
