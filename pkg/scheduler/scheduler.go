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
	MaxThreadID  ThreadID
	ThreadTable  map[ThreadID]Thread
}

func NewRoundRobinScheduler() *RoundRobinScheduler {
	return &RoundRobinScheduler{
		threadQueue:  make([]ThreadID, 0),
		ThreadTable:  make(map[ThreadID]Thread),
		currThreadID: -1,
	}
}

func (r *RoundRobinScheduler) NumCurrent() int {
	return len(r.threadQueue)
}

func (r *RoundRobinScheduler) GetCurrentThreads() []ThreadID {
	return r.threadQueue
}

func (r *RoundRobinScheduler) NewThread(thread Thread) ThreadID {
	curr := r.MaxThreadID
	r.ThreadTable[curr] = thread
	r.threadQueue = append(r.threadQueue, curr)
	r.MaxThreadID += 1
	return curr
}

func (r *RoundRobinScheduler) DeleteThread() {

}

func (r *RoundRobinScheduler) AddThread(thread Thread) {
	r.threadQueue = append(r.threadQueue, r.currThreadID)
	r.ThreadTable[r.currThreadID] = thread
}

func (r *RoundRobinScheduler) GetCurrThreadID() ThreadID {
	return r.currThreadID
}

func (r *RoundRobinScheduler) GetNextThread() (ThreadID, error) {
	if len(r.threadQueue) == 0 {
		return 0, errors.New("No more threads")
	}
	r.currThreadID = r.threadQueue[0]
	if len(r.threadQueue) > 1 {
		r.threadQueue = r.threadQueue[1:]
	} else {
		r.threadQueue = []ThreadID{}
	}
	return r.currThreadID, nil
}
