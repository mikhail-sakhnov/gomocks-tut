package ifaces

type ITaskProcesser interface {
	ProcessTask([]int64) error
}
