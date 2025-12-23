[Client] <--> [TCP Listener] <--> [Connection Handler] <--> [Request Parser] <--> [Router] <--> [Handlers]
                                                                 |
                                                                 v
[Response Generator] <--> [Chunked Encoder] <--> [Packet Validator] <--> [TCP Write]

What all things I have to design here?

step1:
    Design a server that will listen on HTTP port 


step1:

    /* HTTP Format */
    <Method> <Request-URI> <HTTP-Version>
    <Header-Name>: <Header-Value>
    ...
    <Header-Name>: <Header-Value>

    <Request-Body>
    -----------------------------------------------------------------------------------------------------
    Request Line: Starts with the HTTP method (e.g., GET, POST), followed by the request URI (e.g., /path?query=value), and the HTTP version (e.g., HTTP/1.1). Ends with a CRLF (\r\n).

    Headers: Zero or more header lines in the format Name: Value, each ending with CRLF. Common headers include Host, User-Agent, Content-Type, etc.

    Empty Line: A single CRLF to separate headers from the body.

    Body: Optional request body (e.g., for POST requests), containing data like form data or JSON. The body is omitted for methods like GET.

    -----------------------------------------------------------------------------------------------------------

    I have to store all the headers in one filed
    Http method should be stored in another
    Request URI should be getting stored in another filed
    Http version should get in another field
    Request body should get in another 

    type HttpRequest struct{
        Method string
        Url string
        version string
        map<string, string> header
        body string
    }

step2:

    Just design GET feature

    /*HTTP Format*/
    GET /index.html /HTTP/1.1
    Host: www.example.com
    user-Agent: MyBrowser/1.0
