package src

<<<<<<< HEAD
import (
	"fmt"
	"strings"
)

/*
- Statuscode
- StatusText
- Headers
- Body
*/
type HttpResponse struct {
	StatusCode int
	StatusText string
	Header     map[string]string
	Body       string
}

/*We just need statuscode and body to create a new statusReponse*/
func NewHttpResponse(statusCode int, body string) *HttpResponse {
	return &HttpResponse{
		StatusCode: statusCode,
		StatusText: getStatusText(statusCode),
		Header:     make(map[string]string),
		Body:       body,
	}
}

/*
- SetHeader: Now server will set header here and send back to client
*/
func (r *HttpResponse) SetHeader(key, value string) {
	r.Header[key] = value
}

/*
- Now TCP socket can only understand the bytes not the go structs, so let's convert our response and send back to tcp socket
*/

func (r *HttpResponse) ToBytes() []byte {
	var sb strings.Builder

	// First line of response
	sb.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusText))

	// Add Content-Lenght header
	r.Header["Content-Length"] = fmt.Sprintf("%d", len(r.Body))

	// Headers
	for key, value := range r.Header {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	// Empty line + Body
	sb.WriteString("\r\n")
	sb.WriteString(r.Body)

	return []byte(sb.String())
}

func getStatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
=======
type HttpResponse struct {
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
}
