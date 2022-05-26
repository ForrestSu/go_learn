package main

type Sinker struct {
	data string
}

// Write implements interface io.Writer
func (s *Sinker) Write(p []byte) (n int, err error) {
	s.data += string(p)
	return len(p), nil
}

// String implements interface fmt.Stringer
func (s *Sinker) String() string {
	return s.data
}
