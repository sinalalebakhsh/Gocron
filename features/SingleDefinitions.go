package features


type SingleDefinitions struct {
	SingleDefinition map[string]string
}

var OriginSingleDef SingleDefinitions = SingleDefinitions{

	SingleDefinition: map[string]string{
		"build": `
build
    The go build command compiles the source code in the current directory 
    and generates an executable file.`,
	},
}
