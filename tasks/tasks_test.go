package tasks

import "testing"
import "github.com/golang/mock/gomock"
import "github.com/soider/gomocks-tut/mocks"

func TestMeTestDoTasksSchedulesEverythingCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	tasks := [][]int64{
		[]int64{1, 2},
		[]int64{3, 4},
	}
	taskProcessor := mocks.NewMockITaskProcesser(mockCtrl)
	taskProcessor.EXPECT().ProcessTask([]int64{3, 4})
	taskProcessor.EXPECT().ProcessTask([]int64{1, 2})

	DoTasks(tasks, taskProcessor)
}
