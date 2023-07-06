package core

import (
	"github.com/google/uuid"
	"time"
)

type Runner struct {
}

func (r *Runner) Register(task Task) (taskId uint32, err error) {
	time.Sleep(2 * time.Second)
	return uuid.New().ID(), nil
}
