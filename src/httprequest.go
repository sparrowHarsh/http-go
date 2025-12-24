package src

import (
	"bufio"
	"fmt"
<<<<<<< HEAD
	"io"
	"strconv"
=======
	"log"
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
	"strings"
)

type HttpRequest struct {
<<<<<<< HEAD
	Method  string
	Path    string
	Version string
	Header  map[string]string
	Body    string
=======
	method  string
	path    string
	version string
	header  map[string]string
	body    string
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
}

// Create a new HttpRequest
func NewHttpRequest(config HttpRequest) *HttpRequest {
	return &HttpRequest{
<<<<<<< HEAD
		Method:  config.Method,
		Path:    config.Path,
		Version: config.Version,
		Header:  config.Header,
		Body:    config.Body,
=======
		method:  config.method,
		path:    config.path,
		version: config.version,
		header:  config.header,
		body:    config.body,
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
	}
}

/*
	Take input from user and parse it to the local fields
	Now input will come in following format, we have to make get it into the given format
<<<<<<< HEAD
	<Method> <Request-Path> <HTTP-Version>
=======
	<Method> <Request-path> <HTTP-Version>
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
    <Header-Name>: <Header-Value>
    ...
    <Header-Name>: <Header-Value>

    <Request-Body>
*/

<<<<<<< HEAD
func ParseHttpRequest(reader *bufio.Reader) (*HttpRequest, error) {
	// Read the request line (first line)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading request line: %w", err)
	}

	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}

	req := &HttpRequest{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Header:  make(map[string]string),
	}

	// ---- Parse Headers ----
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("error reading header: %w", err)
		}

		line = strings.TrimSpace(line)

		// Empty line marks end of headers
		if line == "" {
			break
=======
func ParseHttpRequest(raw string) (*HttpRequest, error) {
	scanner := bufio.NewScanner(strings.NewReader(raw))

	if !scanner.Scan() {
		log.Println("Empty Request")
		return nil, fmt.Errorf("Empty request")
	}

	// This reads line by line
	line := scanner.Text()
	fmt.Println("First line is", line)

	// split at the comma(' ')
	parts := strings.Split(line, " ")
	fmt.Println("Value of parts", parts)

	// match the first line of HTTP Request
	if len(parts) != 3 {
		log.Println("Invalid request line")
		return nil, fmt.Errorf("Empty request")
	}

	req := &HttpRequest{
		method:  parts[0],
		path:    parts[1],
		version: parts[2],
		header:  make(map[string]string),
	}

	// ---- Parse headers ----
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break // end of headers
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			continue
		}

		key := strings.TrimSpace(headerParts[0])
		value := strings.TrimSpace(headerParts[1])
<<<<<<< HEAD
		req.Header[key] = value
	}

	// ---- Read Body if Content-Length exists ----
	if contentLength, ok := req.Header["Content-Length"]; ok {
		length, err := strconv.Atoi(contentLength)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length: %w", err)
		}

		if length > 0 {
			body := make([]byte, length)
			_, err := io.ReadFull(reader, body)
			if err != nil {
				return nil, fmt.Errorf("error reading body: %w", err)
			}
			req.Body = string(body)
		}
	}

	PrintHttpContent(req)
	return req, nil
}

func PrintHttpContent(req *HttpRequest) {
	fmt.Println("Method :", req.Method)
	fmt.Println("Path   :", req.Path)
	fmt.Println("Version:", req.Version)
	fmt.Println("\nHeaders:")
	for k, v := range req.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	if req.Body != "" {
		fmt.Println("\nBody:", req.Body)
	}
=======

		req.header[key] = value
	}

	return req, nil
}

func PrintHttpContent(req HttpRequest) {
	fmt.Println("Method :", req.method)
	fmt.Println("Path   :", req.path)
	fmt.Println("Version:", req.version)
	fmt.Println("\nHeaders:")
	for k, v := range req.header {
		fmt.Printf("%s: %s\n", k, v)
	}
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
}
