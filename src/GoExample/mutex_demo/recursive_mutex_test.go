package main

import (
	"sync"
	"testing"
)

func TestRecursiveMutexLock(t *testing.T) {
	var mutex RecursiveMutex
	mutex.Lock()
	mutex.Lock()
	mutex.Unlock()
	mutex.Unlock()
}

func TestMutexLock(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	mutex.Lock()
	mutex.Unlock()
	mutex.Unlock()
}
