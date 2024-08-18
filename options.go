package scrubber

// Options holds the configurable parameters for the scrubbing process
type Options struct {
    BufferSize          int
    NumWorkers          int
    MaxConcurrentWrites int64
    CSVDelimiter        rune
    JSONPrettyPrint     bool
}

// DefaultOptions returns an Options struct with default values
func DefaultOptions() Options {
    return Options{
        BufferSize:          4096,
        NumWorkers:          4,
        MaxConcurrentWrites: 1,
        CSVDelimiter:        ',',
        JSONPrettyPrint:     false,
    }
}