package isocodes

import (
	"testing"
)

func TestError_Error(t *testing.T) {
	type tcase struct {
		err  Error
		want string
	}

	tests := map[string]tcase{
		"ErrMarshalJSON":       {err: ErrMarshalJSON, want: "failed to marshal json"},
		"ErrUnmarshalJSON":     {err: ErrUnmarshalJSON, want: "failed to unmarshal json"},
		"ErrInvalidStringCode": {err: ErrInvalidStringCode, want: "invalid string representation of the code"},
		"Custom":               {err: Error("test error"), want: "test error"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tc.err.Error(); got != tc.want {
				t.Errorf("Error() = %v, want %v", got, tc.want)
			}
		})
	}
}
