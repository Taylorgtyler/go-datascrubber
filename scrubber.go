package scrubber

import (
	"io"
	"regexp"
)

// Scrubber interface defines the methods that all scrubbers must implement
type Scrubber interface {
    Scrub(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string) error
    ScrubWithOptions(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string, opts Options) error
}

// ScrubFile is a convenience function that scrubs a file using default options
func ScrubFile(fileType FileType, input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string) error {
    scrubber, err := NewScrubber(fileType)
    if err != nil {
        return err
    }
    return scrubber.Scrub(input, output, patterns, replacement)
}

// ScrubFileWithOptions scrubs a file with custom options
func ScrubFileWithOptions(fileType FileType, input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string, opts Options) error {
    scrubber, err := NewScrubber(fileType)
    if err != nil {
        return err
    }
    return scrubber.ScrubWithOptions(input, output, patterns, replacement, opts)
}