package scheduler

import (
	"strconv"
	"strings"

	"github.com/robfig/cron/v3"
)

type TaskFunc func(args ...interface{}) interface{}

type Task struct {
	Name        string         // 任务名称
	EntryId     cron.EntryID   // cron任务的EntryID
	CyclicTime  *CyclicTimeFmt // 任务周期性执行
	ExeCnt      int            // 任务已经执行次数
	ExeCntLimit int            // 任务执行总次数限制
	Function    TaskFunc       // 任务实事
	Args        []interface{}  // 任务实事的参数
}

// 获取当前已经执行的次数
func (t *Task) CurrentExeCount() int {
	return t.ExeCnt
}

func (task *Task) ToString() string {
	var ret strings.Builder
	ret.WriteString("Name: " + task.Name + ", ")
	ret.WriteString("EntryId: " + strconv.Itoa(int(task.EntryId)) + ", ")
	ret.WriteString("Execution Count: " + strconv.Itoa(task.CurrentExeCount()) + "/" + strconv.Itoa(task.ExeCntLimit) + ", ")
	ret.WriteString("Cycle: " + task.CyclicTime.ToCronFmt())
	return ret.String()
}
