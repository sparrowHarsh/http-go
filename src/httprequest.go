package src

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Header  map[string]string
	Body    string
}

// Create a new HttpRequest
func NewHttpRequest(config HttpRequest) *HttpRequest {
	return &HttpRequest{
		Method:  config.Method,
		Path:    config.Path,
		Version: config.Version,
		Header:  config.Header,
		Body:    config.Body,
	}
}

/*
	Take input from user and parse it to the local fields
	Now input will come in following format, we have to make get it into the given format
	<Method> <Request-Path> <HTTP-Version>
    <Header-Name>: <Header-Value>
    ...
    <Header-Name>: <Header-Value>

    <Request-Body>
*/

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
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			continue
		}

		key := strings.TrimSpace(headerParts[0])
		value := strings.TrimSpace(headerParts[1])
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
}
