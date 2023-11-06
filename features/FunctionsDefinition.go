package features

type FunctionsDefinitions struct {
	MapSingleDefFuncs map[string]string
}

var OriginalSingleDefFunctions = FunctionsDefinitions{
	MapSingleDefFuncs: map[string]string{
		//
		//
		//
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
		//
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
		// Splitting Strings
		"Fields(s)":                     "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"Fields()":                      "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"Fields(string)":                "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"FieldsFunc(s, func)":           "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"FieldsFunc()":                  "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"FieldsFunc(string, function)":  "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"Split(s, sub)":                 "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"Split()":                       "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"Split(string, sub)":            "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"SplitN(s, sub, max)":           "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"SplitN()":                      "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"SplitN(string, sub, max)":      "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"SplitAfter(s, sub)":            "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"SplitAfter()":                  "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"SplitAfter(string, sub)":       "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"SplitAfterN(s, sub, max)":      "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"SplitAfterN()":                 "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"SplitAfterN(string, sub, max)": "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		//
		//
		//
		// 161.Comparing Strings
		"Contains(s, substr)":            "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"Contains()":                     "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"Contains(string, substring)":    "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"ContainsAny(s, substr)":         "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"ContainsAny()":                  "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"ContainsAny(string, substring)": "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"ContainsRune(s, rune)":          "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune()":                 "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune(string, rune)":     "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune(string, r)":        "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune(string, R)":        "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune(s, R)":             "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"ContainsRune(S, R)":             "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"EqualFold(s1, s2)":              "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"EqualFold()":                    "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"EqualFold(string, string)":      "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"EqualFold(string1, string2)":    "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"HasPrefix(s, prefix)":           "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"HasPrefix(s, p)":                "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"HasPrefix(string, p)":           "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"HasPrefix(string, prefix)":      "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"HasPrefix()":                    "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"HasSuffix(s, suffix)":           "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"HasSuffix(s, suf)":              "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"HasSuffix(string, suf)":         "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"HasSuffix(string, sufix)":       "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"HasSuffix()":                    "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"HasSuffix(string, string)":      "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		//
		//
		//
		//179.Altering Strings
		"Replace(s, old, new, n)":           "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"Replace()":                         "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"Replace(1, 2, 3, 4)":               "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"Replace(1,2,3,4)":                  "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"Replace(string, old, new, number)": "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"ReplaceAll(s, old, new)":           "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"ReplaceAll()":                      "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"ReplaceAll(string, old, new)":      "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"Map(func, s)":                      "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"Map()":                             "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"Map(function, string)":             "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		//
		//
		//
		// 183.The Replacer Methods
		"Replace(s)":             "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"Replace(string)":        "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"Replace(S)":             "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"WriteString(writer, s)": "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"WriteString(w, s)":      "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"WriteString(w, str)":    "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"WriteString(w, string)": "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		//
		//
		//
		// 184.Building and Generating Strings
		"Join(slice, sep)":       "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"Join(slice, specified)": "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"Join()":                 "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"Repeat(s, count)":       "Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		"Repeat(str, count)":     "Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		//
		//
		//
		// 186.Building Strings
		"WriteString(s)":      "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"WriteString(string)": "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"WriteString()":       "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"WriteRune(r)":        "WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"WriteRune()":         "WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"WriteByte(b)":        "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"WriteByte(byte)":     "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"WriteByte()":         "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"String()":            "String() 🔔 This method returns the string that has been created by the builder.",
		"Reset()":             "Reset() 🔔 This method resets the string created by the builder.",
		// 🔔
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
		// 🔔
		//
		//
		//
		//
		"Itoa()":      "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"Itoa(val)":   "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"Itoa(value)": "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		// 🔔
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
