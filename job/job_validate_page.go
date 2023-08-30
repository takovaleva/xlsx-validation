package job

import "fmt"

type JobValidatePage struct {
	DefaultJob
}

func NewJobValidatePage() Job {
	return &JobValidatePage{
		DefaultJob: NewDefaultJob(),
	}
}

func (j JobValidatePage) Start() {
	go func() {
		j.startDependencies()

		fmt.Println("Start page validation")
		j.Result <- fmt.Errorf("Page invalid")
		j.Stop()
	}()
}
