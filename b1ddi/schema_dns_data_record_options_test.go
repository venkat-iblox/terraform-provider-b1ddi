package b1ddi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DeepUpdateDataRecordOptions(t *testing.T) {
	rq := require.New(t)
	testData := map[string]struct {
		input       map[string]interface{}
		output      map[string]interface{}
		recordType  string
		errExpected bool
	}{
		"Invalid Options - Bad create_ptr value": {
			input: map[string]interface{}{
				"create_ptr": "asda",
			},
			output:      nil,
			recordType:  "A",
			errExpected: true,
		},
		"Invalid Options - Bad check_rmz value": {
			input: map[string]interface{}{
				"create_ptr": "true",
				"check_rmz":  "lsdclksj",
			},
			output:      nil,
			recordType:  "A",
			errExpected: true,
		},
		"Valid Options": {
			input: map[string]interface{}{
				"create_ptr": "true",
			},
			output: map[string]interface{}{
				"create_ptr": true,
			},
			recordType:  "A",
			errExpected: false,
		},
		"Valid Options - 1": {
			input: map[string]interface{}{
				"create_ptr": "true",
				"check_rmz":  "true",
			},
			output: map[string]interface{}{
				"create_ptr": true,
				"check_rmz":  true,
			},
			recordType:  "A",
			errExpected: false,
		},
		"Valid Options - 2": {
			input: map[string]interface{}{
				"address": "10.0.0.1",
			},
			output: map[string]interface{}{
				"address": "10.0.0.1",
			},
			recordType:  "PTR",
			errExpected: false,
		},
	}

	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			options, err := deepUpdateDataRecordOptions(tc.input, tc.recordType)
			if tc.errExpected {
				rq.Error(err)
				rq.Nil(options)
			} else {
				rq.NoError(err)
				rq.Equal(tc.output, options)
			}
		})
	}
}
