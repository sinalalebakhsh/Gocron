package features

var TitleHTTPClients = []string{
	"ALL HTTP CLIENTS",
	"ALLHTTPCLIENTS",
	"ALL CREATING HTTP CLIENTS",
	"ALLCREATINGHTTPCLIENTS",
}

var OriginalHTTPClients = DataBase{
	Alldatafield: `
436.Creating HTTP Clients
    Putting HTTP Clients in Context

    What are they?
        HTTP requests are used to retrieve data from HTTP servers
    
    Why are they useful?
        HTTP is one of the most widely used protocols and is commonly used to provide
        access to content that can be presented to the user as well as data that is consumed
        programmatically.

    How is it used?
        The features of the net/http package are used to create and send requests and
        process responses.
    
    Are there any pitfalls or limitations?
        These features are well-designed and easy to use, although some features require a
        specific sequence to use.
    
    Are there any alternatives?
        The standard library includes support for other network protocols and also for
        opening and using lower-level network connections. 
        See the 
        
        https://pkg.go.dev/net@go1.17.1 

        https://pkg.go.dev/net
        
        for details of the net package and its subpackages, such as net/smtp,
        for example, which implements the SMTP protocol.


    Chapter Summary:
    Problem                                     Solution
    --------                                    ------------
    Send HTTP requests                          Use the convenience methods for specific HTTP methods
    Configure HTTP requests                     Use the fields and methods defined by the Client struct
    Create a preconfigured request              Use the NewRequest convenience functions
    Use cookies in a request                    Use a cookie jar
    Configure how redirections are processed    Use the CheckRedirect field to register a function that is
                                                invoked to deal with a redirection
    Send multipart forms                        Use the mime/multipart package


    Preparing for This Chapter
    1- go mod init httpclient
    2- printer.go
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
    3- httpclient folder -> file named product.go
        package main
        type Product struct {
            Name, Category string
            Price float64
        }
        var Products = []Product {
            { "Kayak", "Watersports", 279 },
            { "Lifejacket", "Watersports", 49.95 },
            { "Soccer Ball", "Soccer", 19.50 },
            { "Corner Flags", "Soccer", 34.95 },
            { "Stadium", "Soccer", 79500 },
            { "Thinking Cap", "Chess", 16 },
            { "Unsteady Chair", "Chess", 75 },
            { "Bling-Bling King", "Chess", 1200 },
        }
    4- The Contents of the index.html File in the httpclient Folder
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>    
    5- The Contents of the server.go File in the httpclient Folder
        package main
        import (
            "encoding/json"
            "fmt"
            "io"
            "net/http"
            "os"
        )
        func init() {
            http.HandleFunc("/html",
                func(writer http.ResponseWriter, request *http.Request) {
                    http.ServeFile(writer, request, "./index.html")
                })
            http.HandleFunc("/json",
                func(writer http.ResponseWriter, request *http.Request) {
                    writer.Header().Set("Content-Type", "application/json")
                    json.NewEncoder(writer).Encode(Products)
                })
            http.HandleFunc("/echo",
                func(writer http.ResponseWriter, request *http.Request) {
                    writer.Header().Set("Content-Type", "text/plain")
                    fmt.Fprintf(writer, "Method: %v\n", request.Method)
                    for header, vals := range request.Header {
                        fmt.Fprintf(writer, "Header: %v: %v\n", header, vals)
                    }
                    fmt.Fprintln(writer, "----")
                    data, err := io.ReadAll(request.Body)
                    if err == nil {
                        if len(data) == 0 {
                            fmt.Fprintln(writer, "No body")
                        } else {
                            writer.Write(data)
                        }
                    } else {
                        fmt.Fprintf(os.Stdout, "Error reading body: %v\n", err.Error())
                    }
                })
        }
    6- The Contents of the main.go File in the httpclient Folder
        package main
        import (
            "net/http"
        )
        func main() {
            Printfln("Starting HTTP Server")
            http.ListenAndServe(":5000", nil)
        }
    =======================================================================================
    Output:
        The code in httpclient folder will be compiled and executed. 
        Use a web browser to request http://localhost:5000/html and http://localhost:5000/json,
        To see the echo result, request http://localhost:5000/echo
████████████████████████████████████████████████████████████████████████
437.Sending Simple HTTP Requests
    The net/http package provides a set of convenience functions that make basic HTTP requests. 
    The functions are named after the HTTP method of the request they created.
    
    The Convenience Methods for HTTP Requests
    Name                                Description
    -----------------                   -----------------------------
    Get(url)                            This function sends a GET request to the specified HTTP or HTTPS URL. The
                                        results are a Response and an error that reports problems with the request.
    Head(url)                           This function sends a HEAD request to the specified HTTP or HTTPS URL.
                                        A HEAD request returns the headers that would be returned for a GET request.
                                        The results are a Response and an error that reports problems with the request.
    Post(url, contentType, reader)      This function sends a POST request to the specified HTTP or HTTPS URL, with
                                        the specified Content-Type header value. The content for the form is provided
                                        by the specified Reader. The results are a Response and an error that reports
                                        problems with the request.
    PostForm(url, data)                 This function sends a POST request to the specified HTTP or HTTPS URL, with
                                        the Content-Type header set to application/x-www-form-urlencoded. The
                                        content for the form is provided by a map[string][]string. The results are a
                                        Response and an error that reports problems with the request.

████████████████████████████████████████████████████████████████████████
438.Sending a GET Request in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/html")
            if (err == nil) {
                response.Write(os.Stdout)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================================================
    The argument to the Get function is a string that contains the URL to request. 
    The results are a Response value and an error that reports any problems sending the request.

    Output:
        HTTP/1.1 200 OK
        Content-Length: 171
        Accept-Ranges: bytes
        Content-Type: text/html; charset=utf-8
        Date: Thu, 12 Oct 2023 16:17:10 GMT
        Last-Modified: Wed, 11 Oct 2023 12:44:50 GMT
        
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>
    
    ========================================================================================================
    example-2:
    main.go:
    package main
    import (
        "net/http"
        "os"
        // "time"
    )
    func main() {
        // go http.ListenAndServe(":5000", nil)
        // time.Sleep(time.Second)
        response, err := http.Get("https://www.google.com")
        if (err == nil) {
            response.Write(os.Stdout)
        } else {
            Printfln("Error: %v", err.Error())
        }
    }

Output:
    HTTP/2.0 403 Forbidden
    Content-Length: 1579
    Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
    Content-Type: text/html; charset=UTF-8
    Date: Thu, 12 Oct 2023 16:15:23 GMT
    Referrer-Policy: no-referrer
    
    <!DOCTYPE html>
    <html lang=en>
    <meta charset=utf-8>
    <meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
    <title>Error 403 (Forbidden)!!1</title>
    <style>
        *{margin:0;padding:0}html,code{font:15px/22px arial,sans-serif}html{background:#fff;color:#222;padding:15px}body{margin:7% auto 0;max-width:390px;min-height:180px;padding:30px 0 15px}* > body{background:url(//www.google.com/images/errors/robot.png) 100% 5px no-repeat;padding-right:205px}p{margin:11px 0 22px;overflow:hidden}ins{color:#777;text-decoration:none}a img{border:0}@media screen and (max-width:772px){body{background:none;margin-top:0;max-width:none;padding-right:0}}#logo{background:url(//www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png) no-repeat;margin-left:-5px}@media only screen and (min-resolution:192dpi){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat 0% 0%/100% 100%;-moz-border-image:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) 0}}@media only screen and (-webkit-min-device-pixel-ratio:2){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat;-webkit-background-size:100% 100%}}#logo{display:inline-block;height:54px;width:150px}
    </style>
    <a href=//www.google.com/><span id=logo aria-label=Google></span></a>
    <p><b>403.</b> <ins>That's an error.</ins>
    <p>Your client does not have permission to get URL <code>/</code> from this server.  <ins>That's all we know.</ins>

████████████████████████████████████████████████████████████████████████
439.The Fields and Methods Defined by the Response Struct
    Name            Description
    ----------      -------------------------------------
    StatusCode      This field returns the response status code, expressed as an int.
    Status          This field returns a string containing the status description.
    Proto           This field returns a string containing the response HTTP protocol.
    Header          This field returns a map[string][]string that contains the response headers.
    Body            This field returns a ReadCloser, which is a Reader that defines a Close method and
                    which provides access to the response body.
    Trailer         This field returns a map[string][]string that contains the response trailers.
    ContentLength   This field returns the value of the Content-Length header, parsed into an int64 value.
                    TransferEncoding This field returns the set of Transfer-Encoding header values.
    Close           This bool field returns true if the response contains a Connection header set to close,
                    which indicates that the HTTP connection should be closed.
    Uncompressed    This field returns true if the server sent a compressed response that was
                    decompressed by the net/http package.
    Request         This field returns the Request that was used to obtain the response. The Request struct
                    is described in Chapter 24.
    TLS             This field provides details of the HTTPS connection.
    Cookies()       This method returns a []*Cookie, which contains the Set-Cookie headers in the
                    response. The Cookie struct is described in Chapter 24.
    Location()      This method returns the URL from the response Location header and an error that
                    indicates when the response does not contain this header.
    Write(writer)   This method writes a summary of the response to the specified Writer.
████████████████████████████████████████████████████████████████████████
440.Reading the Response Body in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/html")
            if (err == nil && response.StatusCode == http.StatusOK) {
                data, err := io.ReadAll(response.Body)
                if (err == nil) {
                    defer response.Body.Close()
                    os.Stdout.Write(data)
                }
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: in Terminal
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>
████████████████████████████████████████████████████████████████████████
441.Reading and Parsing Data in the main.go
    main.go:
        package main
        import (
            "net/http"
            //"os"
            "time"
            //"io"
            "encoding/json"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/json")
            if (err == nil && response.StatusCode == http.StatusOK) {
                defer response.Body.Close()
                data := []Product {}
                err = json.NewDecoder(response.Body).Decode(&data)
                if (err == nil) {
                    for _, p := range data {
                        Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
                    }
                } else {
                    Printfln("Decode error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: in Terminal
        Name: Kayak,Price: $279.00
        Name: Lifejacket,Price: $49.95
        Name: Soccer Ball,Price: $19.50
        Name: Corner Flags,Price: $34.95
        Name: Stadium,Price: $79500.00
        Name: Thinking Cap,Price: $16.00
        Name: Unsteady Chair,Price: $75.00
        Name: Bling-Bling King,Price: $1200.00
████████████████████████████████████████████████████████████████████████
442.Sending POST Requests
    The Post and PostForm functions are used to send POST requests. 
    The PostForm function encodes a map of values as form data.

    Sending a Form in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            //"encoding/json"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            formData := map[string][]string {
                "name":  { "Kayak "},
                "category": { "Watersports"},
                "price":  { "279"},
            }
            response, err := http.PostForm("http://localhost:5000/echo", formData)
            if (err == nil && response.StatusCode == http.StatusOK) {
                io.Copy(os.Stdout, response.Body)
                defer response.Body.Close()
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: 
        Method: POST
        Header: Accept-Encoding: [gzip]
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Length: [42]
        Header: Content-Type: [application/x-www-form-urlencoded]
        ----
        category=Watersports&name=Kayak+&price=279
████████████████████████████████████████████████████████████████████████
443.Posting a Form Using a Reader
    The Post function sends a POST request 
    to the server and creates the request body by reading content from a Reader.

    Posting from a Reader in the main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            "encoding/json"
            "strings"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            var builder strings.Builder
            err := json.NewEncoder(&builder).Encode(Products[0])
            if (err == nil) {
                response, err := http.Post("http://localhost:5000/echo","application/json",strings.NewReader(builder.String()))
                if (err == nil && response.StatusCode == http.StatusOK) {
                    io.Copy(os.Stdout, response.Body)
                    defer response.Body.Close()
                } else {
                    Printfln("Error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Method: POST
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Length: [54]
        Header: Content-Type: [application/json]
        Header: Accept-Encoding: [gzip]
        ----
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
444.Understanting The Content-Length Header
    This header is set automatically but is included in requests only when it is
    possible to determine how much data will be included in the body in advance. 
    This is done by inspecting the Reader to determine the dynamic type. 
    When the data is stored in memory using the strings.
    Reader, bytes.Reader, or bytes.Buffer type, 
    the built-in len function is used to determine the amount of data, 
    and the result is used to set the Content-Length header.

    For all other types, the Content-Type head is not set, 
    and chunked encoding is used instead, which means that 
    the body is written in blocks of data whose size 
    is declared as part of the request body. 
    This approach allows requests to be sent 
    without needing to read all the data from the Reader just to work
    out how many bytes there are. 
    Chunked encoding is described at 
    
        https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding


████████████████████████████████████████████████████████████████████████
445.Configuring HTTP Client Requests
    The Client struct is used when control is required over an HTTP request 
    and defines the fields and methods described:

    The Client Fields and Methods
    Name                            Description
    -----------------               -------------------------
    Transport                       This field is used to select the transport that will be used to send the HTTP
                                    request. The net/http package provides a default transport.
    CheckRedirect                   This field is used to specify a custom policy for dealing with repeated
                                    redirections, as described in the “Managing Redirections” section.
    Jar                             This field returns a CookieJar, which is used to manage cookies, as
                                    described in the “Working with Cookies” section.
    Timeout                         This field is used to set a timeout for the request, specified as a time.Duration 
    Do(request)                     This method sends the specified Request, returning a Response and an
                                    error that indicates problems sending the request.
    CloseIdleConnections()          This method closes any idle HTTP requests that are currently open and unused.
    Get(url)                        This method is called by the Get function described
    Head(url)                       This method is called by the Head function described
    Post(url, contentType,reader)   This method is called by the Post function described
    PostForm(url, data)             This method is called by the PostForm function described
████████████████████████████████████████████████████████████████████████
446.Useful Request Fields and Methods
    Name                Description
    ------              ------------------------------
    Method              This string field specifies the HTTP method that will be used for the request. The net/
                        http package defines constants for HTTP methods, such as MethodGet and MethodPost.
    URL                 This URL field specifies the URL to which the request will be sent. The URL struct is
    Header              This field is used to specify the headers for the request. The headers are specified in a
                        map[string][]string, and the field will be nil when a Request value is created using the literal struct syntax.
    ContentLength       This field is used to set the Content-Length header using an int64 value.
    TransferEncoding    This field is used to set the Transfer-Encoding header using a slice of strings.
    Body                This ReadCloser field specifies the source for the request body. If you have a Reader
                        that doesn't define a Close method, then the io.NopCloser function can be used to
                        create a ReadCloser whose Close method does nothing.
████████████████████████████████████████████████████████████████████████
447.The Function for Parsing URL Values
    Name            Description
    ---------       ---------------------
    Parse(string)   This method parses a string into a URL. The results are the URL value and an error that
                    indicates problems parsing the string.
████████████████████████████████████████████████████████████████████████
448.Sending a Request in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            "encoding/json"
            "strings"
            "net/url"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            var builder strings.Builder
            err := json.NewEncoder(&builder).Encode(Products[0])
            if (err == nil) {
                reqURL, err := url.Parse("http://localhost:5000/echo")
                if (err == nil) {
                    req := http.Request {
                        Method: http.MethodPost,
                        URL: reqURL,
                        Header: map[string][]string {
                            "Content-Type": { "application.json" },
                        },
                        Body: io.NopCloser(strings.NewReader(builder.String())),
                    }
                    response, err := http.DefaultClient.Do(&req)
                    if (err == nil && response.StatusCode == http.StatusOK) {
                        io.Copy(os.Stdout, response.Body)
                        defer response.Body.Close()
                    } else {
                        Printfln("Request Error: %v", err.Error())
                    }
                } else {
                    Printfln("Parse Error: %v", err.Error())
                }
            } else {
                Printfln("Encoder Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Method: POST
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Type: [application.json]
        Header: Accept-Encoding: [gzip]
        ----
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
449.Using the Convenience Functions to Create a Request
    The net/http Convenience Functions for Creating Requests
    Name                                                    Description
    -------------------------                               ----------------------------
    NewRequest(method, url,reader)                          This function creates a new Reader, configured with the specified method,
                                                            URL, and body. The function also returns an error that indicates problems
                                                            creating the value, including parsing the URL, which is expressed as a string.
    NewRequestWithContext(context, method, url, reader)     This function creates a new Reader that will be sent in the specified context.

████████████████████████████████████████████████████████████████████████
450.Using the Convenience Function in the main.go
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "net/http"
            "os"
            "strings"
            "time"
            //"net/url"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            var builder strings.Builder
            err := json.NewEncoder(&builder).Encode(Products[0])
            if err == nil {
                req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/echo",
                    io.NopCloser(strings.NewReader(builder.String())))
                if err == nil {
                    req.Header["Content-Type"] = []string{"application/json"}
                    response, err := http.DefaultClient.Do(req)
                    if err == nil && response.StatusCode == http.StatusOK {
                        io.Copy(os.Stdout, response.Body)
                        defer response.Body.Close()
                    } else {
                        Printfln("Request Error: %v", err.Error())
                    }
                } else {
                    Printfln("Request Init Error: %v", err.Error())
                }
            } else {
                Printfln("Encoder Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Method: POST
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Type: [application/json]
        Header: Accept-Encoding: [gzip]
        ----
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
451.Working with Cookies
    The Client keeps track of the cookies it receives from the server and automatically includes them in
    subsequent requests.

    example:
    The Contents of the server_cookie.go:
        package main
        import (
            "fmt"
            "net/http"
            "strconv"
        )
        func init() {
            http.HandleFunc("/cookie",
                func(writer http.ResponseWriter, request *http.Request) {
                    counterVal := 1
                    counterCookie, err := request.Cookie("counter")
                    if err == nil {
                        counterVal, _ = strconv.Atoi(counterCookie.Value)
                        counterVal++
                    }
                    http.SetCookie(writer, &http.Cookie{
                        Name: "counter", Value: strconv.Itoa(counterVal),
                    })
                    if len(request.Cookies()) > 0 {
                        for _, c := range request.Cookies() {
                            fmt.Fprintf(writer, "Cookie Name: %v, Value: %v\n",
                                c.Name, c.Value)
                        }
                    } else {
                        fmt.Fprintln(writer, "Request contains no cookies")
                    }
                })
        }
    -----------------------------------
    Changing URL in the main.go:
        package main
        import (
            "io"
            "net/http"
            "os"
            "time"
            // "encoding/json"
            // "strings"
            //"net/url"
            "net/http/cookiejar"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            jar, err := cookiejar.New(nil)
            if err == nil {
                http.DefaultClient.Jar = jar
            }
            for i := 0; i < 3; i++ {
                req, err := http.NewRequest(http.MethodGet,
                    "http://localhost:5000/cookie", nil)
                if err == nil {
                    response, err := http.DefaultClient.Do(req)
                    if err == nil && response.StatusCode == http.StatusOK {
                        io.Copy(os.Stdout, response.Body)
                        defer response.Body.Close()
                    } else {
                        Printfln("Request Error: %v", err.Error())
                    }
                } else {
                    Printfln("Request Init Error: %v", err.Error())
                }
            }
        }
    ====================================================================
    Output: in Terminal
        Request contains no cookies
        Cookie Name: counter, Value: 1
        Cookie Name: counter, Value: 2

████████████████████████████████████████████████████████████████████████
452.The Methods Defined by the CookieJar Interface
    Name                        Description
    ------------------------    ---------------------------------
    SetCookies(url, cookies)    This method stores a *Cookie slice for the specified URL.
    Cookes(url)                 This method returns a *Cookie slice containing the cookies that
                                should be included in a request for the specified URL.


    The net/http/cookiejar package contains an implementation of the CookieJar interface that stores
    cookies in memory. Cookie jars are created with a constructor function
████████████████████████████████████████████████████████████████████████
453.The Cookie Jar Constructor Function in the net/http/cookiejar Package
    Name            Description
    ------------    -------------------------
    New(options)    This function creates a new CookieJar, configured with an Options struct, described
                    next. The function also returns an error that reports problems creating the jar.

    The New function accepts a net/http/cookiejar/Options struct, which is used to configure the
    cookie jar. There is only one Options field, PublicSuffixList, which is used to specify an implementation
    of the interface with the same name, which provides support for preventing cookies from being set too
    widely, which can cause privacy violations. The standard library doesn't contain an implementation of the
    PublicSuffixList interface, but there is an implementation available at 
    
        https://pkg.go.dev/golang.org/x/net/publicsuffix

████████████████████████████████████████████████████████████████████████
454.Creating Separate Clients and Cookie Jars
    A consequence of using the DefaultClient is that all requests share the same cookies, 
    which can be useful, especially since the cookie jar will ensure that each request 
    only includes the cookies that are required for each URL.
████████████████████████████████████████████████████████████████████████
455.Creating Separate Clients in the main.go File
    example:
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            //"encoding/json"
            //"strings"
            //"net/url"
            "net/http/cookiejar"
            "fmt"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            clients := make([]http.Client, 3)
            for index, client := range clients {
                jar, err := cookiejar.New(nil)
                if (err == nil) {
                    client.Jar = jar
                }
                for i := 0; i < 3; i++ {
                    req, err := http.NewRequest(http.MethodGet,
                        "http://localhost:5000/cookie", nil)
                    if (err == nil) {
                        response, err := client.Do(req)
                        if (err == nil && response.StatusCode == http.StatusOK) {
                            fmt.Fprintf(os.Stdout, "Client %v: ", index)
                            io.Copy(os.Stdout, response.Body)
                            defer response.Body.Close()
                        }  else {
                            Printfln("Request Error: %v", err.Error())
                        }
                    } else {
                        Printfln("Request Init Error: %v", err.Error())
                    }
                }
            }
        }
    ====================================================================
    Output:
        Client 0: Request contains no cookies
        Client 0: Cookie Name: counter, Value: 1
        Client 0: Cookie Name: counter, Value: 2
        Client 1: Request contains no cookies
        Client 1: Cookie Name: counter, Value: 1
        Client 1: Cookie Name: counter, Value: 2
        Client 2: Request contains no cookies
        Client 2: Cookie Name: counter, Value: 1
        Client 2: Cookie Name: counter, Value: 2
████████████████████████████████████████████████████████████████████████
456.Sharing a CookieJar in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
            "os"
            "time"
            //"encoding/json"
            //"strings"
            //"net/url"
            "fmt"
            "net/http/cookiejar"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            jar, err := cookiejar.New(nil)
            clients := make([]http.Client, 3)
            for index, client := range clients {
                //jar, err := cookiejar.New(nil)
                if err == nil {
                    client.Jar = jar
                }
                for i := 0; i < 3; i++ {
                    req, err := http.NewRequest(http.MethodGet,
                        "http://localhost:5000/cookie", nil)
                    if err == nil {
                        response, err := client.Do(req)
                        if err == nil && response.StatusCode == http.StatusOK {
                            fmt.Fprintf(os.Stdout, "Client %v: ", index)
                            io.Copy(os.Stdout, response.Body)
                            defer response.Body.Close()
                        } else {
                            Printfln("Request Error: %v", err.Error())
                        }
                    } else {
                        Printfln("Request Init Error: %v", err.Error())
                    }
                }
            }
        }
    ====================================================================
    Output:
        Client 0: Request contains no cookies
        Client 0: Cookie Name: counter, Value: 1
        Client 0: Cookie Name: counter, Value: 2
        Client 1: Cookie Name: counter, Value: 3
        Client 1: Cookie Name: counter, Value: 4
        Client 1: Cookie Name: counter, Value: 5
        Client 2: Cookie Name: counter, Value: 6
        Client 2: Cookie Name: counter, Value: 7
        Client 2: Cookie Name: counter, Value: 8
████████████████████████████████████████████████████████████████████████
457.Managing Redirections
    By default, a Client will stop following redirections after ten requests, 
    but this can be changed by specifying a custom policy.

    example:
    The Contents of the server_redirects.go File in the httpclient Folder:
        package main
        import "net/http"
        func init() {
            http.HandleFunc("/redirect1",
                func(writer http.ResponseWriter, request *http.Request) {
                    http.Redirect(writer, request, "/redirect2",
                        http.StatusTemporaryRedirect)
                })
            http.HandleFunc("/redirect2",
                func(writer http.ResponseWriter, request *http.Request) {
                    http.Redirect(writer, request, "/redirect1",
                        http.StatusTemporaryRedirect)
                })
        }
    ====================================================================
    Sending a Request in the main.go File in the httpclient Folder:
        package main
        import (
            "net/http"
            "os"
            "io"
            "time"
            //"encoding/json"
            //"strings"
            //"net/url"
            //"net/http/cookiejar"
            //"fmt"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            req, err := http.NewRequest(http.MethodGet,
                "http://localhost:5000/redirect1", nil)
            if (err == nil) {
                var response *http.Response
                response, err = http.DefaultClient.Do(req)
                if (err == nil) {
                    io.Copy(os.Stdout, response.Body)
                } else {
                    Printfln("Request Error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Request Error: Get "/redirect1": stopped after 10 redirects
████████████████████████████████████████████████████████████████████████
458.Defining a Custom Redirection Policy in the main.go
    example:
    main.go:
        package main
        import (
            "net/http"
            "os"
            "io"
            "time"
            //"encoding/json"
            //"strings"
            "net/url"
            //"net/http/cookiejar"
            //"fmt"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            http.DefaultClient.CheckRedirect = func(req *http.Request,
                previous []*http.Request) error {
                if len(previous) == 3 {
                    url, _ := url.Parse("http://localhost:5000/html")
                    req.URL = url
                }
                return nil
            }
            req, err := http.NewRequest(http.MethodGet,
                "http://localhost:5000/redirect1", nil)
            if (err == nil) {
                var response *http.Response
                response, err = http.DefaultClient.Do(req)
                if (err == nil) {
                    io.Copy(os.Stdout, response.Body)
                } else {
                    Printfln("Request Error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        <!DOCTYPE html>
        <html>
            <head>
                <title>Pro Go</title>
                <meta name="viewport" content="width=device-width" />
            </head>
            <body>
                <h1>Hello, World</div>
            </body>
        </html>
████████████████████████████████████████████████████████████████████████
459.Creating Multipart Forms:
    The mime/multipart package can be used to create a request body encoded as multipart/form-data, which
    allows a form to safely contain binary data, such as the contents of a file.
████████████████████████████████████████████████████████████████████████
460.The Contents of the server_forms.go File in the httpclient Folder
    example:
    server_forms.go:
        package main
        import (
            "net/http"
            "fmt"
            "io"
        )
        func init() {
            http.HandleFunc("/form",
                func (writer http.ResponseWriter, request *http.Request) {
                    err := request.ParseMultipartForm(10000000)
                    if (err == nil) {
                        for name, vals := range request.MultipartForm.Value {
                            fmt.Fprintf(writer, "Field %v: %v\n", name, vals)
                        }
                        for name, files := range request.MultipartForm.File {
                            for _, file := range files {
                                fmt.Fprintf(writer, "File %v: %v\n", name, file.Filename)
                                if f, err := file.Open(); err == nil {
                                    defer f.Close()
                                    io.Copy(writer, f)
                                }
                            }
                        }
                    } else {
                        fmt.Fprintf(writer, "Cannot parse form %v", err.Error())
                    }
                })
        }
    ====================================================================
    Output:
        <!DOCTYPE html>
        <html>
            <head>
                <title>Pro Go</title>
                <meta name="viewport" content="width=device-width" />
            </head>
            <body>
                <h1>Hello, World</div>
            </body>
        </html>
████████████████████████████████████████████████████████████████████████
461.The multipart.Writer Constructor Function
    Name                    Description
    ---------------         ----------------------------
    NewWriter(writer)       This function creates a new multipart.Writer that writes form data to the specified io.Writer.

████████████████████████████████████████████████████████████████████████
462.The multipart.Writer Methods
    Name                                    Description
    ------------------------                ---------------------------------------
    CreateFormField(fieldname)              This method creates a new form field with the specified name. The results
                                            are an io.Writer that is used to write the field data and an error that
                                            reports problems creating the field.
    CreateFormFile(fieldname, filename)     This method creates a new file field with the specified field name and file
                                            name. The results are an io.Writer that is used to write the field data and
                                            an error that reports problems creating the field.
    FormDataContentType()                   This method returns a string that is used to set the Content-Type request
                                            header and includes the string that denotes the boundaries between the parts of the form.
    Close()                                 This function finalizes the form and writes the terminating boundary that
                                            denotes the end of the form data.
████████████████████████████████████████████████████████████████████████
463.Creating and Sending a Multipart Form in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
            "os"
            "time"
            //"encoding/json"
            //"strings"
            //"net/url"
            //"net/http/cookiejar"
            //"fmt"
            "bytes"
            "mime/multipart"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            var buffer bytes.Buffer
            formWriter := multipart.NewWriter(&buffer)
            fieldWriter, err := formWriter.CreateFormField("name")
            if err == nil {
                io.WriteString(fieldWriter, "Alice")
            }
            fieldWriter, err = formWriter.CreateFormField("city")
            if err == nil {
                io.WriteString(fieldWriter, "New York")
            }
            fileWriter, err := formWriter.CreateFormFile("codeFile", "printer.go")
            if err == nil {
                fileData, err := os.ReadFile("./printer.go")
                if err == nil {
                    fileWriter.Write(fileData)
                }
            }
            formWriter.Close()
            req, err := http.NewRequest(http.MethodPost,
                "http://localhost:5000/form", &buffer)
            req.Header["Content-Type"] = []string{formWriter.FormDataContentType()}
            if err == nil {
                var response *http.Response
                response, err = http.DefaultClient.Do(req)
                if err == nil {
                    io.Copy(os.Stdout, response.Body)
                } else {
                    Printfln("Request Error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Field name: [Alice]
        Field city: [New York]
        File codeFile: printer.go
        package main
        
        import "fmt"
        
        func Printfln(template string, values ...interface{}) {
                fmt.Printf(template+"\n", values...)
        }
    
    Caution Don't use the defer keyword on the call to the Close method; otherwise, the final boundary string
    won't be added to the form until after the request will be sent, producing a form that not all servers will process.
    It is important to call the Close method before sending the request.
████████████████████████████████████████████████████████████████████████


`,
}