package semaphore

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(concurrency int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, concurrency)}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}
