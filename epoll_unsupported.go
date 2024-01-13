//go:build windows
// +build windows

package epoller

// NewPoller creates a new epoll poller.
func NewPoller(connBufferSize int) (*Epoll, error) {
	return nil, errors.New("epoll is not supported on windows")
}
