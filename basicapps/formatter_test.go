package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorFormat(t *testing.T) {

	cases := map[string]struct {
		line, keword, expected string
	}{
		"empty line input": {
			line:     "",
			keword:   "keyword",
			expected: "",
		},
		"colored": {
			line:     "A mote it is to trouble the mind's eye.",
			keword:   "trouble",
			expected: "A mote it is to \x1b[31mtrouble\x1b[0m the mind's eye.\n",
		},
		"not found": {
			line:     "A mote it is to trouble the mind's eye.",
			keword:   "something else",
			expected: "A mote it is to trouble the mind's eye.\n",
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, colorFormat(testCase.keword, testCase.line))
		})
	}
}
