package dsp

import "github.com/Belstowe/distrib-cs-1-autumn/lab2/tasks"

func TotalTime(taskLevels [][]tasks.Task) int {
	tt := 0
	for _, taskLevel := range taskLevels {
		tt += taskLevel[0].T
	}
	return tt
}

func TimeBound(taskLevels [][]tasks.Task, n int) float64 {
	tavg := 0
	for _, taskLevel := range taskLevels {
		for _, task := range taskLevel {
			tavg += task.T * task.R
		}
	}
	return float64(tavg) / float64(n)
}

func FnDeviation(taskLevels [][]tasks.Task, n int) float64 {
	return (float64(TotalTime(taskLevels)) - TimeBound(taskLevels, n)) / TimeBound(taskLevels, n)
}
