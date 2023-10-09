package features

var TitleHTTPServers = []string{
	"ALL HTTP SERVERS",
    "ALLHTTPSERVERS",
    "ALL CREATING HTTP SERVERS",
    "ALLCREATINGHTTPSERVERS",
}


var OriginalHTTPServers = DataBase{
	Alldatafield: `
403.Creating HTTP Servers
    What are they?
    The features described in this chapter make it easy for Go applications to create HTTP servers.

    Why are they useful?
    HTTP is one of the most widely used protocols and is useful for both user-facing
    applications and web services.

    How is it used?
    The features of the net/http package are used to create a server and handle requests.

    Are there any pitfalls or limitations?
    These features are well-designed and easy to use.

    Are there any alternatives?
    The standard library includes support for other network protocols and also for
    opening and using lower-level network connections. 
    See https://pkg.go.dev/net@go1.17.1 for details of the net package and its subpackages, 
    such as net/smtp, for example, which implements the SMTP protocol.

    Problem                                 Solution
    --------                                -------------
    Create an HTTP or HTTPS server          Use the ListenAndServe or ListenAndServeTLS functions
    Inspect an HTTP request                 Use the features of the Request struct
    Produce a response                      Use the ResponseWriter interface or the
                                            convenience functions

    Handle requests to specific URLs        Use the integrated router
    Serve static content                    Use the FileServer and StripPrefix function
    Use a template to produce a response    Write the content to the ResponseWriter
    or produce a JSON response

    Handle form data                        Use the Request methods
    Set or read cookies                     Use the Cookie, Cookies, and SetCookie methods

    Preparing for This Chapter:
    1- go mod init httpserver
    2- printer.go:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
    3- product.go:
        package main
        type Product struct {
            Name, Category string
            Price          float64
        
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
    4- main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
    =================================
    Output:
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
404.Creating a Simple HTTP Server
    example:
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
        =================================
        Output: go run .
        in Web Browser, search localhost:5000/
            Hello, World

████████████████████████████████████████████████████████████████████████
405.Creating the HTTP Listener and Handler
    The net/http Convenience Functions

    The ListenAndServe function starts listening for HTTP requests on a specified network address. 
    The ListenAndServeTLS function does the same for HTTP requests.
    The addresses accepted by the functions can be used to restrict the HTTP server so that it
    only accepts requests on a specific interface or to listen for requests on any interface.

    Name                                            Description
    ---------                                       ----------------------------------------
    ListenAndServe(addr, handler)                   This function starts listening for HTTP requests on a specified
                                                    address and passes requests onto the specified handler.
    ListenAndServeTLS(addr, cert, key, handler)     This function starts listening for HTTPS requests. The arguments are the address

████████████████████████████████████████████████████████████████████████
406.The Method Defined by the Handler Interface
    Name                            Description
    ServeHTTP(writer, request)      This method is invoked to process a HTTP request. The request is described by a
                                    Request value, and the response is written using a ResponseWriter, both of which
                                    are received as parameters.
████████████████████████████████████████████████████████████████████████
407.Inspecting the Request
    The Basic Fields Defined by the Request Struct
    Name        Description
    -------     ------------------------
    Method      This field provides the HTTP method (GET, POST, etc.) as a string. The net/http package defines
                constants for the HTTP methods, such as MethodGet and MethodPost.
    URL         This field returns the requested URL, expressed as a URL value.
    Proto       This field returns a string that indicates the version of HTTP used for the request.
    Host        This field returns a string containing the requested hos.
    Header      This field returns a Header value, which is an alias to map[string][]string and contains the
                request headers. The map keys are the names of the headers, and the values are string slices
                containing the header values.
    Trailer     This field returns a map[string]string that contains any additional headers that are included in
                the request after the body.
    Body        This filed returns a ReadCloser, which is an interface that combines the Read method of the
                Reader interface with the Close method of the Closer interface
████████████████████████████████████████████████████████████████████████
408.Writing Request Fields in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Method: %v", request.Method)
            Printfln("URL: %v", request.URL)
            Printfln("HTTP Version: %v", request.Proto)
            Printfln("Host: %v", request.Host)
            for name, val := range request.Header {
                Printfln("Header: %v, Value: %v", name, val)
            }
            Printfln("---")
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================================
    Output: Compile and execute the project and request http://localhost:5000
        Method: GET
        URL: /
        HTTP Version: HTTP/1.1
        Host: localhost:5000
        Header: Accept, Value: [text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8]
        Header: Accept-Language, Value: [en-US,en;q=0.5]
        Header: Accept-Encoding, Value: [gzip, deflate, br]
        Header: Sec-Fetch-Mode, Value: [navigate]
        Header: Sec-Fetch-Site, Value: [none]
        Header: User-Agent, Value: [Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0]
        Header: Connection, Value: [keep-alive]
        Header: Upgrade-Insecure-Requests, Value: [1]
        Header: Sec-Fetch-Dest, Value: [document]
        Header: Sec-Fetch-User, Value: [?1]
        ---
        Method: GET
        URL: /favicon.ico
        HTTP Version: HTTP/1.1
        Host: localhost:5000
        Header: Accept-Encoding, Value: [gzip, deflate, br]
        Header: Connection, Value: [keep-alive]
        Header: Sec-Fetch-Dest, Value: [image]
        Header: Sec-Fetch-Mode, Value: [no-cors]
        Header: Sec-Fetch-Site, Value: [same-origin]
        Header: User-Agent, Value: [Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0]
        Header: Accept, Value: [image/avif,image/webp,*/*]
        Header: Accept-Language, Value: [en-US,en;q=0.5]
        Header: Referer, Value: [http://localhost:5000/]
        ---
████████████████████████████████████████████████████████████████████████
409.Save All Logs in a txt file:
    Server.go:
        package server
        import (
            "io"
            "log"
            "net/http"
            "os"
        )
        type StringHandler struct {
            Message string
        }
        func MyServer() {
            logFile, err := os.OpenFile("Server/LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer logFile.Close()
            log.SetOutput(logFile)
        
            err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
            if err != nil {
                log.Printf("Error: %v", err.Error())
            }
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            log.Printf("Method: %v", request.Method)
            log.Printf("URL: %v", request.URL)
            log.Printf("HTTP Version: %v", request.Proto)
            log.Printf("Host: %v", request.Host)
            for name, val := range request.Header {
                log.Printf("Header: %v, Value: %v", name, val)
            }
            io.WriteString(writer, sh.Message)
            log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
        }
    Output: Create a TXT file in Server/LogFile/logs.txt and save appendation in to it.
████████████████████████████████████████████████████████████████████████
410.Filtering Requests and Generating Responses
    Useful Fields and Methods Defined by the URL Struct
    The ResponseWriter interface defines the methods that are available when creating a response.
    Name        Description
    -------     -----------------------------
    Scheme      This field returns the scheme component of the URL.
    Host        This field returns the host component of the URL, which may include the port.
    RawQuery    This field returns the query string from the URL. Use the Query method to process the query
                string into a map.
    Path        This field returns the path component of the URL.
    Fragment    This field returns the fragment component of the URL, without the # character.
    Hostname()  This method returns the hostname component of the URL as a string.
    Port()T     his method returns the port component of the URL as a string.
    Query()     This method returns a map[string][]string (a map with string keys and string slice
                values), containing the query string fields.
    User()      This method returns the user information associated with the request.
    String()    This method returns a string representation of the URL.
████████████████████████████████████████████████████████████████████████
411.The ResponseWriter Methods
    Name                Description
    Header()            This method returns a Header, which is an alias to map[string][]string, that can be
                        used to set the response headers.
    WriteHeader(code)   This method sets the status code for the response, specified as an int. The net/http
                        package defines constants for most status codes.
    Write(data)         This method writes data to the response body and implements the Writer interface.
████████████████████████████████████████████████████████████████████████
412.Producing Difference Responses in the main.go
    example:
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            if (request.URL.Path == "/favicon.ico") {
                Printfln("Request for icon detected - returning 404")
                writer.WriteHeader(http.StatusNotFound)
                return
            }
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output: in Terminal
        Request for /
        Request for icon detected - returning 404
████████████████████████████████████████████████████████████████████████
413.Get Logs and Handle request if URL not Exist:
        example:
        main.go:
            package main
            import (
                "io"
                "log"
                "net/http"
                "os"
            )
            type StringHandler struct {
                Message string
            }
            func main() {
                logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                    log.Fatal(err)
                }
                defer logFile.Close()
                log.SetOutput(logFile)
            
                err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
                if err != nil {
                    log.Printf("Error: %v", err.Error())
                }
            }
            func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
                if (request.URL.Path != "/") {
                    Printfln("Request for icon detected - returning 404")
                    writer.WriteHeader(http.StatusNotFound)
                    return
                }
                Printfln("Request for %v", request.URL.Path)
                io.WriteString(writer, sh.Message)
            
             
            
                log.Printf("Method: %v", request.Method)
                log.Printf("URL: %v", request.URL)
                log.Printf("HTTP Version: %v", request.Proto)
                log.Printf("Host: %v", request.Host)
                for name, val := range request.Header {
                    log.Printf("Header: %v, Value: %v", name, val)
                }
                io.WriteString(writer, sh.Message)
                log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
            }
        Output: so we have to logs, one of that is logs about user request,
        another that is about path URL in Terminal
████████████████████████████████████████████████████████████████████████
414.Using the Response Convenience Functions
    Name                                    Description
    -----------------------                 ----------------------------
    Error(writer, message, code)            This function sets the header to the specified code, sets the Content-Type header
                                            to text/plain, and writes the error message to the response. The X-Content-
                                            Type-Options header is also set to stop browsers from interpreting the response as
                                            anything other than text.
    NotFound(writer, request)               This function calls Error and specifies a 404 error code.
    Redirect(writer, request, url, code)    This function sends a redirection response to the specified URL and with the
                                            specified status code.
    ServeFile(writer, request, fileName)    This function sends a response containing the contents of the specified file. The
                                            Content-Type header is set based on the file name but can be overridden by
                                            explicitly setting the header before calling the function. See the “Creating a Static
                                            HTTP Server” section for an example that serves files.
████████████████████████████████████████████████████████████████████████
415.Convenience Functions in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "log"
            "net/http"
            "os"
        )
        type StringHandler struct {
            Message string
        }
        func main() {
            logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer logFile.Close()
            log.SetOutput(logFile)
            err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
            if err != nil {
                log.Printf("Error: %v", err.Error())
            }
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            switch request.URL.Path {
            case "/favicon.ico":
                http.NotFound(writer, request)
            case "/message":
                io.WriteString(writer, sh.Message)
            default:
                http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
            }
            log.Printf("Method: %v", request.Method)
            log.Printf("URL: %v", request.URL)
            log.Printf("HTTP Version: %v", request.Proto)
            log.Printf("Host: %v", request.Host)
            for name, val := range request.Header {
                log.Printf("Header: %v, Value: %v", name, val)
            }
            io.WriteString(writer, sh.Message)
            log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
        }
    Output: Will write in Terminal and File for feture analyzing.
    
████████████████████████████████████████████████████████████████████████
416.Using the Convenience Functions in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            switch request.URL.Path {
            case "/favicon.ico":
                http.NotFound(writer, request)
            case "/message":
                io.WriteString(writer, sh.message)
            default:
                http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
            }
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        every wrong links redirect to specific link.
        uses a switch statement to decide how to respond to a request.
████████████████████████████████████████████████████████████████████████
417.Using the Convenience Routing Handler
    The process of inspecting the URL and selecting a response can produce complex code that is difficult to
    read and maintain. 
    To simplify the process, the net/http package provides a Handler implementation that
    allows matching the URL to be separated from producing a request.

    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{"Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            err := http.ListenAndServe(":5000", nil)
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        don't show wrong search.
████████████████████████████████████████████████████████████████████████
418.The net/http Functions for Creating Routing Rules
    Name                                Description
    -----------------                   ---------------------------------
    Handle(pattern, handler)            This function creates a rule that invokes the specified ServeHTTP method of the
                                        specified Hander for requests that match the pattern.
    HandleFunc(pattern, handlerFunc)    This function creates a rule that invokes the specified function for requests that match
                                        the pattern. The function is invoked with ResponseWriter and Request arguments.
████████████████████████████████████████████████████████████████████████
419.he net/http Functions for Creating Request Handlers
    Name                                        Description
    FileServer(root)                            This function creates a Handler that produces responses using the ServeFile
                                                function. See the “Creating a Static HTTP Server” section for an example that
                                                serves files.
    NotFoundHandler()                           This function creates a Handler that produces responses using the NotFound function.
    RedirectHandler(url, code)                  This function creates a Handler that produces responses using the Redirect function.
    StripPrefix(prefix, handler)                This function creates a Handler that removes the specified prefix from the
                                                request URL and passes on the request to the specified Handler. See the
                                                “Creating a Static HTTP Server” section for details.
    TimeoutHandler(handler, duration, message)  This function passes on the request to the specified Handler but generates an error
                                                response if the response hasn't been produced within the specified duration.

████████████████████████████████████████████████████████████████████████
420.Supporting HTTPS Requests
    The net/http package provides integrated support for HTTPS. 
    To prepare for HTTPS, you will need to add
    two files to the httpserver folder: 
        a certificate file and a private key file.

████████████████████████████████████████████████████████████████████████
421.Getting Certificates for HTTPS
    A good way to get started with HTTPS is with a self-signed certificate, 
    which can be used for development and testing. 
    If you don't already have a self-signed certificate, 
    then you can create one online using sites such as 
    
    https://getacert.com 
    or 
    https://www.selfsignedcertificate.com
    
    both of which will let you create a self-signed certificate easily 
    and without charge.

    Two files are required to use HTTPS, regardless of whether your certificate 
    is self-signed or not. The first is the certificate file, 
    which usually has a cer or cert file extension. 
    The second is the private key file, which usually has a key file extension.

    after ready to deploy:
        https://letsencrypt.org

    The ListenAndServeTLS function is used to enable HTTPS, 
    where the additional arguments specify the
    certificate and private key files, 
    which are named certificate.cer and certificate.key in my project
████████████████████████████████████████████████████████████████████████
422.Enabling HTTPS in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{"Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
        
            go func() {
                err := http.ListenAndServeTLS(":5500", "certificate.cer",
                    "certificate.key", nil)
                if err != nil {
                    Printfln("HTTPS Error: %v", err.Error())
                }
            }()
        
            err := http.ListenAndServe(":5000", nil)
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        HTTPS Error: open certificate.cer: no such file or directory
        Request for /message
████████████████████████████████████████████████████████████████████████
423.Redirecting HTTP Requests to HTTPS
    A common requirement when creating web servers is to redirect HTTP requests to the HTTPS port. 
    This can be done by creating a custom handler
    
    example:
    Redirecting to HTTPS in the main.go:
        package main
        import (
            "net/http"
            "io"
            "strings"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func HTTPSRedirect(writer http.ResponseWriter,
                request *http.Request) {
            host := strings.Split(request.Host, ":")[0]
            target := "https://" + host + ":5500" + request.URL.Path
            if len(request.URL.RawQuery) > 0 {
                target += "?" + request.URL.RawQuery
            }
            http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            go func () {
                err := http.ListenAndServeTLS(":5520", "certificate.cer",
                    "certificate.key", nil)
                if (err != nil) {
                    Printfln("HTTPS Error: %v", err.Error())
                }
            }()
            err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        Not work for my Browser but you have to try, maybe work it for you.
████████████████████████████████████████████████████████████████████████
424.Creating a Static HTTP Server
    The net/http package includes built-in support for responding to requests with the contents of files. 
    To prepare for the static HTTP server, 
    create the httpserver/static folder and add to it a file named index.html

    example:
    static/index.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2">
                Hello, World
            </div>
        </body>
        </html>
    
    store.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2 text-center">
                Products
            </div>
            <table class="table table-sm table-bordered table-striped">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Category</th>
                        <th>Price</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Kayak</td>
                        <td>Watersports</td>
                        <td>$279.00</td>
                    </tr>
                    <tr>
                        <td>Lifejacket</td>
                        <td>Watersports</td>
                        <td>$49.95</td>
                    </tr>
                </tbody>
            </table>
        </body>
        </html>

        The HTML files depend on the Bootstrap CSS package to style the HTML content. Run the command
        shown in here in the httpserver folder to download the Bootstrap CSS file into the static folder.
        (You may have to install the curl command.):

            curl https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css --output static/bootstrap.min.css

        Output:
            ootstrap.min.css
            % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                            Dload  Upload   Total   Spent    Left  Speed
            100  152k    0  152k    0     0   163k      0 --:--:-- --:--:-- --:--:--  163k
████████████████████████████████████████████████████████████████████████
425.Creating the Static File Route
    example:
    Defining a Route in the main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            fsHandler := http.FileServer(http.Dir("./static"))
            http.Handle("/files/", http.StripPrefix("/files", fsHandler))
            err := http.ListenAndServe(":5000", nil)
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        redirect from => https://localhost:5500/files/store.html
                   to => static/store.html
                   
████████████████████████████████████████████████████████████████████████
426.Using Templates to Generate Responses
    example:
    templates/products.html:
        <!DOCTYPE html>
        <html>
        <head>
            <meta name="viewport" content="width=device-width" />
            <title>Pro Go</title>
            <link rel="stylesheet" href="/files/bootstrap.min.css">
        </head>
        <body>
            <h3 class="bg-primary text-white text-center p-2 m-2">Products</h3>
            <div class="p-2">
                <table class="table table-sm table-striped table-bordered">
                    <thead>
                        <tr>
                            <th>Index</th>
                            <th>Name</th>
                            <th>Category</th>
                            <th class="text-end">Price</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range $index, $product := .Data }}
                        <tr>
                            <td>{{ $index }}</td>
                            <td>{{ $product.Name }}</td>
                            <td>{{ $product.Category }}</td>
                            <td class="text-end">
                                {{ printf "$%.2f" $product.Price }}
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
            </div>
        </body>
        </html>
    
    dynamic.go:
        package main
        import (
            "html/template"
            "net/http"
            "strconv"
        )
        type Context struct {
            Request *http.Request
            Data []Product
        }
        var htmlTemplates *template.Template
        func HandleTemplateRequest(writer http.ResponseWriter, request *http.Request) {
            path := request.URL.Path
            if (path == "") {
                path = "products.html"
            }
            t := htmlTemplates.Lookup(path)
            if (t == nil) {
                http.NotFound(writer, request)
            } else {
                err := t.Execute(writer, Context{  request, Products})
                if (err != nil) {
                    http.Error(writer, err.Error(), http.StatusInternalServerError)
                }
            }
        }
        func init() {
            var err error
            htmlTemplates = template.New("all")
            htmlTemplates.Funcs(map[string]interface{} {
                "intVal": strconv.Atoi,
            })
            htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
            if (err == nil) {
                http.Handle("/templates/", http.StripPrefix("/templates/",
                    http.HandlerFunc(HandleTemplateRequest)))
            } else {
                panic(err)
            }
        }

    Output:
        The initialization function loads all the templates with the html extension in the templates folder and
        sets up a route so that requests that start with /templates/ are processed by the HandleTemplateRequest
        function. This function looks up the template, falling back to the products.html file if no file path is
        specified, executes the template, and writes the response.


████████████████████████████████████████████████████████████████████████
427.Understanding Content Type Sniffing
    which implements the MIME Sniffing algorithm defined by: 
        
        https://mimesniff.spec.whatwg.org 
    
    The sniffing process can't detect every content type, 
    but it does well with standard web types, such as HTML, CSS, and JavaScript.
    The DetectContentType function returns a MIME type, 
    which is used as the value for the Content-Type header.

████████████████████████████████████████████████████████████████████████
428.Responding with JSON Data
    JSON responses are widely used in web services, 
    which provide access to an application's data for clients
    that don't want to receive HTML, such as Angular or React JavaScript clients.

    example:
    json.go:
        package main
        import (
            "net/http"
            "encoding/json"
        )
        func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
            writer.Header().Set("Content-Type", "application/json")
            json.NewEncoder(writer).Encode(Products)
        }
        func init() {
            http.HandleFunc("/json", HandleJsonRequest)
        }
    =========================================================================
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

            fsHandler := http.FileServer(http.Dir("./static"))
            http.Handle("/files/", http.StripPrefix("/files", fsHandler))

            err := http.ListenAndServe(":5000", nil)
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================
    Output:
        The initialization function creates a route, 
        which means that requests for /json will be processed by the
        HandleJsonRequest function.
████████████████████████████████████████████████████████████████████████
429.Handling Form Data
    The net/http package provides support for easily receiving and processing form data.

    This template makes use of template variables, expressions, 
    and functions to get the query string from
    the request and select the first index value, 
    which is converted to an int and used to retrieve a Product value
    from the data provided to the template.

    example:
    edit.html:
        <!DOCTYPE html>
        <html>
        <head>
            <meta name="viewport" content="width=device-width" />
            <title>Pro Go</title>
            <link rel="stylesheet" href="/files/bootstrap.min.css">
        </head>
        <body>
            {{ $index := intVal (index (index .Request.URL.Query "index") 0) }}
            {{ if lt $index (len .Data)}}
            {{ with index .Data $index}}
            <h3 class="bg-primary text-white text-center p-2 m-2">Product</h3>
            <form method="POST" action="/forms/edit" class="m-2">
                <div class="form-group">
                    <label>Index</label>
                    <input name="index" value="{{$index}}" class="form-control" disabled />
                    <input name="index" value="{{$index}}" type="hidden" />
                </div>
                <div class="form-group">
                    <label>Name</label>
                    <input name="name" value="{{.Name}}" class="form-control" />
                </div>
                <div class="form-group">
                    <label>Category</label>
                    <input name="category" value="{{.Category}}" class="form-control" />
                </div>
                <div class="form-group">
                    <label>Price</label>
                    <input name="price" value="{{.Price}}" class="form-control" />
                </div>
                <div class="mt-2">
                    <button type="submit" class="btn btn-primary">Save</button>
                    <a href="/templates/" class="btn btn-secondary">Cancel</a>
                </div>
            </form>
            {{ end }}
            {{ else }}
            <h3 class="bg-danger text-white text-center p-2">
                No Product At Specified Index
            </h3>
            {{end }}
        </body>
        </html>
████████████████████████████████████████████████████████████████████████
430.Reading Form Data from Requests
    The Request Form Data Fields and Methods


    Name                Description
    --------            ------------------------------
    Form                This field returns a map[string][]string containing the parsed form data and the
                        query string parameters. The ParseForm method must be called before this field is read.
    PostForm            This field is similar to Form but excludes the query string parameters so that only
                        data from the request body is contained in the map. The ParseForm method must
                        be called before this field is read.
    MultipartForm       This field returns a multipart form represented using the Form struct defined in the
                        mime/multipart package. The ParseMultipartForm method must be called before
                        this field is read.
    FormValue(key)      This method returns the first value for the specified form key and returns the
                        empty string if there is no value. The source of data for this method is the Form
                        field, and calling the FormValue method automatically calls ParseForm or
                        ParseMultipartForm to parse the form.
    PostFormValue(key)  This method returns the first value for the specified form key and returns the
                        empty string if there is no value. The source of data for this method is the PostForm
                        field, and calling the PostFormValue method automatically calls ParseForm or
                        ParseMultipartForm to parse the form.
    FormFile(key)       This method provides access to the first file with the specified key in the form.
                        The results are a File and FileHeader, both of which are defined in the mime/
                        multipart package, and an error. Calling this function causes the ParseForm or
                        ParseMultipartForm functions to be invoked to parse the form.
    ParseForm()         This method parses a form and populates the Form and PostForm fields. The result
                        is an error that describes any parsing problems.
    ParseMultipart      This method parses a MIME multipart form and populates the MultipartForm field.
    Form(max)           The argument specifies the maximum number of bytes to allocate to the form data,
                        and the result is an error that describes any problems processing the form.



    The init function sets up a new route so that the ProcessFormData function handles requests whose
    path is /forms/edit. Within the ProcessFormData function, the request method is checked, and the form
    data in the request is used to create a Product struct and replace the existing data value. In a real project,
    validating the data submitted in the form is essential, but for this chapter I trust that the form contains
    valid data.
    Processing form data:


    https://localhost:5500/templates/edit.html?index=2

    example:
    The Contents of the forms.go File in the httpserver Folder:
        package main
        import (
            "net/http"
            "strconv"
        )
        func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
            if (request.Method == http.MethodPost) {
                index, _ := strconv.Atoi(request.PostFormValue("index"))
                p := Product {}
                p.Name = request.PostFormValue("name")
                p.Category = request.PostFormValue("category")
                p.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
                Products[index] = p
            }
            http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
        }
        func init() {
            http.HandleFunc("/forms/edit", ProcessFormData)
        }
    
    Output:
        without view of upload you can add data to templates URL.
        
████████████████████████████████████████████████████████████████████████
431.Reading Multipart Forms
    example:
    upload.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2 text-center">
                Upload File
            </div>
            <form method="POST" action="/forms/upload" class="p-2"
                    enctype="multipart/form-data">
                <div class="form-group">
                    <label class="form-label">Name</label>
                    <input class="form-control" type="text" name="name">
                </div>
                <div class="form-group">
                    <label class="form-label">City</label>
                    <input class="form-control" type="text" name="city">
                </div>
                <div class="form-group">
                    <label class="form-label">Choose Files</label>
                    <input class="form-control" type="file" name="files" multiple>
                </div>
                <button type="submit" class="btn btn-primary mt-2">Upload</button>
            </form>
        </body>
        </html>
    ====================================================================================
    upload.go:
        package main
        import (
            "fmt"
            "io"
            "net/http"
        )
        func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
            fmt.Fprintf(writer, "Name: %v, City: %v\n", request.FormValue("name"),
                request.FormValue("city"))
            fmt.Fprintln(writer, "------")
            file, header, err := request.FormFile("files")
            if err == nil {
                defer file.Close()
                fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
                for k, v := range header.Header {
                    fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
                }
                fmt.Fprintln(writer, "------")
                io.Copy(writer, file)
            } else {
                http.Error(writer, err.Error(), http.StatusInternalServerError)
            }
        }
        func init() {
            http.HandleFunc("/forms/upload", HandleMultipartForm)
        }
    Output: Search in Browser: http://localhost:5000/files/upload.html
        You Can Upload

████████████████████████████████████████████████████████████████████████
432.The FileHeader Fields and Method
    Name        Description
    -------     ------------------------------
    Name        This field returns a string containing the name of the file.
    Size        This field returns an int64 containing the size of the file.
    Header      This field returns a map[string][]string, which contains the headers for the MIME part that
                contains the file.
    Open()      This method returns a File that can be used to read the content associated with the header, as
                demonstrated in the next section.
████████████████████████████████████████████████████████████████████████
433.Reading and Setting Cookies
    The net/http Function for Setting Cookies
    Name                            Description
    ---------------                 ------------------------------------
    SetCookie(writer, cookie)       This function adds a Set-Cookie header to the specified ResponseWriter. The
                                    cookie is described using a pointer to a Cookie struct, which is described next.



    Cookies can be complex, and care must be taken to configure them correctly. 
    The detail of how cookies work is beyond the scope of this book, 
    but there is a good description available at 
        
        https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies 

    and a detailed breakdown of the cookie
    fields at: 

        https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie.

████████████████████████████████████████████████████████████████████████
434.The Fields Defined by the Cookie Struct
    Name        Description
    -------     --------------------------------
    Name        This field represents the name of the cookie, expressed as a string.
    Value       This field represents the cookie value, expressed as a string.
    Path        This optional field specifies the cookie path.
    Domain      This optional field specifies the host/domain to which the cookie will be set.
    Expires     This field specifies the cookie expiry, expressed as a time.Time value.
    MaxAge      This field specifies the number of seconds until the cookie expires, expressed as an int.
    Secure      When this bool field is true, the client will only send the cookie over HTTPS connections.
    HttpOnly    When this bool field is true, the client will prevent JavaScript code from accessing the cookie.
    SameSite    This field specifies the cross-origin policy for the cookie using the SameSite constants, which
                defines SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, and SameSiteNoneMode.
████████████████████████████████████████████████████████████████████████
435.The Request Methods for Cookies
    Name            Description
    -----------     ----------------
    Cookie(name)    This method returns a pointer to the Cookie value with the specified name and an error
                    that indicates when there is no matching cookie.
    Cookies()       This method returns a slice of Cookie pointers.


    example:
    cookies.go:
        package main
        import (
            "net/http"
            "fmt"
            "strconv"
        )
        func GetAndSetCookie(writer http.ResponseWriter, request *http.Request) {
            counterVal := 1
            counterCookie, err := request.Cookie("counter")
            if (err == nil) {
                counterVal, _ = strconv.Atoi(counterCookie.Value)
                counterVal++
            }
        http.SetCookie(writer, &http.Cookie{
                Name: "counter", Value: strconv.Itoa(counterVal),
            })
            if (len(request.Cookies()) > 0) {
                for _, c := range request.Cookies() {
                    fmt.Fprintf(writer, "Cookie Name: %v, Value: %v", c.Name, c.Value)
                }
            } else {
                fmt.Fprintln(writer, "Request contains no cookies")
            }
        }
        func init() {
            http.HandleFunc("/cookies", GetAndSetCookie)
        }
    ======================================================================================
    Compile and execute the project and use a browser to request http://localhost:5000/cookies
    Output:
        search-1:
            Request contains no cookies
        search-2:
            Cookie Name: counter, Value: 1
        search-3:
            Cookie Name: counter, Value: 2
        ...
	`,
}