package job

type Job interface {
	Start()
	Stop()
	GetChannelInfo() chan error
	AddDependency(jobs ...Job)
}

func NewDefaultJob() DefaultJob {
	return DefaultJob{
		Dependencies: make(map[Job]struct{}),
		Result:       make(chan error),
	}
}

type DefaultJob struct {
	Dependencies map[Job]struct{}
	// здесь нужно хранить результат всегда, а в GetChannelInfo отдавать копии,
	// чтобы можно было несколько раз подписаться на результаты
	Result chan error
}

func (d *DefaultJob) GetChannelInfo() chan error {
	return d.Result
}

func (d *DefaultJob) AddDependency(jobs ...Job) {
	for _, job := range jobs {
		d.Dependencies[job] = struct{}{}
	}
}

func (d *DefaultJob) Stop() {
	select {
	case <-d.Result:
	default:
		close(d.Result)
	}
}

func (d *DefaultJob) startDependencies() {
	for job := range d.Dependencies {
		job.Start()
		for result := range job.GetChannelInfo() {
			d.Result <- result
		}
	}
}
