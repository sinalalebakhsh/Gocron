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
		"IsTitle(rune)": "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"IsTitle()":     "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"IsTitle(r)":    "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"IsTitle(R)":    "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"ToTitle(rune)": "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"ToTitle()":     "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"ToTitle(r)":    "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"ToTitle(R)":    "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		//
		//The strings Functions for Inspecting Strings
		"Count(s, sub)":                    "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"Count()":                          "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"Count(string, sub)":               "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"Index(s, sub)":                    "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"Index()":                          "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"Index(string, sub)":               "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"IndexAny(s, chars)":               "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"IndexAny()":                       "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"IndexAny(string, characters)":     "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"LastIndexAny(s, chars)":           "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"LastIndexAny()":                   "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"LastIndexAny(string, characters)": "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"IndexByte(s, b)":                  "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"IndexByte()":                      "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"IndexByte(string, byte)":          "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"LastIndexByte(s, b)":              "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"LastIndexByte()":                  "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"LastIndexByte(string, byte)":      "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"IndexFunc(s, func)":               "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"IndexFunc()":                      "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"IndexFunc(string, function)":      "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"LastIndexFunc(s, func)":           "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"LastIndexFunc()":                  "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"LastIndexFunc(string, function)":  "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		
		//
		//
		//
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
