package scheduler

type Scheduler interface {
	NumCurrent() int
	GetCurrentThreads() []ThreadID
	NumIdle() int
	GetIdleThreads() []ThreadID
	NewThread() ThreadID
	DeleteThread(id ThreadID) bool
	GetRunThread() (ThreadID, int, error)
	PauseThread(id ThreadID) bool
}

type RoundRobinScheduler struct {
	currentThreads map[ThreadID]bool
	idleThreads    []ThreadID
	maxThreadID    int
	maxTimeQuanta  int
}

func (r *RoundRobinScheduler) NumCurrent() int {
	return len(r.currentThreads)
}

func (r *RoundRobinScheduler) GetCurrentThreads() []ThreadID {
	ret := make([]ThreadID, 0)
	for k, v := range r.currentThreads {

	}
}

func (r *RoundRobinScheduler) NumIdle() int {

}

func (r *RoundRobinScheduler) GetIdleThreads() []ThreadID {

}

func (r *RoundRobinScheduler) NewThread() ThreadID {

}

func (r *RoundRobinScheduler) DeleteThread(id ThreadID) bool {

}

func (r *RoundRobinScheduler) GetRunThread() (ThreadID, int, error) {

}

func (r *RoundRobinScheduler) PauseThread(id ThreadID) (ThreadID, int, error) {

}
