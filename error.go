package packet

type UnsynchronizedError struct {}

func (e UnsynchronizedError) Error() string {
	return "SODIBus Frame Unsynchronized"
}

