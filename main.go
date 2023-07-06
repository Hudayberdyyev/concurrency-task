package main

import (
	"github.com/Hudayberdyyev/concurrency-task/core"
	"log"
)

func main() {
	tasks := make([]core.Task, 5)
	manager := core.NewManager(tasks)

	go func() {
		err := manager.Run()
		if err != nil {
			log.Printf("error when doing tasks: %v\n", err)
		}
	}()

	for progress := range manager.Progress() {
		log.Printf("progress = %d\n", progress)
	}
}
