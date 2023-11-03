package features

var TitleDataBases = []string{
	"ALL WORKING WITH DATABASES",
	"ALLWORKINGWITHDATABASES",
	"ALL DATABASES",
	"ALLDATABASES",
}

var OriginalDataBases = DataBase{
	Alldatafield: `
464.Working with Databases
    There are drivers for a wide range of databases, and a list can be found at 
    https://github.com/golang/go/wiki/sqldrivers

    Putting Working with Databases in Context
    What is it?
    The database/sql package provides features for working with SQL databases.

    Why is it useful?
    Relational databases remain the most effective way of storing large amounts of
    structured data and are used in most large projects.

    How is it used?
    Driver packages provide support for specific databases, while the database/sql
    package provides a set of types that allow databases to be used consistently.
    
    Are there any pitfalls or limitations? 
    These features do not automatically populate struct fields from result rows.

    Are there any alternatives?
    There are third-party packages that build on these features to simplify or enhance their use.

    Preparing for This Chapter:
    1-Initializing the Module
        go mod init data
    2-Add a file named printer.go to the data folder
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
    3-Add a file named main.go
        package main
        func main() {
        Printfln("Hello, Data")
        }
    4-Compiling and Executing the Project
        go run .
    ====================================================================
    Output:
        Hello, Data


    Preparing the Database:
    add a file named products.sql to the data folder:
        DROP TABLE IF EXISTS Categories;

        DROP TABLE IF EXISTS Products;
        
        CREATE TABLE IF NOT EXISTS Categories (
            Id INTEGER NOT NULL PRIMARY KEY,
            Name TEXT
        );
        
        CREATE TABLE IF NOT EXISTS Products (
            Id INTEGER NOT NULL PRIMARY KEY,
            Name TEXT,
            Category INTEGER,
            Price decimal(8, 2),
            CONSTRAINT CatRef FOREIGN KEY(Category) REFERENCES Categories (Id)
        );
        
        INSERT INTO
            Categories (Id, Name)
        VALUES
            (1, "Watersports"),
            (2, "Soccer");
        
        INSERT INTO
            Products (Id, Name, Category, Price)
        VALUES
            (1, "Kayak", 1, 279),
            (2, "Lifejacket", 1, 48.95),
            (3, "Soccer Ball", 2, 19.50),
            (4, "Corner Flags", 2, 34.95);


    Go to https://www.sqlite.org/download.html , look for the precompiled binaries section for your
    operating system, and download the tools package.
    Unpack the zip archive and copy the sqlite3 or sqlite3.exe file into the data folder. Run the command
    Creating the Database command:
        ./sqlite3 products.db ".read products.sql"

████████████████████████████████████████████████████████████████████████
465.Creating the Database command:
    ./sqlite3 products.db ".read products.sql"
████████████████████████████████████████████████████████████████████████
466.Installing a Database Driver
    Run the command:
        go get modernc.org/sqlite
    
    Most database servers are set up separately so that the database driver opens a connection to a separate
    process. SQLite is an embedded database and is included in the driver package, which means no additional
    configuration is required.
████████████████████████████████████████████████████████████████████████
467.Opening a Database
    The standard library provides the database/sql package for working with databases. 

    The functions described here:
    The database/sql Functions for Opening a Database
    Name                        Description
    ----------------            ------------------------------
    Drivers()                   This function returns a slice of strings, each of which contains the name of a database driver.
    Open(driver,connectionStr)  This function opens a database using the specified driver and connection string. The
                                results are a pointer to a DB struct, which is used to interact with the database and an
                                error that indicates problems opening the database.

████████████████████████████████████████████████████████████████████████
468.The Contents of the database.go
    package main
    // The blank identifier is used to import the database driver package, which loads the driver and allows it
    // to register as a provider of the SQL API:
    import (
        "database/sql"
        _ "modernc.org/sqlite"
    )
    func listDrivers() {
        for _, driver := range sql.Drivers() {
            Printfln("Driver: %v", driver)
        }
    }
    func openDatabase() (db *sql.DB, err error) {
        db, err = sql.Open("sqlite", "products.db")
        if err == nil {
            Printfln("Opened database")
        }
        return
    }
    ====================================================================
    Using the DB Struct in the main.go:
        package main
        func main() {
            listDrivers()
            db, err := openDatabase()
            if (err == nil) {
                db.Close()
            } else {
                panic(err)
            }
        }
    ====================================================================
    in Terminal:
        go mod tidy

        Output in Terminal:
            go: finding module for package modernc.org/sqlite
            go: downloading modernc.org/sqlite v1.27.0
            go: found modernc.org/sqlite in modernc.org/sqlite v1.27.0
            go: downloading golang.org/x/sys v0.9.0
            go: downloading modernc.org/ccgo/v3 v3.16.13
            go: downloading modernc.org/libc v1.29.0
            go: downloading modernc.org/mathutil v1.6.0
            go: downloading github.com/mattn/go-sqlite3 v1.14.16
            go: downloading github.com/google/pprof v0.0.0-20221118152302-e6195bd50e26
            go: downloading modernc.org/tcl v1.15.2
            go: downloading github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec
            go: downloading github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
            go: downloading golang.org/x/tools v0.0.0-20201124115921-2c860bdd6e78
            go: downloading modernc.org/cc/v3 v3.40.0
            go: downloading modernc.org/opt v0.1.3
            go: downloading github.com/dustin/go-humanize v1.0.1
            go: downloading github.com/google/uuid v1.3.0
            go: downloading github.com/mattn/go-isatty v0.0.16
            go: downloading modernc.org/memory v1.7.2
            go: downloading golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
            go: downloading golang.org/x/mod v0.3.0
            go: downloading lukechampine.com/uint128 v1.2.0
            go: downloading modernc.org/strutil v1.1.3
            go: downloading modernc.org/token v1.0.1
            go: downloading modernc.org/httpfs v1.0.6
            go: downloading modernc.org/z v1.7.3
            go: downloading modernc.org/ccorpus v1.11.6

    ====================================================================
    Output:
        The main method calls the listDrivers function to print out the names of the loaded drivers and then
        calls the openDatabase function to open the database. Nothing is done with the database yet, but Close
        method is called.
████████████████████████████████████████████████████████████████████████
469.The DB Method for Closing the Database:
    Name        Description
    -------     ----------------------
    Close()     This function closes the database and prevents further operations from being performed.
████████████████████████████████████████████████████████████████████████
470.Executing Statements and Queries
    The DB Methods for Executing SQL Statements
    Name                        Description
    --------------------        -----------------------------
    Query(query,...args)        This method executes the specified query, using the optional placeholder arguments.
                                The results are a Rows struct, which contains the query results, and an error that
                                indicates problems executing the query.
    QueryRow(query, ..args)     This method executes the specified query, using the optional placeholder arguments.
                                The result is a Row struct, which represents the first row from the query results. See
                                the “Executing Queries for Single Rows” section.
    Exec(query,...args)         This method executes statements or queries that do not return rows of data. The
                                method returns a Result, which describes the response from the database, and
                                an error that signals problems with execution. See the “Executing Other Queries” section.

████████████████████████████████████████████████████████████████████████
471.Using Contexts with Databases
    the context package and the Context interface it defines, which is used
    to manage requests as they are processed by a server. All the important methods defined in the
    database/sql package also have versions that accept a Context argument, which is useful if you
    want to take advantage of features like request handling timeouts.

████████████████████████████████████████████████████████████████████████
472.Querying for Multiple Rows
    The Query method executes a query that retrieves one or more rows from the database. The Query method
    returns a Rows struct, which contains the query results and an error that indicates problems. The row data is
    accessed through the methods described in below

    The Rows Struct Methods
    Name                Description
    ----------------    -----------------------------
    Next()              This method advances to the next result row. The result is a bool, which is true when
                        there is data to read and false when the end of the data has been reached, at which point
                        the Close method is automatically called.
    NextResultSet()     This method advances to the next result set when there are multiple result sets in the
                        same database response. The method returns true if there is another set of rows to process.
    Scan(...targets)    This method assigns the SQL values from the current row to the specified variables. The
                        values are assigned via pointers and the method returns an error that indicates when the
                        values cannot be scanned. See the “Understanding the Scan Method” section for details.
    Close()             This method prevents further enumeration of the results and is used when not all of
                        the data is required. There is no need to call this method if the Next method is used to
                        advance until it returns false.
████████████████████████████████████████████████████████████████████████
473.Querying the Database in the main.go
    example:
        package main
        import "database/sql"
        func queryDatabase(db *sql.DB) {
            rows, err := db.Query("SELECT * from Products")
            if err == nil {
                for rows.Next() {
                    var id, category int
                    var name string
                    var price float64
                    rows.Scan(&id, &name, &category, &price)
                    Printfln("Row: %v %v %v %v", id, name, category, price)
                }
            } else {
                Printfln("Error: %v", err)
            }
        }
        func main() {
            //listDrivers()
            db, err := openDatabase()
            if err == nil {
                queryDatabase(db)
                db.Close()
            } else {
                panic(err)
            }
        }
    ====================================================================
    The queryDatabase function performs a simple SELECT query on the Products table with the Query
    method, which produces a Rows result and an error. If the error is nil, a for loop is used to move through
    the result rows by calling the Next method, which returns true if there is a row to process and returns false
    when the end of the data has been reached.
    
	`,
}