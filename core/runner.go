package core

import (
	"github.com/google/uuid"
	"time"
)

const (
	RegisterDurationSec = 1
)

type Runner struct {
}

func (r *Runner) Register(task Task) (taskId uint32, err error) {
	time.Sleep(RegisterDurationSec * time.Second)
	return uuid.New().ID(), nil
}
