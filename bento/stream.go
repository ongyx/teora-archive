package bento

// Stream is a stream of text over a channel.
type Stream struct {
	ch chan string
}

// Read returns the next text from the stream.
// If the stream is closed, nil is returned.
func (s *Stream) Read() *string {
	if v, ok := <-s.ch; ok {
		return &v
	} else {
		return nil
	}
}

// Source runs the function asynchronusly on the stream.
func (s *Stream) Source(fn func(chan<- string)) {
	s.ch = make(chan string)

	go func() {
		fn(s.ch)
		close(s.ch)
	}()
}
