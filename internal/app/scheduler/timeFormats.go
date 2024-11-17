package scheduler

import (
	"strconv"
	"strings"
	"time"
)

type TimeCellType int

const (
	Interval   TimeCellType = 1 // 时间间隔
	Clock      TimeCellType = 2 // 时刻
	ClockScope TimeCellType = 3 // 时刻范围
	Any        TimeCellType = 4 // 任意
)

type TimeCellFmt struct {
	Numbers  []int
	CellType TimeCellType
}

func (timeCell *TimeCellFmt) toCronFmt() string {
	var numStrs []string
	for _, number := range timeCell.Numbers {
		numStrs = append(numStrs, strconv.Itoa(number))
	}
	switch timeCell.CellType {
	case Interval:
		return "*/" + numStrs[0]
	case Clock:
		return numStrs[0]
	case ClockScope:
		return strings.Join(numStrs, ",")
	case Any:
		return "*"
	}
	return ""
}

type CyclicTimeFmt struct {
	Second     *TimeCellFmt
	Minute     *TimeCellFmt
	Hour       *TimeCellFmt
	DayOfMonth *TimeCellFmt
	Month      *TimeCellFmt
	DayOfWeek  *TimeCellFmt
}

func (cyclicTime *CyclicTimeFmt) ToCronFmt() string {
	cellStrs := []string{"*", "*", "*", "*", "*", "*"}
	if cyclicTime.Second != nil {
		second := cyclicTime.Second
		cellStrs[0] = cyclicTime.Second.toCronFmt()
		second.updateLowerUnits(cellStrs, 0)
	}
	if cyclicTime.Minute != nil {
		minute := cyclicTime.Minute
		cellStrs[1] = minute.toCronFmt()
		minute.updateLowerUnits(cellStrs, 1)
	}
	if cyclicTime.Hour != nil {
		hour := cyclicTime.Hour
		cellStrs[2] = hour.toCronFmt()
		hour.updateLowerUnits(cellStrs, 2)
	}
	if cyclicTime.DayOfMonth != nil {
		dayOfMonth := cyclicTime.DayOfMonth
		cellStrs[3] = dayOfMonth.toCronFmt()
		dayOfMonth.updateLowerUnits(cellStrs, 3)
	}
	if cyclicTime.Month != nil {
		month := cyclicTime.Month
		cellStrs[4] = month.toCronFmt()
		month.updateLowerUnits(cellStrs, 4)
	}
	if cyclicTime.DayOfWeek != nil {
		dayOfWeek := cyclicTime.DayOfWeek
		cellStrs[5] = dayOfWeek.toCronFmt()
		dayOfWeek.updateLowerUnits(cellStrs, 5)
	}
	return strings.Join(cellStrs, " ")
}

func (timeCell *TimeCellFmt) updateLowerUnits(cellStrs []string, index int) {
	if timeCell.CellType == Interval {
		now := time.Now()
		if index > 0 {
			cellStrs[0] = strconv.Itoa(now.Second())
		}
		if index > 1 {
			cellStrs[1] = strconv.Itoa(now.Minute())
		}
		if index > 2 {
			cellStrs[2] = strconv.Itoa(now.Hour())
		}
		if index > 3 {
			cellStrs[3] = strconv.Itoa(now.Day())
		}
	}
}
