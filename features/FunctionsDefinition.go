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
		"islower(rune)": "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower()":     "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower(r)":    "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower(runes)":    "IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"tolower(rune)": "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower()":     "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower(r)":    "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower(runes)":    "ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"isupper(rune)": "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper(r)":    "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper(runes)":    "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper()":     "IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"toupper(rune)": "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper()":     "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper(r)":    "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper(runes)":    "ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"istitle(rune)": "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle()":     "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle(r)":    "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle(runes)":    "IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"totitle(rune)": "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle()":     "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle(r)":    "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle(runes)":    "ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		//
		//
		//
		//The strings Functions for Inspecting Strings
		"count(s, sub)":                    "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"count()":                          "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"count(string, sub)":               "Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"index(s, sub)":                    "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"index()":                          "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"index(string, sub)":               "Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexany(s, chars)":               "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"indexany()":                       "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"indexany(string, characters)":     "IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"lastindexAny(s, chars)":           "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"lastindexAny()":                   "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"lastindexAny(string, characters)": "LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"indexbyte(s, b)":                  "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexbyte()":                      "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexbyte(string, byte)":          "IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"lastindexByte(s, b)":              "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"lastindexByte()":                  "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"lastindexByte(string, byte)":      "LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"indexfunc(s, func)":               "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"indexfunc()":                      "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"indexfunc(string, function)":      "IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"lastindexFunc(s, func)":           "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"lastindexFunc()":                  "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"lastindexFunc(string, function)":  "LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		//
		//
		//
		// Splitting Strings
		"fields(s)":                     "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fields()":                      "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fields(string)":                "Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fieldsfunc(s, func)":           "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"fieldsfunc()":                  "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"fieldsfunc(string, function)":  "FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"split(s, sub)":                 "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"split()":                       "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"split(string, sub)":            "Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"splitn(s, sub, max)":           "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitn()":                      "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitn(string, sub, max)":      "SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitafter(s, sub)":            "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitafter()":                  "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitafter(string, sub)":       "SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitaftern(s, sub, max)":      "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"splitaftern()":                 "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"splitaftern(string, sub, max)": "SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		//
		//
		//
		// 161.Comparing Strings
		"contains(s, substr)":            "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"contains()":                     "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"contains(string, substring)":    "Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"containsany(s, substr)":         "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsany()":                  "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsany(string, substring)": "ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsrune(s, rune)":          "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune()":                 "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, rune)":     "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, r)":        "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, R)":        "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(s, R)":             "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(S, R)":             "ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"equalfold(s1, s2)":              "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold()":                    "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold(string, string)":      "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold(string1, string2)":    "EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"hasprefix(s, prefix)":           "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(s, p)":                "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(string, p)":           "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(string, prefix)":      "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix()":                    "HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hassuffix(s, suffix)":           "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(s, suf)":              "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, suf)":         "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, sufix)":       "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix()":                    "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, string)":      "HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		//
		//
		//
		//179.Altering Strings
		"replace(s, old, new, n)":           "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace()":                         "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(1, 2, 3, 4)":               "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(1,2,3,4)":                  "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(string, old, new, number)": "Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replaceall(s, old, new)":           "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"replaceall()":                      "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"replaceall(string, old, new)":      "ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"map(func, s)":                      "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"map()":                             "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"map(function, string)":             "Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		//
		//
		//
		// 183.The Replacer Methods
		"replace(s)":             "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"replace(string)":        "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"replace(str)":             "Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"writestring(writer, s)": "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, s)":      "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, str)":    "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, string)": "WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		//
		//
		//
		// 184.Building and Generating Strings
		"join(slice, sep)":       "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"join(slice, specified)": "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"join()":                 "Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"repeat(s, count)":       "Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		"repeat(str, count)":     "Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		//
		//
		//
		// 186.Building Strings
		"writestring(s)":      "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writestring(string)": "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writestring()":       "WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writerune(r)":        "WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"writerune()":         "WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"writebyte(b)":        "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"writebyte(byte)":     "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"writebyte()":         "WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"string()":            "String() 🔔 This method returns the string that has been created by the builder.",
		"reset()":             "Reset() 🔔 This method resets the string created by the builder.",
		"len()":               "Len() 🔔 This method returns the number of bytes used to store the string created by the builder.",
		// 🔔
		//
		//
		"atoi()":                 "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"atoi(str)":              "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"atoi(string)":           "Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"formatint()":            "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatint(value, base)": "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatint(val, base)":   "FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatuint()":           "FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		"formatuint(val, base)":  "FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		// 🔔
		//
		//
		//
		//
		"itoa()":      "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"itoa(val)":   "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"itoa(value)": "Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		// 🔔
		//
		//
		//
		//
		"formatfloat()":                             "FormatFloat(val, format, precision, size) 🔔 This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"formatfloat(1,2,3,4)":                      "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"formatfloat(val, format, precision, size)": "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		//
		//
		//
		// The Key Reflection Functions
		"typeof(val)":  "TypeOf(val) 🔔 This function returns a value that implements the Type interface, which describes the type of the specified value. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
		"valueof(val)": "ValueOf(val) 🔔 This function returns a Value struct, which allows the specified value to be inspected and manipulated. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
	},
}
