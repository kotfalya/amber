package utils

type Semaphore struct {
	buf chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		make(chan struct{}, size),
	}
}

func (s *Semaphore) Acquire() {
	s.buf <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.buf
}
