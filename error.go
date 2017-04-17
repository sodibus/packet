package packet

// UnsynchronizedError ...
// UnsynchronizedError represents TCP stream does not have a 0xAA leading
type UnsynchronizedError struct{}

func (e UnsynchronizedError) Error() string {
	return "Frame Unsynchronized"
}
