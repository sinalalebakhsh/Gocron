package features

type ReadingandWriting struct {
	allReadingandWriting string
}


var OriginalReadingandWriting ReadingandWriting = ReadingandWriting{
	allReadingandWriting: `
273.Reading and Writing Data
These interfaces are used wherever data is read or written, which means that any
source or destination for data can be treated in much the same way so that writing data to a file, for example,
is just the same as writing data to a network connection.

What are they?
These interfaces define the basic methods required to read and write data.

Why are they useful?
This approach means that just about any data source can be used
in the same way, while still allowing specialized features to be
defined using the composition features.

`,
}