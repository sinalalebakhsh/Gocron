package features

var TitleOfUsingHTMLAndTextTemplates = []string{
	"ALL HTML AND TEMPLATE",
	"ALLHTMLANDTEMPLATE",
}

var OriginalHTMLAndTemplates = DataBase {
	Alldatafield: `
366.Putting HTML and Text Templates in Context
    Question Answer:

    What are they?
    These templates allow HTML and text content to be generated dynamically
    from Go data values.
    
    Why are they useful?
    Templates are useful when large amounts of content are required, such that
    defining the content as strings would be unmanageable.
    
    How are they used?
    The templates are HTML or text files, which are annotated with instructions
    for the template processing engine. When a template is rendered, the
    instructions are processed to generate HTML or text content.

    Are there any pitfalls or limitations?
    The template syntax is counterintuitive and is not checked by the Go
    compiler. This means that care must be taken to use the correct syntax,
    which can be a frustrating process.

    Are there any alternatives?
    Templates are optional, and smaller amounts of content can be produced
    using strings.

    Summary:
    Problem                                         Solution
    ------------------                              -----------------------------------
    Generate an HTML document                       Define an HTML template with actions that
                                                    incorporate data values into the output. Load and
                                                    execute the templates, providing data for the actions.
    Enumerate loaded templates                      Enumerate the results of the Templates method.
    Locate a specific template                      Use the Lookup method.
    Produce dynamic content                         Use a template action.
    Format a data value                             Use the formatting functions.
    Suppress whitespace                             Add hyphens to the template.
    Process a slice                                 Use the slice functions.
    Conditionally execute template content          Use the conditional actions and functions.
    Create a nested template                        Use the define and template actions.
    Define a default template                       Use the block and template actions.
    Create functions for use in a template          Define template functions.
    Disable encoding for function results           Return one of the type aliases defined by the html/template package.
    Store data values for later use in a template   Define template variables.
    Generate a text document                        Use the text/template package.



    Preparing for This Chapter:
    1- go mod init htmltext
    2- printer.go:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
    3- product.go:
        package main
        type Product struct {
            Name, Category string
            Price float64
        }
        var Kayak = Product {
            Name: "Kayak",
            Category: "Watersports",
            Price: 279,
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
        func (p *Product) AddTax() float64 {
            return p.Price * 1.2
        }
        func (p * Product) ApplyDiscount(amount float64) float64 {
            return p.Price - amount
        }
    4- main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
    
████████████████████████████████████████████████████████████████████████
367.Creating HTML Templates
    The html/template package provides support for creating templates that are processed using a data
    structure to generate dynamic HTML output.
    Templates contain static content mixed with expressions that are enclosed in double curly braces,
    known as actions.

    The template uses the simplest action, which is a period (the . character)
    and which prints out the data used to execute the template, 
    which I explain in the next section.

    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
    
    A project can contain multiple templates files.
    extras.html:
        <h1>Extras Template Value: {{ . }}</h1>
    
    The new template uses the same action as the previous example but has different static content to make
    it clear which template has been executed in the next section. Once I have described the basic techniques for
    using templates, I'll introduce more complex template actions.
████████████████████████████████████████████████████████████████████████
368.Loading and Executing Templates
    Using templates is a two-step process. 
    First, the templates files are loaded and processed to create Template values.

    The html/template Functions for Loading Template Files:
    Name                        Description
    ---------------------       --------------------------------------------
    ParseFiles(...files)        This function loads one or more files, which are specified by name. The result
                                is a Template that can be used to generate content and an error that reports
                                problems loading the templates.
    ParseGlob(pattern)          This function loads one or more files, which are selected with a pattern. The
                                result is a Template that can be used to generate content and an error that
                                reports problems loading the templates.

    If you name your template files consistently, 
    then you can use the ParseGlob function to load them with a simple pattern. 
    If you want specific files—or the files are not named consistently—then you can specify
    individual files using the ParseFiles function.
████████████████████████████████████████████████████████████████████████
369.The Template Methods for Selecting and Executing Templates
    Name                                            Description
    ----------------------------                    -------------------------------------------
    Templates()                                     This function returns a slice containing pointers to the Template values that
                                                    have been loaded.
    Lookup(name)                                    This function returns a *Template for the specified loaded template.
    Name()                                          This method returns the name of the Template.
    Execute(writer, data)                           This function executes the Template, using the specified data and writes
                                                    the output to the specified Writer.
    ExecuteTemplate(writer, templateName, data)     This function executes the template with the specified name and data and
                                                    writes the output to the specified Writer.
████████████████████████████████████████████████████████████████████████
370.Loading and Executing a Template
    example:
    main.go:
        package main
        import (
            "fmt"
            "html/template"
            "os"
        )
        func main() {
            t, err := template.ParseFiles("templates/template.html")
            if (err == nil) {
                t.Execute(os.Stdout, &Kayak)
                fmt.Println()
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
371.Loading Multiple Templates
    There are two approaches to working with multiple templates. 
    The first is to create a separate Template value
    for each of them and execute them separately.

    example:
    Using Separate Templates:
    main.go:
        package main
        import (
            "fmt"
            "html/template"
            "os"
        )
        func main() {
            t1, err1 := template.ParseFiles("templates/template.html")
            t2, err2 := template.ParseFiles("templates/extras.html")
            if (err1 == nil && err2 == nil) {
                t1.Execute(os.Stdout, &Kayak)
                os.Stdout.WriteString("\n")
                t2.Execute(os.Stdout, &Kayak)
                os.Stdout.WriteString("\n")
            } else {
                Printfln("Error: %v %v", err1.Error(), err2.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Extras Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
372.Using a Combined Template
    When multiple files are loaded with the ParseFiles, 
    the result is a Template value on which the
    ExecuteTemplate method can be called to execute a specified template. 
    The filename is used as the template name, 
    which means that the templates in this example are named template.html and extras.html.

    You can call the Execute method on the Template returned by the ParseFiles or ParseGlob
    function, and the first template that was loaded will be selected 
    and used to produce the output. 
    Take care when using the ParseGlob function because 
    the first template loaded—and therefore the template that will be
    executed—may not be the file you expect.

    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func main() {
                allTemplates, err1 := template.ParseFiles("templates/template.html",
                    "templates/extras.html")
                if (err1 == nil) {
                    allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
                    os.Stdout.WriteString("\n")
                    allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
                } else {
                    Printfln("Error: %v %v", err1.Error())
                }
            }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Extras Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
373.Enumerating Loaded Templates
    It can be useful to enumerate the templates that have been loaded, 
    especially when using the ParseGlob function, 
    to make sure that all the expected files have been discovered.

    example:
    main.go:
        package main
        import (
            "html/template"
        )
        func main() {
                allTemplates, err := template.ParseGlob("templates/*.html")
            if (err == nil) {
                for _, t := range allTemplates.Templates() {
                    Printfln("Template name: %v", t.Name())
                }
            } else {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        Template name: extras.html
        Template name: template.html
████████████████████████████████████████████████████████████████████████
374.Looking Up a Specific Template
    An alternative to specifying a name is to use the Lookup method to select a template, 
    which is useful when
    you want to pass a template as an argument to a function

    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, &Kayak)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
375.The Template Actions
    Action                      Description
    ------------                -----------------------------------------
    {{ value }}                 This action inserts a data value or the result of an expression into the
    {{ expr }}                  template. A period is used to refer to the data value passed to the Execute or
                                ExecuteTemplate function. See the “Inserting Data Values” section for details.

    {{ value.fieldname }}       This action inserts the value of a struct field. See the “Inserting Data Values” section for details.
    {{ value.method arg }}      This action invokes a method and inserts the result into the template
                                output. Parentheses are not used, and arguments are separated by
                                spaces. See the “Inserting Data Values” section for details.

    {{ func arg }}              This action invokes a function and inserts the result into the output.
                                There are built-in functions for common tasks, such as formatting
                                data values, and custom functions can be defined, as described in the
                                “Defining Template Functions” section.

    {{ expr | value.method }}   Expressions can be chained together using a vertical bar so that the result
    {{ expr | func              of the first expression is used as the last argument in the second expression.
    {{ range value }}           This action iterates through the specified slice and adds the content
    ...                         between the range and end keyword for each element. The actions within
    {{ end }}                   the nested content are executed, with the current element accessible
                                through the period. See the “Using Slices in Templates” section for details.

    {{ range value }}           This action is similar to the range/end combination but defines a section
    ...                         of nested content that is used if the slice contains no elements.
    {{ else }}
    ...
    {{ end }}

    {{ if expr }}               This action evaluates an expression and executes the nested template
    ...                         content if the result is true, as demonstrated in the “Conditionally
    {{ end }}                   Executing Template Content” section. This action can be used with
                                optional else and else if clauses.

    {{ with expr }}             This action evaluates an expression and executes the nested template
    ...                         content if the result isn't nil or the empty string. This action can be used
    {{ end }}                   with optional clauses.

    {{ define "name" }}         This action defines a template with the specified name
    ...
    {{ end }}

    {{ template "name" expr }}  This action executes the template with the specified name and data and
                                inserts the result in the output.

    {{ block "name" expr }}     This action defines a template with the specified name and invokes it
    ...                         with the specified data. This is typically used to define a template that
    {{ end }}                   can be replaced by one loaded from another file, as demonstrated in the
                                “Defining Template Blocks” section.
████████████████████████████████████████████████████████████████████████
376.The Template Expressions for Inserting Values into Templates
    Inserting Data Values

    Expression          Description
    ------------        -----------------------------------------------
    .                   This expression inserts the value passed to the Execute or ExecuteTemplate method into the
                        template output.
    .Field              This expression inserts the value of the specified field into the template output.
    .Method             This expression calls the specified method without arguments and inserts the result into the
                        template output.
    .Method             This expression calls the specified method with the specified argument and inserts the result
    arg                 into the template output.
    call                This expression invokes a struct function field, using the specified arguments, which are
    .Field arg          separated by spaces. The result from the function is inserted into the template output.
    
████████████████████████████████████████████████████████████████████████
377.Inserting Data Values in the template.html
    Unlike Go code, methods are not invoked with parentheses, and arguments
    are simply specified after the name, separated by spaces. 
    It is the responsibility of the developer to ensure
    that arguments are of a type that can be used by the method or function.

    example:
    templates/template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ .Price }}</h1>
        <h1>Tax: {{ .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: 279</h1>
        <h1>Tax: 334.8</h1>
        <h1>Discount Price: 269</h1>
████████████████████████████████████████████████████████████████████████
378.Understanding Contextual Escaping
    Values are automatically escaped to make them safe for inclusion in HTML, CSS, and JavaScript code,
    with the appropriate escaping rules applied based on context. For example, a string value such as
    "It was a <big> boat" used as the text content of an HTML element would be inserted into the
    template as "It was a <big> boat" but as "It was a \u003cbig\u003e boat" when used
    as a string literal value in JavaScript code. Full details of how values are escaped can be found at

    https://golang.org/pkg/html/template.

████████████████████████████████████████████████████████████████████████
379.The Built-in Templates Functions for Formatting Data
    Name        Description
    -------     --------------------------------
    print       This is an alias to the fmt.Sprint function.
    printf      This is an alias to the fmt.Sprintf function.
    println     This is an alias to the fmt.Sprintln function.
    html        This function encodes a value for safe inclusion in an HTML document.
    js          This function encodes a value for safe inclusion in a JavaScript document.
    urlquery    This function encodes a value for use in a URL query string.


    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.3f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.000</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: 269</h1>
████████████████████████████████████████████████████████████████████████
380.Chaining Expressions
    Chaining expressions creates a pipeline for values, 
    which allows the output from one method or function
    to be used as the input for another.

    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.2f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 | printf "$%.2f" }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.00</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: $269.00</h1>
████████████████████████████████████████████████████████████████████████
381.Using Parentheses in html 
    Chaining can be used only for the last argument provided to a function. 
    An alternative approach—and
    one that can be used to set other function arguments is to use parentheses
    
    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.2f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ printf "$%.2f" (.ApplyDiscount 10) }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.00</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: $269.00</h1>
████████████████████████████████████████████████████████████████████████
382.Trimming Whitespace
    HTML isn't sensitive to the whitespace between elements, 
    but whitespace can still cause problems for text content and attribute values, 
    especially when you want to structure the content
    of a template to make it easy to read.

    example:
    template.html
        <h1>
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{ printf "$%.2f" .Price }}
        </h1>
    =============================
    Output:
        <h1>
            Name: Kayak, Category: Watersports, Price,
                $279.00
        </h1>
████████████████████████████████████████████████████████████████████████
383.The minus sign must
    The effect is to remove all of the whitespace to before or after the action.

    The minus sign can be used to trim whitespace, 
    applied immediately after or before the braces
    that open or close an action.

    example:
    template.html:
        <h1>
                Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{ printf "$%.2f" .Price }}
        </h1>

        <h1>
                Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price -}}
        </h1>
    =============================
    Output:
        <h1>
                Name: Kayak, Category: Watersports, Price,
                    $279.00
        </h1>

        <h1>
                Name: Kayak, Category: Watersports, Price,$279.00</h1>
████████████████████████████████████████████████████████████████████████
384.Trimming Additional Whitespace
    The whitespace around the final action has been removed, 
    but there is still a newline character after
    the opening h1 tag because the whitespace trimming applies only to actions.

    example:
    template.html:
        <h1>
            {{- "" -}} Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price -}}
        </h1>
    =============================
    Output:
        <h1>Name: Kayak, Category: Watersports, Price,$279.00</h1>
████████████████████████████████████████████████████████████████████████
385.Slices in Templates
    Template actions can be used to generate content for slices

    example:
    Processing a Slice in the template.html
        {{ range . -}}
            <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price }}</h1>
        {{ end }}
    main.go:
        package main

        import (
            "html/template"
            "os"
        )
        
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
        
            
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Name: Kayak, Category: Watersports, Price,$279.00</h1>
        <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
        <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
        <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
        <h1>Name: Stadium, Category: Soccer, Price,$79500.00</h1>
        <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
        <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
        <h1>Name: Bling-Bling King, Category: Chess, Price,$1200.00</h1>
████████████████████████████████████████████████████████████████████████
386.The Built-in Template Functions for Slices
    Go text templates support the built-in functions

    Name        Description
    ------      --------------
    slice       This function creates a new slice. Its arguments are the original slice, the start index, and the end index.
    index       This function returns the element at the specified index.
    len         This function returns the length of the specified slice.

    example:
    Built-in Functions in the template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range slice . 3 6 -}} 
            <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price }}</h1>
        {{ end }}
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
        <h1>Name: Stadium, Category: Soccer, Price,$79500.00</h1>
        <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
    
████████████████████████████████████████████████████████████████████████
387.Conditionally Executing Template Content
    Actions can be used to conditionally insert content into the output based on the evaluation of their
    expressions.

    The Template Conditional Functions:
        Function            Description
        ------------        -----------------------------------------------
        eq arg1 arg2        This function returns true if arg1 == arg2.
        ne arg1 arg2        This function returns true if arg1 != arg2.
        lt arg1 arg2        This function returns true if arg1 < arg2.
        le arg1 arg2        This function returns true if arg1 <= arg2.
        gt arg1 arg2        This function returns true if arg1 > arg2.
        ge arg1 arg2        This function returns true if arg1 >= arg2.
        and arg1 arg2       This function returns true if both arg1 and arg2 are true.
        not arg1            This function returns true if arg1 is false, and false if it is true.

    example:
    a Conditional Action in the template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price }}</h1>
            {{ end -}}
        {{ end }}
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>        

    Despite the use of the minus sign to trim whitespace, 
    the output is oddly formatted because of the
    way I chose to structure the template.
    
████████████████████████████████████████████████████████████████████████
388.Using the Optional Conditional Actions
    The if action can be used with optional else and else if keywords

    example:
    template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price }}</h1>
            {{ else if gt .Price 1500.00 -}}
                <h1>Expensive Product {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ else -}}
                <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ end -}}
        {{ end }}
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>
                
████████████████████████████████████████████████████████████████████████
389.Creating Named Nested Templates
    The define action is used to create a nested template that can be executed by name, 
    which allows content to
    be defined once and used repeatedly with the template action

    example:
    template.html:
        {{ define "currency" }}{{ printf "$%.2f" . }}{{ end }}
        {{ define "basicProduct" -}}
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- template "currency" .Price }}
        {{- end }}
        {{ define "expensiveProduct" -}}
            Expensive Product {{ .Name }} ({{ template "currency" .Price }})
        {{- end }}
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>{{ template "basicProduct" . }}</h1>
            {{ else if gt .Price 1500.00 -}}
                <h1>{{ template "expensiveProduct" . }}</h1>
            {{ else -}}
                <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ end -}}
        {{ end }}
    =============================
    Output:



        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>


    The define keyword is followed by the template name in quotes, 
    and the template is terminated by the end keyword. 
    The template keyword is used to execute a named template, 
    specifying the template name and a data value:
        ...
        {{- template "currency" .Price }}
        ...
████████████████████████████████████████████████████████████████████████
390.Selecting a Named Template in the main.go
    Using the define and end keywords for the main template content excludes the whitespace used to
    separate the other named templates.

    Adding a Named Template in the template.html:
        {{ define "currency" }}{{ printf "$%.2f" . }}{{ end }}
        {{ define "basicProduct" -}}
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- template "currency" .Price }}
        {{- end }}
        {{ define "expensiveProduct" -}}
            Expensive Product {{ .Name }} ({{ template "currency" .Price }})
        {{- end }}
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            <h1>First product: {{ index . 0 }}</h1>
            {{ range . -}}
                {{ if lt .Price 100.00 -}}
                    <h1>{{ template "basicProduct" . }}</h1>
                {{ else if gt .Price 1500.00 -}}
                    <h1>{{ template "expensiveProduct" . }}</h1>
                {{ else -}}
                    <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
                {{ end -}}
            {{ end }}
        {{- end}}    
    =============================
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Any of the named templates can be executed directly, but I have selected the mainTemplate
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>
████████████████████████████████████████████████████████████████████████
391.Defining Template Blocks
    Template blocks are used to define a template with default content that can be overridden in another
    template file, which requires multiple templates to be loaded and executed together. 

    This is often used to common content, such as a layout.

    The templates must be loaded so that the file that contains the block action is loaded before the file that
    contains the define action that redefines the template. 
    
    When the templates are loaded, the template defined
    in the list.html file redefines the template named body so that the content in the list.html file replaces
    the content in the template.html file.


    example:
    template.html File in the templates Folder
        {{ define "mainTemplate" -}}
            <h1>This is the layout header</h1>
            {{ block "body" . }}
                <h2>There are {{ len . }} products in the source data.</h2>
            {{ end }}
            <h1>This is the layout footer</h1>
        {{ end }}

    Output:
        <h1>This is the layout header</h1>
        
                <h2>There are 8 products in the source data.</h2>

            <h1>This is the layout footer</h1>
    

    When used alone, the output from the template file includes the content in the block. 
    But this content can be redefined by another template file.
    example:
    list.html:
        {{ define "body" }}
            {{ range . }}
                <h2>Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h2>
            {{ end -}}
        {{ end }}
    ========================
    main.go:
        package main

        import (
            "html/template"
            "os"
        )
        
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            
            allTemplates, err := template.ParseFiles("templates/template.html","templates/list.html")
        
        
        
            if err == nil {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>This is the layout header</h1>
        
            
            <h2>Product: Kayak ($279.00)</h2>

            <h2>Product: Lifejacket ($49.95)</h2>

            <h2>Product: Soccer Ball ($19.50)</h2>

            <h2>Product: Corner Flags ($34.95)</h2>

            <h2>Product: Stadium ($79500.00)</h2>

            <h2>Product: Thinking Cap ($16.00)</h2>

            <h2>Product: Unsteady Chair ($75.00)</h2>

            <h2>Product: Bling-Bling King ($1200.00)</h2>

        <h1>This is the layout footer</h1>    
████████████████████████████████████████████████████████████████████████
392.Defining Template Functions
    example:
    main.go File in the htmltext Folder:
        package main
        import (
            "html/template"
            "os"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{} {
                "getCats": GetCategories,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.html")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>This is the layout header</h1>
        
            <h2>There are 8 products in the source data.</h2>

        <h1>This is the layout footer</h1>    
████████████████████████████████████████████████████████████████████████
393.Using a Custom Function in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ . }}</h1>
            {{ end }}
        {{- end }}
    =============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: Watersports</h1>
            <h1>Category: Soccer</h1>
            <h1>Category: Chess</h1>
████████████████████████████████████████████████████████████████████████
394.Creating an HTML Fragment in the main.go
    example:
    main.go:
        ...
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, "<b>p.Category</b>")
                }
            }
            return
        }
        ...
    ===============================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
████████████████████████████████████████████████████████████████████████
395.The Types Aliases Used to Denote Content Types
    Name        Description
    --------    -----------
    CSS         This type denotes CSS content.
    HTML        This type denotes a fragment of HTML.
    HTMLAttr    This type denotes a value that will be used as the value for an HTML attribute.
    JS          This type denotes a fragment of JavaScript code.
    JSStr       This type denotes a value that is intended to appear between quotes in a JavaScript expression.
    Srcset      This type denotes a value that can be used in the srcset attribute of an img element.
    URL         This type denotes a URL.
████████████████████████████████████████████████████████████████████████
396.Returning HTML Content in the main.go
    example:
    main.go:
        ...
        func GetCategories(products []Product) (categories []template.HTML) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, "<b>p.Category</b>")
                }
            }
            return
        }
        ...
    =======================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: <b>p.Category</b></h1>
            <h1>Category: <b>p.Category</b></h1>
            <h1>Category: <b>p.Category</b></h1>
████████████████████████████████████████████████████████████████████████
397.Providing Access to Standard Library Functions
    Adding a Function Mapping in the main.go
    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
            "strings"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string{}
            for _, p := range products {
                if catMap[p.Category] == "" {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{}{
                "getCats": GetCategories,
                "lower":   strings.ToLower,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: Watersports</h1>
            <h1>Category: Soccer</h1>
            <h1>Category: Chess</h1>
████████████████████████████████████████████████████████████████████████
398.Using a Template Function in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ lower . }}</h1>
            {{ end }}
        {{- end }}
    =============================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: watersports</h1>
            <h1>Category: soccer</h1>
            <h1>Category: chess</h1>
████████████████████████████████████████████████████████████████████████
399.Defining Template Variables
    Defining and Using a Template Variable in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            {{ $length := len . }}
            <h1>There are {{ $length }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ lower . }}</h1>
            {{ end }}
        {{- end }}
    =============================================================
    Output:
            
        <h1>There are 8 products in the source data.</h1>
        <h1>Category: watersports</h1>
        <h1>Category: soccer</h1>
        <h1>Category: chess</h1>
████████████████████████████████████████████████████████████████████████
400.Defining and Using a Template Variable in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{- range getCats .  -}}
                {{ if ne ($char := slice (lower .) 0 1) "s"  }}
                    <h1>{{$char}}: {{.}}</h1>
                {{- end }}
            {{- end }}
        {{- end }}
    ==============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
                    <h1>w: Watersports</h1>
                    <h1>c: Chess</h1>
████████████████████████████████████████████████████████████████████████
401.Using Template Variables in Range Actions
    Enumerating a Map in the template.html

    example:
    main.go:
        ...
        func Exec(t *template.Template) error {
            productMap := map[string]Product {}
            for _, p := range Products {
                productMap[p.Name] = p
            }
            return t.Execute(os.Stdout, &productMap)
        }
        ...
    
    template.html:
        {{ define "mainTemplate" -}}
            {{ range $key, $value := . -}}
                <h1>{{ $key }}: {{ printf "$%.2f" $value.Price }}</h1>
            {{ end }}
        {{- end }}
    Output:
        <h1>Bling-Bling King: $1200.00</h1>
            <h1>Corner Flags: $34.95</h1>
            <h1>Kayak: $279.00</h1>
            <h1>Lifejacket: $49.95</h1>
            <h1>Soccer Ball: $19.50</h1>
            <h1>Stadium: $79500.00</h1>
            <h1>Thinking Cap: $16.00</h1>
            <h1>Unsteady Chair: $75.00</h1>
████████████████████████████████████████████████████████████████████████
402.Creating Text Templates
    Loading and Executing a Text Template in the main.go

    The Contents of the template.txt
    example:
    templates/template.txt:
        {{ define "mainTemplate" -}}
            {{ range $key, $value := . -}}
                {{ $key }}: {{ printf "$%.2f" $value.Price }}
            {{ end }}
        {{- end }}
    ---------------------------------------
    main.go:
        package main
        import (
            "text/template"
            "os"
            "strings"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            productMap := map[string]Product {}
            for _, p := range Products {
                productMap[p.Name] = p
            }
            return t.Execute(os.Stdout, &productMap)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{} {
                "getCats": GetCategories,
                "lower": strings.ToLower,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.txt")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    Output:
        Bling-Bling King: $1200.00
            Corner Flags: $34.95
            Kayak: $279.00
            Lifejacket: $49.95
            Soccer Ball: $19.50
            Stadium: $79500.00
            Thinking Cap: $16.00
            Unsteady Chair: $75.00
████████████████████████████████████████████████████████████████████████`,
}

