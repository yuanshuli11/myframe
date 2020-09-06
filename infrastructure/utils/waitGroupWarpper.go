package utils

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

func (w *WaitGroupWrapper) MultiWrap(i int, cb func()) {
	for i > 0 {
		w.Wrap(cb)
		i--
	}
}
