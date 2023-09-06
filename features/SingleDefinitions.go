package features


type SingleDefinitions struct {
    SliceSingleDefinitions []string
	SingleDefinition map[string]string
}

var OriginSingleDef SingleDefinitions = SingleDefinitions{
    SliceSingleDefinitions: []string{
        "build",
        "flag",
        "flag package",

    },

	SingleDefinition: map[string]string{
		"build": `
build:
    The go build command compiles the source code in the current directory 
    and generates an executable file.`,
	
	"flag": `
flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.`,
	
	"flag package": `
flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.`,
    
    "clean":`
clean
The go clean command removes the output produced by the go build command, 
including the executable and any temporary files that were created during the build.`,


},

}
