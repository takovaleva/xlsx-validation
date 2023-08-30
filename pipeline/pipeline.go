package pipeline

import (
	"fmt"
	"xlsx-validation/job"
)

type Pipeline struct {
	Jobs []job.Job
}

func NewPipeline(jobs ...job.Job) Pipeline {
	return Pipeline{
		Jobs: jobs,
	}
}

func (p *Pipeline) Start() {
	for _, job := range p.Jobs {
		job.Start()
		for result := range job.GetChannelInfo() {
			fmt.Println(result)
		}
	}
}
