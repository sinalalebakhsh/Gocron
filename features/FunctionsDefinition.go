package features

type FunctionsDefinitions struct {
	MapSingleDefFuncs map[string]string
}

var OriginalSingleDefFunctions = FunctionsDefinitions{
	MapSingleDefFuncs: map[string]string{
		/*




		 */                       //47.Atoi(str) 49.FormatInt(val, base) 50.FormatUint(val, base)
		"atoi()":                 "47🚀 Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"atoi(str)":              "47🚀 Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"atoi(string)":           "47🚀 Atoi(str) 🔔 This function parses a string into a base 10 int and is equivalent to calling ParseInt(str, 10, 0)",
		"formatint()":            "49🚀 FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatint(value, base)": "49🚀 FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatint(val, base)":   "49🚀 FormatInt(value, base) 🔔 This function returns a string representation of the specified int64 value, expressed in the specified base.",
		"formatuint()":           "50🚀 FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		"formatuint(val, base)":  "50🚀 FormatUint(val, base) 🔔 This function returns a string representation of the specified uint64 value, expressed in the specified base.",
		/*




		 */            //52.Itoa(val)
		"itoa()":      "52🚀 Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"itoa(val)":   "52🚀 Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		"itoa(value)": "52🚀 Itoa(val) 🔔 This function returns a string representation of the specified int value, expressed using base 10.",
		/*




		 */               // 165.Working with Character Case
		"islower(rune)":  "165🚀 IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower()":      "165🚀 IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower(r)":     "165🚀 IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"islower(runes)": "165🚀 IsLower(rune) 🔔 This function returns true if the specified rune is lowercase.",
		"tolower(rune)":  "165🚀 ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower()":      "165🚀 ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower(r)":     "165🚀 ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"tolower(runes)": "165🚀 ToLower(rune) 🔔 This function returns the lowercase rune associated with the specified rune.",
		"isupper(rune)":  "165🚀 IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper(r)":     "165🚀 IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper(runes)": "165🚀 IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"isupper()":      "165🚀 IsUpper(rune) 🔔 This function returns true if the specified rune is uppercase.",
		"toupper(rune)":  "165🚀 ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper()":      "165🚀 ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper(r)":     "165🚀 ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"toupper(runes)": "165🚀 ToUpper(rune) 🔔 This function returns the upper rune associated with the specified rune.",
		"istitle(rune)":  "165🚀 IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle()":      "165🚀 IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle(r)":     "165🚀 IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"istitle(runes)": "165🚀 IsTitle(rune) 🔔 This function returns true if the specified rune is title case.",
		"totitle(rune)":  "165🚀 ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle()":      "165🚀 ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle(r)":     "165🚀 ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		"totitle(runes)": "165🚀 ToTitle(rune) 🔔 This function returns the title case rune associated with the specified rune.",
		/*





		 */                                 // 167.Inspecting Strings || The strings Functions for Inspecting Strings
		"count(s, sub)":                    "167🚀 Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"count()":                          "167🚀 Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"count(string, sub)":               "167🚀 Count(s, sub) 🔔 This function returns an int that reports how many times the specified substring is found in the string s.",
		"index(s, sub)":                    "167🚀 Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"index()":                          "167🚀 Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"index(string, sub)":               "167🚀 Index(s, sub) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexany(s, chars)":               "167🚀 IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"indexany()":                       "167🚀 IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"indexany(string, characters)":     "167🚀 IndexAny(s, chars) 🔔 These functions return the first or last occurrence of any character in the...",
		"lastindexAny(s, chars)":           "167🚀 LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"lastindexAny()":                   "167🚀 LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"lastindexAny(string, characters)": "167🚀 LastIndexAny(s, chars) 🔔 specified string within the string s, or -1 if there is no occurrence.",
		"indexbyte(s, b)":                  "167🚀 IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexbyte()":                      "167🚀 IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"indexbyte(string, byte)":          "167🚀 IndexByte(s, b) 🔔 These functions return the index of the first or last occurrence of a specified",
		"lastindexByte(s, b)":              "167🚀 LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"lastindexByte()":                  "167🚀 LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"lastindexByte(string, byte)":      "167🚀 LastIndexByte(s, b) 🔔 byte within the string s, or -1 if there is no occurrence.",
		"indexfunc(s, func)":               "167🚀 IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"indexfunc()":                      "167🚀 IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"indexfunc(string, function)":      "167🚀 IndexFunc(s, func) 🔔 These functions return the index of the first or last occurrence of the...",
		"lastindexFunc(s, func)":           "167🚀 LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"lastindexFunc()":                  "167🚀 LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		"lastindexFunc(string, function)":  "167🚀 LastIndexFunc(s, func) 🔔  character in the string s for which the specified function returns true, as described in the “Inspecting Strings with Custom Functions” section.",
		/*




		 */                               //161.Comparing Strings
		"contains(s, substr)":            "161🚀 Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"contains()":                     "161🚀 Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"contains(string, substring)":    "161🚀 Contains(s, substr) 🔔 This function returns true if the string s contains substr and false if it does not.",
		"containsany(s, substr)":         "161🚀 ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsany()":                  "161🚀 ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsany(string, substring)": "161🚀 ContainsAny(s, substr) 🔔 This function returns true if the string s contains any of the characters contained in the string substr.",
		"containsrune(s, rune)":          "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune()":                 "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, rune)":     "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, r)":        "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(string, R)":        "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(s, R)":             "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"containsrune(S, R)":             "161🚀 ContainsRune(s, rune) 🔔 This function returns true if the string s contains a specific rune.",
		"equalfold(s1, s2)":              "161🚀 EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold()":                    "161🚀 EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold(string, string)":      "161🚀 EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"equalfold(string1, string2)":    "161🚀 EqualFold(s1, s2) 🔔 This function performs a case-insensitive comparison and returns true of strings s1 and s2 are the same.",
		"hasprefix(s, prefix)":           "161🚀 HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(s, p)":                "161🚀 HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(string, p)":           "161🚀 HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix(string, prefix)":      "161🚀 HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hasprefix()":                    "161🚀 HasPrefix(s, prefix) 🔔 This function returns true if the string s begins with the string prefix.",
		"hassuffix(s, suffix)":           "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(s, suf)":              "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, suf)":         "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, sufix)":       "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix()":                    "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		"hassuffix(string, string)":      "161🚀 HasSuffix(s, suffix) 🔔 This function returns true if the string ends with the string suffix.",
		/*





		 */                              // 169.Splitting Strings
		"fields(s)":                     "169🚀 Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fields()":                      "169🚀 Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fields(string)":                "169🚀 Fields(s) 🔔 This function splits a string on whitespace characters and returns a slice containing the nonwhitespace sections of the string s.",
		"fieldsfunc(s, func)":           "169🚀 FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"fieldsfunc()":                  "169🚀 FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"fieldsfunc(string, function)":  "169🚀 FieldsFunc(s, func) 🔔 This function splits the string s on the characters for which a custom function returns true and returns a slice containing the remaining sections of the string.",
		"split(s, sub)":                 "169🚀 Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"split()":                       "169🚀 Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"split(string, sub)":            "169🚀 Split(s, sub) 🔔 This function splits the string s on every occurrence of the specified substring, returning a string slice. If the separator is the empty string, then the slice will contain strings for each character.",
		"splitn(s, sub, max)":           "169🚀 SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitn()":                      "169🚀 SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitn(string, sub, max)":      "169🚀 SplitN(s, sub, max) 🔔 This function is similar to Split, but accepts an additional int argument that specifies the maximum number of substrings to return. The last substring in the result slice will contain the unsplit portion of the source string.",
		"splitafter(s, sub)":            "169🚀 SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitafter()":                  "169🚀 SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitafter(string, sub)":       "169🚀 SplitAfter(s, sub) 🔔 This function is similar to Split but includes the substring used in the results.",
		"splitaftern(s, sub, max)":      "169🚀 SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"splitaftern()":                 "169🚀 SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		"splitaftern(string, sub, max)": "169🚀 SplitAfterN(s, sub, max) 🔔 This function is similar to SplitAfter, but accepts an additional int argument that specifies the maximum number of substrings to return.",
		/*





		 */                                  // 179.Altering Strings
		"replace(s, old, new, n)":           "179🚀 Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace()":                         "179🚀 Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(1, 2, 3, 4)":               "179🚀 Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(1,2,3,4)":                  "179🚀 Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replace(string, old, new, number)": "179🚀 Replace(s, old, new, n) 🔔 This function alters the string s by replacing occurrences of the string old with the string new. The maximum number of occurrences that will be replaced is specified by the int argument n.",
		"replaceall(s, old, new)":           "179🚀 ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"replaceall()":                      "179🚀 ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"replaceall(string, old, new)":      "179🚀 ReplaceAll(s, old, new) 🔔 This function alters the string s by replacing all occurrences of the string old with the string new. Unlike the Replace function, there is no limit on the number of occurrences that will be replaced.",
		"map(func, s)":                      "179🚀 Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"map()":                             "179🚀 Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		"map(function, string)":             "179🚀 Map(func, s) 🔔 This function generates a string by invoking the custom function for each character in the string s and concatenating the results. If the function produces a negative value, the current character is dropped without a replacement.",
		/*





		 */                       //183.The Replacer Methods
		"replace(s)":             "183🚀 Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"replace(string)":        "183🚀 Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"replace(str)":           "183🚀 Replace(s) 🔔 This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"writestring(writer, s)": "183🚀 WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, s)":      "183🚀 WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, str)":    "183🚀 WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"writestring(w, string)": "183🚀 WriteString(writer, s) 🔔 This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		/*






		 */                       //184.Building and Generating Strings
		"join(slice, sep)":       "184🚀 Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"join(slice, specified)": "184🚀 Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"join()":                 "184🚀 Join(slice, sep) 🔔 This function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"repeat(s, count)":       "184🚀 Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		"repeat(str, count)":     "184🚀 Repeat(s, count) 🔔 This function generates a string by repeating the string s for a specified number of times.",
		/*





		 */                    //186.Building Strings
		"writestring(s)":      "186🚀 WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writestring(string)": "186🚀 WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writestring()":       "186🚀 WriteString(s) 🔔 This method appends the string s to the string being built.",
		"writerune(r)":        "186🚀 WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"writerune()":         "186🚀 WriteRune(r) 🔔 This method appends the character r to the string being built.",
		"writebyte(b)":        "186🚀 WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"writebyte(byte)":     "186🚀 WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"writebyte()":         "186🚀 WriteByte(b) 🔔 This method appends the byte b to the string being built.",
		"string()":            "186🚀 String() 🔔 This method returns the string that has been created by the builder.",
		"reset()":             "186🚀 Reset() 🔔 This method resets the string created by the builder.",
		"len()":               "186🚀 Len() 🔔 This method returns the number of bytes used to store the string created by the builder.",
		"cap()":               "186🚀 Cap() 🔔 This method returns the number of bytes that have been allocated by the builder.",
		"grow(size)":          "186🚀 Grow(size) 🔔 This method increases the number of bytes used allocated by the builder to store the string that is being built.",
		"grow(siz)":           "186🚀 Grow(size) 🔔 This method increases the number of bytes used allocated by the builder to store the string that is being built.",
		"grow(sizes)":         "186🚀 Grow(size) 🔔 This method increases the number of bytes used allocated by the builder to store the string that is being built.",
		/*




		 */                                 //192.Useful Basic Regexp Methods
		"matchstring(s)":                   "192🚀 MatchString(s) 🔔 This method returns true if the string s matches the compiled pattern.",
		"matchstring(str)":                 "192🚀 MatchString(s) 🔔 This method returns true if the string s matches the compiled pattern.",
		"matchstring(string)":              "192🚀 MatchString(s) 🔔 This method returns true if the string s matches the compiled pattern.",
		"matchstring(strings)":             "192🚀 MatchString(s) 🔔 This method returns true if the string s matches the compiled pattern.",
		"findstringindex(s)":               "192🚀 FindStringIndex(s) 🔔 This method returns an int slice containing the location for the left most match made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findstringindex(str)":             "192🚀 FindStringIndex(s) 🔔 This method returns an int slice containing the location for the left most match made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findstringindex(string)":          "192🚀 FindStringIndex(s) 🔔 This method returns an int slice containing the location for the left most match made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findstringindex(strings)":         "192🚀 FindStringIndex(s) 🔔 This method returns an int slice containing the location for the left most match made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findallstringindex(s, max)":       "192🚀 FindAllStringIndex(s, max) 🔔 This method returns a slice of int slices that contain the location for all the matches made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findallstringindex(str, max)":     "192🚀 FindAllStringIndex(s, max) 🔔 This method returns a slice of int slices that contain the location for all the matches made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findallstringindex(string, max)":  "192🚀 FindAllStringIndex(s, max) 🔔 This method returns a slice of int slices that contain the location for all the matches made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findallstringindex(strings, max)": "192🚀 FindAllStringIndex(s, max) 🔔 This method returns a slice of int slices that contain the location for all the matches made by the compiled pattern in the string s. A nil result indicates that no matches were made.",
		"findstring(s)":                    "192🚀 FindString(s) 🔔 This method returns a string containing the left-most match made by the compiled pattern in the string s. An empty string will be returned if no match is made.",
		"findstring(str)":                  "192🚀 FindString(s) 🔔 This method returns a string containing the left-most match made by the compiled pattern in the string s. An empty string will be returned if no match is made.",
		"findstring(string)":               "192🚀 FindString(s) 🔔 This method returns a string containing the left-most match made by the compiled pattern in the string s. An empty string will be returned if no match is made.",
		"findstring(strings)":              "192🚀 FindString(s) 🔔 This method returns a string containing the left-most match made by the compiled pattern in the string s. An empty string will be returned if no match is made.",
		"findallstring(s, max)":            "192🚀 FindAllString(s, max) 🔔 This method returns a string slice containing the matches made by the compiled pattern in the string s. The int argument max specifies the maximum number of matches, with -1 specifying no limit. A nil result is returned if there are no matches.",
		"findallstring(str, max)":          "192🚀 FindAllString(s, max) 🔔 This method returns a string slice containing the matches made by the compiled pattern in the string s. The int argument max specifies the maximum number of matches, with -1 specifying no limit. A nil result is returned if there are no matches.",
		"findallstring(string, max)":       "192🚀 FindAllString(s, max) 🔔 This method returns a string slice containing the matches made by the compiled pattern in the string s. The int argument max specifies the maximum number of matches, with -1 specifying no limit. A nil result is returned if there are no matches.",
		"split(s, max)":                    "192🚀 Split(s, max) 🔔 This method splits the string s using matches from the compiled pattern as separators and returns a slice containing the split substrings.",
		"split(str, max)":                  "192🚀 Split(s, max) 🔔 This method splits the string s using matches from the compiled pattern as separators and returns a slice containing the split substrings.",
		"split(string, max)":               "192🚀 Split(s, max) 🔔 This method splits the string s using matches from the compiled pattern as separators and returns a slice containing the split substrings.",
		/*




		 */                                   // 197.The Regexp Methods for Subexpressions
		"findstringsubmatch(s)":              "197🚀 FindStringSubmatch(s) 🔔 This method returns a slice containing the first match made by the pattern and the text for the subexpressions that the pattern defines.",
		"findstringsubmatch(str)":            "197🚀 FindStringSubmatch(s) 🔔 This method returns a slice containing the first match made by the pattern and the text for the subexpressions that the pattern defines.",
		"findstringsubmatch(string)":         "197🚀 FindStringSubmatch(s) 🔔 This method returns a slice containing the first match made by the pattern and the text for the subexpressions that the pattern defines.",
		"findallstringsubmatch(s, max)":      "197🚀 FindAllStringSubmatch(s, max) 🔔 This method returns a slice containing all the matches and the text for the subexpressions. The int argument is used to specify the maximum number of matches. A value of -1 specifies all matches.",
		"findallstringsubmatch(str, max)":    "197🚀 FindAllStringSubmatch(s, max) 🔔 This method returns a slice containing all the matches and the text for the subexpressions. The int argument is used to specify the maximum number of matches. A value of -1 specifies all matches.",
		"findallstringsubmatch(string, max)": "197🚀 FindAllStringSubmatch(s, max) 🔔 This method returns a slice containing all the matches and the text for the subexpressions. The int argument is used to specify the maximum number of matches. A value of -1 specifies all matches.",
		"findstringsubmatchindex(s)":         "197🚀 FindStringSubmatchIndex(s) 🔔 This method is equivalent to FindStringSubmatch but returns indices rather than substrings. FindAllStringSubmatchIndex",
		"findstringsubmatchindex(str)":       "197🚀 FindStringSubmatchIndex(s) 🔔 This method is equivalent to FindStringSubmatch but returns indices rather than substrings. FindAllStringSubmatchIndex",
		"findstringsubmatchindex(string)":    "197🚀 FindStringSubmatchIndex(s) 🔔 This method is equivalent to FindStringSubmatch but returns indices rather than substrings. FindAllStringSubmatchIndex",
		"numsubexp()":                        "197🚀 NumSubexp() 🔔 This method returns the number of subexpressions.",
		"subexpindex(name)":                  "197🚀 SubexpIndex(name) 🔔 This method returns the index of the subexpression with the specified name or -1 if there is no such subexpression.",
		"subexpindex(names)":                 "197🚀 SubexpIndex(name) 🔔 This method returns the index of the subexpression with the specified name or -1 if there is no such subexpression.",
		"subexpindex(n)":                     "197🚀 SubexpIndex(name) 🔔 This method returns the index of the subexpression with the specified name or -1 if there is no such subexpression.",
		"subexpnames()":                      "197🚀 SubexpNames() 🔔 This method returns the names of the subexpressions, expressed in the order in which they are defined.",
		/*




		 */                                           //199.Replacing Substrings Using a Regular Expression
		"replaceallstring(s, template)":              "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallstring(str, temp)":                "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallstring(string, temp)":             "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallstring(strings, temp)":            "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallstring(strings, template)":        "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallstring(string, template)":         "199🚀 ReplaceAllString(s, template) 🔔 This method replaces the matched portion of the string s with the specified template, which is expanded before it is included in the result to incorporate subexpressions.",
		"replaceallliteralstring(s, sub)":            "199🚀 ReplaceAllLiteralString(s, sub) 🔔 This method replaces the matched portion of the string s with the specified content, which is included in the result without being expanded for subexpressions.",
		"replaceallliteralstring(str, sub)":          "199🚀 ReplaceAllLiteralString(s, sub) 🔔 This method replaces the matched portion of the string s with the specified content, which is included in the result without being expanded for subexpressions.",
		"replaceallliteralstring(string, sub)":       "199🚀 ReplaceAllLiteralString(s, sub) 🔔 This method replaces the matched portion of the string s with the specified content, which is included in the result without being expanded for subexpressions.",
		"replaceallliteralstring(string, specified)": "199🚀 ReplaceAllLiteralString(s, sub) 🔔 This method replaces the matched portion of the string s with the specified content, which is included in the result without being expanded for subexpressions.",
		"replaceallstringfunc(s, func)":              "199🚀 ReplaceAllStringFunc(s, func) 🔔 This method replaces the matched portion of the string s with the result produced by the specified function.",
		"replaceallstringfunc(str, func)":            "199🚀 ReplaceAllStringFunc(s, func) 🔔 This method replaces the matched portion of the string s with the result produced by the specified function.",
		"replaceallstringfunc(string, function)":     "199🚀 ReplaceAllStringFunc(s, func) 🔔 This method replaces the matched portion of the string s with the result produced by the specified function.",
		/*





		 */
		"formatfloat()":                             "FormatFloat(val, format, precision, size) 🔔 This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"formatfloat(1,2,3,4)":                      "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		"formatfloat(val, format, precision, size)": "FormatFloat(val, format, precision, size) 🔔This function returns a string representation of the specified float64 value, expressed using the specified format, precision, and size.",
		/*





		 */             //The Key Reflection Functions
		"typeof(val)":  "TypeOf(val) 🔔 This function returns a value that implements the Type interface, which describes the type of the specified value. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
		"valueof(val)": "ValueOf(val) 🔔 This function returns a Value struct, which allows the specified value to be inspected and manipulated. There is a lot of detail behind the TypeOf and ValueOf functions and their results, and it is easy to lose sight of why reflection can be useful.",
		/*




		 */                          // 203.fmt package
		"print(...vals)":            "203🚀 Print(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out. Spaces are added between values that are not strings.",
		"print()":                   "203🚀 Print(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out. Spaces are added between values that are not strings.",
		"print(...)":                "203🚀 Print(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out. Spaces are added between values that are not strings.",
		"print(...values)":          "203🚀 Print(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out. Spaces are added between values that are not strings.",
		"println(...vals)":          "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"println(vals)":             "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"println()":                 "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"println(values)":           "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"println(...value)":         "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"println(...)":              "203🚀 Println(...vals) 🔔 This function accepts a variable number of arguments and writes out their values to the standard out, separated by spaces and followed by a newline character.",
		"fprint(writer, ...vals)":   "203🚀 Fprint(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer, Spaces are added between values that are not strings.",
		"fprint(w, ...vals)":        "203🚀 Fprint(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer, Spaces are added between values that are not strings.",
		"fprint(w, vals)":           "203🚀 Fprint(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer, Spaces are added between values that are not strings.",
		"fprint()":                  "203🚀 Fprint(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer, Spaces are added between values that are not strings.",
		"fprint(wr, vals)":          "203🚀 Fprint(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer, Spaces are added between values that are not strings.",
		"fprintln(writer, ...vals)": "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(wri, ...vals)":    "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(w, ...vals)":      "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(w, vals)":         "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(w, values)":       "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(w, v)":            "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln()":                "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		"fprintln(wrt, vls)":        "203🚀 Fprintln(writer, ...vals) 🔔 This function writes out a variable number of arguments to the specified writer followed by a newline character. Spaces are added between all values.",
		/*





		 */                            //206.The fmt Functions for Formatting Strings
		"sprintf(t, ...vals)":         "206🚀 Sprintf(t, ...vals) 🔔 This function returns a string, which is created by processing the template t.",
		"sprintf(t, vals)":            "206🚀 Sprintf(t, ...vals) 🔔 This function returns a string, which is created by processing the template t.",
		"sprintf()":                   "206🚀 Sprintf(t, ...vals) 🔔 This function returns a string, which is created by processing the template t.",
		"sprintf(1,...2)":             "206🚀 Sprintf(t, ...vals) 🔔 This function returns a string, which is created by processing the template t.",
		"printf(t, ...vals)":          "206🚀 Printf(t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to the standard out.",
		"printf(t, ...)":              "206🚀 Printf(t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to the standard out.",
		"printf()":                    "206🚀 Printf(t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to the standard out.",
		"printf(1,2)":                 "206🚀 Printf(t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to the standard out.",
		"printf(1,...2)":              "206🚀 Printf(t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to the standard out.",
		"fprintf(writer, t, ...vals)": "206🚀 Fprintf(writer, t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to a Writer, which is described in Chapter 20.",
		"fprintf(write, t, ...vals)":  "206🚀 Fprintf(writer, t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to a Writer, which is described in Chapter 20.",
		"fprintf()":                   "206🚀 Fprintf(writer, t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to a Writer, which is described in Chapter 20.",
		"fprintf(1,2,3)":              "206🚀 Fprintf(writer, t, ...vals) 🔔 This function creates a string by processing the template t. The remaining arguments are used as values for the template verbs. The string is written to a Writer, which is described in Chapter 20.",
		"errorf(t, ...values)":        "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		"errorf(t, ...val)":           "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		"errorf(t, ...vals)":          "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		"errorf()":                    "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		"errorf(1,2)":                 "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		"errorf(1,...2)":              "206🚀 Errorf(t, ...values) 🔔 This function creates an error by processing the template t. The remaining arguments are used as values for the template verbs. The result is an error value whose Error method returns the formatted string.",
		/*




		 */                                    // 220.The fmt Functions for Scanning Strings
		"scan(...vals)":                       "220🚀 Scan(...vals) 🔔 This function reads text from the standard in and stores the space- separated values into specified arguments. Newlines are treated as spaces, and the function reads until it has received values for all of its arguments. The result is the number of values that have been read and an error that describes any problems.",
		"scan(...)":                           "220🚀 Scan(...vals) 🔔 This function reads text from the standard in and stores the space- separated values into specified arguments. Newlines are treated as spaces, and the function reads until it has received values for all of its arguments. The result is the number of values that have been read and an error that describes any problems.",
		"scan()":                              "220🚀 Scan(...vals) 🔔 This function reads text from the standard in and stores the space- separated values into specified arguments. Newlines are treated as spaces, and the function reads until it has received values for all of its arguments. The result is the number of values that have been read and an error that describes any problems.",
		"scan(...2)":                          "220🚀 Scan(...vals) 🔔 This function reads text from the standard in and stores the space- separated values into specified arguments. Newlines are treated as spaces, and the function reads until it has received values for all of its arguments. The result is the number of values that have been read and an error that describes any problems.",
		"scanln(...vals)":                     "220🚀 Scanln(...vals) 🔔 This function works in the same way as Scan but stops reading when it encounters a newline character.",
		"scanln(...)":                         "220🚀 Scanln(...vals) 🔔 This function works in the same way as Scan but stops reading when it encounters a newline character.",
		"scanln()":                            "220🚀 Scanln(...vals) 🔔 This function works in the same way as Scan but stops reading when it encounters a newline character.",
		"scanln(...2)":                        "220🚀 Scanln(...vals) 🔔 This function works in the same way as Scan but stops reading when it encounters a newline character.",
		"scanf(template, ...vals)":            "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"scanf(temp, ...vals)":                "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"scanf(temp, ...values)":              "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"scanf(1,2)":                          "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"scanf(1,...2)":                       "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"scanf()":                             "220🚀 Scanf(template, ...vals) 🔔 This function works in the same way as Scan but uses a template string to select the values from the input it receives.",
		"fscan(reader, ...vals)":              "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscan(reade, ...vals)":               "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscan(reader, ...values)":            "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscan(reader, values)":               "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscan(1,2)":                          "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscan()":                             "220🚀 Fscan(reader, ...vals) 🔔 This function reads space-separated values from the specified reader, which is described in Chapter 20. Newlines are treated as spaces, and the function returns the number of values that have been read and an error that describes any problems.",
		"fscanln(reader, ...vals)":            "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanln(reader, ...values)":          "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanln(reader, values)":             "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanln(read, values)":               "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanln(1,2)":                        "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanln()":                           "220🚀 Fscanln(reader, ...vals) 🔔 This function works in the same way as Fscan but stops reading when it encounters a newline character.",
		"fscanf(reader, template, ...vals)":   "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(reader, template, ...values)": "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(reader, template, values)":    "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(reader, temp, values)":        "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(read, temp, values)":          "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(read, temp, val)":             "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(read, temp, vals)":            "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(1,2,3)":                       "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf(1, 2, 3)":                     "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"fscanf()":                            "220🚀 Fscanf(reader, template, ...vals) 🔔 This function works in the same way as Fscan but uses a template to select the values from the input it receives.",
		"sscan(str, ...vals)":                 "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(str, vals)":                    "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(str, values)":                  "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(string, values)":               "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(string, ...values)":            "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan()":                             "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(1,2)":                          "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscan(1, 2)":                         "220🚀 Sscan(str, ...vals) 🔔 This function scans the specified string for space-separated values, which are assigned to the remaining arguments. The result is the number of values scanned and an error that describes any problems.",
		"sscanf(str, template, ...vals)":      "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(string, template, ...vals)":   "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(string, template, vals)":      "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(string, template, values)":    "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(str, template, values)":       "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(str, temp, values)":           "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(str, temp, vals)":             "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(1,2,3)":                       "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf(1, 2, 3)":                     "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanf()":                            "220🚀 Sscanf(str, template, ...vals) 🔔 This function works in the same way as Sscan but uses a template to select values from the string. Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the",
		"sscanln(str, template, ...vals)":     "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(string, template, ...vals)":  "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(string, template, vals)":     "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(string, template, values)":   "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(str, temp, vals)":            "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(str, template, vals)":        "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(str, template, values)":      "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(1,2,3)":                      "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln(1, 2, 3)":                    "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		"sscanln()":                           "220🚀 Sscanln(str, template, ...vals) 🔔 This function works in the same way as Sscanf but stops scanning the string as soon as a newline character is encountered.",
		/*




		 */                 //226.Useful Functions from the math Package
		"abs(val)":         "226🚀 Abs(val) 🔔 This function returns the absolute value of a float64 value, meaning the distance from zero without considering direction.",
		"abs()":            "226🚀 Abs(val) 🔔 This function returns the absolute value of a float64 value, meaning the distance from zero without considering direction.",
		"ceil(val)":        "226🚀 Ceil(val) 🔔 This function returns the smallest integer that is equal to or greater than the specified float64 value. The result is also a float64 value, even though it represents an integer number.",
		"ceil()":           "226🚀 Ceil(val) 🔔 This function returns the smallest integer that is equal to or greater than the specified float64 value. The result is also a float64 value, even though it represents an integer number.",
		"copysign(x, y)":   "226🚀 Copysign(x, y) 🔔 This function returns a float64 value, which is the absolute value of x with the sign of y.",
		"copysign(1,2)":    "226🚀 Copysign(x, y) 🔔 This function returns a float64 value, which is the absolute value of x with the sign of y.",
		"copysign(1, 2)":   "226🚀 Copysign(x, y) 🔔 This function returns a float64 value, which is the absolute value of x with the sign of y.",
		"copysign()":       "226🚀 Copysign(x, y) 🔔 This function returns a float64 value, which is the absolute value of x with the sign of y.",
		"floor(val)":       "226🚀 Floor(val) 🔔 This function returns the largest integer that is smaller or equal to the specified float64 value. The result is also a float64 value, even though it represents an integer number.",
		"floor()":          "226🚀 Floor(val) 🔔 This function returns the largest integer that is smaller or equal to the specified float64 value. The result is also a float64 value, even though it represents an integer number.",
		"max(x, y)":        "226🚀 Max(x, y) 🔔 This function returns whichever of the specified float64 value is the largest.",
		"max(x,y)":         "226🚀 Max(x, y) 🔔 This function returns whichever of the specified float64 value is the largest.",
		"max(1,2)":         "226🚀 Max(x, y) 🔔 This function returns whichever of the specified float64 value is the largest.",
		"max(1, 2)":        "226🚀 Max(x, y) 🔔 This function returns whichever of the specified float64 value is the largest.",
		"max()":            "226🚀 Max(x, y) 🔔 This function returns whichever of the specified float64 value is the largest.",
		"min(x, y)":        "226🚀 Min(x, y) 🔔 This function returns whichever of the specified float64 value is smallest.",
		"min(x,y)":         "226🚀 Min(x, y) 🔔 This function returns whichever of the specified float64 value is smallest.",
		"min(1,2)":         "226🚀 Min(x, y) 🔔 This function returns whichever of the specified float64 value is smallest.",
		"min(1, 2)":        "226🚀 Min(x, y) 🔔 This function returns whichever of the specified float64 value is smallest.",
		"min()":            "226🚀 Min(x, y) 🔔 This function returns whichever of the specified float64 value is smallest.",
		"mod(x, y)":        "226🚀 Mod(x, y) 🔔 This function returns the remainder of x/y.",
		"mod(x,y)":         "226🚀 Mod(x, y) 🔔 This function returns the remainder of x/y.",
		"mod(1,2)":         "226🚀 Mod(x, y) 🔔 This function returns the remainder of x/y.",
		"mod(1, 2)":        "226🚀 Mod(x, y) 🔔 This function returns the remainder of x/y.",
		"mod()":            "226🚀 Mod(x, y) 🔔 This function returns the remainder of x/y.",
		"pow(x, y)":        "226🚀 Pow(x, y) 🔔 This function returns x raised to the exponent y.",
		"pow(x,y)":         "226🚀 Pow(x, y) 🔔 This function returns x raised to the exponent y.",
		"pow()":            "226🚀 Pow(x, y) 🔔 This function returns x raised to the exponent y.",
		"pow(1,2)":         "226🚀 Pow(x, y) 🔔 This function returns x raised to the exponent y.",
		"pow(1, 2)":        "226🚀 Pow(x, y) 🔔 This function returns x raised to the exponent y.",
		"round(val)":       "226🚀 Round(val) 🔔 This function rounds the specified value to the nearest integer, rounding half values up. The result is a float64 value, even though it represents an integer.",
		"round()":          "226🚀 Round(val) 🔔 This function rounds the specified value to the nearest integer, rounding half values up. The result is a float64 value, even though it represents an integer.",
		"roundtoeven(val)": "226🚀 RoundToEven(val) 🔔 This function rounds the specified value to the nearest integer, rounding half values to the nearest even number. The result is a float64 value, even though it represents an integer.",
		"roundtoeven()":    "226🚀 RoundToEven(val) 🔔 This function rounds the specified value to the nearest integer, rounding half values to the nearest even number. The result is a float64 value, even though it represents an integer.",
		/*




		 */                     //228.Generating Random Numbers
		"seed(s)":              "228 🚀 Seed(s) 🔔 This function sets the seed value using the specified int64 value.",
		"seed()":               "228 🚀 Seed(s) 🔔 This function sets the seed value using the specified int64 value.",
		"float32()":            "228 🚀 Float32() 🔔 This function generates a random float32 value between 0 and 1.",
		"float64()":            "228 🚀 Float64() 🔔 This function generates a random float64 value between 0 and 1.",
		"int()":                "228 🚀 Int() 🔔 This function generates a random int value.",
		"intn(max)":            "228 🚀 Intn(max) 🔔 This function generates a random int smaller than a specified value, as described after the table.",
		"intn()":               "228 🚀 Intn(max) 🔔 This function generates a random int smaller than a specified value, as described after the table.",
		"uint32()":             "228 🚀 UInt32() 🔔 This function generates a random uint32 value.",
		"uint64()":             "228 🚀 UInt64() 🔔 This function generates a random uint64 value.",
		"shuffle(count, func)": "228 🚀 Shuffle(count, func) 🔔 This function is used to randomize the order of elements, as described after the table.",
		/*




		 */                         //234.The Basic Functions for Sorting
		"float64s(slice)":          "234🚀 Float64s(slice) 🔔 This function sorts a slice of float64 values. The elements are sorted in place.",
		"float64s()":               "234🚀 Float64s(slice) 🔔 This function sorts a slice of float64 values. The elements are sorted in place.",
		"float64saresorted(slice)": "234🚀 Float64sAreSorted(slice) 🔔 This function returns true if the elements in the specified float64 slice are in order.",
		"float64saresorted()":      "234🚀 Float64sAreSorted(slice) 🔔 This function returns true if the elements in the specified float64 slice are in order.",
		"ints(slice)":              "234🚀 Ints(slice) 🔔 This function sorts a slice of int values. The elements are sorted in place.",
		"ints()":                   "234🚀 Ints(slice) 🔔 This function sorts a slice of int values. The elements are sorted in place.",
		"intsaresorted(slice)":     "234🚀 IntsAreSorted(slice) 🔔 This function returns true if the elements in the specified int slice are in order.",
		"intsaresorted()":          "234🚀 IntsAreSorted(slice) 🔔 This function returns true if the elements in the specified int slice are in order.",
		"strings(slice)":           "234🚀 Strings(slice) 🔔 This function sorts a slice of string values. The elements are sorted in place.",
		"strings()":                "234🚀 Strings(slice) 🔔 This function sorts a slice of string values. The elements are sorted in place.",
		"stringsaresorted(slice)":  "234🚀 StringsAreSorted(slice) 🔔 This function returns true if the elements in the specified string slice are in order.",
		"stringsaresorted()":       "234🚀 StringsAreSorted(slice) 🔔 This function returns true if the elements in the specified string slice are in order.",
		/*




		 */          //481.Basic Methods Defined by the Type Interface
		"name()":    "481🚀 Name() 🔔 This method returns the name of the type.",
		"pkgpath()": "481🚀 PkgPath() 🔔 This method returns the package path for the type. The empty string is returned for built-in types, such as int and bool.",
		// 🔔
	},
}
