package main

import (
	"github.com/Hudayberdyyev/concurrency-task/core"
	"log"
)

func main() {
	tasks := make([]core.Task, 2)
	manager := core.NewManager(tasks)

	err := manager.Run()
	if err != nil {
		log.Printf("error when doing tasks: %v\n", err)
	}

	progressChan := manager.Progress()
	completed := 0
	for len(progressChan) > 0 {
		<-progressChan
		completed++
	}
	log.Printf("%d tasks done\n", completed)
}
