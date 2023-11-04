package features

type FunctionsDefinitions struct {
	MapSingleDefFuncs map[string]string
}

var OriginalSingleDefFunctions = FunctionsDefinitions{
	MapSingleDefFuncs: map[string]string{

		// The Key Reflection Functions
		"TypeOf(val)":  "TypeOf(val) 🔔 This function returns a value that implements the Type interface, which describes the type of the specified value.",
		"ValueOf(val)": "ValueOf(val) 🔔 This function returns a Value struct, which allows the specified value to be inspected and manipulated.",
	},
}
