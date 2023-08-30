package job

import "fmt"

type JobValidateFirstRow struct {
	DefaultJob
}

func NewJobValidateFirstRow() Job {
	return &JobValidateFirstRow{
		DefaultJob: NewDefaultJob(),
	}
}

func (j JobValidateFirstRow) Start() {
	go func() {
		j.startDependencies()

		fmt.Println("Start first row validation")
		j.Result <- fmt.Errorf("First row invalid")
		j.Stop()
	}()
}
