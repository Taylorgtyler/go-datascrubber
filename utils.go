package scrubber

import (
	"regexp"
)

// CompilePatterns compiles a slice of string patterns into regular expressions
func CompilePatterns(patternStrs []string) ([]*regexp.Regexp, error) {
    patterns := make([]*regexp.Regexp, len(patternStrs))
    for i, str := range patternStrs {
        re, err := regexp.Compile(str)
        if err != nil {
            return nil, err
        }
        patterns[i] = re
    }
    return patterns, nil
}