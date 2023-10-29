package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(works <-chan Task, errChan chan<- struct{}) {
	for work := range works {
		if err := work(); err != nil {
			errChan <- struct{}{}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksChan := make(chan Task)
	var errChanLen int
	var mainErr error
	if m >= 0 {
		errChanLen = m
	} else {
		errChanLen = len(tasks)
	}
	errChan := make(chan struct{}, errChanLen)

	var errCauntVal int
	go func() {
		mu := sync.Mutex{}
		defer close(tasksChan)
		for _, task := range tasks {
			select {
			case <-errChan:
				if m >= 0 {
					mu.Lock()
					errCauntVal++
					if errCauntVal >= m {
						mu.Unlock()
						mainErr = ErrErrorsLimitExceeded
						return
					}
					mu.Unlock()
				}
				tasksChan <- task
			default:
				tasksChan <- task
			}
		}
	}()
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(tasksChan, errChan)
		}()
	}
	wg.Wait()
	return mainErr
}
