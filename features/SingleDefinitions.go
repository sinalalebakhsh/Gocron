package features


type SingleDefinitions struct {
	SingleDefinition map[string]string
}

var OriginSingleDef SingleDefinitions = SingleDefinitions{

	SingleDefinition: map[string]string{
		"build": `
build:
    The go build command compiles the source code in the current directory 
    and generates an executable file.`,
	
	"flag":`
flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.`,
	
	"flag package":`
flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.`},
}
