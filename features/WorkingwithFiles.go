package features

var TitleOfWorkingFiles = []string{
	"WORKING WITH FILES",
	"WORKINGWITHFILES",
	"WORKING WITH FILE",
	"WORKINGWITHFILE",
	"WORK WITH FILE",
	"WORKWITHFILE",
	"WORK WITH FILES",
	"WORKWITHFILES",
}


type WorkingWithFiles struct {
	allWorkWithFiles string
}

var OriginalWorkWithFiles = WorkingWithFiles{
	allWorkWithFiles: `
334.Working with Files
    Putting Working with Files in Context
        
    Answer                                  Question
    ---------------------------             -------------------------------------------------------
    What are they?                          These features provide access to the file system so that files can be read and written.
    Why are they useful?                    Files are used for everything from logging to configuration files.
    How are they used?                      These features are accessed through the os package, which provides platform-
                                            neutral access to the file system.
    Are there any pitfallsor limitations?   Some consideration of the underlying file system must be made, especially when
                                            dealing with paths.
    Are there any alternatives?             Go supports alternative ways of storing data, such as databases, but there are no
                                            alternative mechanisms for accessing files.

    Preparing
    printer.go:
        package main
        import (
            "fmt"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
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
    main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
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
335.The os Package Functions for Reading Files
    Name                Description
    -------------       -------------------------------------
    ReadFile(name)      This function opens the specified file and reads its contents. The results are a byte
                        slice containing the file content and an error indicating problems opening or reading
                        the file.
    Open(name)          This function opens the specified file for reading. The result is a File struct and an
                        error that indicates problems opening the file.

    One of the most common reasons to read a file is to load configuration data. The JSON format is well-
    suited for configuration files because it is simple to process, has good support in the Go standard library
    The Contents of the config.json File in the files Folder:
        {
            "Username": "Alice",
            "AdditionalProducts": [
                {"name": "Hat", "category": "Skiing", "price": 10},
                {"name": "Boots", "category":"Skiing", "price": 220.51 },
                {"name": "Gloves", "category":"Skiing", "price": 40.20 }
            ]
        }
████████████████████████████████████████████████████████████████████████
336.os.ReadFile()
    The LoadConfig function uses the ReadFile function to read the contents of the config.json file. 
    The file will be read from the current working directory when the application is executed, 
    which means that I can open the file just with its name.

    The contents of the file are returned as a byte slice, which is converted to a string and written out. The
    LoadConfig function is invoked by an initialization function, which ensures the configuration file is read.

    example:
    readconfig.go:
        package main
        import "os"
        func LoadConfig() (err error) {
            data, err := os.ReadFile("config.json")
            if err == nil {
                Printfln(string(data))
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            }
        }
    Output:
        {
            "Username": "Alice",
            "AdditionalProducts": [
                {"name": "Hat", "category": "Skiing", "price": 10},
                {"name": "Boots", "category":"Skiing", "price": 220.51 },
                {"name": "Gloves", "category":"Skiing", "price": 40.20 }
            ]
        }
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
337.Decoding the JSON Data
    example:
    readconfig.go:
        package main
        import (
            "encoding/json"
            "os"
            "strings"
        )
        type ConfigData struct {
            UserName           string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            data, err := os.ReadFile("config.json")
            if err == nil {
                decoder := json.NewDecoder(strings.NewReader(string(data)))
                err = decoder.Decode(&Config)
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            } else {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
        Product: Hat, Category: Skiing, Price: $10.00
        Product: Boots, Category: Skiing, Price: $220.51
        Product: Gloves, Category: Skiing, Price: $40.20
████████████████████████████████████████████████████████████████████████
338.os.Open("config.json")
    Using the File Struct to Read a File
    The Open function opens a file for reading and returns a File value, which represents the open file, and an
    error, which is used to indicate problems opening the file. The File struct implements the Reader interface,
    which makes it simple to read and process the example JSON data, without reading the entire file into a byte
    slice.
    The File struct also implements the Closer interface, which defines a Close method.
    
    The defer keyword can be used to call the Close method when the enclosing function completes,
    like this:
        defer file.Close()
    using the defer keyword ensures that the file is closed even when a function returns early.

    example:
    readconfig.go:
        package main
        import (
            "encoding/json"
            "os"
            "time"
        )
        type ConfigData struct {
            UserName           string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            file, err := os.Open("config.json")
            if (err == nil) {
                defer file.Close()
                decoder := json.NewDecoder(file)
                err = decoder.Decode(&Config)
            }
            return
        }
        func init() {
            time.Sleep(time.Second)
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            } else {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
        Product: Hat, Category: Skiing, Price: $10.00
        Product: Boots, Category: Skiing, Price: $220.51
        Product: Gloves, Category: Skiing, Price: $40.20
████████████████████████████████████████████████████████████████████████
339.Methods Defined by the File Struct for Reading at a Specific Location
    Name                        Description
    --------------------        ------------------------------------------
    ReadAt(slice, offset)       This method is defined by the ReaderAt interface and performs a read into the
                                specific slice at the specified position offset in the file.
    Seek(offset, how)           This method is defined by the Seeker interface and moves the offset into
    جستجو کنید                  the file for the next read. The offset is determined by the combination of the
                                two arguments: the first argument specifies the number of bytes to offset,
                                and the second argument determines how the offset is applied—a value of 0
                                means the offset is relative to the start of the file, a value of 1 means the offset
                                is relative to the current read position, and a value of 2 means the offset is
                                relative to the end of the file.
    
    
    Reading from specific locations requires knowledge of the file structure.
    In this example, I know the location of the data I want to read, 
    which allows me to use the ReadAt method to read the username value
    and the Seek method to jump to the start of the product data.
    
    example:
    readconfig.go:
        package main
        import (
            "os"
            "encoding/json"
            //"strings"
        )
        type ConfigData struct {
            UserName string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            file, err := os.Open("config.json")
            if (err == nil) {
                defer file.Close()
                nameSlice := make([]byte, 5)
                file.ReadAt(nameSlice, 19)
                Config.UserName = string(nameSlice)
                file.Seek(55, 0)
                decoder := json.NewDecoder(file)
                err = decoder.Decode(&Config.AdditionalProducts)
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            } else {
                Printfln("Error Loading Config: %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
340.The os Package Function for Writing Files
    Name                                    Description
    -------------------------------         ----------------------------------------
    WriteFile(name,slice, modePerms)        This function creates a file with the specified name, mode, and permissions and
                                            writes the content of the specified byte slice. If the file already exists, its contents
                                            will be replaced with the byte slice. The result is an error that reports any problems
                                            creating the file or writing the data.
    OpenFile(name, flag, modePerms)         The function opens the file with the specified name, using the flags to control how
                                            the file is opened. If a new file is created, then the specified mode and permissions
                                            are applied. The result is a File value that provides access to the file contents and
                                            an error that indicates problems opening the file.
████████████████████████████████████████████████████████████████████████
341.the Write Convenience Function
    The file mode is used to specify special characteristics for the file, 
    but a value of zero is used for regular files, as in the example. 
    You can find a list of the file mode values and their settings at 
    
    https://golang.org/pkg/io/fs/#FileMode


    https://cs.opensource.google/go/go/+/go1.21.1:src/io/fs/fs.go;l=165
    
    type FileMode 
    type FileMode uint32
    A FileMode represents a file's mode and permission bits. 
    The bits have the same definition on all systems, 
    so that information about files can be moved from one system to another portably. 
    Not all bits apply to all systems. 
    The only required bit is ModeDir for directories.
    
    
    const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file; Plan 9 only
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky
        ModeIrregular                                  // ?: non-regular file; nothing else is known about this file
    
        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
    
        ModePerm FileMode = 0777 // Unix permission bits
    )

    The defined file mode bits are the most significant bits of the FileMode. 
    The nine least-significant bits are the standard Unix rwxrwxrwx permissions. 
    The values of these bits should be considered part of the public API and 
    may be used in wire protocols or disk representations: they must not be changed, 
    although new bits might be added.


    example:
    main.go:
        package main
        import (
            "fmt"
            "os"
            "time"
        )
        func main() {
            total := 0.0
            for _, p := range Products {
                total += p.Price
            }
            dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n",time.Now().Format("Mon 15:04:05"), total)
            err := os.WriteFile("output.txt", []byte(dataStr), 0666)
            if err == nil {
                fmt.Println("Output file created")
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output: a file create with this name= output.txt
        Time: Thu 07:27:34, Total: $279.00
████████████████████████████████████████████████████████████████████████
342.Using the File Struct to Write to a File
    The OpenFile function opens a file and returns a File value. 
    Unlike the Open function, the OpenFile function accepts one or 
    more flags that specify how the file should be opened. 
    The flags are defined as constants in the os package,
    Care must be taken with these flags, not all of which are supported by every operating system.

████████████████████████████████████████████████████████████████████████
343.The File Opening Flags
    Name            Description
    --------        ------------------------------------------
    O_RDONLY        This flag opens the file read-only so that it can be read from but not written to.
    O_WRONLY        This flag opens the file write-only so that it can be written to but not read from.
    O_RDWR          This flag opens the file read-write so that it can be written to and read from.
    O_APPEND        This flag will append writes to the end of the file.
    O_CREATE        This flag will create the file if it doesn't exist.
    O_EXCL          This flag is used in conjunction with O_CREATE to ensure that a new file is created. If the file
                    already exists, this flag will trigger an error.
    O_SYNC          This flag enables synchronous writes, such that data is written to the storage device before
                    the write function/method returns.
    O_TRUNC         This flag truncates the existing content in the file.
████████████████████████████████████████████████████████████████████████
344.Writing to a File
    example:
    main.go:
        package main
        import (
            "fmt"
            "time"
            "os"
        )
        func main() {
            total := 0.0
            for _, p := range Products {
                total += p.Price
            }
            dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n",
                time.Now().Format("Mon 15:04:05"), total)
        
                
            file, err := os.OpenFile("output.txt",os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
            
            if (err == nil) {
                defer file.Close()
                file.WriteString(dataStr)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    

    I combined the O_WRONLY flag to open the file for writing, 
    the O_CREATE file to create if it doesn't already
    exist, and the O_APPEND flag to append any written data to the end of the file.
    Output:
        appended to file exist:
            Time: Thu 07:27:34, Total: $279.00
            Time: Thu 08:17:05, Total: $81174.40
            Time: Thu 08:17:09, Total: $81174.40
        
████████████████████████████████████████████████████████████████████████
345.The File Methods for Writing Data
    Name                        Description
    -----------------------     ---------------------------------------------------
    Seek(offset, how)           This method sets the location for subsequent operations.
    Write(slice)                This method writes the contents of the specified byte slice to the file.
                                The results are the number of bytes written and an error that indicates
                                problems writing the data.
    WriteAt(slice, offset)      This method writes the data in the slice at the specified location and is the
                                counterpart to the ReadAt method.
    WriteString(str)            This method writes a string to the file. This is a convenience method that
                                converts the string to a byte slice, invokes the Write method, and returns the
                                results it receives.
████████████████████████████████████████████████████████████████████████
346.Writing JSON Data to a File
    example:
    main.go:
        package main
        import (
            // "fmt"
            // "time"
            "encoding/json"
            "os"
        )
        func main() {
            cheapProducts := []Product{}
            for _, p := range Products {
                if p.Price < 100 {
                    cheapProducts = append(cheapProducts, p)
                }
            }
            file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
            if err == nil {
                defer file.Close()
                encoder := json.NewEncoder(file)
                encoder.Encode(cheapProducts)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        create file = cheap.json
        content of this:
    
            [{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75}]

████████████████████████████████████████████████████████████████████████
347.the Convenience Functions to Create New Files
    The os Package Functions for Creating Files
        The CreateTemp function can be useful, but it is important to understand that the purpose of this
        function is to generate a random filename and that in all other respects, the file that is created is just a
        regular file. The file that is created isn't removed automatically and will remain on the storage device after
        the application has been executed.
    Name                                Description
    -------------------------           --------------------------------------------
    Create(name)                        This function is equivalent to calling OpenFile with the O_RDWR, O_CREATE, and
                                        O_TRUNC flags. The results are the File, which can be used for reading and writing,
                                        and an error that is used to indicate problems creating the file. Note that this
                                        combination of flags means that if a file exists with the specified name, it will be
                                        opened, and its contents will be deleted.
    CreateTemp(dirName, fileName)       This function creates a new file in the directory with the specified name. If the
                                        name is the empty string, then the system temporary directory is used, obtained
                                        using the TempDir function. The file is created with a
                                        name that contains a random sequence of characters, as demonstrated in the text
                                        after the table. The file is opened with the O_RDWR, O_CREATE, and O_EXCL flags.
                                        The file isn't removed when it is closed.
████████████████████████████████████████████████████████████████████████
348.Creating a Temporary File
    The location of the temporary file is specified with a period, meaning the current working directory.
    if the empty string is used, then the file will be created in the default temporary
    directory, which is obtained using the TempDir function described.
    The name of the file can include an asterisk (the * character), 
    and if this is present, the random part of the filename will replace it. 
    If the filename does not contain an asterisk, 
    then the random part of the filename will be added to the end of the name.

    ompile and execute the project, and once execution is complete, you will see a new file in the files
    folder. The file in my project is named tempfile-1732419518.json, but your filename will be different, and
    you will see a new file and a unique name each time the program is executed.

    example:
    main.go:
        package main
        import (
            "os"
            "encoding/json"
        )
        func main() {
            cheapProducts := []Product {}
            for _, p := range Products {
                if (p.Price < 100) {
                    cheapProducts = append(cheapProducts, p)
                }
            }
            file, err := os.CreateTemp(".", "tempfile-*.json")
            if (err == nil) {
                defer file.Close()
                encoder := json.NewEncoder(file)
                encoder.Encode(cheapProducts)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
    tempfile-1982129407.json:
        [{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75}]
████████████████████████████████████████████████████████████████████████
349.Working with File Paths
    If you want to read and write files in
    other locations, then you must specify file paths. The issue is that not all of the operating systems that Go
    supports express file paths in the same way. For example, the path to a file named mydata.json in my home
    directory on a Linux system might be expressed like this:
        /home/adam/mydata.json

    where the path to the same in my home directory is expressed like this:
        C:\Users\adam\mydata.json
████████████████████████████████████████████████████████████████████████
350.The Common Location Functions Defined by the os Package
    Name                Description
    --------------      ------------------------------------
    Getwd()             This function returns the current working directory, expressed as a string, and an
                        error that indicates problems obtaining the value.
    UserHomeDir()       This function returns the user's home directory and an error that indicates problems
                        obtaining the path.
    UserCacheDir()      This function returns the default directory for user-specific cached data and an error
                        that indicates problems obtaining the path.
    UserConfigDir()     This function returns the default directory for user-specific configuration data and an
                        error that indicates problems obtaining the path.
    TempDir()           This function returns the default directory for temporary files and an error that
                        indicates problems obtaining the path.
████████████████████████████████████████████████████████████████████████
351.The path/filepath Functions for Paths
    Name                    Description
    ------------            ------------------------------------------
    Abs(path)               This function returns an absolute path, which is useful if you have a relative path,
                            such as a filename.
    IsAbs(path)             This function returns true if the specified path is absolute.
    Base(path)              This function returns the last element from the path.
    Clean(path)             This function tidies up path strings by removing duplicate separators and relative references.
    Dir(path)               This function returns all but the last element of the path.
    EvalSymlinks(path)      This function evaluates a symbolic link and returns the resulting path.
    Ext(path)               This function returns the file extension from the specified path, which is
                            assumed to be the suffix following the final period in the path string.
    FromSlash(path)         This function replaces each forward slash with the platform's file separator character.
    ToSlash(path)           This function replaces the platform's file separator with forward slashes.
    Join(...elements)       This function combines multiple elements using the platform's file separator.
    Match(pattern, path)    This function returns true if the path is matched by the specified pattern.
    Split(path)             This function returns the components on either side of the final path separator in
                            the specified path.
    SplitList(path)         This function splits a path into its components, which are returned as a string slice.
    VolumeName(path)        This function returns the volume component of the specified path or the empty
                            string if the path does not contain a volume.
████████████████████████████████████████████████████████████████████████
352.Working with a Path
    example:
    main.go
        package main
        import (
            // "fmt"
            // "time"
            "os"
            //"encoding/json"
            "path/filepath"
        )
        func main() {
            path, err := os.UserHomeDir()
            if (err == nil) {
                path = filepath.Join(path, "MyApp", "MyTempFile.json")
            }
            Printfln("Full path: %v", path)
            Printfln("Volume name: %v", filepath.VolumeName(path))
            Printfln("Dir component: %v", filepath.Dir(path))
            Printfln("File component: %v", filepath.Base(path))
            Printfln("File extension: %v", filepath.Ext(path))
        }
    Output:
        Username: Alice
        Full path: /home/sina/MyApp/MyTempFile.json
        Volume name: 
        Dir component: /home/sina/MyApp
        File component: MyTempFile.json
        File extension: .json
    
    received on the Windows machine:
        Username: Alice
        Full path: C:\Users\adam\MyApp\MyTempFile.json
        Volume name: C:
        Dir component: C:\Users\adam\MyApp
        File component: MyTempFile.json
        File extension: .json
████████████████████████████████████████████████████████████████████████
353.The os Package Functions for Managing Files and Directories
    Name                            Description
    -----------------               ------------------------------------------
    Chdir(dir)                      This function changes the current working directory to the specified directory.
                                    The result is an error that indicates problems making the change.
    Mkdir(name, modePerms)          This function creates a directory with the specified name and mode/
                                    permissions. The result is an error that is nil if the directory is created or that
                                    describes a problem if one arises.
    MkdirAll(name, modePerms)       This function performs the same task as Mkdir but creates any parent directories
                                    in the specified path.
    MkdirTemp(parentDir, name)      This function is similar to CreateTemp but creates a directory rather than a file.
                                    A random string is added to the end of the specified name or in place of an
                                    asterisk, and the new directory is created within the specified parent. The results
                                    are the name of the directory and an error indicating problems.
    Remove(name)                    This function removes the specified file or directory. The result is an error that
                                    describes any problems that arise.
    RemoveAll(name)                 This function removes the specified file or directory. If the name specifies a
                                    directory, then any children it contains are also removed. The result is an error
                                    that describes any problems that arise.
    Rename(old, new)                This function renames the specified file or folder. The result is an error that
                                    describes any problems that arise.
    Symlink(old, new)               This function creates a symbolic link to the specified file. The result is an error
                                    that describes any problems that arise.
████████████████████████████████████████████████████████████████████████
354.Creating Directories
    example:
    main.go:
        package main
        import (
            // "fmt"
            // "time"
            "encoding/json"
            "os"
            "path/filepath"
        )
        func main() {
            path, err := os.UserHomeDir()
            if err == nil {
                path = filepath.Join(path, "0-Repo/TEST-2/MyApp", "MyTempFile.json")
            }
            Printfln("Full path: %v", path)
            err = os.MkdirAll(filepath.Dir(path), 0766)
            if err == nil {
                file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
                if err == nil {
                    defer file.Close()
                    encoder := json.NewEncoder(file)
                    encoder.Encode(Products)
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        print this:
            Username: Alice
            Full path: /home/sina/0-Repo/TEST-2/MyApp/MyTempFile.json

        Create this:
            MyTempFile.json in that directory with content:
                [{"Name":"Kayak","Category":"Watersports","Price":279},{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Stadium","Category":"Soccer","Price":79500},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75},{"Name":"Bling-Bling King","Category":"Chess","Price":1200}]
████████████████████████████████████████████████████████████████████████
355.ReadDir(name)
    The os Package Function for Listing Directories
    This function reads the specified directory and returns a DirEntry slice, each of which
    describes an item in the directory.
    The result of the ReadDir function is a slice of values that implement the DirEntry interface, which
    defines the methods.
████████████████████████████████████████████████████████████████████████
356.The Methods Defined by the DirEntry Interface
    Name        Description
    --------    --------------------------
    Name()      This method returns the name of the file or directory described by the DirEntry value.
    IsDir()     This method returns true if the DirEntry value represents a directory.
    Type()      This method returns a FileMode value, which is an alias to uint32, which describes the file more
                and the permissions of the file or directory represented by the DirEntry value.
    Info()      This method returns a FileInfo value that provides additional details about the file or directory
                represented by the DirEntry value.
████████████████████████████████████████████████████████████████████████
357.Useful Methods Defined by the FileInfo Interface
    Name            Description
    ----------      ------------------------------------
    Name()          This method returns a string containing the name of the file or directory.
    Size()          This method returns the size of the file, expressed as an int64 value.
    Mode()          This method returns the file mode and permission settings for the file or directory.
    ModTime()       This method returns the last modified time of the file or directory.
████████████████████████████████████████████████████████████████████████
358.Stat(path)
    The os Package Function for Inspecting a File
    This function accepts a path string. It returns a FileInfo value that describes the file and an
    error, which indicates problems inspecting the file.
████████████████████████████████████████████████████████████████████████
359.Enumerating Files
    example:
    main.go:
        package main
        import (
            "os"
        )
        func main() {
            path, err := os.Getwd()
            if err == nil {
                dirEntries, err := os.ReadDir(path)
                if err == nil {
                    for _, dentry := range dirEntries {
                        Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
                    }
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Entry name: .git, IsDir: true
        Entry name: .vscode, IsDir: true
        Entry name: README.md, IsDir: false
        Entry name: U.sh, IsDir: false
        Entry name: cheap.json, IsDir: false
        Entry name: config.json, IsDir: false
        Entry name: go.mod, IsDir: false
        Entry name: main.go, IsDir: false
        Entry name: output.txt, IsDir: false
        Entry name: printer.go, IsDir: false
        Entry name: product.go, IsDir: false
        Entry name: readconfig.go, IsDir: false
████████████████████████████████████████████████████████████████████████
360.Determining Whether a File Exists تعیین اینکه آیا یک فایل وجود دارد یا خیر
    Checking Whether a File Exists:
    main.go:
        package main
        import (
            "os"
        )
        func main() {
            targetFiles := []string { "no_such_file.txt", "config.json" }
            for _, name := range targetFiles {
                info, err := os.Stat(name)
                if os.IsNotExist(err) {
                    Printfln("File does not exist: %v", name)
                } else if err != nil  {
                    Printfln("Other error: %v", err.Error())
                } else {
                    Printfln("File %v, Size: %v", info.Name(), info.Size())
                }
            }
        }
    Output:
        Username: Alice
        File does not exist: no_such_file.txt
        File config.json, Size: 253
████████████████████████████████████████████████████████████████████████
361.The path/filepath Function for Locating Files with a Pattern
    Name                        Description
    --------------------        -------------------------------
    Match(pattern, name)        This function matches a single path against a pattern. The results are a bool,
                                which indicates if there is a match, and an error, which indicates problems
                                with the pattern or with performing the match.
    Glob(pathPatten)            This function finds all the files that match the specified pattern. The results
                                are a string slice containing the matched paths and an error that indicates
                                problems with performing the search.
████████████████████████████████████████████████████████████████████████
362.The Search Pattern Syntax for the path/filepath Functions
    Term        Description
    --------    -------------------
     *          This term matches any sequence of characters, excluding the path separator.
     ?          This term matches any single character, excluding the path separator.
     [a-Z]      This term matches any character in the specified range.
████████████████████████████████████████████████████████████████████████
363.Locating Files
    example:
    main.go:
        package main
        import (
            "os"
            "path/filepath"
        )
        func main() {
            path, err := os.Getwd()
            if err == nil {
                matches, err := filepath.Glob(filepath.Join(path, "*.json"))
                if err == nil {
                    for _, m := range matches {
                        Printfln("Match: %v", m)
                    }
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Match: /home/sina/0-Repo/TEST-2/cheap.json
        Match: /home/sina/0-Repo/TEST-2/config.json
████████████████████████████████████████████████████████████████████████
364.The Function Provided by the path/filepath Package
    Name                        Description
    -----------------------     ----------------------------------
    WalkDir(directory, func)    This function calls the specified function for each file and directory in the
                                specified directory.
████████████████████████████████████████████████████████████████████████
365.Walking a Directory قدم زدن در یک فهرست
    example:
    main.go:
        package main
        import (
            "os"
            "path/filepath"
        )
        func callback(path string, dir os.DirEntry, dirErr error) (err error) {
            info, _ := dir.Info()
            Printfln("Path %v, Size: %v", path, info.Size())
            return
        }
        func main() {
            path, err := os.Getwd()
            if err == nil {
                err = filepath.WalkDir(path, callback)
            } else {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Path /home/sina/0-Repo/TEST-2, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.git, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.git/COMMIT_EDITMSG, Size: 4
        Path /home/sina/0-Repo/TEST-2/.git/FETCH_HEAD, Size: 92
        Path /home/sina/0-Repo/TEST-2/.git/refs/remotes/origin/main, Size: 41
        Path /home/sina/0-Repo/TEST-2/.git/refs/tags, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.vscode, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.vscode/extensions.json, Size: 79
        Path /home/sina/0-Repo/TEST-2/.vscode/settings.json, Size: 405
        Path /home/sina/0-Repo/TEST-2/README.md, Size: 9
        Path /home/sina/0-Repo/TEST-2/U.sh, Size: 68
        Path /home/sina/0-Repo/TEST-2/cheap.json, Size: 487
        Path /home/sina/0-Repo/TEST-2/config.json, Size: 253
        Path /home/sina/0-Repo/TEST-2/go.mod, Size: 22
        Path /home/sina/0-Repo/TEST-2/main.go, Size: 8579
        Path /home/sina/0-Repo/TEST-2/output.txt, Size: 109
        Path /home/sina/0-Repo/TEST-2/printer.go, Size: 129
        Path /home/sina/0-Repo/TEST-2/product.go, Size: 474
        Path /home/sina/0-Repo/TEST-2/readconfig.go, Size: 704
████████████████████████████████████████████████████████████████████████
`,
}