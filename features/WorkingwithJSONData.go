package features

var TitleOfSlice = []string{
	"JSON",
	"JSONS",
	"JSONDATA",
	"JSON DATA",
	"WORKWITHJSONDATA",
	"WORK WITH JSON DATA",
	"WORKINGWITHJSONDATA",
	"WORKING WITH JSON DATA",
	"WORKING JSON DATA",
	"WORKINGJSONDATA",
	"JSON DATAS",
	"JSONDATAS",
	"DATA JSON",
	"DATAJSON",
}

type WorkingwithJSONData struct {
	allJSONData string
}

var OriginalJSONData = WorkingwithJSONData{
	allJSONData: `
302.Working with JSON Data
    Putting Working with JSON Data in Context

    What is it?
    JSON data is the de facto standard for exchanging data, especially in HTTP applications.

    Why is it useful?
    JSON is simple enough to be supported by any language but can represent
    relatively complex data.

    How is it used?
    The encoding/json package provides support for encoding and decoding JSON data.

    Are there any pitfalls or limitations?
    Not all Go data types can be represented in JSON, which requires the developer
    to be mindful of how Go data types will be expressed.

    Are there any alternatives?
    There are many other data encodings available, some of which are supported by
    the Go standard library.

    The safest approach is to define a map with string keys and empty interface values, which ensures that
    all the key-value pairs in the JSON data can be decoded into the map


    Problem                     Solution
    -----------                 ------------------------
    Encode JSON                 dataCreate an Encoder with a Writer and invoke the Encode method
    Control struct encoding     Use JSON struct tags or implement the Mashaler interface
    Decode JSON data            Create a Decoder with a Reader and invoke the Decode method
    Control struct decoding     Use JSON struct tags or implement the Unmarshaler interface
████████████████████████████████████████████████████████████████████████
303.The encoding/json Constructor Functions for JSON Data
    Name                    Description
    -------------------     --------------------------------------------
    NewEncoder(writer)      This function returns an Encoder, which can be used to encode JSON data and
                            write it to the specified Writer.
    NewDecoder(reader)      This function returns a Decoder, which can be used to read JSON data from the
                            specified Reader and decode it.
████████████████████████████████████████████████████████████████████████
304.The Functions for Creating and Parsing JSON Data
    Name                            Description
    ------------------------        -----------------------------------------------
    Marshal(value)                  This function encodes the specified value as JSON. The results are the
                                    JSON content expressed in a byte slice and an error, which indicates any
                                    encoding problems.
    Unmarshal(byteSlice, val)       This function parses JSON data contained in the specified slice of bytes
                                    and assigns the result to the specified value.
████████████████████████████████████████████████████████████████████████
305.The Encoder Methods
    The NewEncoder constructor function is used to create an Encoder, which can be used to write JSON data to a Writer.

    Name                            Description
    -------------------------       --------------------------------------------------
    Encode(val)                     This method encodes the specified value as JSON and writes it to the Writer.
    SetEscapeHTML(on)               This method accepts a bool argument that, when true, encodes
                                    characters that would be dangerous in HTML to be escaped. The default
                                    behavior is to escape these characters.
    SetIndent(prefix, indent)       This method specifies a prefix and indentation that is applied to the name
                                    of each field in the JSON output.
████████████████████████████████████████████████████████████████████████
306.Expressing the Basic Go Data Types in JSON
    Data                TypeDescription
    ----------------    ------------------------------------
    bool                Go bool values are expressed as JSON true or false.
    string              Go string values are expressed as JSON strings. By default, unsafe HTML
                        characters are escaped.
    float32, float64    Go floating-point values are expressed as JSON numbers.
    int, int<size>      Go integer values are expressed as JSON numbers.
    uint, uint<size>    Go integer values are expressed as JSON numbers.
    byte                Go bytes are expressed as JSON numbers.
    rune                Go runes are expressed as JSON numbers.
    nil                 The Go nil value is expressed as the JSON null value.
    Pointers            The JSON encoder follows pointers and encodes the value at the pointer's location.
████████████████████████████████████████████████████████████████████████
307.encoding/json
    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var b bool = true
            var str string = "Hello"
            var fval float64 = 99.99
            var ival int = 200
            var pointer *int = &ival
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            for _, val := range []interface{}{b, str, fval, ival, pointer} {
                encoder.Encode(val)
            }
            fmt.Print(writer.String())
        }
    Output:
        true
        "Hello"
        99.99
        200
        200
████████████████████████████████████████████████████████████████████████
308.Encoding Slices and Arrays
    Example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
            numbers := [3]int{10, 20, 30}
            var byteArray [5]byte
            copy(byteArray[0:], []byte(names[0]))
            byteSlice := []byte(names[0])
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            
            encoder.Encode(names)
            encoder.Encode(numbers)
            encoder.Encode(byteArray)
            encoder.Encode(byteSlice)
            fmt.Print(writer.String())
    }
    Output:
        ["Kayak","Lifejacket","Soccer Ball"]
        [10,20,30]
        [75,97,121,97,107]
        "S2F5YWs="
████████████████████████████████████████████████████████████████████████
309.Encoding Maps
    Go maps are encoded as JSON objects, with the map keys used as the object keys. The values contained in
    the map are encoded based on their type.
    Maps can also be useful for creating custom JSON representations of Go data.


    encoder.Encode()
    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            m := map[string]float64{
                "Kayak":      279,
                "Lifejacket": 49.95,
            }
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            encoder.Encode(m)
            fmt.Print(writer.String())
        }
    Output:
            {"Kayak":279,"Lifejacket":49.95}

████████████████████████████████████████████████████████████████████████
310.Encoding Structs
    The Encoder expresses struct values as JSON objects, using the exported struct field names as the object’s
    keys and the field values as the object's values

    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            encoder.Encode(Kayak)
            fmt.Print(writer.String())
        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
311.Effect of Promotion in JSON in Encoding
    When a struct defines an embedded field that is also a struct, the fields of the embedded struct are promoted
    and encoded as though they are defined by the enclosing type.
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product
            Discount float64
        }
    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct {
                Product: &Kayak,
                Discount: 10.50,
            }
            encoder.Encode(&dp)
            fmt.Print(writer.String())
        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279,"Discount":10.5}
████████████████████████████████████████████████████████████████████████
312.Customizing the JSON Encoding of Structs
    How a struct is encoded can be customized using struct tags, which are string literals that follow fields. Struct
    tags are part of the Go support for reflection, 
    that tags follow fields and can be used to alter two aspects of how a field is encoded in JSON.

    example:
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product ^json:"product"^           // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64
        }
    
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279},"Discount":10.5}

████████████████████████████████████████████████████████████████████████
313.Omitting a Field حذف یک فیلد
    The Encoder skips fields decorated with a tag that specifies a hyphen (the - character) for the name
    The new tag tells the Encoder to skip the Discount field when creating the JSON representation of a
    DIscountedProduct value.

    exampe:
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product ^json:"product"^           // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^         // ------------------------->   Use the symbol ^ above the Tab button instead

        }        
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
████████████████████████████████████████████████████████████████████████
314.Omitting Unassigned Fields
    By default, the JSON Encoder includes struct fields, even when they have not been assigned a value
    
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            encoder.Encode(&dp)
            dp2 := DiscountedProduct { Discount: 10.50 }
            encoder.Encode(&dp2)
            fmt.Print(writer.String())
        }
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
        {"product":null}

    To omit a nil field, the omitempty keyword is added to the tag for the field
    discount.go:
        package main
        type DiscountedProduct struct {
            // *Product ^json:"product"^                   // ------------------------->   Use the symbol ^ above the Tab button instead

            *Product ^json:"product,omitempty"^            // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^
        }
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
        {}


    To skip a nil field without changing the name or field promotion, specify the omitempty keyword without a name
    discount.go:
        package main
        type DiscountedProduct struct {
            // *Product ^json:"product"^
            // *Product ^json:"product,omitempty"^
            *Product ^json:",omitempty"^                    // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^                     // ------------------------->   Use the symbol ^ above the Tab button instead

        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279}
        {}
████████████████████████████████████████████████████████████████████████
315.Forcing Fields to be Encoded as Strings
    Struct tags can be used to force a field value to be encoded as a string, overriding the normal encoding for
    the field type
    
    example:
    discount.go:
     package main
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^
            // Discount float64 ^json:"-"^              // ------------------------->   Use the symbol ^ above the Tab button                

            Discount float64 ^json:",string"^           // ------------------------->   Use the symbol ^ above the Tab button    

        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"}
        {"Discount":"10.5"}
████████████████████████████████████████████████████████████████████████
316.Encoding Interfaces
    The JSON encoder can be used on values assigned to interface variables, but it is the dynamic type that
    is encoded.

    No aspect(جنبه) of the interface is used to adapt the JSON, and all the exported fields of each value in the slice
    are included in the JSON. This can be a useful feature, but care must be taken when decoding this kind of
    JSON, because each value can have a different set of fields

    example:
    interface.go:
        package main
        type Named interface{ GetName() string }
        type Person struct{ PersonName string }
        func (p *Person) GetName() string { return p.PersonName }
        func (p *DiscountedProduct) GetName() string { return p.Name }

    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            namedItems := []Named { &dp, &Person{ PersonName: "Alice"}}
            encoder.Encode(namedItems)
            fmt.Print(writer.String())
        }
    Output:
        [{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"},{"PersonName":"Alice"}]

████████████████████████████████████████████████████████████████████████
317.The Marshaler Method
    Creating Completely Custom JSON Encodings
    The Encoder checks to see whether a struct implements the Marshaler interface, which denotes a type that
    has a custom encoding and which defines the method.

    Name                Description
    --------------      ------------------------------------------
    MarshalJSON()       This method is invoked to create a JSON representation of a value and returns a byte
                        slice containing the JSON and an error indicating encoding problems.
████████████████████████████████████████████████████████████████████████
318.json.Marshal()
    The MarshalJSON method can generate JSON in any way that suits the project.

    I define a map with string keys and use the
    empty interface for the values. This allows me to build the JSON by adding key-value pairs to the map
    and then pass the map to the Marshal function, which uses the built-in support
    to encode each of the values contained in the map.

    example:
    discount.go:
        package main
        import "encoding/json"
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^
            Discount float64 ^json:",string"^
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if dp.Product != nil {
                m := map[string]interface{}{
                    "product": dp.Name,
                    "cost":    dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }

    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            namedItems := []Named { &dp, &Person{ PersonName: "Alice"}}
            encoder.Encode(namedItems)
            fmt.Print(writer.String())
        }
        
    Output:
        [{"cost":268.5,"product":"Kayak"},{"PersonName":"Alice"}]
    
████████████████████████████████████████████████████████████████████████
319.Decoding JSON Data
    The NewDecoder constructor function creates a Decoder, which can be used to decode JSON data obtained
    from a Reader.

    The Decoder Methods:

    Name                        Description
    ---------------------       --------------------------------------------
    Decode(value)               This method reads and decodes data, which is used to create the specified
                                value. The method returns an error that indicates problems decoding the
                                data to the required type or EOF.
    DisallowUnknownFields()     By default, when decoding a struct type, the Decoder ignores any key in
                                the JSON data for which there is no corresponding struct field. Calling this
                                method causes the Decode to return an error, rather than ignoring the key.
    UseNumber()                 By default, JSON number values are decoded into float64 values. Calling
                                this method uses the Number type instead, as described in the “Decoding
                                Number Values” section.
████████████████████████████████████████████████████████████████████████
320.Decoding Basic Data Types
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
            "asd/asd"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button    

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        asd.Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                asd.Printfln("Decoded (%T): %v", val, val)
            }
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (float64): 200
    
████████████████████████████████████████████████████████████████████████
321.Decoding Number Values
    The Methods Defined by the Number Type
    Name            Description
    ---------       ------------------------------------
    Int64()         This method returns the decoded value as a int64 and an error that indicates if the value
                    cannot be converted.
    Float64()       This method returns the decoded value as a float64 and an error that indicates if the
                    value cannot be converted.
    String()        This method returns the unconverted string from the JSON data.

    Not all JSON number values can be expressed as
    Go int64 values, so this is the method that is typically called first. 
    If attempting to convert to an integer
    fails, then the Float64 method can be called. 
    If a number cannot be converted to either Go type, then the
    String method can be used to get the unconverted string from the JSON data.
████████████████████████████████████████████████████████████████████████
322.Decoding Numbers
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button 

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                if num, ok := val.(json.Number); ok {
                    if ival, err := num.Int64(); err == nil {
                        Printfln("Decoded Integer: %v", ival)
                    } else if fpval, err := num.Float64(); err == nil {
                        Printfln("Decoded Floating Point: %v", fpval)
                    } else {
                        Printfln("Decoded String: %v", num.String())
                    }
                } else {
                    Printfln("Decoded (%T): %v", val, val)
                }
            }
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (float64): 200
████████████████████████████████████████████████████████████████████████
323.Specifying Types for Decoding
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button 

            var bval bool
            var sval string
            var fpval float64
            var ival int
            vals := []interface{}{&bval, &sval, &fpval, &ival}
            decoder := json.NewDecoder(reader)
            for i := 0; i < len(vals); i++ {
                err := decoder.Decode(vals[i])
                if err != nil {
                    Printfln("Error: %v", err.Error())
                    break
                }
            }
            Printfln("Decoded (%T): %v", bval, bval)
            Printfln("Decoded (%T): %v", sval, sval)
            Printfln("Decoded (%T): %v", fpval, fpval)
            Printfln("Decoded (%T): %v", ival, ival)
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (int): 200
████████████████████████████████████████████████████████████████████████
324.Decoding Arrays
    The Decoder processes arrays automatically, but care must be taken because JSON allows arrays to contain
    values of different types, which conflicts with the strict type rules enforced by Go.

    The source JSON data contains two arrays, one of which contains only numbers and one of which mixes
    numbers and strings. The Decoder doesn’t try to figure out if a JSON array can be represented using a single
    Go type and decodes every array into an empty interface slice:
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^[10,20,30]["Kayak","Lifejacket",279]^)     // ------------------------->   Use the symbol ^ above the Tab button 

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                Printfln("Decoded (%T): %v", val, val)
            }
        }
    Output:
        Decoded ([]interface {}): [10 20 30]
        Decoded ([]interface {}): [Kayak Lifejacket 279]
████████████████████████████████████████████████████████████████████████
325.Specifying the Decoded Array Type
    The second array contains a mix of values, which means that I have to specify
    the empty interface as the target type. The literal slice syntax is awkward when using the empty interface
    because two sets of braces are required:
    ...
    mixed := []interface{} {}


    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^[10,20,30]["Kayak","Lifejacket",279]^) // ------------------------->   Use the symbol ^ above the Tab button 

            ints := []int {}
            mixed := []interface{} {}
            vals := []interface{} { &ints, &mixed}
            decoder := json.NewDecoder(reader)
            for i := 0; i < len(vals); i++ {
                err := decoder.Decode(vals[i])
                if err != nil {
                    Printfln("Error: %v", err.Error())
                    break
                }
            }
            Printfln("Decoded (%T): %v", ints, ints)
            Printfln("Decoded (%T): %v", mixed, mixed)
        }
    Output:
        Decoded ([]int): [10 20 30]
        Decoded ([]interface {}): [Kayak Lifejacket 279]
████████████████████████████████████████████████████████████████████████
326.Decoding Maps
    The safest approach is to define a map with string keys and empty interface values, which ensures that
    all the key-value pairs in the JSON data can be decoded into the map
    JavaScript objects are expressed as key-value pairs, which makes it easy to decode them into Go maps
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^{"Kayak" : 279, "Lifejacket" : 49.95}^)    // ------------------------->   Use the symbol ^ above the Tab button 
 
            m := map[string]interface{}{}
            decoder := json.NewDecoder(reader)
            err := decoder.Decode(&m)
            if err != nil {
                Printfln("Error: %v", err.Error())
            } else {
                Printfln("Map: %T, %v", m, m)
                for k, v := range m {
                    Printfln("Key: %v, Value: %v", k, v)
                }
            }
        }
    Output:
        Map: map[string]interface {}, map[Kayak:279 Lifejacket:49.95]
        Key: Kayak, Value: 279
        Key: Lifejacket, Value: 49.95
████████████████████████████████████████████████████████████████████████
327.a Specific Value Type 
    A single JSON object can be used for multiple data types as values, but if you know in advance that you
    will be decoding a JSON object that has a single value type, then you can be more specific when defining the
    map into which the data will be decoded.

    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^{"Kayak" : 279, "Lifejacket" : 49.95}^)    // ------------------------->   Use the symbol ^ above the Tab button 
 
            m := map[string]float64 {}
            decoder := json.NewDecoder(reader)
            err := decoder.Decode(&m)
            if err != nil {
                Printfln("Error: %v", err.Error())
            } else {
                Printfln("Map: %T, %v", m, m)
                for k, v := range m {
                    Printfln("Key: %v, Value: %v", k, v)
                }
            }
        }
    Output:
        Map: map[string]float64, map[Kayak:279 Lifejacket:49.95]
        Key: Kayak, Value: 279
        Key: Lifejacket, Value: 49.95
████████████████████████████████████████████████████████████████████████
328.Decoding Structs

    The Decoder decodes the JSON object and uses the keys to set the values of the exported struct fields.
    The capitalization of the fields and JSON keys don't have to match, and the Decoder will ignore any JSON
    key for which there isn't a struct field and ignore any struct field for which there is no JSON key.
    The JSON objectscontain different capitalization and have more or fewer keys than the Product struct
    fields. The Decoder processes the data as best as it can.

    example:
    main.go:
        package main
        import (
            "strings"
            "encoding/json"
            "io"
        )
        func main() {
            reader := strings.NewReader(^    // ------------------------->   Use the symbol ^ above the Tab button 
 
                {"Name":"Kayak","Category":"Watersports","Price":279}
                {"Name":"Lifejacket","Category":"Watersports" }
                {"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
            ^)                                          // ------------------------->   Use the symbol ^ above the Tab button 
 
            decoder := json.NewDecoder(reader)
            for {
                var val Product
                err := decoder.Decode(&val)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                } else {
                    Printfln("Name: %v, Category: %v, Price: %v",
                        val.Name, val.Category, val.Price)
                }
            }
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279
        Name: Lifejacket, Category: Watersports, Price: 0
        Name: Canoe, Category: Watersports, Price: 100

████████████████████████████████████████████████████████████████████████
329.Decoding to Interface Types
    As I explained earlier in the chapter, the JSON encoder deals with interfaces by encoding the value
    using the exported fields of the dynamic type. This is because JSON deals with key-value pairs and
    has no way to express methods. As a consequence, you cannot decode directly to an interface variable
    from JSON. Instead, you must decode to a struct or map and then assign the value that is created to an
    interface variable.
████████████████████████████████████████████████████████████████████████
330.Disallowing Unused Keys غیر مجاز کردن
    By default, the Decoder will ignore JSON keys for which there is no corresponding struct field. This behavior
    can be changed by calling the DisallowUnknownFields method

    example:
    Disallowing Unused Keys in the main.go
        ...
        decoder := json.NewDecoder(reader)
        decoder.DisallowUnknownFields()
        ...
    output:
        Name: Kayak, Category: Watersports, Price: 279
        Name: Lifejacket, Category: Watersports, Price: 0
        Error: json: unknown field "inStock"
████████████████████████████████████████████████████████████████████████
331.Struct Tags 
    The tag applied to the Discount field tells the Decoder that the value for this field should be obtained
    from the JSON key named offer and that the value will be parsed from a string, instead of the JSON
    number that would usually be expected for a Go float64 value.

    example:
    discount.go:
        package main
        import "encoding/json"
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^                 // ------------------------->   Use the symbol ^ above the Tab button 
 
            Discount float64 ^json:"offer,string"^       // ------------------------->   Use the symbol ^ above the Tab button 
 
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if (dp.Product != nil) {
                m := map[string]interface{} {
                    "product": dp.Name,
                    "cost": dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }    
    main.go:
        package main
        import (
            "strings"
            "encoding/json"
            "io"
        )
        func main() {
            reader := strings.NewReader(^
                {"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}^)   // ------------------------->   Use the symbol ^ above the Tab button 
 
            decoder := json.NewDecoder(reader)
            for {
                var val DiscountedProduct
                err := decoder.Decode(&val)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                } else {
                    Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
                        val.Name, val.Category, val.Price, val.Discount)
                }
            }
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279, Discount: 10
████████████████████████████████████████████████████████████████████████
332.Creating Completely Custom JSON Decoders
    The Unmarshaler Method
    Name                            Description
    ------------------------        ------------------------------------------
    UnmarshalJSON(byteSlice)        This method is invoked to decode JSON data contained in the specified
                                    byte slice. The result is an error indicating encoding problems.
████████████████████████████████████████████████████████████████████████
333.Defining a Custom Decoder
    This implementation of the UnmarshalJSON method uses the Unmarshal method to decode the JSON
    data into a map and then checks the type of each value required for the DiscountedProduct struct.


    example:
    discount.go:
        package main
        import (
            "encoding/json"
            "strconv"
        )
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^               // ------------------------->   Use the symbol ^ above the Tab button 
    
            Discount float64 ^json:"offer,string"^         // ------------------------->   Use the symbol ^ above the Tab button 
    
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if dp.Product != nil {
                m := map[string]interface{}{
                    "product": dp.Name,
                    "cost":    dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }
        func (dp *DiscountedProduct) UnmarshalJSON(data []byte) (err error) {
            mdata := map[string]interface{}{}
            err = json.Unmarshal(data, &mdata)
            if dp.Product == nil {
                dp.Product = &Product{}
            }
            if err == nil {
                if name, ok := mdata["Name"].(string); ok {
                    dp.Name = name
                }
                if category, ok := mdata["Category"].(string); ok {
                    dp.Category = category
                }
                if price, ok := mdata["Price"].(float64); ok {
                    dp.Price = price
                }
                if discount, ok := mdata["Offer"].(string); ok {
                    fpval, fperr := strconv.ParseFloat(discount, 64)
                    if fperr == nil {
                        dp.Discount = fpval
                    }
                }
            }
            return
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279, Discount: 10
`,
}
