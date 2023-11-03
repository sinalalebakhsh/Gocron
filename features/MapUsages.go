package features

type AllMapUsages struct {
	SingleDef map[string]string
}

var MapUsage = AllMapUsages{
	SingleDef: map[string]string{
		"godoc":`Usage:

godoc [flag]

The flags are:

-v
	verbose mode
-timestamps=true
	show timestamps with directory listings
-index
	enable identifier and full text search index
	(no search box is shown if -index is not set)
-index_files=""
	glob pattern specifying index files; if not empty,
	the index is read from these files in sorted order
-index_throttle=0.75
	index throttle value; a value of 0 means no time is allocated
	to the indexer (the indexer will never finish), a value of 1.0
	means that index creation is running at full throttle (other
	goroutines may get no time while the index is built)
-index_interval=0
	interval of indexing; a value of 0 sets it to 5 minutes, a
	negative value indexes only once at startup
-play=false
	enable playground
-links=true
	link identifiers to their declarations
-write_index=false
	write index to a file; the file name must be specified with
	-index_files
-maxresults=10000
	maximum number of full text search results shown
	(no full text index is built if maxresults <= 0)
-notes="BUG"
	regular expression matching note markers to show
	(e.g., "BUG|TODO", ".*")
-goroot=$GOROOT
	Go root directory
-http=addr
	HTTP service address (e.g., '127.0.0.1:6060' or just ':6060')
-templates=""
	directory containing alternate template files; if set,
	the directory may provide alternative template files
	for the files in $GOROOT/lib/godoc
-url=path
	print to standard output the data that would be served by
	an HTTP request for path
-zip=""
	zip file providing the file system to serve; disabled if empty
`,
	},
}