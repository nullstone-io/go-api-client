package find

import "sync"

type onceError struct {
	sync.Once
}

func (oe *onceError) Do(fn func() error) error {
	var err error
	oe.Once.Do(func() {
		err = fn()
	})
	return err
}
