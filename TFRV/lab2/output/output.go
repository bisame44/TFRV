package output

import "github.com/bisame44/TFRV/tree/main/TFRV/lab2/tasks"

type AlgoEfficiency struct {
	PerformanceTime   float64
	ScheduleTotalTime int
	ScheduleDeviation float64
	Schedule          [][]tasks.Task
}
