package scrubber

import (
	"fmt"

	"github.com/taylorgtyler/go-datascrubber/csv"
	"github.com/taylorgtyler/go-datascrubber/json"
	"github.com/taylorgtyler/go-datascrubber/text"
)

// FileType represents the type of file to be scrubbed
type FileType int

const (
    TextFile FileType = iota
    CSVFile
    JSONFile
)

// NewScrubber creates a new scrubber based on the file type
func NewScrubber(fileType FileType) (Scrubber, error) {
    switch fileType {
    case TextFile:
        return &text.TextScrubber{}, nil
    case CSVFile:
        return &csv.CSVScrubber{}, nil
    case JSONFile:
        return &json.JSONScrubber{}, nil
    default:
        return nil, fmt.Errorf("unsupported file type: %v", fileType)
    }
}