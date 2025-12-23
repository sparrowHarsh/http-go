package src

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

type HttpRequest struct {
	method  string
	path    string
	version string
	header  map[string]string
	body    string
}

// Create a new HttpRequest
func NewHttpRequest(config HttpRequest) *HttpRequest {
	return &HttpRequest{
		method:  config.method,
		path:    config.path,
		version: config.version,
		header:  config.header,
		body:    config.body,
	}
}

/*
	Take input from user and parse it to the local fields
	Now input will come in following format, we have to make get it into the given format
	<Method> <Request-path> <HTTP-Version>
    <Header-Name>: <Header-Value>
    ...
    <Header-Name>: <Header-Value>

    <Request-Body>
*/

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
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			continue
		}

		key := strings.TrimSpace(headerParts[0])
		value := strings.TrimSpace(headerParts[1])

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
}
