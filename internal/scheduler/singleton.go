package scheduler

import "sync"

var (
	scheduler *Scheduler
	sOnce     sync.Once
)

func InitScheduler() {
	sOnce.Do(
		func() {
			if scheduler == nil {
				scheduler = NewScheduler()
				go scheduler.TaskSignalListner()
				go scheduler.TaskPoolListener()
			}
		},
	)
}

func GetScheduler() *Scheduler {
	return scheduler
}
