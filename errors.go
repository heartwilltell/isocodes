package isocodes

const (
	// ErrMarshalJSON - indicates an error in the process
	// of marshalling code to json.
	ErrMarshalJSON Error = "failed to marshal json"

	// ErrUnmarshalJSON - indicates an error in the process
	// of unmarshalling code to json.
	ErrUnmarshalJSON Error = "failed to unmarshal json"

	// ErrInvalidStringCode - indicates an error in the process
	// of converting string representation to code type.
	ErrInvalidStringCode Error = "invalid string representation of the code"
)

// Error represents package level errors.
type Error string

func (e Error) Error() string { return string(e) }
