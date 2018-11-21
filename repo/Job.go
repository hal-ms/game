package repo

import (
	"errors"

	"github.com/hal-ms/game/model"
)

var Job = jobRepo{job: model.Job{
	Jobs: []string{
		"cook",
		"pianist",
		"programmer",
		"carpenter",
		"priest",
	}, Job: "",
}}

type jobRepo struct {
	job model.Job
}

func (j *jobRepo) Get() model.Job {
	return j.job
}

func (j *jobRepo) Job(s string) {
	j.job.Job = s
}

func (j *jobRepo) Exist(s string) (bool, error) {
	for _, v := range j.Get().Jobs {
		if s == v {
			return true, nil
		}
	}
	return false, errors.New("Not Exist Job")
}
