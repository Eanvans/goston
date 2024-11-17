package scheduler

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestScheduler(test *testing.T) {
	myScheduler := NewScheduler()
	return
	// CellType == Interval 的只能有一个单位
	// 比如”每1分15秒执行一次“，应该在Second设置Interval==75，而不是同时在Second设置Interval==15并且在Minute设置Interval==1
	myScheduler.AddTask("每2秒", &CyclicTimeFmt{
		Second: &TimeCellFmt{Numbers: []int{2}, CellType: Interval},
	}, 8, func(args ...interface{}) (ret interface{}) {
		log.Println(args[0].(string))
		return
	})
	myScheduler.AddTask("每1分", &CyclicTimeFmt{
		Minute: &TimeCellFmt{Numbers: []int{1}, CellType: Interval},
	}, 8, func(args ...interface{}) (ret interface{}) {
		log.Println(args[0].(string))
		return
	})
	myScheduler.AddTask("每1个月", &CyclicTimeFmt{
		Month: &TimeCellFmt{Numbers: []int{1}, CellType: Interval},
	}, 8, func(args ...interface{}) (ret interface{}) {
		log.Println(args[0].(string))
		return
	})
	myScheduler.AddTask("每天凌晨03:36:15", &CyclicTimeFmt{
		Second: &TimeCellFmt{Numbers: []int{15}, CellType: Clock},
		Minute: &TimeCellFmt{Numbers: []int{36}, CellType: Clock},
		Hour:   &TimeCellFmt{Numbers: []int{3}, CellType: Clock},
	}, 8, func(args ...interface{}) (ret interface{}) {
		log.Println(args[0].(string))
		return
	})

	myScheduler.Start()
	time.Sleep(2 * time.Second)
	fmt.Println(myScheduler.ListTasks())
	time.Sleep(30 * time.Second)
	myScheduler.RemoveTask("每3秒")
	fmt.Println(myScheduler.ListTasks())
	time.Sleep(5 * time.Minute)
}
