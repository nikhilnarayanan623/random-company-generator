package usecase

import (
	"runtime"
	"sync"
)

func FanIn[T any](done <-chan struct{}, streamsChannels ...<-chan T) <-chan T {

	fanInEmpChan := make(chan T, runtime.NumCPU())

	go func() {

		defer close(fanInEmpChan)
		var wg sync.WaitGroup

		for i := range streamsChannels {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					case fanInEmpChan <- <-streamsChannels[i]:
					}
				}
			}()
		}
		wg.Wait()
	}()

	return fanInEmpChan
}

// This is a generic function which take a done channel and a callback function.
// and will return a channel and you can infinitely receive callback function's return values through the channel. until the done channel close
func Generator[T any](done <-chan struct{}, fn func() T) <-chan T {

	stream := make(chan T, runtime.NumCPU())

	go func() {
		defer close(stream)

		// call the call back function to generate employee until the done channel close
		for {
			select {
			case <-done:
				return
			case stream <- fn(): // create and employee and send it the channel
			}
		}
	}()

	return stream
}

// This function to generate the callback function's data in multiple goroutines(expecting runtime.NumCPU).
// will return a channel all goroutines generated data will be on there.
// once closing the done channel or sending value to done channel will stop all the generation
func MultiGenerator[T any](done <-chan struct{}, numOfGoroutinesNeedToRun int, fn func() T) <-chan T {

	// channel to receive the data for the user
	stream := make(chan T, runtime.NumCPU())

	go func() {
		// close the stream when all the generating goroutines get done the job
		defer close(stream)
		var wg sync.WaitGroup

		// fire n number of goroutines to run generate data
		for i := 1; i <= numOfGoroutinesNeedToRun; i++ {

			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-done: // one's the done channel notify to stop then return
						return
						// else  send new data to the stream
					case stream <- fn():

					}
				}
			}()
		}
		// wait for all goroutines to complete.
		wg.Wait()
	}()

	return stream
}
