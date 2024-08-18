package json

import (
	"io"
	"regexp"

	"github.com/taylorgtyler/scrubber"
)

type JSONScrubber struct{}

func (s *JSONScrubber) Scrub(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string) error {
    return s.ScrubWithOptions(input, output, patterns, replacement, scrubber.DefaultOptions())
}

func (s *JSONScrubber) ScrubWithOptions(input io.Reader, output io.Writer, patterns []*regexp.Regexp, replacement string, opts scrubber.Options) error {
    // Implement JSON scrubbing logic here
    // Use the options to configure the scrubbing process, including opts.JSONPrettyPrint
    return nil
}