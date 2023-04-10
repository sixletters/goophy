package scheduler

import (
	"errors"
)

type Scheduler interface {
	NumCurrent() int
	NewThread() ThreadID
	DeleteThread(id ThreadID) bool
	GetNextThread() (ThreadID, error)
}

type RoundRobinScheduler struct {
	threadChan   chan ThreadID
	threadQueue  []ThreadID
	currThreadID ThreadID
	MaxThreadID  ThreadID
	ThreadTable  map[ThreadID]Thread
}

func NewRoundRobinScheduler() *RoundRobinScheduler {
	return &RoundRobinScheduler{
		threadQueue:  make([]ThreadID, 0),
		threadChan:   make(chan ThreadID, 100),
		ThreadTable:  make(map[ThreadID]Thread),
		currThreadID: -1,
	}
}

func (r *RoundRobinScheduler) NumCurrent() int {
	return len(r.threadChan)
	// return len(r.threadQueue)
}

func (r *RoundRobinScheduler) GetCurrentThreads() []ThreadID {
	return r.threadQueue
}

func (r *RoundRobinScheduler) NewThread(thread Thread) ThreadID {
	curr := r.MaxThreadID
	r.ThreadTable[curr] = thread
	r.threadChan <- curr
	// r.threadQueue = append(r.threadQueue, curr)
	r.MaxThreadID += 1
	return curr
}

func (r *RoundRobinScheduler) AddThread(thread Thread) {
	r.threadChan <- r.currThreadID
	// r.threadQueue = append(r.threadQueue, r.currThreadID)
	r.ThreadTable[r.currThreadID] = thread
}

func (r *RoundRobinScheduler) GetCurrThreadID() ThreadID {
	return r.currThreadID
}

func (r *RoundRobinScheduler) GetNextThread() (ThreadID, error) {
	if len(r.threadChan) == 0 {
		return 0, errors.New("no more threads")
	}
	r.currThreadID = <-r.threadChan
	// if len(r.threadChan) > 1 {
	// 	r.threadQueue = r.threadQueue[1:]
	// } else {
	// 	r.threadQueue = []ThreadID{}
	// }
	return r.currThreadID, nil
}
