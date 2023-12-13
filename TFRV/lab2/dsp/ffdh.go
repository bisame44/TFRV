package dsp

import (
	"fmt"

	"github.com/bisame44/TFRV/tree/main/TFRV/lab2/tasks"
)

func FFDH(records []tasks.Task, n int) ([][]tasks.Task, error) {
	oversize := make([]tasks.Task, 0)
	for _, record := range records {
		if record.R > n {
			oversize = append(oversize, record)
		}
	}
	if len(oversize) != 0 {
		return nil, fmt.Errorf("only having %d processes, while tasks %v require more", n, oversize)
	}

	taskLevels := make([][]tasks.Task, 0)
	levelTakenProc := make([]int, 0)
outerLoop:
	for _, task := range records {
		for i := range taskLevels {
			if levelTakenProc[i]+task.R <= n {
				levelTakenProc[i] += task.R
				taskLevels[i] = append(taskLevels[i], task)
				continue outerLoop
			}
		}
		taskLevels = append(taskLevels, []tasks.Task{task})
		levelTakenProc = append(levelTakenProc, task.R)
	}

	return taskLevels, nil
}
