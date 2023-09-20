package features

var TitleOfReadingWriting = []string{
	"READING",
	"WRITING",
    "READING AND WRITING DATA",
    "READINGANDWRITINGDATA",
	"READINGWRITINGDATA",
	"READING WRITING",
	"READINGWRITING",
	"READING WRITING DATAS",
	"READINGWRITINGDATAS",
	"READING & WRITING DATA",
	"READING&WRITINGDATA",
	"READING & WRITING DATAS",
	"READING&WRITINGDATAS",
	"READING & WRITING",
	"READING&WRITING",
}

type ReadingandWriting struct {
	allReadingandWriting string
}

var OriginalReadingandWriting ReadingandWriting = ReadingandWriting{
	allReadingandWriting: `
273.Reading and Writing Data
    These interfaces are used wherever data is read or written, which means that any
    source or destination for data can be treated in much the same way so that writing data to a file, for example,
    is just the same as writing data to a network connection.

    What are they?
    These interfaces define the basic methods required to read and write data.

    Why are they useful?
    This approach means that just about any data source can be used
    in the same way, while still allowing specialized features to be
    defined using the composition features.

    How is it used?
    The io package defines these interfaces, but the implementations
    are available from a range of other packages

    Are there any pitfalls or limitations?
    These interfaces don't entirely hide the detail of sources or
    destinations for data and additional methods are often required,
    provided by interfaces that build on Reader and Writer.

    Are there any alternatives?
    The use of these interfaces is optional, but they are hard to avoid
    because they are used throughout the standard library.

    The Reader and Writer interfaces are defined by the io package and 
    provide abstract ways to read and write data, 
    without being tied to where the data is coming from or going to.

    Preparing for This Chapter:
    product.go:
        package main
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    279,
        }
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
    printer.go:
        package main
        import (
            "fmt"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
    main.go:
        package main
        func main() {
            Printfln("Product: %v, Price : %v", Kayak.Name, Kayak.Price)
        }
████████████████████████████████████████████████████████████████████████
274.The Reader interface
    The Reader interface doesn't include any detail about where data comes from or how it is obtained—it
    just defines the Read method. The details are left to the types that implement the interface, and there are
    reader implementations in the standard library for different data sources.

    defines a single method:
    Name                Description
    --------------      -------------------------
    Read(byteSlice)     This method reads data into the specified []byte. The method returns the number of
                        bytes that were read, expressed as an int, and an error.
████████████████████████████████████████████████████████████████████████
275.io.Reader package
    example:
        package main
        import (
            "io"
            "strings"
        )
        func processData(reader io.Reader) {
            b := make([]byte, 2)
            for {
                count, err := reader.Read(b)
                if count > 0 {
                    Printfln("Read %v bytes: %v", count, string(b[0:count]))
                }
                if err == io.EOF {
                    break
                }
            }
        }
        func main() {
            r := strings.NewReader("Kayak")
            processData(r)
        }
    Output:
        Read 2 bytes: Ka
        Read 2 bytes: ya
        Read 1 bytes: k
████████████████████████████████████████████████████████████████████████
276.Writer interface
    Write(byteSlice)
        This method writes the data from the specified byte slice. The method returns the
        number of bytes that were written and an error. The error will be non-nil if the
        number of bytes written is less than the length of the slice.
        The Writer interface doesn't include any details of how the written data is stored, transmitted, or
        processed, all of which is left to the types that implement the interface.
████████████████████████████████████████████████████████████████████████
277.io.Writer
    example:
        package main
        import(
            "io"
            "strings"
            "asd/asd"
        )
        func processData(reader io.Reader, writer io.Writer) {
            b := make([]byte, 2)
            for {
                count, err := reader.Read(b)
                if count > 0 {
                    writer.Write(b[0:count])
                    asd.Printfln("Read %v bytes: %v", count, string(b[0:count]))
                }
                if err == io.EOF {
                    break
                }
            }
        }
        func main() {
            r := strings.NewReader("Kayak")
            var builder strings.Builder
            processData(r, &builder)
            asd.Printfln("String builder contents: %s", builder.String())
        }
    Output:
        Read 2 bytes: Ka
        Read 2 bytes: ya
        Read 1 bytes: k
        String builder contents: Kayak
████████████████████████████████████████████████████████████████████████
278.io.EOF
    The io package defines a special error named EOF, 
    which is used to signal when the Reader reaches the end of the data. 
    If the error result from the Read function is equal to the EOF error, 
    then I break out of the for
    loop that has been reading data from the Reader:
    ...
    if err == io.EOF {
        break
    }
    ...
████████████████████████████████████████████████████████████████████████
279.the Utility Functions for Readers and Writers
    توابع مفید برای خوانندگان و نویسندگان
    
    Functions in the io Package for Readng and Writing Data
    Name                            Description
    ------------------------        -------------------------------------------------------
    Copy(w, r)                      This function copies data from a Reader to a Writer until EOF is returned or
                                    another error is encountered. The results are the number of bytes copies and an
                                    error used to describe any problems.
    CopyBuffer(w, r, buffer)        This function performs the same task as Copy but reads the data into the
                                    specified buffer before it is passed to the Writer.
    CopyN(w, r, count)              This function copies count bytes from the Reader to the Writer. The results are
                                    the number of bytes copies and an error used to describe any problems.
    ReadAll(r)                      This function reads data from the specified Reader until EOF is reached. The
                                    results are a byte slice containing the read data and an error, which is used to
                                    describe any problems.
    ReadAtLeast(r, byteSlice, min)  This function reads at least the specified number of bytes from the reader,
                                    placing them into the byte slice. An error is reported if fewer bytes than specified
                                    are read.
    ReadFull(r, byteSlice)          This function fills the specified byte slice with data. The result is the number
                                    of bytes read and an error. An error will be reported if EOF was encountered
                                    before enough bytes to fill the slice were read.
    WriteString(w, str)             This function writes the specified string to a writer.
████████████████████████████████████████████████████████████████████████
280.io.Copy(writer, reader)
    example:
        package main
        import(
            "io"
            "asd/asd"
            "strings"
        )
        func processData(reader io.Reader, writer io.Writer) {
            count, err := io.Copy(writer, reader)
                if (err == nil) {
                    asd.Printfln("Read %v bytes", count)
                } else {
                    asd.Printfln("Error: %v", err.Error())
                }
        }
        func main() {

            r := strings.NewReader("Kayak .")
            var builder strings.Builder
            processData(r, &builder)
            asd.Printfln("String builder contents: %s", builder.String())
        }
    Output:
        Read 7 bytes
        String builder contents: Kayak .
████████████████████████████████████████████████████████████████████████
281.The io Package Functions for Specialized Readers and Writers
    Name                        Description
    -----------                 ----------------------------------------
    Pipe()                      This function returns a PipeReader and a PipeWriter, which can be used to connect
                                functions that require a Reader and a Writer, as described in the “Using Pipes” section.
    MultiReader(...readers)     This function defines a variadic parameter that allows an arbitrary number of Reader
                                values to be specified. The result is a Reader that passes on the content from each of
                                its parameters in the sequence they are defined, as described in the “Concatenating
                                Multiple Readers” section.
    MultiWriter(...writers)     This function defines a variadic parameter that allows an arbitrary number of Writer
                                values to be specified. The result is a Writer that sends the same data to all the
                                specified writers, as described in the “Combining Multiple Writers” section.
    LimitReader(r, limit)       This function creates a Reader that will EOF after the specified number of bytes, as
                                described in the “Limiting Read Data” section.
████████████████████████████████████████████████████████████████████████
282.Pipe
    Pipes are used to connect code that consumes data through a Reader 
    and code that produces code through a Writer.

    The GenerateData function defines a Writer parameter, which it uses to write bytes from a string.
    example:
    data.go:
        package main
        import (
            "io"
            "asd/asd"
        )
        func GenerateData(writer io.Writer) {
            data := []byte("Kayak, Lifejacket")
            writeSize := 4
            for i := 0; i < len(data); i += writeSize {
                    end := i + writeSize;
                    if (end > len(data)) {
                        end = len(data)
                    }
                    count, err := writer.Write(data[i: end])
                    asd.Printfln("Wrote %v byte(s): %v", count, string(data[i: end]))
                    if (err != nil)  {
                        asd.Printfln("Error: %v", err.Error())
                    }
                }
            }
        func ConsumeData(reader io.Reader) {
            data := make([]byte, 0, 10)
            slice := make([]byte, 2)
            for {
                count, err := reader.Read(slice)
                if (count > 0) {
                    asd.Printfln("Read data: %v", string(slice[0:count]))
                    data = append(data, slice[0:count]...)
                }
                if (err == io.EOF) {
                    break
                }
            }
            asd.Printfln("Read data: %v", string(data))
        }

    Notice the parentheses at the end of this statement. These are required when creating a goroutine for an
    anonymous function, but it is easy to forget them.
    main.go:
        package main
        import (
            "io"
        )
        func main() {
        
            pipeReader, pipeWriter := io.Pipe()
            go func() {
                GenerateData(pipeWriter)
                pipeWriter.Close()
            }()
            ConsumeData(pipeReader)
        }



    The output highlights the fact that pipes are synchronous. The GenerateData function calls the writer’s
    Write method and then blocks until the data is read. This is why the first message in the output is from the
    reader: the reader is consuming the data two bytes at a time, which means that two read operations are
    required before the initial call to the Write method, which is used to send four bytes, completes, and the
    message from the GenerateData function is displayed.
    Output:
        Read data: Ka
        Read data: ya
        Wrote 4 byte(s): Kaya
        Read data: k,
        Read data:  L
        Wrote 4 byte(s): k, L
        Read data: if
        Read data: ej
        Wrote 4 byte(s): ifej
        Read data: ac
        Read data: ke
        Wrote 4 byte(s): acke
        Read data: t
        Wrote 1 byte(s): t
        Read data: Kayak, Lifejacket
████████████████████████████████████████████████████████████████████████
283.PipeReader and a PipeWriter
    The io.Pipe function returns a PipeReader and a PipeWriter. The PipeReader and PipeWriter structs
    implement the Closer interface
    Name        Description
    -------     -----------------------
    Close()     This method closes the reader or writer. The details are implementation specific, but, in
                general, any subsequent reads from a closed Reader will return zero bytes and the EOF error,
                while any subsequent writes to a closed Writer will return an error.

    The PipeReader struct implements the Reader interface, which means I can use it as the argument to
    the ConsumeData function. The ConsumeData function is executed in the main goroutine, which means that
    the application won't exit until the function completes.
    The effect is that data is written into the pipe using the PipeWriter and read from the pipe using the
    PipeReader. When the GenerateData function is complete, the Close method is called on the PipeWriter,
    which causes the next read by the PipeReader to produce EOF.
████████████████████████████████████████████████████████████████████████
284.io.MultiReader()
    example:
    main.go:
        package main
        import (
            "io"
            "strings"
        )
        func main() {
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            ConsumeData(concatReader)
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Li
        Read data: fe
        Read data: ja
        Read data: ck
        Read data: et
        Read data: Ca
        Read data: no
        Read data: e
        Read data: KayakLifejacketCanoe
████████████████████████████████████████████████████████████████████████
285.io.MultiWriter()
example:
    package main
    import (
        "io"
        "strings"
        "asd/asd"
    )
    func main() {
        var w1 strings.Builder
        var w2 strings.Builder
        var w3 strings.Builder
        combinedWriter := io.MultiWriter(&w1, &w2, &w3)
        GenerateData(combinedWriter)
        asd.Printfln("Writer #1: %v", w1.String())
        asd.Printfln("Writer #2: %v", w2.String())
        asd.Printfln("Writer #3: %v", w3.String())
    }
Output:
    Wrote 4 byte(s): Kaya
    Wrote 4 byte(s): k, L
    Wrote 4 byte(s): ifej
    Wrote 4 byte(s): acke
    Wrote 1 byte(s): t
    Writer #1: Kayak, Lifejacket
    Writer #2: Kayak, Lifejacket
    Writer #3: Kayak, Lifejacket
████████████████████████████████████████████████████████████████████████
286.io.TeeReader(concatReader, &writer)
    Echoing Reads to a Writer

    The TeeReader function returns a Reader that echoes the data that it receives to a Writer.

    example:
        package main
        import (
            "asd/asd"
            "io"
            "strings"
        )
        func main() {
        
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            var writer strings.Builder
            teeReader := io.TeeReader(concatReader, &writer)
            ConsumeData(teeReader)
            asd.Printfln("Echo data: %v", writer.String())
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Li
        Read data: fe
        Read data: ja
        Read data: ck
        Read data: et
        Read data: Ca
        Read data: no
        Read data: e
        Read data: KayakLifejacketCanoe
        Echo data: KayakLifejacketCanoe
████████████████████████████████████████████████████████████████████████
287.io.LimitReader(concatReader, 5)
    The LimitReader function is used to restrict the amount of data that can be obtained from a Reader

    example:
        package main
        import (
            "io"
            "strings"
        )
        func main() {
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            limited := io.LimitReader(concatReader, 5)
            ConsumeData(limited)
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Kayak
████████████████████████████████████████████████████████████████████████
288.bufio
    The bufio package provides support for adding buffers to readers and writers.

    example:
        This code defined a struct type named CustomReader that acts as a wrapper around a Reader. 
        The implementation of the Read method generates output that reports how much data is read and
        how many read operations are performed overall. 
        custom.go:
            package main
            import (
                "io"
                "asd/asd"
            )
            type CustomReader struct {
                reader    io.Reader
                readCount int
            }
            func NewCustomReader(reader io.Reader) *CustomReader {
                return &CustomReader{reader, 0}
            }
            func (cr *CustomReader) Read(slice []byte) (count int, err error) {
                count, err = cr.reader.Read(slice)
                cr.readCount++
                asd.Printfln("Custom Reader: %v bytes", count)
                if err == io.EOF {
                    asd.Printfln("Total Reads: %v", cr.readCount)
                }
                return
            }

        main.go:
            package main
            import (
                "asd/asd"
                "io"
                "strings"
            )
            func main() {
                text := "It was a boat. A small boat."
                var reader io.Reader = NewCustomReader(strings.NewReader(text))
                var writer strings.Builder
                slice := make([]byte, 5)
                for {
                    count, err := reader.Read(slice)
                    if count > 0 {
                        writer.Write(slice[0:count])
                    }
                    if err != nil {
                        break
                    }
                }
                asd.Printfln("Read data: %v", writer.String())
            }
        The NewCustomreader function is used to create a CustomReader that reads from a string and uses a for
        loop to consume the data using a byte slice.
        Output:
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 3 bytes
            Custom Reader: 0 bytes
            Total Reads: 7
            Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
289.bufio Functions for Creating Buffered Readers
    Name                        Description
    ----------------------      ---------------------------------
    NewReader(r)                This function returns a buffered Reader with the default buffer size (which is
                                4,096 bytes at the time of writing).
    NewReaderSize(r, size)      This function returns a buffered Reader with the specified buffer size.
████████████████████████████████████████████████████████████████████████
290.bufio.NewReader(reader)
    The default buffer size is 4,096 bytes, which means that the buffered reader was able to read all the data
    in a single read operation, plus an additional read to produce the EOF result. Introducing the buffer reduces
    the overhead associated with the read operations, albeit at the cost of the memory used to buffer the data.

    example:
        package main
        import (
            "io"
            "strings"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var reader io.Reader = NewCustomReader(strings.NewReader(text))
            var writer strings.Builder
            slice := make([]byte, 5)
            reader = bufio.NewReader(reader)
            for {
                count, err := reader.Read(slice)
                if (count > 0) {
                    writer.Write(slice[0:count])
                }
                if (err != nil) {
                    break
                }
            }
            Printfln("Read data: %v", writer.String())
        }
Output:
    Custom Reader: 28 bytes
    Custom Reader: 0 bytes
    Total Reads: 2
    Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
291.Additional Buffered Reader Methods
    The NewReader and NewReaderSize functions return bufio.Reader values, which implement the io.
    Reader interface and which can be used as drop-in wrappers for other types of Reader methods, seamlessly
    introducing a read buffer.


████████████████████████████████████████████████████████████████████████
292.The Methods Defined by the Buffered Reader
    Name                Description
    --------------      ------------------------------------------
    Buffered()          This method returns an int that indicates the number of bytes that can be read from the buffer.
    Discard(count)      This method discards the specified number of bytes.
    Peek(count)         This method returns the specified number of bytes without removing them from the
                        buffer, meaning they will be returned by subsequent calls to the Read method.
    Reset(reader)       This method discards the data in the buffer and performs subsequent reads from the
                        specified Reader.
    Size()              This method returns the size of the buffer, expressed int.
████████████████████████████████████████████████████████████████████████
293.buffered.Read(slice)
    example:
        package main
        import (
            "io"
            "strings"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var reader io.Reader = NewCustomReader(strings.NewReader(text))
            var writer strings.Builder
            slice := make([]byte, 5)
            buffered := bufio.NewReader(reader)
            for {
                count, err := buffered.Read(slice)
                if (count > 0) {
                    Printfln("Buffer size: %v, buffered: %v",
                        buffered.Size(), buffered.Buffered())
                    writer.Write(slice[0:count])
                }
                if (err != nil) {
                        break
                    }
            }
            Printfln("Read data: %v", writer.String())
        }
    Output:
        Custom Reader: 28 bytes
        Buffer size: 4096, buffered: 23
        Buffer size: 4096, buffered: 18
        Buffer size: 4096, buffered: 13
        Buffer size: 4096, buffered: 8
        Buffer size: 4096, buffered: 3
        Buffer size: 4096, buffered: 0
        Custom Reader: 0 bytes
        Total Reads: 2
        Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
294.bufio Functions for Creating Buffered Writers
    Name                        Description
    -----------------------     --------------------------------------
    NewWriter(w)                This function returns a buffered Writer with the default buffer size (which is
                                4,096 bytes at the time of writing).
    NewWriterSize(w, size)      This function returns a buffered Writer with the specified buffer size.
████████████████████████████████████████████████████████████████████████
295.Methods Defined by the bufio.Writer Struct
    Name                Description
    --------------      -----------------------------------
    Available()         This method returns the number of available bytes in the buffer.
    Buffered()          This method returns the number of bytes that have been written to the buffer.
    Flush()             This method writes the contents of the buffer to the underlying Writer.
    Reset(writer)       This method discards the data in the buffer and performs subsequent writes to the
                        specified Writer.
    Size()              This method returns the capacity of the buffer in bytes.
████████████████████████████████████████████████████████████████████████
296.Defining a Custom Writer example
    The NewCustomWriter constructor wraps a Writer with a CustomWriter struct, which reports on its
    write operations.
    custom.go:
        package main
        import (
            "asd/asd"
            "io"
        )
        type CustomReader struct {
            reader    io.Reader
            readCount int
        }
        func NewCustomReader(reader io.Reader) *CustomReader {
            return &CustomReader{reader, 0}
        }
        func (cr *CustomReader) Read(slice []byte) (count int, err error) {
            count, err = cr.reader.Read(slice)
            cr.readCount++
            asd.Printfln("Custom Reader: %v bytes", count)
            if err == io.EOF {
                asd.Printfln("Total Reads: %v", cr.readCount)
            }
            return
        }
        type CustomWriter struct {
            writer     io.Writer
            writeCount int
        }
        func NewCustomWriter(writer io.Writer) *CustomWriter {
            return &CustomWriter{writer, 0}
        }
        func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
            count, err = cw.writer.Write(slice)
            cw.writeCount++
            asd.Printfln("Custom Writer: %v bytes", count)
            return
        }
        func (cw *CustomWriter) Close() (err error) {
            if closer, ok := cw.writer.(io.Closer); ok {
                closer.Close()
            }
            asd.Printfln("Total Writes: %v", cw.writeCount)
            return
        }    
    main.go:
        package main
        import (
            "strings"
            "asd/asd"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            var writer = NewCustomWriter(&builder)
            for i := 0; true; {
                end := i + 5
                if end >= len(text) {
                    writer.Write([]byte(text[i:]))
                    break
                }
                writer.Write([]byte(text[i:end]))
                i = end
            }
            asd.Printfln("Written data: %v", builder.String())
        }
    Output:
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 3 bytes
        Written data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
297.Using a Buffered Writer in the main.go example
    example:
    main.go:
        package main
        import (
            "strings"
            "asd/asd"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
            for i := 0; true; {
                end := i + 5
                if end >= len(text) {
                    writer.Write([]byte(text[i:]))
                    writer.Flush()
                    break
                }
                writer.Write([]byte(text[i:end]))
                i = end
            }
            asd.Printfln("Written data: %v", builder.String())
        }
    Output:
        Custom Writer: 20 bytes
        Custom Writer: 8 bytes
        Written data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
298.Scanning from a Reader
    example:
    main.go
        package main
        import (
            "io"
            "strings"
            "asd/asd"
            "fmt"
        )
        func scanFromReader(reader io.Reader, template string,
            vals ...interface{}) (int, error) {
            return fmt.Fscanf(reader, template, vals...)
        }
        func main() {
            reader := strings.NewReader("Kayak Watersports $279.00")
            var name, category string
            var price float64
            scanTemplate := "%s %s $%f"
            _, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
            if err != nil {
                asd.Printfln("Error: %v", err.Error())
            } else {
                asd.Printfln("Name: %v", name)
                asd.Printfln("Category: %v", category)
                asd.Printfln("Price: %.2f", price)
            }
        }
    Output:
        Name: Kayak
        Category: Watersports
        Price: 279.00
████████████████████████████████████████████████████████████████████████
299.Scanning Gradually اسکن به تدریج
    example:
    main.go
        package main
        import (
            "io"
            "strings"
            "asd/asd"
            "fmt"
        )
        func scanSingle(reader io.Reader, val interface{}) (int, error) {
            return fmt.Fscan(reader, val)
        }
        func main() {
            reader := strings.NewReader("Kayak Watersports $279.00")
            for {
                var str string
                _, err := scanSingle(reader, &str)
                if err != nil {
                    if err != io.EOF {
                        asd.Printfln("Error: %v", err.Error())
                    }
                    break
                }
                asd.Printfln("Value: %v", str)
            }
        }
    Output:
        Value: Kayak
        Value: Watersports
        Value: $279.00
████████████████████████████████████████████████████████████████████████
300.Writing Formatted Strings to a Writer
    The fmt package also provides functions for writing formatted strings to a Writer
    The writeFormatted function uses the fmt.Fprintf function to write a string formatted with a template
    to a Writer.

    example:
    main.go:
        package main
        import (
            "io"
            "strings"
            "fmt"
        )
        func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
            fmt.Fprintf(writer, template, vals...)
        }
        func main() {
            var writer strings.Builder
            template := "Name: %s, Category: %s, Price: $%.2f"
            writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
            fmt.Println(writer.String())
        }
    Output:
        Name: Kayak, Category: Watersports, Price: $279.00
████████████████████████████████████████████████████████████████████████
301.strings.Replacer struct
    The strings.Replacer struct can be used to perform replacements on a string and output the modified
    result to a Writer.

    example:
        package main
        import (
            "fmt"
            "io"
            "strings"
        )
        func writeReplaced(writer io.Writer, str string, subs ...string) {
            replacer := strings.NewReplacer(subs...)
            replacer.WriteString(writer, str)
        }
        func main() {
            text := "It was a boat. A small boat."
            subs := []string{"boat", "kayak", "small", "huge"}
            var writer strings.Builder
            writeReplaced(&writer, text, subs...)
            fmt.Println(writer.String())
        }    
    Output:
        It was a kayak. A huge kayak.
`,
}
