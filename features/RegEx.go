package features

var TitleOfRegEx = []string{
	"ALL REGEX",
	"ALL REGULAR EXPRESSION",
}

type AllRegEx struct {
	allRegex string
}

var OriginalAllRegex AllRegEx = AllRegEx{
	allRegex: `

189.Regular Expressions
    The regular expressions used in this section perform basic matches, but the regexp package
    supports an extensive pattern syntax, which is described at https://pkg.go.dev/regexp/syntax@go1.17.1.
    
    The Basic Functions Provided by the regexp Package
    Function                Description
    -----------------       --------------------------------------------------------------------------
    Match(pattern, b)       This function returns a bool that indicates whether a pattern is matched by
                            the byte slice b.
    
    MatchString(patten, s)  This function returns a bool that indicates whether a pattern is matched by
                            the string s.
    
    Compile(pattern)        This function returns a RegExp that can be used to perform repeated pattern
                            matching with the specified pattern.
    
    MustCompile(pattern)    This function provides the same feature as Compile but panics, 
                            if the specified pattern cannot be compiled.
    
    example:
        package main
        import (
            "fmt"
            //"strings"
            "regexp"
        )
        func main() {
            description := "A boat for one person"
            match, err := regexp.MatchString("[A-z]oat", description)
            if (err == nil) {
                fmt.Println("Match:", match)
            } else {
                fmt.Println("Error:", err)
            }
        }
    Output:
        Match: true
    
    

████████████████████████████████████████████████████████████████████████
160.String Processing and Regular Expressions
    What are they?
    String processing includes a wide range of operations, from trimming
    whitespace to splitting a string into components. Regular expressions
    are patterns that allow string matching rules to be concisely defined.

    Why are they useful?
    These operations are useful when an application needs to process
    string values. A common example is processing HTTP requests.

    How are they used?
    These features are contained in the strings and regexp packages,
    which are part of the standard library.

    Are there any pitfalls or limitations?
    There are some quirks in the way that some of these operations are
    performed, but they mostly behave as you would expect.

    Are there any alternatives?
    The use of these packages is optional, and they do not have
    to be used. That said, there is little point in creating your own
    implementations of these features since the standard library is well-
    written and thoroughly tested.

    Problem                     Solution
    -------------------         -------------------------------------------------------------------
    Compare strings             Use the Contains, EqualFold, or Has* function in the strings package

    Convert string case         Use the ToLower, ToUpper, Title, or ToTitle function in the
                                strings package

    Check or change             Use the functions provided by the unicode package
    character case

    Find content in strings     Use the functions provided by the strings or regexp package

    Split a string              Use the Fields or Split* function in the strings and regexp packages

    Join strings                Use the Join or Repeat function in the strings package

    Trim characters from        Use the Trim* functions in the strings package
    a string

    Perform a substitution      Use the Replace* or Map function in the strings package,
    تعویض انجام دهید            use a Replacer, or use the Replace* functions in the regexp package

    Efficiently build a         Use the Builder type in the strings package
    string
████████████████████████████████████████████████████████████████████████
161.Comparing Strings
    The strings Functions for Comparing Strings

    Function                    Description
    -------------               ------------------------
    Contains(s, substr)         This function returns true if the string s contains substr and false if it does not.

    ContainsAny(s, substr)      This function returns true if the string s contains any of the characters
                                contained in the string substr.

    ContainsRune(s, rune)       This function returns true if the string s contains a specific rune.

    EqualFold(s1, s2)           This function performs a case-insensitive comparison and returns true of
                                strings s1 and s2 are the same.

    HasPrefix(s, prefix)        This function returns true if the string s begins with the string prefix.

    HasSuffix(s, suffix)        This function returns true if the string ends with the string suffix.

    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            product := "Kayak"
            fmt.Println("Contains:", strings.Contains(product, "yak"))
            fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
            fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
            fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
            fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
            fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
        }
    Output:
        Contains: true
        ContainsAny: true
        ContainsRune: true
        HasPrefix: true
        HasSuffix: true
        EqualFold: true
████████████████████████████████████████████████████████████████████████
162.Using The Byte-Oriented Functions
    example:
        package main
        import (
            "fmt"
            "strings"
            "bytes"
        )
        func main() {
            price := "€100"
            fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
            fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
                []byte { 226, 130 }))
        }
    Output:
        Strings Prefix: true
        Bytes Prefix: true
████████████████████████████████████████████████████████████████████████
163.Converting String Case
    The Case Functions in the strings Package
    Function        Description
    --------------  ------------------------------------------------------
    ToLower(str)    This function returns a new string containing the characters in the specified string
                    mapped to lowercase.

    ToUpper(str)    This function returns a new string containing the characters in the specified string
                    mapped to lowercase.

    Title(str)      This function converts the specific string so that the first character of each word is
                    uppercase and the remaining characters are lowercase.

    ToTitle(str)    This function returns a new string containing the characters in the specified string
                    mapped to title case.

    Care must be taken with the Title and ToTitle functions, which don't work the way you might expect.
    The Title function returns a string that is suitable for use as a title, but it treats all words the same

    Creating a Title:
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for sailing"
            fmt.Println("Original:", description)
            fmt.Println("Title:", strings.Title(description))
        fmt.Println("Title:", strings.ToTitle(description))
        }
    Output:
        Original: A boat for sailing
        Title: A Boat For Sailing
        Title: A BOAT FOR SAILING
████████████████████████████████████████████████████████████████████████
164.Title Case
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            specialChar := "\u01c9"
            fmt.Println("Original:", specialChar, []byte(specialChar))
            upperChar := strings.ToUpper(specialChar)
            fmt.Println("Upper:", upperChar, []byte(upperChar))
            titleChar := strings.ToTitle(specialChar)
            fmt.Println("Title:", titleChar, []byte(titleChar))
        }
    Output:
        Original: ǉ [199 137]
        Upper: Ǉ [199 135]
        Title: ǈ [199 136]
████████████████████████████████████████████████████████████████████████
165.Working with Character Case
    The unicode package provides functions that can be used to determine or change the case of individual
    characters
    Functions in the unicode Package for Character Case
    Function        Description
    -------------   -----------------------------------------------------------------
    IsLower(rune)   This function returns true if the specified rune is lowercase.
    ToLower(rune)   This function returns the lowercase rune associated with the specified rune.
    IsUpper(rune)   This function returns true if the specified rune is uppercase.
    ToUpper(rune)   This function returns the upper rune associated with the specified rune.
    IsTitle(rune)   This function returns true if the specified rune is title case.
    ToTitle(rune)   This function returns the title case rune associated with the specified rune.
████████████████████████████████████████████████████████████████████████
166.the Rune Case Functions
    example:
        package main
        import (
            "fmt"
            "unicode"
        )
        func main() {
            product := "Kayak"
            for _, char := range product {
                fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
            }
        }
    Output:
        K Upper case: true
        a Upper case: false
        y Upper case: false
        a Upper case: false
        k Upper case: false        
████████████████████████████████████████████████████████████████████████
167.Inspecting Strings
    The strings Functions for Inspecting Strings
    Function                Description
    ------------            --------------------------------------------------
    Count(s, sub)           This function returns an int that reports how many times the specified
                            substring is found in the string s.
    Index(s, sub)           These functions return the index of the first or last occurrence of a specified
    LastIndex(s, sub)       substring string within the string s, or -1 if there is no occurrence.

    IndexAny(s, chars)      These functions return the first or last occurrence of any character in the
    LastIndexAny(s, chars)  specified string within the string s, or -1 if there is no occurrence.

    IndexByte(s, b)         These functions return the index of the first or last occurrence of a specified
    LastIndexByte(s, b)     byte within the string s, or -1 if there is no occurrence.

    IndexFunc(s, func)      These functions return the index of the first or last occurrence of the
    LastIndexFunc(s, func)  character in the string s for which the specified function returns true, as
                            described in the “Inspecting Strings with Custom Functions” section.
    
    example:
        package main
        import (
            "fmt"
            "strings"
            //"unicode"
        )
        func main() {
            description := "A boat for one person"
            fmt.Println("Count:", strings.Count(description, "o"))
            fmt.Println("Index:", strings.Index(description, "o"))
            fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
            fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
            fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
            fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
        }
    Output:
        Count: 4
        Index: 3
        LastIndex: 19
        IndexAny: 2
        LastIndex: 19
        LastIndexAny: 4
████████████████████████████████████████████████████████████████████████
168.IndexFunc and LastIndexFunc functions 
    Inspecting Strings with Custom Functions
    The IndexFunc and LastIndexFunc functions use a custom function to inspect strings, using custom functions

    Custom functions receive a rune and return a bool result that indicates if the character meets the
    desired condition. The IndexFunc function invokes the custom function for each character in the string until
    a true result is obtained, at which point the index is returned.
    The isLetterB variable is assigned a custom function that receives a rune and returns true if the rune
    is a uppercase or lowercase B. The custom function is passed to the strings.IndexFunc function
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            isLetterB := func (r rune) bool {
                return r == 'B' || r == 'b'
            }
            fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
        }
    Output:
        IndexFunc: 2
████████████████████████████████████████████████████████████████████████
169.Splitting Strings
    The Functions for Splitting Strings in the strings Package
    Function                    Description
    ---------------             --------------------------------
    Fields(s)                   This function splits a string on whitespace characters and returns a slice
                                containing the nonwhitespace sections of the string s.

    FieldsFunc(s, func)         This function splits the string s on the characters for which a custom function
                                returns true and returns a slice containing the remaining sections of the string.

    Split(s, sub)               This function splits the string s on every occurrence of the specified substring,
                                returning a string slice. If the separator is the empty string, then the slice will
                                contain strings for each character.

    SplitN(s, sub, max)         This function is similar to Split, but accepts an additional int argument that
                                specifies the maximum number of substrings to return. The last substring in the
                                result slice will contain the unsplit portion of the source string.

    SplitAfter(s, sub)          This function is similar to Split but includes the substring used in the results.


    SplitAfterN(s, sub, max)    This function is similar to SplitAfter, but accepts an additional int argument
                                that specifies the maximum number of substrings to return.

    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            splits := strings.Split(description, " ")
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
            splitsAfter := strings.SplitAfter(description, " ")
            for _, x := range splitsAfter {
                fmt.Println("SplitAfter >>" + x + "<<")
            }
        }
    Output:
        Split >>A<<
        Split >>boat<<
        Split >>for<<
        Split >>one<<
        Split >>person<<
        SplitAfter >>A <<
        SplitAfter >>boat <<
        SplitAfter >>for <<
        SplitAfter >>one <<
        SplitAfter >>person<<
████████████████████████████████████████████████████████████████████████
170.SplitN and SplitAfterN functions 
    Restricting the Number of Results
    The SplitN and SplitAfterN functions accept an int argument that specifies the maximum number of
    results that should be included in the results
    Restricting the Results
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            splits := strings.SplitN(description, " ", 3)
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
        }
    Output:
        Split >>A<<
        Split >>boat<<
        Split >>for one person<<
████████████████████████████████████████████████████████████████████████
171.strings.SplitN function
    Splitting on Whitespace Characters
    One limitation of the Split, SplitN, SplitAfter, and SplitAfterN functions is they do not deal with
    repeated sequences of characters, which can be a problem when splitting a string on whitespace characters

    The words in the source string are double-spaced, but the SplitN function splits only on the first space
    character, which produces odd results.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splits := strings.SplitN(description, " ", 3)
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
        }
    Output:
        Split >>This<<
        Split >><<
        Split >>is  double  spaced<<    
████████████████████████████████████████████████████████████████████████
172.Fields Function
    The Fields function doesn't support a limit on the number of results 
    but does deal with the double spaces properly.
    The Fields function has a better approach, which is to split on any character for
    which the IsSpace function in the unicode package returns true.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splits := strings.Fields(description)
            for _, x := range splits {
                fmt.Println("Field >>" + x + "<<")
            }
        }
    Output:
        Field >>This<<
        Field >>is<<
        Field >>double<<
        Field >>spaced<<
████████████████████████████████████████████████████████████████████████
173.FieldsFunc function
    Splitting Using a Custom Function to Split Strings
    The FieldsFunc function splits a string by passing each character to a custom function and splitting when
    that function returns true
    The custom function receives a rune and returns true if that rune should cause the string to split.
    The FieldsFunc function is smart enough to deal with repeated characters
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splitter := func(r rune) bool {
                return r == ' '
            }
            splits := strings.FieldsFunc(description, splitter)
            for _, x := range splits {
                fmt.Println("Field >>" + x + "<<")
            }
        }
    Output:
        Field >>This<<
        Field >>is<<
        Field >>double<<
        Field >>spaced<<
████████████████████████████████████████████████████████████████████████
174.Trimming Strings
    The Functions for Trimming Strings in the strings Package
    
    Function                Description
    ------------            ---------------------------------------
    TrimSpace(s)            This function returns the string s without leading or trailing whitespace characters.

    Trim(s, set)            This function returns a string from which any leading or trailing characters
                            contained in the string set are removed from the string s.

    TrimLeft(s, set)        This function returns the string s without any leading character contained
                            in the string set. This function matches any of the specified characters—use
                            the TrimPrefix function to remove a complete substring.

    TrimRight(s, set)       This function returns the string s without any trailing character contained
                            in the string set. This function matches any of the specified characters—use
                            the TrimSuffix function to remove a complete substring.

    TrimPrefix(s, prefix)   This function returns the string s after removing the specified prefix string.
                            This function removes the complete prefix string—use the TrimLeft
                            function to remove characters from a set.

    TrimSuffix(s, suffix)   This function returns the string s after removing the specified suffix string.
                            This function removes the complete suffix string—use the TrimRight
                            function to remove characters from a set.

    TrimFunc(s, func)       This function returns the string s from which any leading or trailing
                            character for which a custom function returns true are removed.

    TrimLeftFunc(s, func)   This function returns the string s from which any leading character for
                            which a custom function returns true are removed.

    TrimRightFunc(s, func)  This function returns the string s from which any trailing character for
                            which a custom function returns true are removed.
████████████████████████████████████████████████████████████████████████
175.TrimSpace function
    Trimming Whitespace
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            username := " Alice"
            trimmed := strings.TrimSpace(username)
            fmt.Println("Trimmed:", ">>" + trimmed + "<<")
        }
    Output:
        Trimmed: >>Alice<<
████████████████████████████████████████████████████████████████████████
176.Trim, TrimLeft, and TrimRight functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            trimmed := strings.Trim(description, "Anor ")
            fmt.Println("Trimmed:>>"+ trimmed+ "<<")
        }
    Ourput:
        Trimmed:>>boat for one pers<<
████████████████████████████████████████████████████████████████████████
177.TrimPrefix and TrimSuffix functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            prefixTrimmed := strings.TrimPrefix(description, "A boat ")
            wrongPrefix := strings.TrimPrefix(description, "A hat ")
            fmt.Println("Trimmed:", prefixTrimmed)
            fmt.Println("Not trimmed:", wrongPrefix)
        }
    Output:
        Trimmed: for one person
        Not trimmed: A boat for one person
████████████████████████████████████████████████████████████████████████
178.TrimFunc, TrimLeftFunc, and TrimRightFunc functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            trimmer := func(r rune) bool {
                return r == 'A' || r == 'n'
            }
            trimmed := strings.TrimFunc(description, trimmer)
            fmt.Println("Trimmed:", trimmed)
        }
    Output:
        Trimmed:  boat for one perso
████████████████████████████████████████████████████████████████████████
179.Altering Strings
    تغییر رشته‌ها 
    
    The Functions for Altering Strings in the strings Package
    Function                    Description
    ----------------            -----------------------------------
    Replace(s, old, new, n)     This function alters the string s by replacing occurrences of the string old with the
                                string new. The maximum number of occurrences that will be replaced is specified by
                                the int argument n.

    ReplaceAll(s, old, new)     This function alters the string s by replacing all occurrences of the string old with
                                the string new. Unlike the Replace function, there is no limit on the number of
                                occurrences that will be replaced.

    Map(func, s)                This function generates a string by invoking the custom function for each character in
                                the string s and concatenating the results. If the function produces a negative value,
                                the current character is dropped without a replacement.
████████████████████████████████████████████████████████████████████████
180.Replace and ReplaceAll functions
    The Replace function allows a maximum number of changes to be specified, 
    while the ReplaceAll function will replace all the
    occurrences of the substring it finds
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
            replace := strings.Replace(text, "boat", "canoe", 1)
            replaceAll := strings.ReplaceAll(text, "boat", "truck")
            fmt.Println("Replace:", replace)
            fmt.Println("Replace All:", replaceAll)
        }
    Output:
        Replace: It was a canoe. A small boat.
        Replace All: It was a truck. A small truck.
████████████████████████████████████████████████████████████████████████
181.Map function
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
        text := "It was a boat. A small boat."
        mapper := func(r rune) rune {
                if r == 'b' {
                    return 'c'
                }
                return r
            }
            mapped := strings.Map(mapper, text)
            fmt.Println("Mapped:", mapped)
        }
    Output:
        Mapped: It was a coat. A small coat.
████████████████████████████████████████████████████████████████████████
182.NewReplacer function
    The strings package exports a struct type named Replacer that is used to replace strings
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat.   111"
            replacer := strings.NewReplacer("boat", "kayak", 
            "small", "huge",
            "111", "222",
        )
            replaced := replacer.Replace(text)
            fmt.Println("Replaced:", replaced)
        }
    Output:
        Replaced: It was a kayak. A huge kayak.   222
████████████████████████████████████████████████████████████████████████
183.The Replacer Methods
    Name                    Description
    -------------------     --------------------------------------------
    Replace(s)              This method returns a string for which all the replacements specified with the
                            constructor have been performed on the string s.

    WriteString(writer, s)  This method is used to perform the replacements specified with the constructor
                            and write the results to an io.Writer
████████████████████████████████████████████████████████████████████████
184.Building and Generating Strings
    The strings Functions for Generating Strings
    Function            Description
    ----------------    ----------------------------------------------------------------------------------------
    Join(slice, sep)    This function combines the elements in the specified string slice, with the specified
                        separator string placed between elements.

    Repeat(s, count)    This function generates a string by repeating the string s for a specified number of times.
████████████████████████████████████████████████████████████████████████
185.Join and Repeat functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
            elements := strings.Fields(text)
            joined := strings.Join(elements, "--")
            fmt.Println("Joined:", joined)
            esplited := strings.Split(text, " ")
            fmt.Printf("%q\n",esplited)
        }
    Output:
        Joined: It--was--a--boat.--A--small--boat.
        ["It" "was" "a" "boat." "A" "small" "boat."]
████████████████████████████████████████████████████████████████████████
186.Building Strings
    The strings.Builder Methods
    Name                Description
    ---------------     --------------------------------------------
    WriteString(s)      This method appends the string s to the string being built.
    WriteRune(r)        This method appends the character r to the string being built.
    WriteByte(b)        This method appends the byte b to the string being built.
    String()            This method returns the string that has been created by the builder.
    Reset()             This method resets the string created by the builder.
    Len()               This method returns the number of bytes used to store the string created by the builder.
    Cap()               This method returns the number of bytes that have been allocated by the builder.
    Grow(size)          This method increases the number of bytes used allocated by the builder to store the
                        string that is being built.
████████████████████████████████████████████████████████████████████████
187.builder.String()
    Creating the string using the Builder is more efficient than using the concatenation operator on regular
    string values, especially if the Grow method is used to allocate storage in advance.
    Care must be taken to use pointers when passing Builder values to and from functions and
    methods; otherwise, the efficiency gains will be lost when the Builder is copied.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            for _, sub := range strings.Fields(text) {
                if (sub == "small") {
                    builder.WriteString("very ")
                }
                builder.WriteString(sub)
                builder.WriteRune(' ')
            }
            fmt.Println("String:", builder.String())
        }
        
`,
}
