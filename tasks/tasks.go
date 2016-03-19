package tasks

import "fmt"
import "github.com/soider/gomocks-tut/ifaces"
import "sync"
import "math/rand"
import "time"

type MyProcesser struct{}

func (p *MyProcesser) ProcessTask(task []int64) error {
	fmt.Println("Processing task", task)
	return nil
}

func DoTasks(tasks [][]int64, processer ifaces.ITaskProcesser) {
	rand.Seed(time.Now().UTC().UnixNano())
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task []int64) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			processer.ProcessTask(task)
		}(task)
	}
	wg.Wait()

}
