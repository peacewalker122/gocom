package util

import (
	"gocom/util/counter"
	"sync"
)

func freqWorker(done chan struct{}, rw *sync.RWMutex, file []byte, count counter.Maps, worker int) {
	chunkSize := len(file) / worker
	for i := 0; i < worker; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == worker-1 {
			end = len(file)
		}
		go func(start, end int) {
			rw.Lock()
			for _, b := range file[start:end] {
				count.Add(b)
			}
			rw.Unlock()
			done <- struct{}{}
		}(start, end)
	}
}
