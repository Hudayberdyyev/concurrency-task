package core

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tasks := make([]Task, 2)
	manager := NewManager(tasks)

	numCpu := runtime.GOMAXPROCS(0)
	startTime := time.Now()

	err := manager.Run()
	assert.NoError(t, err)

	elapsedTime := time.Since(startTime).Seconds()
	maxRequiredTime := float64((len(tasks)+1)*RegisterDurationSec) / float64(numCpu)
	assert.LessOrEqual(t, elapsedTime, maxRequiredTime)

	progressCount := 0
	for len(manager.progress) > 0 {
		<-manager.progress
		progressCount++
	}

	close(manager.progress)

	assert.Equal(t, progressCount, len(tasks))
}
