package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

// Used as a replacement of ENUMs
type TaskStatus string

const (
	InProgress TaskStatus = "In Progress"
	Pending    TaskStatus = "Pending"
	Completed  TaskStatus = "Completed"
)

// Task
type Task struct {
	id          int
	description string
	status      TaskStatus
	createdAt   time.Time
}

func (t *Task) isValidStatus() error {
	switch t.status {
	case InProgress, Pending, Completed:
		return nil
	default:
		return errors.New("invalid task status")
	}
}

// TaskQueue 
type TaskQueue struct {
	tasks []Task
	mutex sync.Mutex // for creating thread safe methods
}

func (tq *TaskQueue) AddTask(task Task) bool {
	err := task.isValidStatus()
	if err != nil {
		log.Fatal(err)
		return false
	}
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	tq.tasks = append(tq.tasks, task)
	return true
}

func (tq *TaskQueue) Retrieve(id int) (*Task, error) {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	for i := range tq.tasks {
		if tq.tasks[i].id == id {
			return &tq.tasks[i], nil // return original task pointer
		}
	}
	return &Task{}, errors.New("Task not found")
}
func (tq *TaskQueue) UpdateTask(id int, newTask Task) {
	err := newTask.isValidStatus()
	if err != nil {
		log.Fatal(err)
	}
	task, err := tq.Retrieve(id)
	if err != nil {
		log.Fatal("No Task found to update")
	}
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	task.description = newTask.description
	task.status = newTask.status
}

// Worker
type Worker struct{ 
	id int
}

// Scheduler
type Scheduler struct {
	TaskQueue chan Task
	Workers   []Worker
}

func main() {
	task := Task{
		id:          1,
		description: "BTP500 - DSA Assignment 1",
		status:      "In Progress",
		createdAt:   time.Now(),
	}

	tasksForTest := TaskQueue{}
	tasksForTest.AddTask(task)
	reTask, err := tasksForTest.Retrieve(1)
	if err != nil {
		log.Fatal(err)
	}
	print(reTask.description, "\n")
	task2 := Task{
		id:          2,
		description: "BTH545 - UI/UX workshop 6",
		status:      "Pending",
		createdAt:   time.Now(),
	}
	tasksForTest.AddTask(task2)
	// tasksForTest.UpdateTask(1, task2)

}