package text

import (
	"io"
	"regexp"

	scrubber "github.com/taylorgtyler/go-datascrubber"
)

type TextScrubber struct{}

func (s *TextScrubber) Scrub(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string) error {
    return s.ScrubWithOptions(input, output, patterns, replacement, scrubber.DefaultOptions())
}

func (s *TextScrubber) ScrubWithOptions(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string, opts scrubber.Options) error {
    // Implement text scrubbing logic here
    // Use the options to configure the scrubbing process
    return nil
}