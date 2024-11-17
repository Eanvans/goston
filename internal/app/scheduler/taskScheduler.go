package scheduler

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/robfig/cron/v3"
)

type TaskScheduler struct {
	mu    sync.Mutex
	cron  *cron.Cron
	tasks map[string]*Task
}

func NewScheduler() *TaskScheduler {
	return &TaskScheduler{
		cron:  cron.New(cron.WithSeconds()),
		tasks: map[string]*Task{},
	}
}

func (scheduler *TaskScheduler) Start() {
	scheduler.cron.Start()
	log.Println("Task Scheduler Started")
}

func (scheduler *TaskScheduler) Stop() {
	scheduler.cron.Stop()
	log.Println("Task Scheduler Stopped")
}

// AddTask execLimit = -1 为无限次数
func (s *TaskScheduler) AddTask(taskName string, timeInterval *CyclicTimeFmt, execLimit int, taskFunc TaskFunc) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 如果已经存在同名任务，则不再重复加入
	if _, exist := s.tasks[taskName]; exist {
		return false
	}
	// 把 task 加入 scheduler.tasks
	task := &Task{
		Name:        taskName,
		ExeCntLimit: execLimit,
		CyclicTime:  timeInterval,
		Function:    taskFunc,
		Args:        []interface{}{taskName},
	}

	s.tasks[taskName] = task
	// 把 task 加入 scheduler.cron
	entryId, err := s.cron.AddFunc(task.CyclicTime.ToCronFmt(), func() {
		// 如果超出执行次数，则不再执行，并且删除任务
		if task.ExeCntLimit != -1 && task.ExeCnt >= task.ExeCntLimit {
			s.RemoveTask(taskName)
			log.Println(taskName + " has reached execution count " + strconv.Itoa(task.ExeCnt) + "/" + strconv.Itoa(task.ExeCntLimit) + ", will no longer execute")
			return
		}
		// 正常执行，并记录执行次数
		task.Function(task.Args...)
		task.ExeCnt++
	})
	if err != nil {
		return false
	}
	// 更新 EntryId
	task.EntryId = entryId
	return true
}

func (s *TaskScheduler) RemoveTask(taskName string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 如果不存在名为taskName的任务，则无法删除
	if _, exist := s.tasks[taskName]; !exist {
		return false
	}
	task := s.tasks[taskName]
	// 将 task 移出 scheduler.cron
	s.cron.Remove(task.EntryId)
	// 将 task 移出 scheduler.tasks
	delete(s.tasks, taskName)
	return true
}

func (scheduler *TaskScheduler) ListTasks() string {
	var ret strings.Builder
	ret.WriteString("Scheduled Tasks:\n")
	for _, task := range scheduler.tasks {
		ret.WriteString(task.ToString() + "\n")
	}
	return ret.String()
}
