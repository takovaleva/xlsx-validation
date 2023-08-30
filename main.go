package main

import (
	"xlsx-validation/job"
	"xlsx-validation/pipeline"
)

func main() {
	rowJob := job.NewJobValidateFirstRow()
	pageJob := job.NewJobValidatePage()
	pageJob.AddDependency(rowJob)

	pipe := pipeline.NewPipeline(pageJob)
	pipe.Start()
}
