package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	testData := map[string]struct {
		input        map[string]interface{}
		output       bool
		outputExists bool
		key          string
		errExpected  bool
	}{"error parsing": {
		input: map[string]interface{}{
			"key1": "qwert",
		},
		key:          "key1",
		output:       false,
		outputExists: true,
		errExpected:  true,
	},
		"valid parsing": {
			input: map[string]interface{}{
				"key1": "true",
			},
			key:          "key1",
			output:       true,
			outputExists: true,
			errExpected:  false,
		},
		"key doesn't exist": {
			input: map[string]interface{}{
				"key1": "qw1ert",
			},
			key:          "key2",
			output:       false,
			outputExists: false,
			errExpected:  false,
		},
	}
	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			out, exists, err := ToBool(tc.input, tc.key)
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.outputExists, exists)
			assert.Equal(t, tc.output, out)
		})
	}
}

func TestToInt(t *testing.T) {
	testData := map[string]struct {
		input        map[string]interface{}
		output       int
		outputExists bool
		key          string
		errExpected  bool
	}{
		"error parsing": {
			input: map[string]interface{}{
				"key1": "qwert",
			},
			key:          "key1",
			output:       0,
			outputExists: true,
			errExpected:  true,
		},
		"valid parsing": {
			input: map[string]interface{}{
				"key1": "1",
			},
			key:          "key1",
			output:       1,
			outputExists: true,
			errExpected:  false,
		},
		"key doesn't exist": {
			input: map[string]interface{}{
				"key1": "qw1ert",
			},
			key:          "key2",
			output:       0,
			outputExists: false,
			errExpected:  false,
		},
	}

	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			out, exists, err := ToInt(tc.input, tc.key)
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.outputExists, exists)
			assert.Equal(t, tc.output, out)
		})
	}
}
