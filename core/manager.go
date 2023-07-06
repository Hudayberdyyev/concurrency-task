package core

import "sync"

type Manager struct {
	runner    Runner
	tasks     []Task
	progress  chan uint8
	completed uint8
}

func NewManager(tasks []Task) *Manager {
	return &Manager{
		tasks:    tasks,
		runner:   Runner{},
		progress: make(chan uint8, len(tasks)),
	}
}

func (m *Manager) Run() error {
	wg := sync.WaitGroup{}
	for index, task := range m.tasks {
		taskId, err := m.runner.Register(task)
		m.completed++
		wg.Add(1)
		go func(completed uint8) {
			defer wg.Done()
			m.progress <- completed
		}(m.completed)
		if err != nil {
			return err
		}
		m.tasks[index].taskId = taskId
	}

	wg.Wait()
	return nil
}

func (m *Manager) Progress() <-chan uint8 {
	return m.progress
}
