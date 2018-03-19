package verrors // New returns an error that formats as the given text.

func New(text string) error {
	return &VError{text}
}

// errorString is a trivial implementation of error.
type VError struct {
	s string
}

func (e *VError) Error() string {
	return e.s
}
