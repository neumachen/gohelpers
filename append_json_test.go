package paratils

import (
	"fmt"
	"reflect"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAppendJSON(t *testing.T) {
	tests := []struct {
		originalBytes []byte
		bytesToAdd    []byte
		appendedBytes []byte
	}{
		{
			originalBytes: []byte(`{}`),
			bytesToAdd:    []byte(`{}`),
			appendedBytes: nil,
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{}`),
			appendedBytes: []byte(`{"test":"originalBytes"}`),
		},
		{
			originalBytes: []byte(`{}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"bytesToAdd"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"originalBytes","test":"bytesToAdd"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd", "nested": {"moo":"moo"}}`),
			appendedBytes: []byte(`{"test":"originalBytes","test":"bytesToAdd", "nested": {"moo":"moo"}}`),
		},
	}

	for i := range tests {
		actual := AppendJSON(tests[i].originalBytes, tests[i].bytesToAdd)
		assert.Assert(
			t,
			reflect.DeepEqual(tests[i].appendedBytes, actual),
			fmt.Sprintf(
				"not requal a: %s != b: %s\n",
				string(tests[i].appendedBytes),
				string(actual),
			),
		)
	}

}
