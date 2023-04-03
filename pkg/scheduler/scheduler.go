package scheduler

import (
	"errors"
)

type Scheduler interface {
	NumCurrent() int
	GetCurrentThreads() []ThreadID
	NewThread() ThreadID
	DeleteThread(id ThreadID) bool
	GetNextThread() (ThreadID, error)
}

type RoundRobinScheduler struct {
	threadQueue  []ThreadID
	currThreadID ThreadID
	ThreadTable  map[ThreadID]Thread
}

func NewRoundRobinScheduler() *RoundRobinScheduler {
	return &RoundRobinScheduler{
		threadQueue: make([]ThreadID, 0),
		ThreadTable: make(map[ThreadID]Thread),
	}
}

func (r *RoundRobinScheduler) NumCurrent() int {
	return len(r.threadQueue)
}

func (r *RoundRobinScheduler) GetCurrentThreads() []ThreadID {
	return r.threadQueue
}

func (r *RoundRobinScheduler) NewThread(thread Thread) ThreadID {
	curr := r.currThreadID
	r.threadQueue = append(r.threadQueue, curr)
	r.currThreadID += 1
	return curr
}

func (r *RoundRobinScheduler) DeleteThread() {

}

func (r *RoundRobinScheduler) AddThread(thread Thread) {
	r.threadQueue = append(r.threadQueue, r.currThreadID)
	r.ThreadTable[r.currThreadID] = thread
	r.currThreadID += 1
	return
}

func (r *RoundRobinScheduler) GetNextThread() (ThreadID, error) {
	if len(r.threadQueue) == 0 {
		return 0, errors.New("NO more threads")
	}
	r.currThreadID = r.threadQueue[0]
	if len(r.threadQueue) > 1 {
		r.threadQueue = r.threadQueue[1:]
	} else {
		r.threadQueue = []ThreadID{}
	}
	return r.currThreadID, nil
}
