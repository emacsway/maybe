package maybe

// AoS implements the Maybe monad for a slice of strings.
type AoS struct {
	just []string
	err  error
}

// NewAoS constructs a "Just" AoS from a given slice of strings.
func NewAoS(s []string) AoS {
	return AoS{just: s}
}

// ErrAoS constructs a "Nothing" AoS from a given error.
func ErrAoS(e error) AoS {
	return AoS{err: e}
}

// Bind applies a function that takes a slice of strings and returns an AoS.
func (m AoS) Bind(f func(s []string) AoS) AoS {
	if m.err != nil {
		return m
	}

	return f(m.just)
}

// BindS applies a function that takes a slice of strings and returns an S.
func (m AoS) BindS(f func(s []string) S) S {
	if m.err != nil {
		return ErrS(m.err)
	}

	return f(m.just)
}

// Unbox returns the underlying slice of strings value or error.
func (m AoS) Unbox() ([]string, error) {
	return m.just, m.err
}