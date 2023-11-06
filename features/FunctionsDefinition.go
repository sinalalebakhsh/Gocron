package features

type FunctionsDefinitions struct {
	MapSingleDefFuncs map[string]string
}

var OriginalSingleDefFunctions = FunctionsDefinitions{
	MapSingleDefFuncs: map[string]string{
		//  Working with Character Case
		"IsLower(rune)": "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"IsLower()":     "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"IsLower(r)":    "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"IsLower(R)":    "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"ToLower(rune)": "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"ToLower()":     "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"ToLower(r)":    "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"ToLower(R)":    "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"IsUpper(rune)": "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"IsUpper(r)":    "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"IsUpper(R)":    "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"IsUpper()":     "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"ToUpper(rune)": "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"ToUpper()":     "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"ToUpper(r)":    "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"ToUpper(R)":    "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		//
		//
		//
		//
		"Atoi()":                 "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"Atoi(str)":              "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"Atoi(string)":           "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"FormatInt()":            "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"FormatInt(value, base)": "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"FormatInt(val, base)":   "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"FormatUint()":           "FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		"FormatUint(val, base)":  "FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		//
		//
		//
		//
		"Itoa()":      "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"Itoa(val)":   "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"Itoa(value)": "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		//
		//
		//
		//
		"FormatFloat()":                             "FormatFloat(val, format, precision, size) 🔔 This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"FormatFloat(1,2,3,4)":                      "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"FormatFloat(val, format, precision, size)": "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		//
		//
		//
		// The Key Reflection Functions
		"TypeOf(val)":  "TypeOf(val) 🔔 This function returns a value that implements the Type interface, which describes the type of the specified value. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
		"ValueOf(val)": "ValueOf(val) 🔔 This function returns a Value struct, which allows the specified value to be inspected and manipulated. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
	},
}
