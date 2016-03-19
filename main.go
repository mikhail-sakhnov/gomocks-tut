package main

import "github.com/soider/gomocks-tut/tasks"
import "fmt"

type MyProcesser struct{}

func (p *MyProcesser) ProcessTask(task []int64) error {
	fmt.Println("Processing task", task)
	return nil
}

func main() {
	tasksSlice := [][]int64{
		[]int64{1, 2},
		[]int64{3, 4},
	}
	tasks.DoTasks(tasksSlice, &MyProcesser{})
}
